<template>
  <div :class="fields.is_vertical ? 'flex flex-col space-y-5' : 'flex flex-row'">
    <div v-for="child in obj.children" :key="child.data.id" class="flex-auto">
      <component :is="child.data.type" :obj="child"></component>
    </div>
  </div>
</template>

<script lang="ts">
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import { defineComponent, onMounted, ref } from "vue";
import UIText from "@/components/workflowEditor/Dashboard/UIText.vue";
import UIGauge from "@/components/workflowEditor/Dashboard/UIGauge.vue";
import { UIListType } from "@/models/Data/Dashboard/UITypes";

export default defineComponent({
  name: "UIList",
  components: {
    UIText,
    UIGauge,
  },
  props: {
    obj: {
      required: true,
      type: DashboardElement,
    },
  },
  setup(props) {
    const fields = ref<UIListType>(props.obj.data.fields as UIListType);

    return {
      fields,
    };
  },
});
</script>

<style></style>
