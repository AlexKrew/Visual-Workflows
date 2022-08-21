<template>
  <div :class="fields.is_vertical ? 'flex flex-col space-y-5' : 'flex flex-row'">
    <div v-for="child in obj.children" :key="child.data.id" class="flex-auto">
      <UIElement :obj="child" />
    </div>
  </div>
</template>

<script lang="ts">
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import { defineComponent, onMounted, ref } from "vue";
import { UIListType } from "@/models/Data/Dashboard/UITypes";
import UIElement from "./UIElement.vue";

export default defineComponent({
  name: "UIList",
  components: { UIElement },
  props: {
    obj: {
      required: true,
      type: DashboardElement,
    },
  },
  setup(props) {
    const fields = ref<UIListType>(props.obj.data.fields as UIListType);

    return { fields };
  },
});
</script>

<style></style>