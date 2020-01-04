
<template>
  <v-app id="inspire">
    <v-app-bar
      app
      color="white"
      dense
    >
      <v-toolbar-title class="pa-3">INUOJ</v-toolbar-title>
        <v-btn router :to="{path : '/register'}" text class="pa-2">회원가입</v-btn>
        <v-btn router :to="{path : '/login'}" text class="pa-0">로그인</v-btn>
    </v-app-bar>
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
  export default {
    data: () => ({

    }),
  }
</script>
