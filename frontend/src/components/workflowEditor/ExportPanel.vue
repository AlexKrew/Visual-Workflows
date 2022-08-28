<template>
  <div class="modal-mask">
    <div class="modal-wrapper">
      <div class="modal-container">
        <textarea v-model="textAreaValue" class="border w-full h-5/6" rows="20"></textarea>
        <div class="flex justify-center m-5 space-x-5">
          <button
            @click="onImport()"
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded flex-1"
          >
            <p class="text-center">Import</p>
          </button>
          <button
            @click="onCancel()"
            class="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded flex-1"
          >
            <p class="text-center">Cancel</p>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import GridData from "@/models/Data/GridData";
import { defineComponent, onMounted, ref } from "vue";
import { emitter } from "../util/Emittery";

export default defineComponent({
  components: {},
  props: {},
  setup(props, ctx) {
    const textAreaValue = ref<string>("");

    onMounted(() => {
      textAreaValue.value = JSON.stringify(GridData.workflow, null, 4);
    });

    function onImport() {
      GridData.loadWorkflow(JSON.parse(textAreaValue.value));
      emitter.emit("UpdateWorkflowEditor", undefined);
    }

    function onCancel() {
      emitter.emit("OpenImportExportModal", false);
    }

    return {
      textAreaValue,
      onImport,
      onCancel,
    };
  },
});
</script>

<style>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 50%;
  height: 50%;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
}
</style>
