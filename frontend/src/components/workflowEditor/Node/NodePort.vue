<template>
  <div>
    <!-- Port and Title -->
    <div class="fill-width flex relative">
      <div
        ref="portRef"
        :id="portModel.id"
        class="circle bg-blue-700 absolute transform translate-y-1"
        :class="{ right: '0px' }"
        :style="[
          { width: portModel.portSize + 'px', height: portModel.portSize + 'px' },
          portModel.data.is_input
            ? 'left: -' + portModel.portSize / 2 + 'px;'
            : 'right: -' + portModel.portSize / 2 + 'px;',
        ]"
      ></div>
      <span class="justify-center flex-auto mx-3" :class="[portModel.data.is_input ? 'text-left' : 'text-right']">{{
        portModel.data.label
      }}</span>
    </div>

    <!-- Default Text Field -->
    <div v-if="portModel.data.hasDefaultField" class="px-2">
      <textarea
        ref="textAreaRef"
        class="bg-gray-200 w-full px-1"
        v-model="textAreaValue"
        :placeholder="portModel.data.defaultPlaceholder"
        :style="[{ resize: 'none', height: textAreaHeight + 'px', minHeight: '24px' }]"
      ></textarea>
    </div>
  </div>
</template>

<script lang="ts">
import Vector2 from "@/components/util/Vector";
import GridModel from "@/models/GridModel";
import EdgeModel from "@/models/Node/EdgeModel";
import NodeModel from "@/models/Node/NodeModel";
import PortModel from "@/models/Node/PortModel";
import { InteractEvent } from "@interactjs/types";
import interact from "interactjs";
import { defineComponent, nextTick, onMounted, onUnmounted, ref, watch } from "vue";
import { emitter } from "@/components/util/Emittery";
import GridData from "@/models/Data/GridData";

export default defineComponent({
  components: {},
  props: {
    portModel: {
      type: PortModel,
      required: true,
    },
  },
  setup(props) {
    const portRef = ref<HTMLInputElement>();
    const textAreaRef = ref<HTMLInputElement>();

    const textAreaValue = ref(props.portModel.data.defaultValue);
    const textAreaHeight = ref(24);

    const node = props.portModel.parent as NodeModel;
    const grid = props.portModel.parent?.parent as GridModel;

    onMounted(() => {
      if (!props.portModel.parent?.parent) return;
      setPortPos();
      interact(`#${props.portModel.id}`)
        .draggable({})
        .on("dragstart", onDragStart)
        .on("dragmove", onDragMove)
        .on("dragend", onDragEnd)
        .dropzone({})
        .on("drop", onDrop);
    });

    onUnmounted(() => {
      interact(`#${props.portModel.id}`).unset();
    });

    emitter.on("PortsUpdatePos", (parent) => {
      if (parent == props.portModel.parent) setPortPos();
    });

    // Resize Text Area
    watch(textAreaValue, () => {
      props.portModel.setDefaultValue(textAreaValue.value);

      let oldHeight = textAreaHeight.value;
      textAreaHeight.value = 0; // Change to 0 to get accurate ScrollHeight to shrink textArea, pretty stupid System,

      nextTick(function () {
        // Wait one Tick for Style to take effect
        let newHeight = textAreaRef.value?.scrollHeight;
        if (newHeight) {
          textAreaHeight.value = newHeight; // Change Height to real value
          if (oldHeight != newHeight) emitter.emit("PortsUpdatePos", props.portModel.parent as NodeModel);
        }
      });
    });

    // Init Port Pos
    function setPortPos() {
      if (!grid || !node || !portRef.value) return;
      const rect: DOMRect = portRef.value.getBoundingClientRect();

      const posAbs = new Vector2(rect.x + rect.width / 2, rect.y + rect.height / 2);
      const pos = Vector2.subtract(posAbs, node.posGridCell, grid.posRel);

      props.portModel.setPos(pos);
    }

    //#region InteractJS
    function onDragStart(event: InteractEvent) {
      if (!grid) return;
      if (props.portModel.data.is_input) {
        const connection: EdgeModel | undefined = grid.getEdge(undefined, props.portModel.id);
        if (connection) {
          connection.setPortIn(undefined);
          grid.setTmp(connection.data.id);
        }
      } else {
        (grid as GridModel).addEdge(EdgeModel.NewEdgeFromPort(props.portModel), true);
      }
    }

    function onDragMove(event: InteractEvent) {
      if (!grid) return;
      if (grid.tmpEdgeIndex >= 0) {
        grid.edges[grid.tmpEdgeIndex].setMousePos(new Vector2(event.clientX, event.clientY));
      }
    }

    function onDragEnd() {
      if (!props.portModel.parent?.parent) return;
      grid.resetTmp(true);
    }

    function onDrop(event: InteractEvent) {
      if (!grid) return;

      // event.target         = the Element on which it gets dropped
      // event.relatedTarget  = the Element which dropped
      if (grid.getPortByID(event.target.id)?.data.is_input) {
        const connection = grid.getTmpEdge();
        connection.setPortIn(grid.getPortByID(event.target.id));
        grid.resetTmp();
      } else {
        grid.resetTmp(true);
      }
    }
    //#endregion

    return {
      portRef,
      textAreaRef,
      textAreaValue,
      textAreaHeight,
    };
  },
});
</script>

<style>
.circle {
  border-radius: 2500px;
  -webkit-border-radius: 2500px;
  -moz-border-radius: 2500px;
}
</style>
