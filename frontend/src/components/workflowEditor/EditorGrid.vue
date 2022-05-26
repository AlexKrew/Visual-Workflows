<template>
  <div class="drag grid bg-gray-400">
    <!-- Nodes -->
    <div v-for="node in grid.nodes" :key="node.id">
      <EditorNode class="drag node snap" :node-model="node" :id="node.id" />
    </div>
    <!-- Lines -->
    <NodeLine :curve="testCurve"></NodeLine>
  </div>
</template>

<script lang="ts">
import { onMounted, ref, defineComponent } from "vue";
import interact from "interactjs";
import { Vector2 } from "@/components/util/Vector";
import { InteractUtil } from "@/components/util/InteractUtil";
import EditorNode from "@/components/workflowEditor/Node/NodeComponent.vue";
import NodeLine from "@/components/workflowEditor/Node/NodeLine.vue";
import { BezierCurve } from "@/components/util/BezierCurve";
import GridModel from "@/models/GridModel";
import TestModels from "@/models/TestModels";

export default defineComponent({
  components: {
    EditorNode,
    NodeLine,
  },
  setup() {
    const grid = ref<GridModel>(TestModels.getGrid());
    var pos: Vector2 = new Vector2(0, 0);
    var gridPos: Vector2 = new Vector2(0, 0);
    var testCurve = ref(
      new BezierCurve(
        new Vector2(100, 100),
        new Vector2(130, 130),
        new Vector2(330, 130),
        new Vector2(300, 100)
      )
    );

    onMounted(() => {
      interact(".drag").draggable({}).on("dragmove", onDragMove);
      console.log(grid.value.pos);
      var t = grid.value;
      t.pos = new Vector2(100, 100);
      console.log(grid.value.pos);
    });

    function onDragMove(event: any) {
      pos = new Vector2(
        (parseFloat(event.target.getAttribute("posX")) || 0) + event.dx,
        (parseFloat(event.target.getAttribute("posY")) || 0) + event.dy
      );

      if (event.target.classList.contains("snap")) {
        gridPos = InteractUtil.updateGridPos(pos, 20);
        InteractUtil.translateElem(gridPos, event);
      } else {
        InteractUtil.translateElem(pos, event);
      }

      event.target.setAttribute("posX", pos.x);
      event.target.setAttribute("posY", pos.y);

      testCurve.value = new BezierCurve(
        new Vector2(200, 200),
        new Vector2(130, 130),
        new Vector2(330, 130),
        gridPos
      );
    }

    return {
      testCurve,
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
