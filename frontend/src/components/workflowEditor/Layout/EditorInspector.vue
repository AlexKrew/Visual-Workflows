<template>
  <Card class="bg-gray-100 absolute w-2/12 z-10 p-3 top-20 right-5">
    <div class="w-full">
      <button @click="onDeploy" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-full">
        <p class="text-center">Deploy</p>
      </button>
    </div>
  </Card>
</template>

<script lang="ts">
import { workflowInstancesService } from "@/api";
import Card from "@/components/CardComponent.vue";
import GridData from "@/models/Data/GridData";
export default {
  components: {
    Card,
  },
  setup() {
    async function onDeploy() {
      console.log(GridData.workflow);
      await workflowInstancesService.updateWorkflow(
        GridData.workflow.id,
        JSON.parse(JSON.stringify(GridData.workflow))
      );
    }

    return {
      onDeploy,
    };
  },
};
</script>

<style></style>
