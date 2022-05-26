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
import { defineComponent } from "vue";
import NodeModel from "@/models/NodeModel";

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
});
</script>

<style></style>
