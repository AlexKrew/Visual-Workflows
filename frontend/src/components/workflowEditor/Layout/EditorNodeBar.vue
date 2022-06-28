<template>
  <Card class="bg-slate-200 absolute w-2/12 h-full z-10">
    <p>Node Bar</p>
    <div v-for="category in categorys" :key="category">
      <h1 class="text-xl font-bold">{{category}}</h1>
      <div v-for="node in nodesFromCategory(category)" :key="node.id" class="m-5">
        <Card :id="node.id">
          <p class="text-center">{{node.title}}</p>
        </Card>
      </div>
    </div>
  </Card>
</template>

<script lang="ts">
import Card from "@/components/CardComponent.vue";
import NodeModel from "@/models/Node/NodeModel";
import TestModels from "@/models/Debug/TestModels";
export default {
  components: {
    Card,
  },
  setup() {
    let categorys = TestModels.nodeCategorys;
    let nodes = TestModels.nodes;

    function nodesFromCategory(category: string): NodeModel[]{
      return nodes.filter((node) => node.category == category);
    }

    return {
      categorys,
      nodes,
      nodesFromCategory
    }
  }
};
</script>

<style></style>
