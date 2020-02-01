<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-form
        ref="form"
        style="min-width:43%"
      >
        <v-text-field
          v-model="id"
          :counter="20"
          label="아이디"
          maxlength="20"
          required
          @keyup.enter="login()"
        ></v-text-field>

        <v-text-field
          v-model="pass"
          :counter="20"
          type="password"
          label="비밀번호"
          maxlength="20"
          required
          @keyup.enter="login()"
        ></v-text-field>

        <v-row>
          <v-btn
            color="success"
            class="ml-4"
            @click="login()"
          >로그인
          </v-btn>
          <v-spacer></v-spacer>
          <v-col class="pt-2">
            <a class="hover-line" @click="$router.push({path: '/reset'})">비밀번호를 잊어버리셨나요?</a>
          </v-col>
        </v-row>
      </v-form>
    </v-row>
  </v-container>
</template>

<script>
  export default{
    data: () => ({
      id: "",
      pass: "",
    }),
    methods: {
      login() {
        this.$axios.post('/api/login', {id:this.id, pass:this.$sha256(this.pass)})
        .then(res => {
          localStorage.setItem('token', res.data.token)
          this.$f.getUserValid() // $f.userId, admin 정보 갱신
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
