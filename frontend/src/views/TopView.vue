<template>
  <v-container style="height: 100%">
    <v-row class="TopRow" align="start" no-gutters>
      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
        <div>
          <h2>TOP</h2>
        </div>
      </v-col>
      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
        <!-- navigation of sign in -->
        <login-card />
      </v-col>
    </v-row>

    <v-overlay :model-value="overlay" class="align-center justify-center">
      <v-progress-circular indeterminate size="64" width="12" color="info"></v-progress-circular>
    </v-overlay>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue"
import { FirebaseAuth } from "../main"
import { userStore } from "../pinia/userStore"
import { Api } from "../apis/Api"
import router from "../router"
import LoginCard from "../components/LoginCard.vue"

export default defineComponent({
  name: "TopView",
  setup() {
      const overlay = ref(false);
      const userStoreObj = userStore();
      onMounted(async () => {
          overlay.value = true;
          await FirebaseAuth.onAuthStateChanged(async (user) => {
              if (user) {
                  const token = await user.getIdToken();
                  userStoreObj.setJwt(token);
                  router.replace({ name: "Home" });
              }
          });
          overlay.value = false;
      });

      return {
          overlay,
      };
  },
  components: { LoginCard }
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