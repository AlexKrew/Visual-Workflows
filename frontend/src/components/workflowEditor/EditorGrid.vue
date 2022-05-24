<template>
  <div class="drag bg-gray-400">
    <!-- <EditorNode class="drag snap" />
    <EditorNode class="drag snap" /> -->
    <NodeLine :curve="testCurve"></NodeLine>
  </div>
</template>

<script lang="ts">
import { onMounted, ref, defineComponent } from 'vue'
import interact from "interactjs";
import { Vector2 } from "@/components/util/Vector"
import { InteractUtil } from "@/components/util/InteractUtil";
import EditorNode from '@/components/workflowEditor/Node/Node.vue';
import NodeLine from '@/components/workflowEditor/Node/NodeLine.vue';
import {BezierCurve} from '@/components/util/BezierCurve';

export default defineComponent({
  components: {
    EditorNode,
    NodeLine,
  },
  setup() {
    var pos: Vector2 = new Vector2(0, 0);
    var gridPos: Vector2 = new Vector2(0, 0);

    var testCurve = new BezierCurve(new Vector2(100, 100), new Vector2(130, 130), new Vector2(330, 130), new Vector2(300, 100));

    onMounted(() => {
      interact(".drag")
        .draggable({})
        .on('dragmove', onDragMove)
    });

    function onDragMove(event: any) {
      pos = new Vector2(
        (parseFloat(event.target.getAttribute('posX')) || 0) + event.dx,
        (parseFloat(event.target.getAttribute('posY')) || 0) + event.dy
      );

      if (event.target.classList.contains("snap")) {
        gridPos = InteractUtil.updateGridPos(pos, 20);
        InteractUtil.translateElem(gridPos, event);
      } else {
        InteractUtil.translateElem(pos, event);
      }

      event.target.setAttribute("posX", pos.x);
      event.target.setAttribute("posY", pos.y);
    };

    return{
      testCurve
    }
  }
})
</script>

<style>
</style>
