<template>
  <v-container>

    <v-card elevation="4">
      <v-row>
        <v-col cols="8">
          <v-text-field
            class="pt-3 px-6"
            label="제목"
            v-model="title"
            maxlength="50"
            counter="50"
            placeholder=" "
            outlined
          ></v-text-field>
        </v-col>
        <v-spacer></v-spacer>
        <v-btn @click="save()" class="mt-4 mr-6" color="success">업로드</v-btn>
      </v-row>
      <v-col class="pt-0" cols="8">
        <v-row>
          <v-text-field
            class="pl-3 pr-12 ml-3"
            label="문제번호"
            v-model="prob_no"
            placeholder=" "
            dense
            outlined
          ></v-text-field>
          <v-col class="pr-12 mt-0 pt-0" cols="6">
            <v-select
              v-if="$f.admin==2"
              :items="$store.state.category0"
              label="분류"
              dense
              outlined
              v-model="category"
            ></v-select>
            <v-select
              v-else
              :items="$store.state.category2"
              label="분류"
              dense
              outlined
              v-model="category"
            ></v-select>
          </v-col>
        </v-row>
      </v-col>
    </v-card>


    <textEditor ref="textEdit" place="내용을 입력해주세요."></textEditor>
    <codeEditor ref="codeEdit" readOnly="3"></codeEditor>
    <div class="pb-6"></div>
  </v-container>
</template>

<script>
import codeEditor from '../semiViews/codeEditor.vue'
import textEditor from '../semiViews/textEditor.vue'

export default{
  components: {
    codeEditor,
    textEditor,
  },
  data: () => ({
    title: "",
    prob_no: "",
    category: "질문",
  }),
  methods: {
    save() {
      this.$f.getUserValid()
      .then(re => {
        if (re == null) {
          this.$router.push({path: "/login"})
          return
        }
        var flag = confirm("글을 업로드하면 다시 수정하거나 삭제할 수 없습니다.\n정말 업로드하시겠습니까?")
        if (flag == false) return

        this.$axios.post('/api/board/new/post', {
          title: this.title,
          content: this.$refs.textEdit.content,
          code: this.$refs.codeEdit.code,
          id: this.$f.userId,
          prob_no: Number(this.$f.modifyString(this.prob_no)),
          category: this.$store.state.categoryOrd[this.category],
        }, this.$f.makeHeaderObject())
        .then(res => {
          this.$router.push('/board')
        }).catch(err => { this.$f.malert() })
      })
    },
  },
}
</script>
