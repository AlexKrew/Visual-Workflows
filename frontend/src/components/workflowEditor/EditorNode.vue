<template>
  <div>
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
import Card from "@/components/util/Card.vue";

export default {
  components: {
    Card
  },
  setup() {

  },
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

.draggable {
  position: absolute;
  width: 200px;
}
</style>
