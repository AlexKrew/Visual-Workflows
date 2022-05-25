<template>
  <div class="flex flex-row fill-height">
    <EditorNodeBar class="basis-1/6" />
    <EditorGrid id="EditorGrid" class="absolute" />
    <EditorInspector class="absolute right-0 w-2/12 h-full" />
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
#EditorGrid {
  left: 16%;
  width: 10000px;
  height: 10000px;
}
</style>
