<template>
  <v-container>
    <v-card style="width: 100%; height: 100%">
      <v-card-title>Sign In</v-card-title>
      <v-card-content>
        <div class="align-self-center">
          <v-icon icon="fa fa-plus" />
          <v-btn icon @click="SignInWithGoogle()">
            <v-icon icon="mdi-google"></v-icon>
          </v-btn>
        </div>
      </v-card-content>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import { defineComponent } from "vue"
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth"
import { FirebaseAuth } from "../main"
import { userStore } from "../pinia/userStore"
import router from "../router"
import { Api } from "../apis/Api"

export default defineComponent({
  name: "LoginCard",
  setup() {
    const userStoreObj = userStore()

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
      SignInWithGoogle
    }
  }
})
</script>