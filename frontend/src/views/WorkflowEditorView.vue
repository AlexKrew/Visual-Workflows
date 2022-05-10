<template>
  <div class="flex flex-row fill-height">
    <EditorNodeBar class="basis-2/12"/>
    <EditorGrid class="basis-8/12"/>
    <EditorInspector class="basis-2/12"/>
  </div>
</template>

<script>
  import EditorNodeBar from "@/components/workflowEditor/EditorNodeBar.vue";
  import EditorGrid from "@/components/workflowEditor/EditorGrid.vue";
  import EditorInspector from "@/components/workflowEditor/EditorInspector.vue";

  export default {
    components: {
      EditorNodeBar,
      EditorGrid,
      EditorInspector
    }
  }
</script>

<style>
  
</style>