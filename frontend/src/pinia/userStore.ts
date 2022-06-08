import { defineStore } from "pinia";

export const userStore = defineStore("userStore", {
  state: () => {
    return {
      jwt: "",
    };
  },
  getters: {
    getJWT: (state) => state.jwt,
  },
  actions: {
    setJwt(newJWT: string) {
      this.jwt = newJWT.trim();
    },
    resetJWT() {
      this.jwt = "";
    },
  },
});
