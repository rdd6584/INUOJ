<template>
  <v-container>
    <v-row>
        <v-layout>
          <v-flex style="text-align:center;">
            <h1>안녕하세요!</h1>
            <h1>인천대학교 알고리즘 커뮤니티에 오신 것을 환영합니다.</h1>

            <div v-if="$f.userId != ''">건의사항 및 문의 가입탈퇴는 메일로 보내주세요. inuojteam@gmail.com</div>
          </v-flex>
        </v-layout>
    </v-row>
  </v-container>
</template>

<script>
  export default {

  }
</script>
