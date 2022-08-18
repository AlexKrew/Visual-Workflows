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
import { UIGaugeType } from "@/models/Data/Dashboard/UITypes"
import { defineComponent, onMounted, ref } from "vue";
import VueGauge from "vue-gauge";

export default defineComponent({
  name: "UIGauge",
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
    const fields = ref<UIGaugeType>(props.obj.data.fields as UIGaugeType);

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
