<template>
  <div class="circle" @click.passive="onClick">

  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue"

export default defineComponent({
  setup() {
    function onClick(){
      console.log("Click")
    }

    function onRelease(){
      console.log("Release")
    }

    return{
      onClick,
      onRelease
    }
  }
})
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