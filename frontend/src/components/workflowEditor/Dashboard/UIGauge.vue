<template>
  <div>
    <h2 v-if="fields.label" class="font-bold">{{ fields.label }}:</h2>
    <VueGauge
      :refid="id"
      :options="{
        centralLabel: fields.value.toString(),
        needleValue: getValueInPercent(fields.min_value, fields.max_value, fields.value),
        rangeLabel: [fields.min_value.toString(), fields.max_value.toString()],
        arcDelimiters: [99.9],
      }"
    />
  </div>
</template>

<script lang="ts">
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import { UIGaugeType } from "@/models/Data/Dashboard/UITypes";
import { defineComponent, onBeforeMount, onMounted, ref } from "vue";
import VueGauge from "vue-gauge";
import { uuid } from "vue-uuid";

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
    const id = ref<string>("Gauge-" + uuid.v4());

    onBeforeMount(() => {
      if (!fields.value.min_value) fields.value.min_value = 0;
      if (!fields.value.max_value) fields.value.max_value = 100;
      if (!fields.value.value) fields.value.value = fields.value.min_value;
    });

    function getValueInPercent(min: number, max: number, value: number): number {
      return ((value - min) / (max - min)) * 100;
    }

    return {
      fields,
      getValueInPercent,
      id,
    };
  },
});
</script>

<style></style>