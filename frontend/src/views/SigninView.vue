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
import { defineComponent } from "vue"
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth"
import { FirebaseAuth } from "../main"
import router from "../router"

export default defineComponent({
  name: "SignInView",

  setup() {

    const SignInWithGoogle = () => {
      console.log("SignInWithGoogle")
      const provider = new GoogleAuthProvider()
      signInWithPopup(FirebaseAuth, provider)
        .then(() => {
          router.replace({ name: "Home" })
        })
        .catch((err) => {
          console.table(err)
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
