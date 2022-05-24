<template>
  <div>
    <svg id="svgID" width="1000" height="1000" xmlns="http://www.w3.org/2000/svg">
      <path :d="curve.toSVGString()" stroke="black" fill="transparent" />s
    </svg>
  </div>
</template>

<script lang="ts">
import { BezierCurve } from "@/components/util/BezierCurve";
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    curve: {
      type: BezierCurve,
      required: false,
      default: new BezierCurve(),
    },
  },
  setup(props) {
    console.log(props.curve.toSVGString());
  },
});
</script>

<style>
.svgID{
  position: absolute,
}
</style>
