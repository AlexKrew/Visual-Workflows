<template>
  <div class="fill-width flex">
    <div ref="portRef" :id="portModel.id" class="circle fill-width bg-gray-400" :class="circlePosClass">
      <!-- <div
        v-if="isDragging"
        class="circle fill-width absolute drop-number bg-blue-700"
        :style="{ left: mousePosRel.x + 'px', top: mousePosRel.y + 'px' }"
      ></div> -->
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

    onMounted(() => {
      if (!props.portModel.node?.grid) return;
      setPortPos();
      interact(`#${props.portModel.id}`)
        .draggable({})
        .on("dragstart", onDragStart)
        .on("dragmove", onDragMove)
        .on("dragend", onDragEnd)
        .dropzone({})
        .on("drop", onDrop);
      // .on("dragenter", onDragEnter);
    });

    function setPortPos() {
      if (!props.portModel.node?.grid) return;
      if (portRef.value) {
        const rect: DOMRect = portRef.value.getBoundingClientRect();

        const posAbs = new Vector2(rect.x + rect.width / 2, rect.y + rect.height / 2);
        const pos = Vector2.subtract(posAbs, props.portModel.node.gridPos, props.portModel.node.grid.posAbs);

        props.portModel.setPos(pos);
      }
    }

    function onDragStart(event: InteractEvent) {
      if (!props.portModel.node?.grid) return;
      if (props.portModel.isInput) {
        const connection: NodeConnectionModel | undefined = props.portModel.node.grid.getConnection(
          undefined,
          props.portModel.id
        );
        if (connection) {
          connection.setPortIn(undefined);
          props.portModel.node.grid.setTmp(connection.id);
        }
      } else {
        props.portModel.node.grid.addConnection(
          new NodeConnectionModel(props.portModel, undefined, new Vector2(event.clientX, event.clientY)),
          true
        );
      }
    }

    function onDragMove(event: InteractEvent) {
      if (!props.portModel.node?.grid) return;
      if (props.portModel.node.grid.tmpConnectionIndex >= 0) {
        props.portModel.node.grid.connections[props.portModel.node.grid.tmpConnectionIndex].setMousePos(
          new Vector2(event.clientX, event.clientY)
        );
      }
    }

    function onDragEnd(event: InteractEvent) {
      if (!props.portModel.node?.grid) return;
      props.portModel.node.grid.resetTmp(true);
    }

    function onDrop(event: InteractEvent) {
      if (!props.portModel.node?.grid) return;
      // event.target         = the Element on which it gets dropped
      // event.relatedTarget  = the Element which dropped
      if (props.portModel.node.grid.getPortByID(event.target.id)?.isInput) {
        const connection = props.portModel.node.grid.getTmpConnection();
        connection.setPortIn(props.portModel.node.grid.getPortByID(event.target.id));
        props.portModel.node.grid.resetTmp();
      } else {
        props.portModel.node.grid.resetTmp(true);
      }
    }

    return {
      circlePosClass,
      portRef,
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
