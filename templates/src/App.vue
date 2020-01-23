<template>
  <v-app id="inspire">
    <v-app-bar
      app
      color="white"
    >
      <v-btn elevation="0" width="140" height="45" rotuer :to="{path : '/'}" class="pa-0">
        <v-img
          src="./assets/logo1.jpg"
          aspect-ratio="1"
          max-width="140"
          max-height="45"
        ></v-img>
      </v-btn>

      <v-row class="pr-5" justify="end">
        <div v-if="$f.userId === ''">
          <v-btn router :to="{path : '/register'}" text>회원가입</v-btn>
          <v-btn router :to="{path : '/login'}" text>로그인</v-btn>
        </div>

        <div v-if="$f.userId !== ''">
          <v-btn class="my-case" text router :to="{path : '/profile/' + $f.userId}">{{$f.userId}}</v-btn text>
          <v-btn @click="logout()" text>로그아웃</v-btn>
          <v-btn router :to="{path : '/list'}" text>문제 목록</v-btn>
          <v-btn router :to="{path:'/status'}" text>채점 현황</v-btn>
        </div>

        <div v-if="$f.admin == '1' || $f.admin == '2'">
          <v-btn router :to="{path : '/sup/list'}" text>문제 관리</v-btn>
        </div>
      </v-row>
    </v-app-bar>
    <v-content>
      <router-view :key="$route.fullPath"/>
    </v-content>
  </v-app>
</template>

<style>
    .my-case {
      text-transform: none;
    }
    .centered-input >>> input {
      text-align: center;
    }
</style>

<script>
  export default {
    created(){
      this.$f.getUserValid()
      this.$forceUpdate()
    },
    data: () => ({

    }),
    methods: {
      logout() {
        localStorage.removeItem('token')
        this.$f.userId = ""
        this.$f.admin = ""
        this.$forceUpdate()
      }
    }
  }
</script>
