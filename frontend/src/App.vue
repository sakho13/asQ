<template>
  <v-app :theme="'mainLightTheme'">
    <v-app-bar color="primary">
      <v-app-bar-title>WithMe</v-app-bar-title>
    </v-app-bar>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Api } from "../apis/Api"

export default defineComponent({
  name: 'App',

  setup() {
    Api.hello()
      .then((res) => {
        console.table(res)
      })
      .catch((err) => {
        console.log(err)
      })

  }
})
</script>
