<template>
  <div class="flex flex-row fill-height">
    <EditorNodeBar />
    <EditorGrid/>
    <EditorInspector />
  </div>
</template>

<script lang="ts">
import EditorNodeBar from "@/components/workflowEditor/EditorNodeBar.vue";
import EditorGrid from "@/components/workflowEditor/EditorGrid.vue";
import EditorInspector from "@/components/workflowEditor/EditorInspector.vue";

export default {
  components: {
    EditorNodeBar,
    EditorGrid,
    EditorInspector,
  },
};
</script>

<style>
</style>
