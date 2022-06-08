<template>
  <v-container class="SignInView d-flex justify-center">
    <div class="align-self-center">
      <v-icon icon="fa fa-plus" />
      <v-btn icon @click="SignInWithGoogle()">
        <v-icon icon="mdi-google"></v-icon>
      </v-btn>
    </div>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue"
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth"
import { FirebaseAuth } from "../main"
import router from "../router"
import { Api } from "../../apis/Api"
import { userStore } from "../pinia/userStore"
import { storeToRefs } from "pinia"

export default defineComponent({
  name: "SignInView",

  setup() {
    const userStoreObj = userStore()
    const refUserStoreObj = storeToRefs(userStoreObj)

    onMounted(() => {
      FirebaseAuth.onAuthStateChanged(async (user) => {
        if (user) {
          const token = await user.getIdToken()
          userStoreObj.setJwt(token)
          router.replace({ name: "Home" })
        }
      })
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
      SignInWithGoogle,
    }
  }
})
</script>

<style lang="scss" scoped>
.SignInView {
  height: 100%;
}
</style>
