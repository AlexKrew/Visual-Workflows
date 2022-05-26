<template>
  <div :id="nodeModel.id" class="max-w-[200px] absolute w-[200px]">
    <Card>
      <h2 class="text-center">Node Title</h2>
      <div class="w-full" v-for="port in nodeModel.ports" :key="port.id">
        <NodeConnector :port-model="port" />
      </div>
    </Card>
  </div>
</template>

<script lang="ts">
import Card from "@/components/util/CardComponent.vue";
import NodeConnector from "./NodePort.vue";
import { defineComponent, onMounted } from "vue";
import NodeModel from "@/models/NodeModel";
import interact from "interactjs";
import Vector2 from "@/components/util/Vector";
import InteractUtil from "@/components/util/InteractUtil";
import { InteractEvent } from "@interactjs/types";

export default defineComponent({
  components: {
    Card,
    NodeConnector,
  },
  props: {
    nodeModel: {
      type: NodeModel,
      required: true,
    },
  },
  setup(props) {
    onMounted(() => {
      interact(`#${props.nodeModel.id}`).draggable({}).on("dragmove", onDragMove);
    });

    function onDragMove(event: InteractEvent) {
      props.nodeModel.addPos(new Vector2(event.dx, event.dy));
      InteractUtil.translateElem(props.nodeModel.gridPos, event);
    }
  },
});
</script>

<style></style>
