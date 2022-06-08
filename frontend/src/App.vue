<template>
  <v-app :theme="'mainLightTheme'">
    <v-app-bar color="primary">

      <v-app-bar-nav-icon @click="drawerOpening = !drawerOpening"></v-app-bar-nav-icon>

      <v-app-bar-title class="AppBarTitle" @click="jumpTopPage()">WithMe</v-app-bar-title>
      <v-spacer />
    </v-app-bar>

    <v-navigation-drawer v-model="drawerOpening" bottom temporary>
      <v-list nav>
        <template v-if="refUserStoreObj.getJWT.value === ''">
          <v-list-item>
            <router-link :to="{ name: 'SignIn' }">Sign In</router-link>
          </v-list-item>
        </template>
        <template v-else>
          <!--  -->
        </template>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { storeToRefs } from 'pinia'
import { defineComponent, ref } from 'vue'
import { Api } from "../apis/Api"
import { userStore } from './pinia/userStore'
import router from './router'
import { RouterLink, RouterView } from "vue-router"
import { FirebaseAuth } from './main'

export default defineComponent({
  name: 'App',

  components: {
    RouterLink,
    RouterView,
  },

  setup() {
    const drawerOpening = ref(false)
    const userStoreObj = userStore()
    const refUserStoreObj = storeToRefs(userStoreObj)

    // 疎通確認
    Api.hello()
      .then((res) => {
        // console.table(res)
      })

    const signout = async () => {
      await FirebaseAuth.signOut()
    }

    const jumpTopPage = () => {
      router.push({ name: "Top" })
    }

    return {
      drawerOpening,
      refUserStoreObj,

      jumpTopPage,
      signout,
    }
  }
})
</script>

<style lang="scss" scoped>
.AppBarTitle {
  cursor: pointer;
}
</style>
