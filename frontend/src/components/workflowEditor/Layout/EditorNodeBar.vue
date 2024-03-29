<template>
  <Card class="absolute w-2/12 z-10 top-4 left-3" header="Node Bar" :collapsible="true">
    <div class="p-3">
      <div v-for="category in categorys" :key="category">
        <h1 class="text-xl font-bold">{{ category }}</h1>
        <div v-for="node in nodesFromCategory(category)" :key="node.id" :id="node.id" :ref="(el) => cards.push(el)">
          <Card class="m-5">
            <p class="text-center">{{ node.data.name }}</p>
          </Card>
        </div>
      </div>
      <NodeComponent v-if="curDragNode" :key="curDragNode.id" :node-model="curDragNode" />
    </div>
  </Card>
</template>

<script lang="ts">
import Card from "@/components/CardComponent.vue";
import NodeModel from "@/models/Node/NodeModel";
import GridData from "@/models/Data/GridData";
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
    let categorys = GridData.nodeCategorys;
    let nodes = GridData.nodes;
    let grid = GridData.grid;
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
      });
    });

    function nodesFromCategory(category: string): NodeModel[] {
      return nodes.filter((node) => node.data.category == category);
    }

    function addNode(node: NodeModel) {
      curDragNode.value = node.clone() as NodeModel;
      curDragNode.value.setParent(GridData.grid);

      const div = cards.value.find((div) => div.id == node.id);
      if (div) {
        const rect = div.getBoundingClientRect();
        curDragNode.value.setPos(new Vector2(rect.x, rect.y -64));  //64 because of navbar TODO

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
        grid.addChild(curDragNode.value, true);
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
