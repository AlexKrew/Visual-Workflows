<template>
  <div class="drag bg-gray-400">
    <!-- Nodes -->
    <div v-for="node in grid.nodes" :key="node.id">
      <EditorNode class="drag snap" :node-model="node" />
    </div>
    <!-- Lines -->
    <!-- <NodeLine :curve="testCurve"></NodeLine> -->
  </div>
</template>

<script lang="ts">
import { onMounted, ref, defineComponent } from "vue";
import interact from "interactjs";
import { Vector2 } from "@/components/util/Vector";
import InteractUtil from "@/components/util/InteractUtil";
import EditorNode from "@/components/workflowEditor/Node/NodeComponent.vue";
import NodeLine from "@/components/workflowEditor/Node/NodeLine.vue";
import { BezierCurve } from "@/components/util/BezierCurve";
import GridModel from "@/models/GridModel";
import TestModels from "@/models/TestModels";

export default defineComponent({
  components: {
    EditorNode,
  },
  setup() {
    const grid = ref<GridModel>(TestModels.getGrid());

    onMounted(() => {
      interact(".drag").draggable({}).on("dragmove", onDragMove);
    });

    function onDragMove(event: any) {
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
