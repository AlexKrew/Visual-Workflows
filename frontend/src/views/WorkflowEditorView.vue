<template>
  <div class="flex flex-row fill-height HighlightDisabled" :key="updateKey">
    <div v-if="isLoading">
      <p>Loading</p>
    </div>

    <template v-else>
      <EditorNodeBar />
      <EditorGrid />
      <EditorInspector />
      <ExportPanel v-if="openModal" />
    </template>
  </div>
</template>

<script lang="ts">
import EditorNodeBar from "@/components/workflowEditor/Layout/EditorNodeBar.vue";
import EditorGrid from "@/components/workflowEditor/Layout/EditorGrid.vue";
import EditorInspector from "@/components/workflowEditor/Layout/EditorInspector.vue";
import { emitter } from "@/components/util/Emittery";
import { defineComponent, onBeforeMount, onMounted, ref } from "vue";
import EmptyWorkflowJSON from "../models/Data/JSON/EmptyWorkflow.json";
import testJSON from "../test/test1.json";
import GridData from "@/models/Data/GridData";
import { workflowInstancesService } from "@/api";
import ExportPanel from "@/components/workflowEditor/ExportPanel.vue";

export default defineComponent({
  components: {
    EditorNodeBar,
    EditorGrid,
    EditorInspector,
    ExportPanel,
  },
  props: {
    workflowId: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    let updateKey = ref(0);
    const isLoading = ref(true);
    const openModal = ref(false);

    onBeforeMount(async () => {
      // Load all default Nodes from Nodes.json
      GridData.loadDefaultData();

      // Load Empty Initial Workflow
      const workflowJSON = await workflowInstancesService.loadWorkflow(props.workflowId);
      // GridData.loadWorkflow(JSON.parse(JSON.stringify(testJSON)));
      GridData.loadWorkflow(JSON.parse(JSON.stringify(workflowJSON)));
      isLoading.value = false;
      emitter.emit("UpdateNavBar", [1, props.workflowId]);
    });

    onMounted(() => {
      // Loads current Workflow
      // let json: JSON = JSON.parse(JSON.stringify(workflowJSON));
      // TestModels.grid = GridModel.fromJSON(json);
      // emitter.emit("UpdateWorkflowEditor");
    });

    emitter.on("UpdateWorkflowEditor", () => {
      updateKey.value++;
    });

    emitter.on("OpenImportExportModal", (open) => {
      openModal.value = open;
    });

    return {
      updateKey,
      isLoading,
      openModal,
    };
  },
});
</script>

<style></style>