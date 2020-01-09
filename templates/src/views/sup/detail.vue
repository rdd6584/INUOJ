<template>
  <v-app>
    <v-container>
        <v-btn @click="test()">테스트</v-btn>
        <v-card elevation="4">
          <v-col cols="6">
            <v-text-field
              class="pt-3 px-3"
              label="문제 제목"
              :value="title"
              placeholder=" "
              outlined
            ></v-text-field>
          </v-col>
          <v-row>
            <v-col cols="3">
              <v-text-field
                class="px-3 ml-3"
                label="시간 제한"
                :value="t_limit"
                suffix="ms"
                dense
                outlined
              ></v-text-field>
            </v-col>
            <v-col cols="3">
              <v-text-field
                class="mr-3"
                label="메모리 제한"
                :value="m_limit"
                suffix="mb"
                dense
                outlined
              ></v-text-field>
            </v-col>
          </v-row>
        </v-card>
      <textEditor
        title="본문"
        ref="desc1"
        place="수식은 Latex를 이용합니다."
      ></textEditor>

      <textEditor
        title="입력"
        ref="desc2"
        place="입력 형식을 설명하세요."
      ></textEditor>

      <textEditor
        title="출력"
        ref="desc3"
        place="출력 형식을 설명하세요."
      ></textEditor>
        <v-row class="pt-5">
          <v-col cols="6">
            <v-card elevation="4">
              <v-textarea
               outlined
               label="예제 입력"
               v-model="samplein[0]"
               ></v-textarea>
            </v-card>
          </v-col>

          <v-col cols="6">
            <v-card elevation="4">
              <v-textarea
               outlined
               label="예제 출력"
               v-model="sampleout[0]"
              ></v-textarea>
            </v-card>
          </v-col>
        </v-row>
        <v-row justify="center">
            <v-btn @click="save()" class="mb-6" color="success">저장</v-btn>
        </v-row>

        <div class="pt-3"></div>
        <v-card min-height="100" elevation="4">
          <v-row class="pt-5 mx-5">
            <v-file-input
              v-model="files"
              color="deep-purple accent-4"
              counter
              label="File input"
              multiple
              placeholder="Select your files"
              prepend-icon="mdi-paperclip"
              outlined
              :show-size="1000"
            >
              <template v-slot:selection="{ index, text }">
                <v-chip
                  v-if="index < 2"
                  color="deep-purple accent-4"
                  dark
                  label
                  small
                >
                  {{ text }}
                </v-chip>
                <span
                  v-else-if="index === 2"
                  class="overline grey--text text--darken-3 mx-2"
                >
                  +{{ files.length - 2 }} File(s)
                </span>
              </template>
              </v-file-input>
              <v-btn class="ma-5" color="success" @click="upload()">파일 업로드</v-btn>
            </v-row>
          </v-card>
      </v-container>
  </v-app>
</template>

<script>
import textEditor from "../../semiViews/textEditor.vue"
  export default {
    components: {
      textEditor,
    },
    data: () => ({
       files: [],
       ori_no: 0,
       stat: 0,
       owner: "",
       title: "",
       t_limit: "1000",
       m_limit: "512",
       samplein: [""],
       sampleout: [""],
     }),
     async created() {

     },
     methods: {
       test() {
         console.log(this.samplein[0])
       },
       async upload() {
         await this.$f.getUserValid()
         .then(
           re => {
             if (re===null) {
               this.$router.push({path:'/login'})
               return
             }

             if (this.files.length % 2 == 1) { alert("입출력 파일 쌍이 맞지 않습니다"); return}
             var ord = [], mavi = -1, a = "", b = "", c = "", d = "", fl = 0, ff = 0
             var visit = []
             for (var i = 0; i < this.files.length; i++) visit[i] = 0
             for (var i = 0; i < this.files.length; i++) {
               mavi = -1
               for (var j in this.files) {
                 if (visit[j] == 0 && (mavi == -1 || this.files[mavi].name > this.files[j].name))
                  mavi = j
               }

               visit[mavi] = 1
               ord[i] = mavi
             }
             for (var i = 0; i < ord.length; i += 2) {
               a = "", b = "", c = "", d = "", fl = 0, ff = 0

               for (var j of this.files[ord[i]].name) {
                 if (j == '.') fl++
                 else if (fl === 0) a += j
                 else b += j
               }
               for (var j of this.files[ord[i + 1]].name) {
                 if (j == '.') ff++
                 else if (ff === 0) c += j
                 else d += j
               }

               if (fl == 1 && ff == 1 && a == c && b == "in" && d == "out");
               else { alert("입출력 파일 쌍이 맞지 않습니다"); return }
             }

             var formData = new FormData()
             for (var i = 0; i < this.files.length; i += 2) {
               formData.append('input', this.files[ord[i]])
               formData.append('output', this.files[ord[i + 1]])
             }

            var token = this.$f.makeHeaderObject()
             alert("완료 메시지가 나올 때까지 창을 떠나지 마세요.")
             this.$axios.post('/api/problem/upload/data', formData, token)
              .then( res => { "문제 업로드가 완료되었습니다." })
              .catch( err => { alert("파일 전송 오류")
            })
           })
       },
       async save() {
         await this.$f.getUserValid()
         .then(
           re => {
             if (re===null) {
               this.$router.push({path:'/login'})
               return
             }

             if (this.$f.isOnlyNum(this.t_limit) || this.t_limit[0] === '0') {
                 alert("시간 제한을 확인하세요")
                 return
             }
             if (this.$f.isOnlyNum(this.m_limit) || this.m_limit[0] === '0') {
                alert("메모리 제한을 확인하세요")
                return
             }
             var formdata = new FormData()
             formdata.append('ori_no', this.ori_no)
             formdata.append('owner', this.owner)
             formdata.append('stat', this.stat)
             formdata.append('title', this.title)
             formdata.append('t_limit', this.t_limit)
             formdata.append('m_limit', this.m_limit)
             formdata.append('description', this.$refs.desc1.content)
             formdata.append('description', this.$refs.desc2.content)
             formdata.append('description', this.$refs.desc3.content)
             for (var i of this.samplein) formdata.append('samplein', i)
             for (var i of this.sampleout) formdata.append('sampleout', i)

             var token = this.$f.makeHeaderObject()
             this.$axios.post('/api/problem/upload/desc', formdata, token)
               .then(res => {alert('저장이 완료되었습니다.')})
               .catch(err => {this.$f.malert()})
           })
       },
     },
  }
</script>
