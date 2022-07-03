<template>
  <v-app :theme="'mainLightTheme'">
    <v-app-bar color="primary">

      <v-app-bar-nav-icon @click="drawerOpening = !drawerOpening"></v-app-bar-nav-icon>

      <v-app-bar-title class="AppBarTitle" @click="jumpTopPage()">ASキュー</v-app-bar-title>
      <v-spacer />
    </v-app-bar>

    <v-navigation-drawer v-model="drawerOpening" bottom temporary>
      <template v-if="refUserStoreObj.getJWT.value === ''">
        <v-list nav>
          <v-list-item>
            <router-link :to="{ name: 'About' }">About</router-link>
          </v-list-item>
        </v-list>
      </template>
      <template v-else>
        <v-list nav>
          <v-list-item>
            <div class="pa-2">
              <v-btn block @click="signout()">Sign Out</v-btn>
            </div>
          </v-list-item>
        </v-list>
      </template>
    </v-navigation-drawer>

    <v-main>
      <router-view />
    </v-main>

  </v-app>
</template>

<script lang="ts">
import { storeToRefs } from 'pinia'
import { defineComponent, onMounted, ref } from 'vue'
import { useTheme } from 'vuetify'
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
    const theme = useTheme()


    onMounted(async () => {
      const date = new Date()
      const hour = date.getHours()
      console.log(hour, theme.name.value)
      // if (hour <= 6 || hour >= 18) {
      //   console.log("dark")
      //   theme.name.value = 'mainDarkTheme'
      // } else {
      //   theme.name.value = 'mainLightTheme'
      // }
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
