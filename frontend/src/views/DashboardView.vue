<template>
  <div v-if="canvas && canvas.children.length > 0" class="relative w-full h-full p-5 bg-gray-500">
      <component :is="canvas.children[0].data.type" :obj="canvas.children[0]"></component>
  </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, onMounted, ref } from "vue";
import testDashboard from "../test/testDashboard.json";
import DElement from "@/models/Data/Dashboard/DashboardElement";
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import text_element from "@/components/workflowEditor/Dashboard/UIText.vue";
import list_element from "@/components/workflowEditor/Dashboard/UIList.vue";

export default defineComponent({
  components: {text_element, list_element},
  props: {},
  setup(props, ctx) {
    const canvas = ref<DashboardElement>();

    onBeforeMount(() => {
      canvas.value = new DElement(JSON.parse(JSON.stringify(testDashboard.canvas)));
    });

    return {
      canvas,
    };
  },
});
</script>

<style></style>
