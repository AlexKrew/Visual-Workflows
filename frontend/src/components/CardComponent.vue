<template>
  <div class="bg-white rounded-lg shadow-sm shadow-gray-500 relative" :class="isVisible ? heightClass : ''">
    <button v-if="collapsible" @click="isVisible = !isVisible" class="absolute top-3 right-3" :class="{'rotate-180': isVisible}">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 1000 1000"
        width="15"
        height="15"
      >
        <path
          d="M990,919.7c0,39.4-31.9,71.4-71.3,71.4c-19.8,0-37.7-8.1-50.7-21.1l-0.1,0.1L499.4,601.4l-360.9,361c-13,17.4-33.8,28.7-57.2,28.7c-39.4,0-71.3-32-71.3-71.4c0-19.5,7.8-37.2,20.5-50.1l-0.5-0.5l368.5-368.7l-0.1-0.1l39.7-39.7c18.4-18.4,36.5-30.3,60.6-30.3c24.1,0,45.1,13.4,62,30.3l39.7,39.7l-0.1,0.1l365.9,366C980.8,879.5,990,898.5,990,919.7z M81.3,569.6c23.4,0,44.1-11.3,57.2-28.7l360.9-361l368.5,368.7l0.1-0.1c12.9,13,30.8,21.1,50.7,21.1c39.4,0,71.3-32,71.3-71.4c0-21.2-9.2-40.2-23.9-53.2L600.3,79l0.1-0.1l-39.7-39.7c-16.9-16.9-37.9-30.3-62-30.3c-24.1,0-42.2,11.9-60.6,30.3l-39.7,39.7l0.1,0.1L30,447.7l0.5,0.5C17.8,461.1,10,478.7,10,498.3C10,537.7,41.9,569.6,81.3,569.6z"
        />
      </svg>
    </button>
    <h2 v-if="header" class="font-bold text-center p-2">{{header}}</h2>
    <slot v-if="isVisible"></slot>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  props: {
    header: {
      type: String,
      required: false,
    },
    collapsible: {
      type: Boolean,
      required: false,
    },
    heightClass: {
      type: String,
      required: false
    }
  },
  setup(props) {
    const isVisible = ref<boolean>(true);

    return {
      isVisible,
    };
  },
});
</script>

<style></style>
