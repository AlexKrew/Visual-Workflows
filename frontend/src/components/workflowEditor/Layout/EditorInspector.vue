<template>
  <div class="absolute w-2/12 z-10 top-20 right-3 h-full">
    <Card class="p-3 mb-3">
      <div class="w-full flex">
        <button
          @click="onSave"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded flex-1 mr-3"
        >
          <p class="text-center">Save</p>
        </button>
        <button @click="onStart" class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded flex-1">
          <p class="text-center">Start</p>
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
import { onBeforeUnmount, onMounted, ref } from "vue";
export default {
  components: {
    Card,
  },
  setup() {
    let logs = ref<LogType[]>(GridData.logs);
    const logElements = ref<HTMLDivElement[]>([]);

    const connection = new WebSocket("ws://localhost:8000/workflow/websocket");
    connection.onmessage = (event) => {
      let json: any = JSON.parse(event.data);
      logs.value.push({
        id: json["id"],
        time: new Date(json["timestamp"]),
        message: json["message"],
      });
      console.log(json);
    };

    onBeforeUnmount(() => {
      connection.close();
    });

    async function onSave() {
      await workflowInstancesService.updateWorkflow(
        GridData.workflow.id,
        JSON.parse(JSON.stringify(GridData.workflow))
      );
    }

    async function onStart() {
      // await onDeploy();
      await workflowInstancesService.startWorkflow(GridData.workflow.id);
    }

    return {
      onSave,
      onStart,
      logs,
      DateTime,
      logElements,
    };
  },
};
</script>

<style></style>
