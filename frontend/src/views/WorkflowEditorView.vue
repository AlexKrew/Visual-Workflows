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
import { onMounted, ref } from "vue";
import workflowJSON from "../test/testWorkflow.json";
import TestModels from "@/models/Debug/TestModels";
import GridModel from "@/models/GridModel";

export default {
  components: {
    EditorNodeBar,
    EditorGrid,
    EditorInspector,
  },
  setup() {
    let updateKey = ref(0);

    onMounted(() => {
      let json: JSON = JSON.parse(JSON.stringify(workflowJSON));
      TestModels.grid = GridModel.fromJSON(json);
      emitter.emit("UpdateWorkflowEditor");
      console.log(TestModels.grid);
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
