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
} from "components/util/Vector"
import {
  InteractUtil
} from "components/util/InteractUtil";

export default {
  components: {
    Node
  },
  data() {
    return {
      pos: new Vector2(0, 0)
    }
  },
  mounted() {
    const draggable = this.$refs.draggable;
    interact(draggable)
      .draggable({})
      .on('dragmove', this.onDragMove)
  },
  methods: {
    onDragMove(event) {
      this.pos.add(event.dx, event.dy);
      InteractUtil.translateElem(this.pos, event);
    },
  },
}
</script>

<style>
</style>
