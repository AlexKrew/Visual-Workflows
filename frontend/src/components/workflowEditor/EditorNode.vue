<template>
<div ref="draggable" id="draggable">
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
import {InteractUtil} from "@/components/util/InteractUtil.js";
import {Vector2} from "@/components/util/Vector.js";

export default {
  components: {
    Card
  },
  data() {
    return {
      pos: new Vector2(0, 0),
      gridPos: new Vector2(0, 0),
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
      this.pos.add(event.dx, event.dy)
      this.gridPos = InteractUtil.updateGridPos(this.pos, 50);
      InteractUtil.translateElem(this.gridPos, event);
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
#draggable{
  position: absolute;
  width: 200px;
}
</style>
