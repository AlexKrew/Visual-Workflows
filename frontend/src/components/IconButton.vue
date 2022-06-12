<template>
  <button
    type="button"
    class="inline-flex items-center p-1.5 border border-transparent rounded shadow-sm text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
    :class="{'bg-gray-500 hover:bg-gray-600': disabled}"
    :disabled="disabled"
  >
    <slot />
  </button>
</template>

<script lang="ts">
export default {
  props: {
    disabled: {
      type: Boolean,
      required: true,
    },
  }
}
</script>
