<template>
  <v-container>
    <v-row v-if="token==''" justify="center" align="center">
      <v-form
        ref="form"
        v-model="valid"
        style="min-width:43%"
      >
        <v-text-field
          v-model="id"
          :rules="idRules"
          :counter="20"
          label="아이디"
          maxlength="20"
          @keyup.enter="sendMail()"
          required
        ></v-text-field>
        <v-text-field
          v-model="email"
          :rules="emailRules"
          label="E-mail"
          maxlength="50"
          @keyup.enter="sendMail()"
          required
        ></v-text-field>
        <v-btn
          :disabled="!valid"
          color="success"
          class="mr-4"
          @click="sendMail()"
        >
          확인
        </v-btn>
        <p class="pt-3">확인버튼을 누르면 비밀번호 초기화 인증 메일이 전송됩니다.</p>
      </v-form>
    </v-row>

    <v-row v-else justify="center" align="center">
      <v-form
        ref="form"
        v-model="valid"
        style="min-width:43%"
      >
        <v-text-field
          v-model="pass"
          :rules="passRules"
          type="password"
          label="비밀번호"
          @keyup.enter="changePass()"
          maxlength="20"
          required
        ></v-text-field>

        <v-text-field
          v-model="passConf"
          :rules="passConfRules"
          type="password"
          label="비밀번호 확인"
          @keyup.enter="changePass()"
          maxlength="20"
          required
        ></v-text-field>
        <v-btn
          :disabled="!valid"
          color="success"
          class="mr-4"
          @click="changePass()"
        >
          비밀번호 변경
        </v-btn>
        <p class="pt-3">이메일이 인증되었습니다. 새로운 비밀번호를 입력하세요.</p>
      </v-form>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    data: () => ({
      valid: true,
      id: '',
      pass: '',
      passRules: [
        v => !!v || '비밀번호를 입력하세요',
        v => (/^(?=.*[a-zA-Z])(?=.*\d).{8,20}$/.test(v) && v.length >= 8 && v.length <= 20) || '비밀번호 8~20글자이며 알파벳과 숫자를 반드시 포함해야 됩니다.',
      ],
      passConf: '',
      email: '@inu.ac.kr',
      token: '',
    }),
    created() {
      if (typeof this.$route.query.token !== 'undefined') this.token = this.$route.query.token
    },
    methods: {
      sendMail () {
          this.$axios.post('/api/reset/pass', {
            id : this.id,
            email : this.email,
          }).then(res => {
            if (res.data.status == 'ok') alert("비밀번호 초기화 이메일을 전송하였습니다.")
            else if (res.data.status == 'no info') alert("일치하는 정보가 없습니다.")
            else if (res.data.status == 'fail') alert('오늘 가능한 횟수를 초과하였습니다. 24시간 뒤에 다시 시도해주세요.')
          }).catch(err => {
            alert('이메일을 전송 도중 오류가 발생하였습니다.')
          })
      },
      changePass () {
          this.$axios.post('/api/reset/pass/done', {
            token : this.token,
            pass : this.$sha256(this.pass),
          }).then(res => {
            if (res.data.status == 'ok') { this.$router.push('/login') }
            else { alert("토큰이 만료되었습니다.") }
          }).catch(err => {
            alert('비밀번호 변경 도중 오류가 발생했습니다.')
          })
      },
    },
    computed: {
      passConfRules() {
        const rules=[v => !!v || '비밀번호를 확인하세요',]
        rules.push(this.pass === this.passConf || '비밀번호를 확인하세요')
        return rules
      },
    },
  }
</script>
