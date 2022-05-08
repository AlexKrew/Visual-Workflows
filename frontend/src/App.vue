<template>
  <v-app>
    <v-app-bar>
    </v-app-bar>
    <v-main>
      <v-container fluid class="bg-red-300 fill-height ma-0">
        <router-view/>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import NavBar from "./components/NavBar.vue"

export default {
  components: {
    NavBar
  }
}
</script>
