<template>
  <Card class="bg-slate-200 absolute w-2/12 h-full z-10">
    <p>Node Bar</p>
    <div v-for="category in categorys" :key="category">
      <h1 class="text-xl font-bold">{{ category }}</h1>
      <div v-for="node in nodesFromCategory(category)" :key="node.id" :id="node.id" :ref="(el) => cards.push(el)">
        <Card class="m-5">
          <p class="text-center">{{ node.title }}</p>
        </Card>
      </div>
    </div>
    <NodeComponent v-if="curDragNode" :key="curDragNode.id" :node-model="curDragNode" />
  </Card>
</template>

<script lang="ts">
import Card from "@/components/CardComponent.vue";
import NodeModel from "@/models/Node/NodeModel";
import TestModels from "@/models/Debug/TestModels";
import { onMounted, onUnmounted, ref } from "vue";
import interact from "interactjs";
import { InteractEvent } from "@interactjs/types";
import NodeComponent from "../Node/NodeComponent.vue";
import Vector2 from "@/components/util/Vector";
import { emitter } from "@/components/util/Emittery";
export default {
  components: {
    Card,
    NodeComponent,
  },
  setup() {
    let categorys = TestModels.nodeCategorys;
    let nodes = TestModels.nodes;
    let grid = TestModels.grid;
    const curDragNode = ref<NodeModel | undefined>();
    const cards = ref<HTMLDivElement[]>([]);

    onMounted(() => {
      nodes.forEach((node) => {
        interact(`#${node.id}`)
          .draggable({})
          .on("dragstart", onDragStart)
          .on("dragmove", onDragMove)
          .on("dragend", onDragEnd);
        // .on("drop", onDrop);
      });
    });

    onUnmounted(() => {
      nodes.forEach((node) => {
        interact(`#${node.id}`).unset();
      })
    });

    function nodesFromCategory(category: string): NodeModel[] {
      return nodes.filter((node) => node.category == category);
    }

    function addNode(node: NodeModel) {
      curDragNode.value = node.clone() as NodeModel;
      curDragNode.value.setParent(grid);

      const div = cards.value.find((div) => div.id == node.id);
      if (div) {
        const rect = div.getBoundingClientRect();
        curDragNode.value.setPos(new Vector2(rect.x, rect.y));

        cards.value.splice(0, cards.value.length);
      }
    }

    //#region InteractJS
    function onDragStart(event: InteractEvent) {
      const selectedNode = nodes.find((node) => {
        return node.id == event.target.id;
      });
      if (selectedNode) {
        addNode(selectedNode);
      } else {
        throw new Error("No Node Found with ID: " + event.target.id);
      }
    }

    function onDragMove(event: InteractEvent) {
      if (curDragNode.value) {
        curDragNode.value.addPos(new Vector2(event.dx, event.dy));
      }
    }

    function onDragEnd() {
      if (curDragNode.value) {
        curDragNode.value.addPos(grid.posRel.negateReturn());
        grid.addChildren(curDragNode.value);
        curDragNode.value = undefined;
        emitter.emit("UpdateWorkflowEditor");
      }
    }
    //#endregion

    return {
      categorys,
      nodes,
      nodesFromCategory,
      curDragNode,
      cards,
    };
  },
};
</script>

<style></style>
