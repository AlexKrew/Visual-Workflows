<template>
    <!-- <svg id="svgID" width="1000" height="1000" xmlns="http://www.w3.org/2000/svg"> -->
    <path :id="connection.id" :d="connection.getCurve().toSVGString()" stroke="black" fill="transparent" class="absolute top-0 left-0"/>
    <!-- </svg> -->
</template>

<script lang="ts">
import EdgeModel from "@/models/Node/EdgeModel";
import { defineComponent, onMounted } from "vue";

export default defineComponent({
  props: {
    connection: {
      type: EdgeModel,
      required: true,
    },
  },
  setup() {
    onMounted(() => {
      //
    });
  },
});
</script>

<style>
.svgID {
  position: absolute;
}
</style>
