<template>
  <v-container style="height: 100%">
    <v-row class="TopRow" align="start" no-gutters>
      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
        <div>
          <h2>TOP</h2>
        </div>
      </v-col>
      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
        <v-container class="SignInView d-flex justify-center">
          <h2>Sign In</h2>
          <div class="align-self-center">
            <v-icon icon="fa fa-plus" />
            <v-btn icon @click="SignInWithGoogle()">
              <v-icon icon="mdi-google"></v-icon>
            </v-btn>
          </div>
        </v-container>
      </v-col>
    </v-row>

    <v-overlay :model-value="overlay" class="align-center justify-center">
      <v-progress-circular indeterminate size="64" width="12" color="info"></v-progress-circular>
    </v-overlay>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue"
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth"
import { FirebaseAuth } from "../main"
import { userStore } from "../pinia/userStore"
import { Api } from "../apis/Api"
import router from "../router"

export default defineComponent({
  name: "TopView",
  setup() {
    const overlay = ref(false)
    const userStoreObj = userStore()

    onMounted(() => {
      overlay.value = true

      FirebaseAuth.onAuthStateChanged(async (user) => {
        if (user) {
          const token = await user.getIdToken()
          userStoreObj.setJwt(token)
          router.replace({ name: "Home" })
        }
      })
      overlay.value = false
    })

    const SignInWithGoogle = () => {
      const provider = new GoogleAuthProvider()
      signInWithPopup(FirebaseAuth, provider)
        .then(async (res) => {
          const token = await res.user.getIdToken()

          const apiRes = await Api.createUser({
            firebase_jwt: token,
          })

          if (apiRes.result_flg === 1) {
            userStoreObj.setJwt(token)
            router.replace({ name: "Home" })
          } else {
            userStoreObj.resetJWT()
            router.replace({ name: "Top" })
          }
        })
    }

    return {
      overlay,

      SignInWithGoogle,
    }
  }
})
</script>

<style lang="scss" scoped>
.SignInView {
  height: 100%;
}

.TopRow {
  height: 100%;
}
</style>