<template>
  <div>
    <h2 class="font-bold">
    {{data.label}}:
    </h2>
    <VueGauge
      :options="{
        centralLabel: data.value.toString(),
        needleValue: getValueInPercent(data.min_value, data.max_value, data.value),
        rangeLabel: [data.min_value.toString(), data.max_value.toString()],
        arcDelimiters: [99.9]
      }"
      class="min-w"
    ></VueGauge>
  </div>
</template>

<script lang="ts">
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import { UIGauge } from "@/models/Data/Dashboard/UITypes";
import { numberLiteralTypeAnnotation } from "@babel/types";
import { defineComponent, onMounted, ref } from "vue";
import VueGauge from "vue-gauge";

export default defineComponent({
  name: "gauge-element",
  components: {
    VueGauge,
  },
  props: {
    obj: {
      required: true,
      type: DashboardElement,
    },
  },
  setup(props, ctx) {
    const data = ref<UIGauge>(props.obj.data as UIGauge);

    onMounted(() => {
      console.log("Gauge");
    });

    function getValueInPercent(min: number, max: number, value: number):number {
      return (value -min)/(max-min)*100
    }

    return {
      data,
      getValueInPercent
    };
  },
});
</script>

<style></style>
