<template>
  <v-container>
    <v-row justify="center" align="center">
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
            @keyup="valCheck(1)"
            required
          ></v-text-field>

          <v-text-field
            v-model="pass"
            :rules="passRules"
            type="password"
            label="비밀번호"
            maxlength="20"
            required
          ></v-text-field>

          <v-text-field
            v-model="passConf"
            :rules="passConfRules"
            type="password"
            label="비밀번호 확인"
            maxlength="20"
            required
          ></v-text-field>

          <v-text-field
            v-model="email"
            :rules="emailRules"
            label="E-mail"
            maxlength="50"
            @keyup="valCheck(2)"
            required
          ></v-text-field>
          <v-btn
            :disabled="!valid"
            color="success"
            class="mr-4"
            @click="register()"
          >
            회원가입
          </v-btn>
        </v-form>
      </v-card>
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
      idCheck: false,
      emailCheck: false,
    }),
    methods: {
      register () {
          this.$axios.post('/api/regi-done', {
            id : this.id,
            pass : this.$sha256(this.pass),
            email : this.email,
          }).then(res => {
            if (res.data.status == 'ok') {
              this.$router.push({path:'/auth?id=' + this.id})
            }
            else alert('회원가입 도중 오류가 발생했습니다.')
          }).catch(err => {
            alert('회원가입 도중 오류가 발생했습니다.')
          })
      },
      valCheck(req) {
        var query=""
        if (req == 1) query = `id=${this.id}`
        else if (req == 2) query = `email=${this.email}`
        this.$axios.get(`/api/register/valid?` + query)
        .then(res => {
          if (req == 1) this.idCheck = res.data.status
          else if (req == 2) this.emailCheck = res.data.status
        })
      },
    },
    computed: {
      idRules() {
        const rules=[v => !!v || 'ID를 입력하세요',]
        rules.push(v => (/^[a-zA-Z0-9]+/.test(v) && v.length >= 2 && v.length <= 20) || '아이디는 2~20글자이며 알파벳과 숫자로만 구성됩니다.',)
        rules.push(this.idCheck || '중복되는 ID가 있습니다.')
        return rules
      },
      emailRules() {
        const rules=[v => !!v || '이메일을 입력하세요',]
        rules.push(v => /^[a-zA-Z0-9\_]+@inu\.ac\.kr/.test(v) || '이메일은 @inu.ac.kr로 끝나야 합니다.')
        rules.push(this.emailCheck || '중복되는 이메일이 있습니다.')
        return rules
      },
      passConfRules() {
        const rules=[v => !!v || '비밀번호를 확인하세요',]
        rules.push(this.pass === this.passConf || '비밀번호를 확인하세요')
        return rules
      },
    },
  }
</script>
