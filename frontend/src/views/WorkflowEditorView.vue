<template>
  <div class="flex flex-row fill-height">
    <EditorNodeBar />
    <EditorGrid/>
    <EditorInspector />
  </div>
</template>

<script lang="ts">
import EditorNodeBar from "@/components/workflowEditor/Layout/EditorNodeBar.vue";
import EditorGrid from "@/components/workflowEditor/Layout/EditorGrid.vue";
import EditorInspector from "@/components/workflowEditor/Layout/EditorInspector.vue";

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
