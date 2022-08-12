<template>
  <div :class="data.is_vertical ? 'space-y-5' : 'flex'">
    <div v-for="element in obj.children" :key="element.data.id" class="flex-auto">
      <component :is="element.data.type" :obj="element"></component>
    </div>
  </div>
</template>

<script lang="ts">
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import { defineComponent, onMounted, ref } from "vue";
import text_element from "@/components/workflowEditor/Dashboard/UIText.vue";
import list_element from "@/components/workflowEditor/Dashboard/UIList.vue";
import { UIList } from "@/models/Data/Dashboard/UITypes";

export default defineComponent({
  components: {
    text_element,
    // list_element,
  },
  props: {
    obj: {
      required: true,
      type: DashboardElement,
    },
  },
  setup(props) {
    const data = ref<UIList>(props.obj.data as UIList);

    return {
      data,
    };
  },
});
</script>

<style></style>
