<template>
  <component :is="obj.data.type" :obj="obj" :key="obj.data.id"></component>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import UIList from "./UIList.vue";
import UIText from "./UIText.vue";
import UIGauge from "./UIGauge.vue";

export default defineComponent({
  name: "UIElement",
  components: { UIList, UIText, UIGauge },
  props: {
    obj: {
      required: true,
      type: DashboardElement,
    },
  },
  setup(props, ctx) {
    //
  },
});
</script>

<style></style>