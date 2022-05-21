<template>
<div ref="draggable" class="bg-gray-400">
  <Node />
  <Node />
  <Node />
</div>
</template>

<script lang="ts">
import interact from "interactjs";
import Node from "@/components/workflowEditor/EditorNode.vue"
import {
  Vector2
} from "@/components/util/Vector"
import {
  InteractUtil
} from "@/components/util/InteractUtil";
import { ref } from 'vue'

var pos = new Vector2(0, 0);

function init() {
  const draggable = ref("draggable").value;
  interact(draggable)
    .draggable({})
    .on('dragmove', onDragMove)
}

function onDragMove(event: any) {
  pos.add(event.dx, event.dy);
  InteractUtil.translateElem(pos, event);
};

export default {
  components: {
    Node
  },
  mounted() {
    init();
  }
}
</script>

<style>
</style>
