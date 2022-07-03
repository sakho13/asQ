<template>
  <v-container>
    <div>
      <h1>HOME</h1>
    </div>
    <v-overlay :model-value="overlay" class="align-center justify-center">
      <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import { Api } from '../apis/Api';
import { userStore } from '../pinia/userStore';
import router from '../router';

export default defineComponent({
  name: 'HomeView',

  setup() {
    const userStoreObj = userStore()
    const overlay = ref(false)

    onMounted(async () => {
      overlay.value = true

      const res = await Api.createUser({
        firebase_jwt: userStoreObj.jwt,
      })
      if (res.result_flg === 1) {
        console.log(res.response)
      } else {
        router.replace({ name: "SignIn" })
      }

      overlay.value = false
    })

    return {
      overlay,
    }
  }
});
</script>
