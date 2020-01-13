<template>
  <v-app id="inspire">
    <v-app-bar
      app
      color="white"
    >
      <v-btn height="50" rotuer :to="{path : '/'}" class="pa-0" x-large outlined>
        <v-img
          src="./assets/logo1.jpg"
          aspect-ratio="1"
          max-width="140"
          max-height="45"
        ></v-img>
      </v-btn>
      <v-row class="pr-5" justify="end">
        <v-btn v-if="$f.userId !== ''" text>{{$f.userId}}</v-btn text>
        <v-btn v-if="$f.userId === ''" router :to="{path : '/register'}" text>회원가입</v-btn>
        <v-btn v-if="$f.userId === ''" router :to="{path : '/login'}" text>로그인</v-btn>
        <v-btn v-if="$f.userId !== ''" @click="logout()" text>로그아웃</v-btn>

        <v-btn router :to="{path : '/list'}" text>문제 목록</v-btn>
        <v-btn v-if="$f.userId !== ''" router :to="{path:'/status'}" text>채점 현황</v-btn>
        <v-btn v-if="$f.userId !== ''" router :to="{path : '/sup/list'}" text>문제 관리</v-btn>
        <v-btn router :to="{name:'test'}" text>테스트</v-btn>
      </v-row>
    </v-app-bar>
    <v-content>
      <router-view :key="$route.fullPath"/>
    </v-content>
  </v-app>
</template>

<style>
    .centered-input >>> input {
      text-align: center
    }
</style>

<script>
  export default {
    created(){
      this.$f.getUserValid()
    },
    data: () => ({

    }),
    methods: {
      logout() {
        localStorage.removeItem('token')
        this.$f.userId = ""
        this.$forceUpdate()
      }
    }
  }
</script>
