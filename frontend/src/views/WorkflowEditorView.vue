<template>
  <v-row class="fill-height">
    <EditorNodeBar class="w-3/12"/>
    <EditorGrid class="w-1/2"/>
    <EditorInspector class="w-3/12"/>
  </v-row>
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