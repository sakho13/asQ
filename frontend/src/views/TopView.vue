<template>
  <v-container>
    <h3>Here is TOP page!</h3>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue"
import { FirebaseAuth } from "../main"
import { userStore } from "../pinia/userStore"
import router from "../router"

export default defineComponent({
  name: "TopView",
  setup() {
    onMounted(() => {
      FirebaseAuth.onAuthStateChanged(async (user) => {
        if (user) {
          const token = await user.getIdToken()
          const userStoreObj = userStore()
          userStoreObj.setJwt(token)
          router.replace({ name: "Home" })
        }
      })
    })
  }
})
</script>
