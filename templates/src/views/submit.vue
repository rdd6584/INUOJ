<template>
  <v-container>
    <codeEditor
      @send="send()"
      readOnly="0"
      ref="editor"></codeEditor>
      <div class="pb-12"></div>
  </v-container>
</template>

<script>
import codeEditor from '../semiViews/codeEditor.vue'
export default{
  components: {
    codeEditor,
  },
  methods: {
    send() {
      console.log(this.$route.params.num)
      console.log(this.$route)

      this.$f.getUserValid()
      .then(re => {
        if (re == null) {
          this.$router.push({path:'/login'})
          return
        }
        this.$axios.post('/api/problem/submit', {
          code: this.$refs.editor.code,
          lang: this.$store.state.lang.indexOf(this.$refs.editor.choice),
          id: this.$f.userId,
          prob_no: Number(this.$route.params.num)
        }, this.$f.makeHeaderObject())
        .then(res => {
          this.$router.push({path:'/status?prob_no=' + this.$route.params.num})
        }).catch(err => {
          if (err.response.status == 404) {
            alert("존재하지 않는 문제입니다.")
            return
          }
          this.$f.malert()}
        )
      })
    }
  },
}
</script>
