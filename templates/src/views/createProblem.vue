<template>
  <v-app>
    <v-container>
        <v-btn @click="test">테스트</v-btn>
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
        :value="description[0]"
        place="수식은 Latex를 이용합니다."
      ></textEditor>

      <textEditor
        title="입력"
        :value="description[1]"
        place="입력 형식을 설명하세요."
      ></textEditor>

      <textEditor
        title="출력"
        :value="description[2]"
        place="출력 형식을 설명하세요."
      ></textEditor>
        <v-row class="pt-5">
          <v-col cols="6">
            <v-card elevation="4">
              <v-textarea
               outlined
               label="예제 입력"
               ></v-textarea>
            </v-card>
          </v-col>

          <v-col cols="6">
            <v-card elevation="4">
              <v-textarea
               outlined
               label="예제 출력"
               ></v-textarea>
            </v-card>
          </v-col>
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
import textEditor from "../semiViews/textEditor.vue"
  export default{
    components: {
      textEditor,
    },
    data: () => ({
       files: [],
       title: "",
       description: ["", "", ""],
       t_limit: 1000,
       m_limit: 512,
       samplein: [""],
       sampleout: [""],
     }),
     methods: {
       test() {
         console.log(this.t_limit)
       },
       upload() {
         var formData = new FormData()
         formData.append(this.files)

         this.$axios.post('/ttt',
           formData,
           { headers:
             {'Content-Type': 'multipart/form-data'}
           }).then(
            res => { console.log(res) })
          .catch(err => { alert("파일 전송 오류")
        })
       },
     },
  }
</script>
