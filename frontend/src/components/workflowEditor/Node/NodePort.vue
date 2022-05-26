<template>
  <div class="fill-width flex">
    <div ref="portRef" :id="portModel.id" class="circle fill-width" :class="circlePosClass">
      <!-- <NodeConnection v-if="portModel.tmpConnection" :connection="portModel.tmpConnection" class="relative right-0"/> -->
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
import NodeConnection from "./NodeConnection.vue";

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
      setPortPos();

      interact(`#${props.portModel.id}`)
        .draggable({})
        .on("dragstart", onDragStart)
        .on("dragmove", onDragMove)
        .on("dragend", onDragEnd);
    });

    function setPortPos() {
      if (portRef.value) {
        const rect: DOMRect = portRef.value.getBoundingClientRect();

        const posAbs = new Vector2(rect.x + (rect.width/2), rect.y + (rect.height/2));
        const pos = Vector2.subtract(posAbs, props.portModel.node.gridPos, props.portModel.node.grid.posAbs);

        console.log(posAbs)
        console.log(pos)
        
        props.portModel.setPos(pos);
      }
    }

    function onDragStart(event: InteractEvent) {
      if (props.portModel.isInput) {
        // TODO Input
      } else {
        const connection = new NodeConnectionModel(
          props.portModel,
          undefined,
          new Vector2(event.clientX, event.clientY)
        );
        props.portModel.setTmpConnection(connection);
      }
    }
    function onDragMove(event: InteractEvent) {
      if (props.portModel.isInput) {
        // ToDo Input
      } else {
        const mousePos = new Vector2(event.clientX, event.clientY);
        props.portModel.tmpConnection?.setMousePos(mousePos);
      }
    }
    function onDragEnd(event: InteractEvent) {
      // props.portModel.setTmpConnection(null);
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
  width: 20px;
  height: 20px;
  -webkit-border-radius: 25px;
  -moz-border-radius: 25px;
  border-radius: 25px;
  background: grey;
}
</style>
