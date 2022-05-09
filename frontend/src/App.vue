<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated class="text-dark">
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          aria-label="Menu"
          icon="menu"
        />

        <q-toolbar-title>
          With Me
          <q-icon :color="apiServerIsOn ? 'red': 'grey'" name="api"/>
        </q-toolbar-title>

        <!-- <div>Quasar v{{ $q.version }}</div> -->
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import {
  defineComponent,
  watch,
  onMounted,
  ref
} from 'vue'
import { useQuasar } from "quasar"
import { Api } from "../apis/Api"

export default defineComponent({
  name: "LayoutDefault",

  setup() {
    const $q = useQuasar()
    const apiServerIsOn = ref(false)

    onMounted(async () => {
      const res = await Api.hello()
      if (res.result_flg === 1) {
        apiServerIsOn.value = true
      }
    })


    watch(
      () => $q.dark.isActive,
      (val) => {
        console.log(val ? 'On dark mode' : 'On light mode')
      }
    )

    return {
      apiServerIsOn,
    }
  }
})
</script>
