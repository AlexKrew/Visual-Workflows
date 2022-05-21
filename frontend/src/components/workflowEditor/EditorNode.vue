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

<script lang="ts">
import interact from "interactjs";
import Card from "@/components/util/Card.vue";
import {
  InteractUtil
} from "@/components/util/InteractUtil";
import {
  Vector2
} from "@/components/util/Vector";
import {
  ref
} from 'vue'

var pos = new Vector2(0, 0);
var gridPos = new Vector2(0, 0);

function init() {
  const draggable = ref("draggable").value;
  interact(draggable)
    .draggable({})
    .on('dragmove', onDragMove)
}

function onDragMove(event: any) {
  pos.add(event.dx, event.dy)
  gridPos = InteractUtil.updateGridPos(pos, 50);
  InteractUtil.translateElem(gridPos, event);
}

export default {
  components: {
    Card
  },
  mounted() {
    init();
  },
  methods: {},
  computed: {}
}
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

#draggable {
  position: absolute;
  width: 200px;
}
</style>
