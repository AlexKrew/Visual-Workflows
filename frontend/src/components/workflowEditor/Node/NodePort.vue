<template>
  <div :id="portModel.id" class="fill-width flex">
    <div class="circle justify-start" v-if="portModel.isInput"></div>
    <span class="justify-center flex-auto">{{ portModel.title }}</span>
    <div class="circle justify-end" v-if="!portModel.isInput"></div>
  </div>
</template>

<script lang="ts">
import NodePortModel from "@/models/NodePortModel";
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    portModel: {
      type: NodePortModel,
      required: true,
    },
  },
});
</script>

<style>
.circle {
  width: 20px;
  height: 20px;
  -webkit-border-radius: 25px;
  -moz-border-radius: 25px;
  border-radius: 25px;
  background: grey;
}
</style>
