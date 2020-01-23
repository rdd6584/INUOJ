<template>
  <v-container>

  </v-container>
</template>

<script>
import textEditor from '../semiViews/textEditor.vue'
import codeEditor from '../semiViews/codeEditor.vue'

export default{
  components: {
    textEditor,
    codeEditor,
  },
  created() {

  },
  data: () => ({
    post_no: 0,
    title: "",
    content: "",
    code: "",
    id: "",
    category: "", // 변환해서 주기
    prob_no: 0,
    cmt_no: 0,
    cmt_list: [],
    post_time: "",
  }),
}
</script>
