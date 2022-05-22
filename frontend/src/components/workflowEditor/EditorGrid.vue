<template>
  <div class="drag bg-gray-400">
    <EditorNode class="drag" />
    <EditorNode class="drag" />
  </div>
</template>

<script lang="ts">
import { onMounted, ref } from 'vue'
import interact from "interactjs";
import { Vector2 } from "@/components/util/Vector"
import { InteractUtil } from "@/components/util/InteractUtil";
import EditorNode from '@/components/workflowEditor/EditorNode.vue';
import { Interactable } from '@interactjs/types';

export default {
  components: {
    Node,
    EditorNode
  },
  setup() {
    var pos: Vector2 = new Vector2(0, 0);
    var gridPos: Vector2 = new Vector2(0, 0);

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
      gridPos = InteractUtil.updateGridPos(pos, 20);
      InteractUtil.translateElem(gridPos, event);

      event.target.setAttribute("posX", pos.x);
      event.target.setAttribute("posY", pos.y);
    };
  }
}
</script>

<style>
</style>
