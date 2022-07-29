<template>
  <div class="absolute w-2/12 z-10 top-20 right-3 h-full">
    <Card class="p-3 mb-3">
      <div class="w-full">
        <button @click="onDeploy" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-full">
          <p class="text-center">Deploy</p>
        </button>
      </div>
    </Card>
    <Card :collapsible="true" header="Debug Console" height-class="h-1/2">
      <div class="overflow-y-auto h-[90%]">
        <div v-for="log in logs" :key="log.id" class="p-2">
          <p class="font-semibold">[{{ DateTime.getTime(log.time) }}]</p>
          <p>{{ log.message }}</p>
        </div>
      </div>
    </Card>
  </div>
</template>

<script lang="ts">
import { workflowInstancesService } from "@/api";
import Card from "@/components/CardComponent.vue";
import DateTime from "@/components/util/DateTime";
import GridData from "@/models/Data/GridData";
import { LogType } from "@/models/Data/Types";
import { ref } from "vue";
export default {
  components: {
    Card,
  },
  setup() {
    let logs = ref<LogType[]>(GridData.logs);
    const logElements = ref<HTMLDivElement[]>([]);

    async function onDeploy() {
      await workflowInstancesService
        .updateWorkflow(GridData.workflow.id, JSON.parse(JSON.stringify(GridData.workflow)))
        .then(() => {
          logs.value.push({
            id: "",
            time: new Date(),
            message: "Deployed",
          });
        });
    }

    return {
      onDeploy,
      logs,
      DateTime,
      logElements,
    };
  },
};
</script>

<style></style>
