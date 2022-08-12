<template>
  <div class="">
    <span class="p-2">
      {{ data.label }}
    </span>
    <span class="p-2 border border-black rounded-lg">
      {{ data.value }}
    </span>
  </div>
</template>

<script lang="ts">
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import { UIText } from "@/models/Data/Dashboard/UITypes";
import { defineComponent, onMounted, ref } from "vue";

export default defineComponent({
  props: {
    obj: {
      required: true,
      type: DashboardElement,
    },
  },
  setup(props, ctx) {
    const data = ref<UIText>(props.obj.data as UIText);

    return {
      data,
    };
  },
});
</script>

<style></style>
