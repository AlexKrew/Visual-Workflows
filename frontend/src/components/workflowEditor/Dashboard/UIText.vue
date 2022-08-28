<template>
  <div class="">
    <span v-if="fields.label" class="p-2" :style="{ fontSize: fields.font_size + 'px' }">
      {{ fields.label }}
    </span>
    <span v-if="fields.value" class="p-2 border border-black rounded-lg">
      {{ fields.value }}
    </span>
  </div>
</template>

<script lang="ts">
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import { UITextType } from "@/models/Data/Dashboard/UITypes";
import { defineComponent, onMounted, ref } from "vue";

export default defineComponent({
  name: "UIText",
  props: {
    obj: {
      required: true,
      type: DashboardElement,
    },
  },
  setup(props, ctx) {
    const fields = ref<UITextType>(props.obj.data.fields as UITextType);
    if (!fields.value.font_size || fields.value.font_size == 0) fields.value.font_size = 16;

    return {
      fields,
    };
  },
});
</script>

<style></style>