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
          portModel.isInput ? 'left: -' + portModel.portSize / 2 + 'px;' : 'right: -' + portModel.portSize / 2 + 'px;',
        ]"
      ></div>
      <span class="justify-center flex-auto mx-3" :class="[portModel.isInput ? 'text-left' : 'text-right']">{{
        portModel.title
      }}</span>
    </div>

    <!-- Default Text Field -->
    <div v-if="portModel.hasDefaultField" class="px-2">
      <textarea
        ref="textAreaRef"
        class="bg-gray-200 w-full px-1"
        v-model="textAreaValue"
        :placeholder="portModel.placeholder"
        :style="[{ resize: 'none', height: portModel.textAreaScrollHeight + 'px', minHeight: '24px'}]"
      ></textarea>
    </div>
  </div>
</template>

<script lang="ts">
import Vector2 from "@/components/util/Vector";
import GridModel from "@/models/GridModel";
import NodeConnectionModel from "@/models/Node/NodeConnectionModel";
import NodeModel from "@/models/Node/NodeModel";
import NodePortModel from "@/models/Node/NodePortModel";
import { InteractEvent } from "@interactjs/types";
import interact from "interactjs";
import { defineComponent, nextTick, onMounted, ref, watch } from "vue";

export default defineComponent({
  components: {},
  props: {
    portModel: {
      type: NodePortModel,
      required: true,
    },
  },
  setup(props) {
    const portRef = ref<HTMLInputElement>();
    const textAreaRef = ref<HTMLInputElement>();

    const textAreaValue = ref("");

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

    // Resize Text Area
    watch(textAreaValue, () => {
      props.portModel.changeTextAreaHeight(24, true); // Change to 0 to get accurate ScrollHeight to shrink textArea, pretty stupid System,
      nextTick(function () {  // Wait one Tick for Style to take effect
        let newHeight = textAreaRef.value?.scrollHeight;
        if (newHeight) {
          props.portModel.changeTextAreaHeight(newHeight, true);  // Change Height to real value
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
      if (props.portModel.isInput) {
        const connection: NodeConnectionModel | undefined = grid.getConnection(undefined, props.portModel.id);
        if (connection) {
          connection.setPortIn(undefined);
          grid.setTmp(connection.id);
        }
      } else {
        (grid as GridModel).addConnection(
          new NodeConnectionModel(props.portModel, undefined, new Vector2(event.clientX, event.clientY)),
          true
        );
      }
    }

    function onDragMove(event: InteractEvent) {
      if (!grid) return;
      if (grid.tmpConnectionIndex >= 0) {
        grid.connections[grid.tmpConnectionIndex].setMousePos(new Vector2(event.clientX, event.clientY));
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
      if (grid.getPortByID(event.target.id)?.isInput) {
        const connection = grid.getTmpConnection();
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
