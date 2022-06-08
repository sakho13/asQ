<template>
  <v-container>
    <h1>HOME</h1>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { Api } from '../../apis/Api';
import { userStore } from '../pinia/userStore';
import router from '../router';

export default defineComponent({
  name: 'HomeView',

  setup() {
    const userStoreObj = userStore()

    onMounted(async () => {
      const res = await Api.createUser({
        firebase_jwt: userStoreObj.getJWT
      })
      if (res.result_flg === 1 && res.response) {
        console.log(res.response.through)
      } else {
        router.replace({ name: "SignIn" })
      }
    })
  }
});
</script>
