<template>
  <li class="hover:bg-gray-50">
    <div class="flex items-center px-4 py-4 sm:px-6">
      <div class="min-w-0 flex-1 flex items-center">
        <div class="min-w-0 flex-1 px-4 md:grid md:grid-cols-3 md:gap-4">
          <div>
            <p class="text-sm font-medium text-blue-600 truncate">{{ workflow.name }}</p>
            <p class="mt-2 flex items-center text-sm text-gray-500">
              <!-- TODO: Add matching leading icon based on the status -->
              {{ workflow.status }}
            </p>
          </div>

          <div class="flex flex-row items-center justify-center space-x-5"></div>

          <div class="flex flex-row items-center justify-end">
            <button
              type="button"
              class="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-blue-700 bg-blue-100 hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              @click="$emit('clickEditor')"
            >
              Editor
            </button>
          </div>
        </div>
      </div>
    </div>
  </li>
</template>

<script lang="ts">
import WorkflowControls from "./WorkflowControls.vue";
export default {
  emits: ["clickEditor"],

  components: {},

  props: {
    workflow: {
      type: Object,
      required: true,
    },
  },

  setup(props: any) {
    return {};
  },
};
</script>