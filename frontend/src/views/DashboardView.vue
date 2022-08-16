<template>
  <div :key="updateKey">
    <div v-if="canvas && canvas.children.length > 0" class="relative w-full h-full p-5 bg-gray-300">
      <component :is="canvas.children[0].data.type" :obj="canvas.children[0]"></component>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, onMounted, ref } from "vue";
import testDashboard from "../test/testDashboard.json";
import DashboardElement from "@/models/Data/Dashboard/DashboardElement";
import text_element from "@/components/workflowEditor/Dashboard/UIText.vue";
import list_element from "@/components/workflowEditor/Dashboard/UIList.vue";
import DashboardModel from "@/models/Data/Dashboard/DashboardModel";
import { emitter } from "@/components/util/Emittery";
import { UIText, UpdateField } from "@/models/Data/Dashboard/UITypes";

export default defineComponent({
  components: { text_element, list_element },
  props: {},
  setup(props, ctx) {
    const canvas = ref<DashboardElement>(DashboardModel.canvas);
    const updateKey = ref(0);

    onBeforeMount(() => {
      DashboardModel.canvas = new DashboardElement(JSON.parse(JSON.stringify(testDashboard.canvas)));
      canvas.value = DashboardModel.canvas;

      const element = DashboardModel.getElementByID("C1-L1-L1-T2");
    });

    emitter.on("UpdateDashboard", () => {
      updateKey.value++;
    });

    const connection = new WebSocket("ws://localhost:8000/websocket");
    connection.onmessage = (event) => {
      let json: any = JSON.parse(event.data);

      if(json["type"] == "field_updated") DashboardModel.updateFields(...(json["data"] as UpdateField[]))
      if(json["type"] == "rebuild_ui") DashboardModel.canvas = new DashboardElement(JSON.parse(JSON.stringify(json["data"])));
    }

    return {
      canvas,
      updateKey,
    };
  },
});
</script>

<style></style>
