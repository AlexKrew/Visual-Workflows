<template>
  <div
    :id="nodeModel.id"
    class="max-w-[200px] absolute w-[200px]"
    :style="{ left: nodeModel.posGridCell.x + 'px', top: nodeModel.posGridCell.y + 'px' }"
  >
    <Card>
      <h2 class="text-center">{{ nodeModel.title }}</h2>
      <div class="w-full">
        <NodePort v-for="port in ports" :key="port.id" :port-model="port" />
      </div>
      <div v-if="nodeModel.addablePorts.length > 0" class="w-full flex justify-center">
        <button
          class="w-full p-1 m-3 bg-blue-500 hover:bg-blue-700 text-white font-bold rounded"
          @click="onAddAddablePorts()"
        >
          ADD
        </button>
      </div>
    </Card>
  </div>
</template>

<script lang="ts">
import Card from "@/components/CardComponent.vue";
import NodePort from "./NodePort.vue";
import { defineComponent, onMounted, onUnmounted, ref } from "vue";
import NodeModel from "@/models/Node/NodeModel";
import interact from "interactjs";
import Vector2 from "@/components/util/Vector";
import { InteractEvent } from "@interactjs/types";
import NodePortModel from "@/models/Node/NodePortModel";

export default defineComponent({
  components: {
    Card,
    NodePort,
  },
  props: {
    nodeModel: {
      type: NodeModel,
      required: true,
    },
  },
  setup(props) {
    const ports = ref<NodePortModel[]>(props.nodeModel.children as NodePortModel[]);

    onMounted(() => {
      interact(`#${props.nodeModel.id}`).draggable({}).on("dragmove", onDragMove);
    });

    onUnmounted(() => {
      interact(`#${props.nodeModel.id}`).unset();
    });

    function onDragMove(event: InteractEvent) {
      props.nodeModel.addPos(new Vector2(event.dx, event.dy));
    }

    function onAddAddablePorts() {
      props.nodeModel.addAddablePorts();
    }

    return {
      onAddAddablePorts,
      ports,
    };
  },
});
</script>

<style></style>
