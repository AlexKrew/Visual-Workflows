<template>
<div ref="draggable">
  <Card class="max-w-[200px]">
    <h2 class="text-center">Node Title</h2>
    <v-divider></v-divider>
    <div>
      <div id="circle"></div>
    </div>
  </Card>
</div>
</template>

<script>
import interact from "interactjs";
import Card from "@/components/util/Card.vue";

export default {
  components: {
    Card
  },
  data() {
    return {
      xPos: 0,
      yPos: 0,
    }
  },
  mounted() {
    const draggable = this.$refs.draggable;
    interact(draggable)
      .draggable({})
      .on('dragmove', this.onDragMove)
  },
  unmounted() {},
  methods: {
    onDragMove(event) {
      this.xPos += event.dx
      this.yPos += event.dy
      event.target.style.transform = 'translate(' + this.xPos + 'px, ' + this.yPos + 'px)'
    },
  },
  computed: {}
};
</script>

<style>
#circle {
  width: 20px;
  height: 20px;
  -webkit-border-radius: 25px;
  -moz-border-radius: 25px;
  border-radius: 25px;
  background: grey;
}
</style>
