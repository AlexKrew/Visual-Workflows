<template>
  <div>
    <div v-if="canvas && canvas.children.length > 0" class="relative w-full h-full p-5 bg-gray-300" :key="updateKey">
      <UIElement :obj="canvas.children[0]"></UIElement>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, getCurrentInstance, onBeforeMount, onMounted, ref, watch } from "vue";
import testDashboard from "../test/testDashboard.json";
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import DashboardModel from "@/models/Data/Dashboard/DashboardModel";
import { emitter } from "@/components/util/Emittery";
import { UITextType, UpdateFieldType } from "@/models/Data/Dashboard/UITypes";
import { dashboardInstanceService } from "@/api";
import UIElement from "@/components/workflowEditor/Dashboard/UIElement.vue";

export default defineComponent({
  components: { UIElement },
  props: {
    workflowId: {
      required: true,
      type: String,
    },
  },
  setup(props, ctx) {
    const canvas = ref<DashboardElement>(DashboardModel.canvas);
    let updateKey = ref<number>(0);

    onBeforeMount(async () => {
      const c = await dashboardInstanceService.getDashboard(props.workflowId);
      DashboardModel.setCanvas(new DashboardElement(c["canvas"]));
      emitter.emit("UpdateNavBar", [2, props.workflowId]);
    });

    emitter.on("UpdateDashboard", () => {
      canvas.value = DashboardModel.canvas;
      updateKey.value++;
    });

    const connection = new WebSocket("ws://localhost:8000/dashboard/websocket");
    connection.onmessage = (event) => {
      console.log("Dashboard Websocket", event);
      let json: any = JSON.parse(event.data);
      if (json["type"] == "field_updated") DashboardModel.updateFields(...(json["data"] as UpdateFieldType[]));
      if (json["type"] == "rebuild_ui")
        DashboardModel.canvas = new DashboardElement(JSON.parse(JSON.stringify(json["data"]["canvas"])));
    };

    return {
      canvas,
      updateKey,
    };
  },
});
</script>

<style></style>