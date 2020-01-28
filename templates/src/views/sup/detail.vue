<template>
  <v-container>
      <v-card elevation="4">
        <v-col cols="6">
          <v-text-field
            class="pt-3 px-3"
            label="문제 제목"
            v-model="title"
            hide-details
            placeholder=" "
            outlined
          ></v-text-field>
        </v-col>
        <v-row>
          <v-col cols="3">
            <v-text-field
              class="px-3 ml-3"
              label="시간 제한"
              v-model="t_limit"
              suffix="ms"
              dense
              outlined
            ></v-text-field>
          </v-col>
          <v-col cols="3">
            <v-text-field
              class="mr-3"
              label="메모리 제한"
              v-model="m_limit"
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

      <v-row v-for="(val, i) in samplein" class="pt-5">
        <v-col cols="6" >
          <v-card elevation="4">
            <v-textarea
             outlined
             hide-details
             :label="'예제' + (i + 1) + ' 입력'"
             v-model="samplein[i]"
             ></v-textarea>
          </v-card>
        </v-col>

        <v-col cols="6">
          <v-card elevation="4">
            <v-textarea
             outlined
             hide-details
             :label="'예제' + (i + 1) + ' 출력'"
             v-model="sampleout[i]"
            ></v-textarea>
          </v-card>
        </v-col>
      </v-row>

      <v-row justify="center">
        <v-btn @click="samplein.push(''); sampleout.push('')" class="mx-4 mb-6" color="success">예제 추가</v-btn>
        <v-btn @click="save()" class="mb-6" color="success">저장</v-btn>
      </v-row>

      <div class="pt-3"></div>
      <v-card min-height="100" elevation="4">
        <v-card-title>채점 데이터</v-card-title>
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

        <v-card v-if="datas.length != 0" class="mt-1 px-3">
          <v-row>
            <v-col class="py-0" cols="2" v-for="i in datas">
              <v-checkbox hide-details class="px-1 py-0" v-model="selected" :label="i" :value="i"></v-checkbox>
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-btn class="mx-3" color="success" @click="selectAll()">전체선택</v-btn>
            <v-btn class="mx-3" color="success" @click="del()">삭제</v-btn>
          </v-row>
        </v-card>

        <v-card v-else class="mt-1">
          <v-row justify="center">
            <v-card-subtitle>Data set is empty</v-card-subtitle>
          </v-row>
        </v-card>
    </v-container>
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
       datas: [],
       selected: [],
     }),
     mounted() {
       this.ori_no = Number(this.$route.params.ori_no)
       this.$axios.get("/api/bdmin/detail/" + this.ori_no, this.$f.makeHeaderObject())
       .then(res => {
         this.t_limit = res.data.t_limit + ""
         this.m_limit = res.data.m_limit + ""
         this.title = res.data.title
         this.owner = res.data.owner
         this.stat = res.data.stat

         this.$refs.desc1.content = res.data.description[0]
         this.$refs.desc2.content = res.data.description[1]
         this.$refs.desc3.content = res.data.description[2]

         if (res.data.samplein) this.samplein = res.data.samplein
         if (res.data.sampleout) this.sampleout = res.data.sampleout
         if (res.data.datas) this.datas = res.data.datas
       })
       .catch(err => {
         if (err.response.status == 404) {
           this.$router.push('/wrongaccess')
           return
         }
         this.$f.malert()})
     },
     methods: {
       selectAll() {
         if (this.selected.length == this.datas.length)
           this.selected.splice(0, this.selected.length)
         else this.selected = this.datas.slice()
       },
       del() {
          this.$f.getUserValid()
          .then(re => {
            if (re === null) {
              this.$router.push("/login")
              return
            }
            this.$axios.post("/api/bdmin/discard/data", {
              ori_no : this.ori_no,
              files : this.selected,
            }, this.$f.makeHeaderObject())
            .then(res => {
              for (var i of this.selected)
                this.datas.splice(this.datas.indexOf(i), 1)
              this.selected.splice(0, this.selected.length)
            })
            .catch(err => this.$f.malert())
          })
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

             var count = {}
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
               if (typeof count[a] !== 'undefined') { alert("중복된 파일명이 존재합니다."); return }
               else count[a] = 1
             }

             for (var i of this.datas)
              if (typeof count[i] !== 'undefined') { alert("중복된 파일명이 이미 존재합니다."); return }

             var formData = new FormData()
             formData.append('ori_no', this.ori_no)
             for (var i = 0; i < this.files.length; i += 2) {
               formData.append('input', this.files[ord[i]])
               formData.append('output', this.files[ord[i + 1]])
             }

            var token = this.$f.makeHeaderObject()
             alert("완료 메시지가 나올 때까지 창을 떠나지 마세요.")
             this.$axios.post('/api/bdmin/upload/data', formData, token)
              .then(
                res => { alert("데이터가 업로드 되었습니다.")
                for (var i = 0; i < this.files.length; i += 2) {
                  var len = this.files[ord[i]].name.length
                  this.datas.push(this.files[ord[i]].name.substr(0, len - 3))
                }
                this.files = []
              })
              .catch( err => { this.$f.malert() })
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

             if (this.$f.isOnlyNum(this.t_limit) === false || this.t_limit[0] === '0') {
                 alert("시간 제한을 확인하세요")
                 return
             }
             if (this.$f.isOnlyNum(this.m_limit) === false || this.m_limit[0] === '0') {
                alert("메모리 제한을 확인하세요")
                return
             }
             var formdata = new FormData()
             formdata.append('ori_no', this.ori_no)
             formdata.append('owner', this.owner)
             formdata.append('title', this.title)
             formdata.append('t_limit', this.t_limit)
             formdata.append('m_limit', this.m_limit)
             formdata.append('description', this.$refs.desc1.content)
             formdata.append('description', this.$refs.desc2.content)
             formdata.append('description', this.$refs.desc3.content)

             for (var i in this.samplein) {
               if (this.samplein[i] != "" && this.sampleout[i] != "") {
                 formdata.append('samplein', this.samplein[i])
                 formdata.append('sampleout', this.sampleout[i])
               }
             }

             var token = this.$f.makeHeaderObject()
             this.$axios.post('/api/bdmin/upload/desc', formdata, token)
               .then(res => {alert('저장이 완료되었습니다.')})
               .catch(err => {this.$f.malert()})
           })
       },
     },
  }
</script>
