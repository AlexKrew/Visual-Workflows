<template>
  <div
    id="EditorGrid"
    class="drag absolute bg-gray-400 top"
    :style="{ left: grid.posRel.x + 'px', top: grid.posRel.y + 'px' }"
  >
    <!-- Connections -->
    <svg id="svgID" width="10000" height="10000" xmlns="http://www.w3.org/2000/svg" class="absolute top-0 left-0">
      <NodeConnection v-for="connection in grid.connections" :key="connection.id" :connection="connection" />
    </svg>
    <!-- Nodes -->
    <div v-for="node in grid.children" :key="node.id">
      <NodeComponent :node-model="node" />
    </div>
  </div>
</template>

<script lang="ts">
import { onMounted, ref, defineComponent } from "vue";
import interact from "interactjs";
import Vector2 from "@/components/util/Vector";
import NodeComponent from "../Node/NodeComponent.vue";
import GridModel from "@/models/GridModel";
import TestModels from "@/models/Debug/TestModels";
import { InteractEvent } from "@interactjs/types";
import NodeConnection from "../Node/NodeConnection.vue";

export default defineComponent({
  components: {
    NodeComponent,
    NodeConnection,
  },
  setup() {
    const grid = ref<GridModel>(TestModels.grid);

    onMounted(() => {
      interact("#EditorGrid").draggable({}).on("dragmove", onDragMove);
    });

    function onDragMove(event: InteractEvent) {
      grid.value.addPos(new Vector2(event.dx, event.dy));
    }

    return {
      grid,
    };
  },
});
</script>

<style>
.drag {
  position: absolute;
}
#EditorGrid {
  width: 10000px;
  height: 10000px;
}
</style>