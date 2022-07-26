<template>
  <Card class="bg-gray-200 absolute right-0 w-2/12 h-full z-10">
    <p>Inspector</p>
    <button @click="onDeploy">Deploy</button>
  </Card>
</template>

<script lang="ts">
import Card from "@/components/CardComponent.vue";
import GridData from "@/models/Data/GridData";
export default {
  components: {
    Card,
  },
  setup() {
    function onDeploy(){
      console.log(GridData.workflow);
    }

    return{
      onDeploy
    }
  }
};
</script>

<style></style>
