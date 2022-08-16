<template>
  <div>
    <h2 class="font-bold">
    {{fields.label}}:
    </h2>
    <VueGauge
      :options="{
        centralLabel: fields.value.toString(),
        needleValue: getValueInPercent(fields.min_value, fields.max_value, fields.value),
        rangeLabel: [fields.min_value.toString(), fields.max_value.toString()],
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
    const fields = ref<UIGauge>(props.obj.data.fields as UIGauge);

    function getValueInPercent(min: number, max: number, value: number):number {
      return (value -min)/(max-min)*100
    }

    return {
      fields,
      getValueInPercent
    };
  },
});
</script>

<style></style>
