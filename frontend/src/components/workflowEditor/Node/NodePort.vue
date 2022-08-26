<template>
  <div :key="updateKey">
    <!-- Port and Title -->
    <div class="fill-width flex relative">
      <!-- Trigger Port -->
      <div
        v-if="portModel.data.is_trigger"
        ref="portRef"
        :id="portModel.id"
        class="arrow-right absolute transform translate-y-1"
        :class="[{ right: '0px' }]"
        :style="[
          { width: portModel.portSize + 'px', height: portModel.portSize + 'px' },
          portModel.data.is_input
            ? 'left: -' + portModel.portSize / 3 + 'px;'
            : 'right: -' + portModel.portSize / 2 + 'px;',
        ]"
      ></div>

      <!-- Normal Port -->
      <div
        v-else
        ref="portRef"
        :id="portModel.id"
        class="circle absolute transform translate-y-1"
        :class="{ right: '0px' }"
        :style="[
          { width: portModel.portSize + 'px', height: portModel.portSize + 'px' },
          { backgroundColor: portColor },
          portModel.data.is_input
            ? 'left: -' + portModel.portSize / 2 + 'px;'
            : 'right: -' + portModel.portSize / 2 + 'px;',
        ]"
      ></div>

      <!-- Port Label -->
      <span class="justify-center flex-auto mx-3" :class="[portModel.data.is_input ? 'text-left' : 'text-right']">{{
        portModel.data.label
      }}</span>
    </div>

    <!-- Default Text Field -->
    <div
      v-if="
        portModel.data.hasDefaultField && !portModel.data.defaultFieldHidden && portModel.data.datatype != 'BOOLEAN'
      "
      class="px-2 pb-3"
    >
      <textarea
        ref="textAreaRef"
        class="bg-gray-200 w-full px-1"
        v-model="textAreaValue"
        :placeholder="portModel.data.defaultPlaceholder"
        :style="[{ resize: 'none', height: textAreaHeight + 'px', minHeight: '24px' }]"
      ></textarea>
    </div>

    <!-- Default Checkbox -->
    <div
      v-if="
        portModel.data.hasDefaultField && !portModel.data.defaultFieldHidden && portModel.data.datatype == 'BOOLEAN'
      "
      class="px-2 pb-3"
    >
      <input type="checkbox" v-model="checkboxValue" />
    </div>

    <!-- Default Select -->
    <div v-if="portModel.data.options" class="px-2 pb-3">
      <select class="bg-gray-200 w-full px-1" v-model="selectValue">
        <option disabled value>Select Method</option>
        <option v-for="option in portModel.data.options" :key="option">{{ option }}</option>
      </select>
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
import { EdgeType } from "@/models/Data/Types";
import { Datatype, DatatypeColors, Datatypes } from "@/models/Data/DataTypes";

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

    const textAreaValue = ref(props.portModel.data.default_value.value);
    const textAreaHeight = ref(24);
    const portColor = ref<string>(DatatypeColors[props.portModel.data.datatype]);
    const selectValue = ref<string>("");
    const checkboxValue = ref<boolean>(props.portModel.data.default_value.value as boolean);

    const node = props.portModel.parent as NodeModel;
    const grid = props.portModel.parent?.parent as GridModel;

    const updateKey = ref(0);

    onMounted(() => {
      if (!props.portModel.parent?.parent) return;
      if (!portColor.value) portColor.value = DatatypeColors[Datatype.ANY]; // set Default color for old workflows TODO delete
      if (props.portModel.data.datatype == Datatype.BOOLEAN && !checkboxValue.value) checkboxValue.value = false;

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

    emitter.on("UpdatePort", (id) => {
      if (props.portModel.data.id == id) {
        portColor.value = DatatypeColors[props.portModel.data.datatype];
        updateKey.value++;
      }
    });

    // Resize Text Area
    watch(textAreaValue, () => {
      props.portModel.setDefaultValue(textAreaValue.value, props.portModel.data.datatype as Datatype);

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

    watch(selectValue, () => {
      props.portModel.setDefaultValue(selectValue.value, Datatype.STRING);

      if (props.portModel.data.identifier == "parse") {
        let node = props.portModel.parent as NodeModel;
        let port = node.children.find((child) => (child as PortModel).data.identifier == "output") as PortModel;
        port.data.datatype = selectValue.value as Datatype;

        emitter.emit("UpdatePort", port.data.id);
      }
    });

    watch(checkboxValue, () => {
      props.portModel.setDefaultValue(checkboxValue.value, Datatype.BOOLEAN);
    });

    // Init Port Pos
    function setPortPos() {
      if (!grid || !node || !portRef.value) return;
      const rect: DOMRect = portRef.value.getBoundingClientRect();

      const posAbs = new Vector2(rect.x + rect.width / 2, rect.y + rect.height / 2);
      const pos = Vector2.subtract(posAbs, node.posGridCell, grid.posRel, new Vector2(0, 64)); //64 is the Size of the Navbar, quick and dirty

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
          props.portModel.setDefaultFieldHidden(false);
        }
      } else {
        (grid as GridModel).addEdge(EdgeModel.NewEdgeFromPort(props.portModel), true);
      }
    }

    function onDragMove(event: InteractEvent) {
      if (!grid) return;
      if (grid.tmpEdgeIndex >= 0) {
        grid.edges[grid.tmpEdgeIndex].setMousePos(new Vector2(event.clientX, event.clientY - 64)); // quick and dirty
      }
    }

    function onDragEnd() {
      if (!props.portModel.parent?.parent) return;
      grid.resetTmp(true);
    }

    // event.target         = the Element on which it gets dropped
    // event.relatedTarget  = the Element which dropped
    function onDrop(event: InteractEvent) {
      if (!grid || !event.relatedTarget) return;

      const portIn = props.portModel;
      const portOut = grid.getPortByID(event.relatedTarget.id);

      if (portOut && portIn.data.is_input && Datatypes.allowedConnection(portOut.data.datatype, portIn.data.datatype)) {
        // Connected to an Input
        const connection = grid.getTmpEdge();
        // Connected Port Found
        connection.setPortIn(portIn);

        // Check if this connection is a duplicate
        let edge: EdgeType[] = GridData.grid.data.edges.filter(
          (edge) =>
            edge.origin.port_id == connection.data.origin.port_id &&
            edge.target.port_id == connection.data.target.port_id
        );
        if (edge.length > 1) {
          // Edge is Duplicate -> Delete
          grid.resetTmp(true);
        } else {
          // Save Edge
          portIn.setDefaultFieldHidden(true);
          grid.resetTmp();
        }
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
      selectValue,
      portColor,
      checkboxValue,
      updateKey,
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
.arrow-right {
  width: 0;
  height: 0;
  border-top: 10px solid transparent;
  border-bottom: 10px solid transparent;

  border-left: 15px solid rgb(54, 54, 54);
}
</style>