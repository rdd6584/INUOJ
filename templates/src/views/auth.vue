<template>
  <v-container>
    <h1>이메일 인증을 완료하세요!</h1>
    <h1>버튼을 누르면 가입하신 이메일로 인증 메일이 재전송됩니다.</h1>
    <h1>메일이 수신되지 않는 경우, 메일차단내역을 확인하세요.</h1>
    <v-btn @click="sendAuth()">이메일 재전송</v-btn>
  </v-container>
</template>

<script>
export default {
  created() {
    if (typeof this.$route.query.email !== 'undefined'
      && typeof this.$route.query.token !== 'undefined') {
        this.$axios.post('/api/emailauth', {
          token: this.$route.query.token,
          email: this.$route.query.email,
        })
        .then(res => {
          if (res.data.status == 'ok')
            this.$router.push({path:'/login'})
          else alert("만료된 인증입니다")
        })
        .catch(err => {alert("예기치 못한 오류입니다.")})
      }
  },
  methods: {
    sendAuth() {
      this.$axios.post('/api/sendauth', {
        id : this.$route.query.id,
      })
      .then(res => {
        if (res.data.status == 'done') { alert("이미 인증된 회원입니다.") }
        else if (res.data.status == 'fail') { alert("인증 요청 횟수를 초과하였습니다.\n24시간 뒤에 회원가입을 다시 요청하세요.") }
        else { alert("이메일이 전송되었습니다.") }
      })
      .catch(err => {alert("예기치 못한 오류가 발생했습니다")})
    },
  },
}
</script>
