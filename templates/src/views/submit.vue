<template>
  <v-container>
    <codeEditor
      @send="send()"
      submit='true'
      ref="editor"></codeEditor>
  </v-container>
</template>

<script>
// POST /api/problem/submit, body : { prob_no, code, id, lang }
import codeEditor from '../semiViews/codeEditor.vue'
export default{
  components: {
    codeEditor,
  },
  methods: {
    send() {
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
          prob_no: this.$route.params.num
        }, this.$f.makeHeaderObject())
        .then(res => {
          this.$router.push({path:'/status?prob_no=' + this.$f.params.num})
        }).catch(err => {this.$f.malert()})
      })
    }
  },
}
</script>
