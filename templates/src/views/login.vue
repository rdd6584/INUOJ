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
        >로그인
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
        this.$axios.post('/api/login', {id:this.id, pass:this.pass})
        .then(res => {
          localStorage.setItem('token', res.data.token)
          this.$router.push({path:'/'})
        })
        .catch(err => {
          if (err.response.data.status == "failLogin") alert("아이디와 비밀번호를 확인하세요")
          else if (err.response.data.status == "failAuth") {
            this.$router.push({path:"/auth?id=" + this.id})
          }
          else alert("예기치 못한 오류입니다")
        })
      },
    },
  }
</script>
