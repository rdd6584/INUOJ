
<template>
  <v-app id="inspire">
    <v-app-bar
      app
      color="white"
      dense
    >
      <v-btn router :to="{path : '/'}" class="pa-3">LOGO</v-btn>
      <v-btn v-if="$f.getUserValid() != true" router :to="{path : '/register'}" text>회원가입</v-btn>
      <v-btn v-if="$f.getUserValid() != true" router :to="{path : '/login'}" text>로그인</v-btn>
      <v-btn v-if="$f.getUserValid() == true" text
      router :to= "{
        name : 'users',
        params: {
          userId: $store.state.userID
        },
      }">{{$store.state.userID}}</v-btn>

      <v-btn v-if="$f.getUserValid() == true" text
      router :to= "{
        name : 'status',
        query: {
          prob_no: 0, id, lang, result, page
        },
      }">채점 현황</v-btn>
      <v-btn router :to="{path : '/test'}" text class="pa-0">테스트</v-btn>
    </v-app-bar>
    <v-content>
      <router-view></router-view>
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
    data: () => ({

    }),
  }
</script>
