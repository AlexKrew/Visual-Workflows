<template>
  <div id="Editor" class="w-full h-full absolute">
    <div
      id="EditorGrid"
      class="drag bg-gray-500"
      :style="{ left: grid.posRel.x + 'px', top: grid.posRel.y + 'px' }"
      @scroll.prevent
    >
      <!-- Connections -->
      <svg id="svgID" width="10000" height="10000" xmlns="http://www.w3.org/2000/svg">
        <NodeEdge v-for="edge in grid.edges" :key="edge.data.id" :edge="edge" />
      </svg>
      <!-- Nodes -->
      <NodeComponent v-for="node in nodes" :key="node.id" :node-model="node" />
    </div>
  </div>
</template>

<script lang="ts">
import { onMounted, ref, defineComponent, onUnmounted } from "vue";
import interact from "interactjs";
import Vector2 from "@/components/util/Vector";
import NodeComponent from "../Node/NodeComponent.vue";
import GridModel from "@/models/GridModel";
import GridData from "@/models/Data/GridData";
import { InteractEvent } from "@interactjs/types";
import NodeEdge from "../Node/NodeEdge.vue";
import NodeModel from "@/models/Node/NodeModel";

export default defineComponent({
  components: {
    NodeComponent,
    NodeEdge: NodeEdge,
  },
  setup() {
    const grid = ref<GridModel>(GridData.grid);
    const nodes = ref<NodeModel[]>(grid.value.children as NodeModel[]);

    onMounted(() => {
      interact("#EditorGrid")
        .draggable({
          modifiers: [
            interact.modifiers.restrict({
              restriction: "#Editor",
              elementRect: { top: 0, left: 0, bottom: 1, right: 1 },
            }),
          ],
        })
        .on("dragmove", onDragMove);
    });

    onUnmounted(() => {
      interact("#EditorGrid").unset();
    });

    function onDragMove(event: InteractEvent) {
      grid.value.addPos(new Vector2(event.dx, event.dy));
    }

    return {
      grid,
      nodes,
    };
  },
});
</script>

<style>
.drag {
  position: absolute;
  width: 10000px;
  height: 10000px;
}
</style>
