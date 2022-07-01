<template>
  <v-container></v-container>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import { FirebaseAuth } from '../main'
import router from '../router'

export default  defineComponent({
  name: "SingOut",
  setup() {
    onMounted(async () => {
      await FirebaseAuth.signOut()
      router.replace({ name: "Top" })
    })
  }
})
</script>