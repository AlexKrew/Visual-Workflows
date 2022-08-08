<template>
  <div
    :id="nodeModel.id"
    class="max-w-[200px] absolute w-[200px]"
    :style="{ left: nodeModel.posGridCell.x + 'px', top: nodeModel.posGridCell.y + 'px' }"
  >
    <Card>
      <!-- Delete Node Button -->
      <button @click="onDeleteNode" class="absolute top-1 right-1">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100" width="18px" height="18px">
          <path
            fill="#e15b64"
            d="M80.75 5h-61.5C11.393 5 5 11.393 5 19.25v61.5C5 88.607 11.393 95 19.25 95h61.5C88.607 95 95 88.607 95 80.75v-61.5C95 11.393 88.607 5 80.75 5zM69.445 62.374a5 5 0 1 1-7.072 7.071L50 57.071 37.626 69.445c-.977.976-2.256 1.464-3.536 1.464s-2.559-.488-3.536-1.464a5 5 0 0 1 0-7.071L42.929 50 30.555 37.626a5 5 0 1 1 7.071-7.071L50 42.929l12.374-12.374a5.001 5.001 0 0 1 7.071 7.071L57.071 50l12.374 12.374z"
            style="fill: rgb(225, 91, 100)"
          ></path>
          <path
            fill="#fff"
            d="M69.445 30.555a5.001 5.001 0 0 0-7.071 0L50 42.929 37.626 30.555a5.001 5.001 0 0 0-7.071 7.071L42.929 50 30.555 62.374a5 5 0 1 0 7.072 7.071L50 57.071l12.374 12.374c.977.976 2.256 1.464 3.536 1.464s2.559-.488 3.536-1.464a5 5 0 0 0 0-7.071L57.071 50l12.374-12.374a5 5 0 0 0 0-7.071z"
            style="fill: rgb(255, 255, 255)"
          ></path>
        </svg>
      </button>

      <!-- Node Name -->
      <h2 v-show="!showNameField" @click="openField(true)" class="text-center font-bold py-1">
        {{ nodeModel.data.name }}
      </h2>
      <input
        ref="nameFieldRef"
        type="text"
        v-model="name"
        v-show="showNameField"
        @focus="openField(true)"
        @blur="openField(false)"
        class="border rounded-lg p-1 font-bold w-[90%] text-center"
      />

      <!-- Ports -->
      <div class="w-full">
        <div v-for="port in ports" :key="port.id">
          <!-- Seperator with Deletion for Addable Ports -->
          <div v-if="checkLastGroupID(port.data)" class="relative w-full my-4">
            <div class="w-11/12 border-t border-gray-400 mx-auto"></div>
            <button @click="onDeleteAddablePorts(port.data.group_id)" class="absolute right-2 top-1 mb-2">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100" width="18px" height="18px">
                <path
                  fill="#e15b64"
                  d="M80.75 5h-61.5C11.393 5 5 11.393 5 19.25v61.5C5 88.607 11.393 95 19.25 95h61.5C88.607 95 95 88.607 95 80.75v-61.5C95 11.393 88.607 5 80.75 5zM69.445 62.374a5 5 0 1 1-7.072 7.071L50 57.071 37.626 69.445c-.977.976-2.256 1.464-3.536 1.464s-2.559-.488-3.536-1.464a5 5 0 0 1 0-7.071L42.929 50 30.555 37.626a5 5 0 1 1 7.071-7.071L50 42.929l12.374-12.374a5.001 5.001 0 0 1 7.071 7.071L57.071 50l12.374 12.374z"
                  style="fill: rgb(225, 91, 100)"
                ></path>
                <path
                  fill="#fff"
                  d="M69.445 30.555a5.001 5.001 0 0 0-7.071 0L50 42.929 37.626 30.555a5.001 5.001 0 0 0-7.071 7.071L42.929 50 30.555 62.374a5 5 0 1 0 7.072 7.071L50 57.071l12.374 12.374c.977.976 2.256 1.464 3.536 1.464s2.559-.488 3.536-1.464a5 5 0 0 0 0-7.071L57.071 50l12.374-12.374a5 5 0 0 0 0-7.071z"
                  style="fill: rgb(255, 255, 255)"
                ></path>
              </svg>
            </button>
          </div>

          <!-- Port -->
          <NodePort :port-model="port" />
        </div>
      </div>

      <!-- Add Addable Ports Button -->
      <div v-if="nodeModel.data.addablePorts.length > 0" class="w-full flex justify-center">
        <button
          class="w-full p-1 m-3 bg-blue-500 hover:bg-blue-700 text-white font-bold rounded"
          @click="onAddAddablePorts()"
        >
          ADD
        </button>
      </div>
    </Card>
  </div>
</template>

<script lang="ts">
import Card from "@/components/CardComponent.vue";
import NodePort from "./NodePort.vue";
import { defineComponent, nextTick, onMounted, onUnmounted, ref } from "vue";
import NodeModel from "@/models/Node/NodeModel";
import interact from "interactjs";
import Vector2 from "@/components/util/Vector";
import { InteractEvent } from "@interactjs/types";
import PortModel from "@/models/Node/PortModel";
import GridData from "@/models/Data/GridData";
import { PortType } from "@/models/Data/Types";
import { emitter } from "@/components/util/Emittery";

export default defineComponent({
  components: {
    Card,
    NodePort,
  },
  props: {
    nodeModel: {
      type: NodeModel,
      required: true,
    },
  },
  setup(props) {
    const ports = ref<PortModel[]>(props.nodeModel.children as PortModel[]);
    let lastGroupID = "";

    const showNameField = ref(false);
    const nameFieldRef = ref<HTMLInputElement>();
    const name = ref<string>(props.nodeModel.data.name);

    onMounted(() => {
      interact(`#${props.nodeModel.id}`).draggable({}).on("dragmove", onDragMove);
    });

    onUnmounted(() => {
      interact(`#${props.nodeModel.id}`).unset();
    });

    function onDragMove(event: InteractEvent) {
      props.nodeModel.addPos(new Vector2(event.dx, event.dy));
    }

    function onAddAddablePorts() {
      lastGroupID = "";
      props.nodeModel.addAddablePorts();
    }

    function onDeleteNode() {
      if (!props.nodeModel.parent) return;

      props.nodeModel.parent.removeChild(props.nodeModel.id);
    }

    function onDeleteAddablePorts(groupID: string) {
      lastGroupID = "";
      props.nodeModel.removeAddablePorts(groupID);
    }

    function checkLastGroupID(port: PortType): boolean {
      if (port.added && lastGroupID != port.group_id) {
        lastGroupID = port.group_id;
        return true;
      }
      return false;
    }

    //#region Editable Name Field
    function openField(bool: boolean) {
      showNameField.value = bool;

      if (bool) {
        nextTick(function () {
          nameFieldRef.value?.focus();
        });
      }
      else{
        props.nodeModel.setName(name.value);
      }
    }
    //#endregion

    return {
      onAddAddablePorts,
      onDeleteNode,
      checkLastGroupID,
      onDeleteAddablePorts,
      ports,
      showNameField,
      nameFieldRef,
      openField,
      name,
    };
  },
});
</script>

<style></style>
