<template>
  <div class="flex flex-row fill-height" :key="updateKey">
    <EditorNodeBar />
    <EditorGrid/>
    <EditorInspector />
  </div>
</template>

<script lang="ts">
import EditorNodeBar from "@/components/workflowEditor/Layout/EditorNodeBar.vue";
import EditorGrid from "@/components/workflowEditor/Layout/EditorGrid.vue";
import EditorInspector from "@/components/workflowEditor/Layout/EditorInspector.vue";
import { emitter } from "@/components/util/Emittery";
import { ref } from "vue";

export default {
  components: {
    EditorNodeBar,
    EditorGrid,
    EditorInspector,
  },
  setup() {
    let updateKey = ref(0);

    emitter.on("UpdateWorkflowEditor", () => {
      updateKey.value++;
    });

    return {
      updateKey,
    };
  },
};
</script>

<style></style>
