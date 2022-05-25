<template>
  <div class="max-w-[200px] absolute w-[200px]">
    <Card>
      <h2 class="text-center">Node Title</h2>
      <v-divider></v-divider>
      <div>
        <NodeConnector />
      </div>
    </Card>
  </div>
</template>

<script lang="ts">
import Card from "@/components/util/CardComponent.vue";
import NodeConnector from "./NodeConnector.vue";
import { defineComponent } from "vue";

export default defineComponent({
  components: {
    Card,
    NodeConnector,
  },
});
</script>

<style></style>
