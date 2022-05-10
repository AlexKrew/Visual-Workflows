<template>
<div ref="draggable" class="bg-gray-400">
  <Node />
  <Node />
  <Node />
</div>
</template>

<script>
import Node from "@/components/workflowEditor/EditorNode.vue"
import interact from "interactjs";

export default {
  components: {
    Node
  },
  data() {
    return {
      xPos: 0,
      yPos: 0,
    }
  },
  mounted() {
    const draggable = this.$refs.draggable;
    console.log(draggable.clientHeight)
    interact(draggable)
      .draggable({})
      .on('dragmove', this.onDragMove)
  },
  methods: {
    onDragMove(event) {
      this.xPos += event.dx
      this.yPos += event.dy
      event.target.style.transform = 'translate(' + this.xPos + 'px, ' + this.yPos + 'px)'
    },
  },
}
</script>

<style>
</style>
