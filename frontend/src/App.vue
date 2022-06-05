<template>
  <v-app :theme="'mainLightTheme'">
    <v-app-bar color="primary">

      <v-app-bar-nav-icon @click="drawerOpening = !drawerOpening"></v-app-bar-nav-icon>

      <v-app-bar-title class="AppBarTitle" @click="jumpTopPage()">WithMe</v-app-bar-title>
      <v-spacer />
    </v-app-bar>

    <v-navigation-drawer v-model="drawerOpening" bottom temporary>
      <v-list nav>
        <v-list-item>
          <router-link :to="{name: 'SignIn'}">Sign In</router-link>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { Api } from "../apis/Api"
import router from './router'

export default defineComponent({
  name: 'App',

  setup() {
    const drawerOpening = ref(false)

    // 疎通確認
    Api.hello()
      .then((res) => {
        console.table(res)
      })
      .catch((err) => {
        console.log(err)
      })


    const jumpTopPage = () => {
      router.push({ name: "Top" })
    }

    return {
      drawerOpening,

      jumpTopPage,
    }
  }
})
</script>

<style lang="scss" scoped>
.AppBarTitle {
  cursor: pointer;
}
</style>
