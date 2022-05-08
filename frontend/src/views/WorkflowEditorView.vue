<template>
  <EditorNodeBar/>
  <EditorGrid/>
  <EditorInspector/>
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