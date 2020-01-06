<template>
  <v-row justify="center" align="center">
      <v-form
        ref="form"
        style="min-width:500px"
      >
        <v-text-field
          v-model="id"
          :counter="20"
          label="아이디"
          maxlength="20"
          required
        ></v-text-field>

        <v-text-field
          v-model="pass"
          :counter="20"
          type="password"
          label="비밀번호"
          maxlength="20"
          required
        ></v-text-field>

        <v-btn
          color="success"
          class="mr-4"
          @click="login()"
        >
          로그인
        </v-btn>
      </v-form>
    </v-card>
  </v-row>
</template>

<script>
  export default{
    data: () => ({
      id: "",
      pass: "",
    }),
    methods: {
      login() {
        this.$axios.post('/api/login/valid', {id:this.id, pass:this.pass})
        .then(res => {
          if (res.data.status) {
            localStorage.setItem('token', res.data.token)
            localStorage.setItem('auth', res.data.auth)
          } else alert("아이디와 비밀번호를 확인해주세요.")
        })
        .catch(err => {alert("예기치 못한 오류가 발생했습니다.")})
      },
    },
  }
</script>
