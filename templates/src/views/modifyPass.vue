<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-card elevation="3" min-width="50%">
        <v-card-title>비밀번호 변경</v-card-title>
        <v-divider></v-divider>
        <v-form
          class="ma-5"
          ref="form"
          v-model="valid"
        >
          <v-text-field
            v-model="nowPass"
            type="password"
            label="현재 비밀번호"
            maxlength="20"
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
          <v-btn
            :disabled="!valid"
            color="success"
            class="mr-4"
            @click="save()"
          >
            변경
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
      pass: '',
      nowPass: '',
      passRules: [
        v => !!v || '비밀번호를 입력하세요',
        v => (/^(?=.*[a-zA-Z])(?=.*\d).{8,20}$/.test(v) && v.length >= 8 && v.length <= 20) || '비밀번호 8~20글자이며 알파벳과 숫자를 반드시 포함해야 됩니다.',
      ],
      passConf: '',
    }),
    methods: {
      save () {
        this.$axios.post('/api/edit/password', {
          id : this.$f.userId,
          pass : this.$sha256(this.nowPass),
          newpass : this.$sha256(this.pass),
        }, this.$f.makeHeaderObject()).then(res => {
          if (res.data.status == 'fail') alert("비밀번호를 확인해주세요.")
          else {
            this.$router.push({path:'/profile/' + this.$f.userId})
            alert('비밀번호가 변경되었습니다.')
          }
        }).catch(err => { this.$f.malert() })
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
