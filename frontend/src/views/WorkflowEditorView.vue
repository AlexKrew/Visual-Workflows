<template>
  <div class="flex flex-row fill-height" :key="updateKey">
    <EditorNodeBar />
    <EditorGrid />
    <EditorInspector />
  </div>
</template>

<script lang="ts">
import EditorNodeBar from "@/components/workflowEditor/Layout/EditorNodeBar.vue";
import EditorGrid from "@/components/workflowEditor/Layout/EditorGrid.vue";
import EditorInspector from "@/components/workflowEditor/Layout/EditorInspector.vue";
import { emitter } from "@/components/util/Emittery";
import { onBeforeMount, onMounted, ref } from "vue";
import EmptyWorkflowJSON from "../models/Data/JSON/EmptyWorkflow.json"
import GridData from "@/models/Data/GridData";

export default {
  components: {
    EditorNodeBar,
    EditorGrid,
    EditorInspector,
  },
  setup() {
    let updateKey = ref(0);

    onBeforeMount(() => {
      // Load all default Nodes from Nodes.json
      GridData.loadDefaultNodes();

      // Load Empty Initial Workflow
      GridData.loadWorkflow(JSON.parse(JSON.stringify(EmptyWorkflowJSON)));

      console.log(GridData.grid)

    });

    onMounted(() => {
      // Loads current Workflow
      // let json: JSON = JSON.parse(JSON.stringify(workflowJSON));
      // TestModels.grid = GridModel.fromJSON(json);
      // emitter.emit("UpdateWorkflowEditor");
    }),
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
