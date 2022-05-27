<template>
  <div class="fill-width flex">
    <div ref="portRef" :id="portModel.id" class="circle fill-width bg-gray-400" :class="circlePosClass">
      <div
        v-if="isDragging"
        class="circle fill-width absolute drop-number bg-blue-700"
        :style="{ left: mousePosRel.x + 'px', top: mousePosRel.y + 'px' }"
      ></div>
    </div>
    <span class="justify-center flex-auto">{{ portModel.title }}</span>
  </div>
</template>

<script lang="ts">
import Vector2 from "@/components/util/Vector";
import NodeConnectionModel from "@/models/NodeConnectionModel";
import NodePortModel from "@/models/NodePortModel";
import { InteractEvent } from "@interactjs/types";
import interact from "interactjs";
import { defineComponent, onMounted, ref } from "vue";

export default defineComponent({
  components: {},
  props: {
    portModel: {
      type: NodePortModel,
      required: true,
    },
  },
  setup(props) {
    let circlePosClass = props.portModel.isInput ? "" : "absolute right-0";
    const portRef = ref<HTMLInputElement>();
    let isDragging = ref<boolean>(false);
    let mousePosRel = ref<Vector2>(new Vector2(0, 0));
    let portOut: NodePortModel | null = null; // Reference to PortOut when Dragging from Input Port

    onMounted(() => {
      setPortPos();
      interact(`#${props.portModel.id}`)
        .draggable({})
        .on("dragstart", onDragStart)
        .on("dragmove", onDragMove)
        .on("dragend", onDragEnd)
        .dropzone({})
        .on("drop", onDrop)
        .on("dragenter", onDragEnter);
    });

    function setPortPos() {
      if (portRef.value) {
        const rect: DOMRect = portRef.value.getBoundingClientRect();

        const posAbs = new Vector2(rect.x + rect.width / 2, rect.y + rect.height / 2);
        const pos = Vector2.subtract(posAbs, props.portModel.node.gridPos, props.portModel.node.grid.posAbs);

        props.portModel.setPos(pos);
      }
    }

    //#region +++++ DragHandler +++++
    function onDragStart(event: InteractEvent) {
      if (props.portModel.isInput) {
        if (props.portModel.connections.length <= 0) return;
        console.log("#");
        portOut = props.portModel.connections[0].portOut;
        portOut.moveConnectionToTmp(props.portModel.connections[0].id);
        isDragging.value = true;
        props.portModel.clearConnections();
      } else {
        props.portModel.setTmpConnection(
          new NodeConnectionModel(props.portModel, undefined, new Vector2(event.clientX, event.clientY))
        );
        isDragging.value = true;
      }
    }
    function onDragMove(event: InteractEvent) {
      const mousePos = new Vector2(event.clientX, event.clientY);
      mousePosRel.value = Vector2.subtract(mousePos, props.portModel.gridPos, props.portModel.node.grid.posAbs);
      if (props.portModel.isInput) {
        if(!portOut) return;
        portOut.tmpConnection?.setMousePos(mousePos);
      } else {
        props.portModel.tmpConnection?.setMousePos(mousePos);
      }
    }
    function onDragEnd(event: InteractEvent) {
      isDragging.value = false;
      if (props.portModel.isInput) {
        if(!portOut) return;
        portOut.setTmpConnection(null);
      } else {
        props.portModel.setTmpConnection(null);
      }
    }
    //#endregion +++++ +++++ +++++

    //#region +++++ Dropzone Handler +++++
    function onDragEnter(event: InteractEvent) {
      console.log("Hey");
    }
    function onDrop(event: InteractEvent) {
      if (!event.relatedTarget) return;

      let portOut = props.portModel.node.grid.getPortByID(event.relatedTarget.id);
      if (!portOut || !portOut.tmpConnection) return;

      portOut.saveTmpConnection(props.portModel);
    }
    //#endregion +++++ +++++ +++++

    return {
      circlePosClass,
      portRef,
      isDragging,
      mousePosRel,
    };
  },
});
</script>

<style>
.circle {
  width: 15px;
  height: 15px;
  -webkit-border-radius: 25px;
  -moz-border-radius: 25px;
  border-radius: 25px;
}
</style>
