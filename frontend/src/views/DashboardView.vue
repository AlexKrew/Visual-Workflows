<template>
  <div>
    <div v-if="canvas && canvas.children.length > 0" class="relative w-full h-full p-5 bg-gray-300" :key="updateKey">
      <component :is="canvas.children[0].data.type" :obj="canvas.children[0]"></component>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, getCurrentInstance, onBeforeMount, onMounted, ref, watch } from "vue";
import testDashboard from "../test/testDashboard.json";
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import UIText from "@/components/workflowEditor/Dashboard/UIText.vue";
import UIList from "@/components/workflowEditor/Dashboard/UIList.vue";
import UIGauge from "@/components/workflowEditor/Dashboard/UIGauge.vue";
import DashboardModel from "@/models/Data/Dashboard/DashboardModel";
import { emitter } from "@/components/util/Emittery";
import { UITextType, UpdateFieldType } from "@/models/Data/Dashboard/UITypes";
import { dashboardInstanceService } from "@/api";

export default defineComponent({
  components: { UIText, UIList, UIGauge },
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
    });

    emitter.on("UpdateDashboard", () => {
      canvas.value = DashboardModel.canvas;
      updateKey.value++;
    });

    const connection = new WebSocket("ws://localhost:8000/dashboard/websocket");
    connection.onmessage = (event) => {
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