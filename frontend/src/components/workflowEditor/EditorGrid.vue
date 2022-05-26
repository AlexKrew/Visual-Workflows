<template>
  <div class="drag bg-gray-400">
    <!-- Nodes -->
    <div v-for="node in grid.nodes" :key="node.id">
      <EditorNode :node-model="node" />
    </div>
    <!-- Lines -->
    <!-- <NodeLine :curve="testCurve"></NodeLine> -->
  </div>
</template>

<script lang="ts">
import { onMounted, ref, defineComponent } from "vue";
import interact from "interactjs";
import Vector2 from "@/components/util/Vector";
import InteractUtil from "@/components/util/InteractUtil";
import EditorNode from "@/components/workflowEditor/Node/NodeComponent.vue";
import GridModel from "@/models/GridModel";
import TestModels from "@/models/TestModels";
import { InteractEvent } from "@interactjs/types";

export default defineComponent({
  components: {
    EditorNode,
  },
  setup() {
    const grid = ref<GridModel>(TestModels.getGrid());

    onMounted(() => {
      interact(".drag").draggable({}).on("dragmove", onDragMove);
    });

    function onDragMove(event: InteractEvent) {
      grid.value.addPos(new Vector2(event.dx, event.dy));
      InteractUtil.translateElem(grid.value.pos, event);
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
</style>
