<template>
  <v-container>
    <v-card>
      <v-row class="pl-8 pt-4">
        <a @click="$router.push({path:'/problem/' + prob_no})"><h4>{{prob_no}} - {{prob_title}}</h4></a>
        <h4 class="pl-3">{{category}}</h4>
      </v-row>
      <v-row class="pl-4">
        <v-card-title>{{title}}</v-card-title>
        <v-spacer></v-spacer>
      </v-row>

    </v-card>
    <v-card class="mt-1 pl-4">
      <v-row class="pl-4">
        <v-card-title class="px-0 pb-0">본문</v-card-title>
        <v-spacer/>
        <v-col class="pr-6" style="text-align:right;" align-self="end">
          <a @click="$router.push({path: '/profile/' + id})">{{id}} </a>
          <div>{{post_time}}</div>
        </v-col>
      </v-row>
      <v-divider></v-divider>
      <div class="pt-6"/>
      <span v-html="content"/>
      <div class="py-6"/>
    </v-card>

    <div class="pt-1" v-if="codeExist==true">
      <codeEditor ref="codeView" readOnly="2"/>
    </div>

    <div class="py-6">
      <v-divider></v-divider>
    </div>


    <v-card v-for="(val, i) in cmt_list" class="mt-1 pl-4">
      <v-row class="pl-4">
        <v-card-title class="px-0 pb-0">댓글</v-card-title>
        <v-spacer/>
        <v-col class="pr-6" style="text-align:right;" align-self="end">
          <a @click="$router.push({path: '/profile/' + val.id})">{{val.id}} </a>
          <div>{{val.cmt_time}}</div>
        </v-col>
      </v-row>
      <v-divider></v-divider>
      <div class="pt-6"/>
      <span v-html="val.comment"/>
      <div class="py-6"/>
    </v-card>

    <textEditor ref="textEdit" place="댓글을 작성해주세요."></textEditor>
    <v-row justify="center">
      <v-btn @click="save()" color="success">작성</v-btn>
    </v-row>
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
  mounted() {
    this.post_no = this.$route.params.num
    
    this.$axios.get('/api/board/view/' + this.post_no, this.$f.makeHeaderObject())
    .then(res => {
      this.title = res.data.title
      this.content = res.data.content
      if (res.data.code != '') this.$refs.codeView.code = res.data.code
      else this.codeExist = false

      this.id = res.data.id
      this.category = this.$store.state.category[res.data.category]
      this.prob_no = Number(res.data.prob_no)
      this.prob_title = res.data.prob_title
      if (res.data.cmt_list) this.cmt_list = res.data.cmt_list
      this.post_time = res.data.post_time
    }).catch(err => {this.$f.malert()})
  },
  data: () => ({
    post_no: 0,
    title: "",
    content: "",
    id: "",
    category: "", // 변환해서 주기
    prob_no: 0,
    prob_title: "",
    cmt_list: [],   // post_no, id, comment, cmt_time
    post_time: "",
    codeExist: true,
  }),
  methods: {
    save() {
      this.$f.getUserValid()
      .then(re => {
        if (re == null) {
          this.$router.push('/login')
          return
        }
        this.$axios.post('/api/board/new/comment', {
          post_no: this.post_no,
          id: this.$f.userId,
          content: this.$refs.textEdit.content,
        }, this.$f.makeHeaderObject())
        .then(res => {
          this.cmt_list.push({
            post_no: this.post_no,
            id: this.$f.userId,
            comment: this.$refs.textEdit.content})
          this.$refs.textEditor.content = ""
        }).catch(err => {this.$f.malert()})
      })
    },
  },
}
</script>
