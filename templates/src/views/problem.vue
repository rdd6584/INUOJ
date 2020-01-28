<template>
  <v-container>
    <v-card class="px-5" elevation="4">
      <v-row class="py-3">
        <v-btn color="success" class="mx-3" router :to="{path: '/submit/' + prob_no}">제출</v-btn>
        <v-spacer></v-spacer>
        <v-btn dark class="mx-3" router :to="{path: '/status?prob_no=' + prob_no}">채점현황</v-btn>
        <v-btn dark class="mx-3" router :to="{path: '/board?&prob_no=' + prob_no}">게시판</v-btn>
      </v-row>
      <v-row>
        <h2>{{title}}</h2>
      </v-row>
      <v-simple-table dense>
        <template v-slot:default>
          <thead>
            <tr>
              <th class="text-left">시간 제한</th>
              <th class="text-left">메모리 제한</th>
              <th class="text-left">맞은 사람</th>
              <th class="text-left">시도한 사람</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{{ t_limit }} ms</td>
              <td>{{ m_limit }} MB</td>
              <td>{{ accept }} 명</td>
              <td>{{ attempt }} 명</td>
            </tr>
          </tbody>
        </template>
      </v-simple-table>

      <v-row class="pt-3">
        <span class="viewHtml" v-html="description[0]"></span>
      </v-row>

      <v-row>
        <h3 style="color:rgb(0, 98, 184);" class="mt-1">입력</h3>
      </v-row>
      <v-row>
        <span class="viewHtml" v-html="description[1]"></span>
      </v-row>

      <v-row>
        <h3 style="color:rgb(0, 98, 184);" class="mt-1">출력</h3>
      </v-row>
      <v-row>
        <span class="viewHtml" v-html="description[2]"></span>
      </v-row>

      <v-divider class="mt-6"></v-divider>
      <v-row class="px-3" v-for="(val, i) in samplein">
        <v-col cols="6" class="pl-0 ml-0">
          <v-row class="px-4 pb-1">
            <h4 style="color:rgb(0, 98, 184);" class="mt-1">예제{{i + 1}} 입력</h4>
            <v-spacer></v-spacer>
            <v-btn dark small v-clipboard="samplein[i]">복사</v-btn>
          </v-row>
          <v-card class="py-0 my-0" color="rgb(245,245,245)" elevation="0">
            <v-textarea
             outlined
             dense
             auto-grow
             hide-details
             readonly
             v-model="samplein[i]"
             ></v-textarea>
          </v-card>
        </v-col>

        <v-col cols="6" class="mr-0 pr-0">
          <v-row class="px-4 pb-1">
            <h4 style="color:rgb(0, 98, 184);" class="mt-1">예제{{i + 1}} 출력</h4>
            <v-spacer></v-spacer>
            <v-btn dark small v-clipboard="sampleout[i]">복사</v-btn>
          </v-row>
          <v-card color="rgb(245,245,245)" elevation="0">
            <v-textarea
             outlined
             dense
             auto-grow
             hide-details
             readonly
             v-model="sampleout[i]"
            ></v-textarea>
          </v-card>
        </v-col>
      </v-row>
      <div class="py-2"></div>
    </v-card>
  </v-container>
</template>

<style scoped>
.viewHtml >>> img { max-width: 100%; max-height: 70%;}
.viewHtml >>> .ql-align-left { text-align: left; }
.viewHtml >>> .ql-align-center { text-align: center; }
.viewHtml >>> .ql-align-right { text-align: right; }
</style>

<script>
  export default{
    data: () => ({
       prob_no: 0,
       stat: 0,
       result: 0,
       owner: "",
       title: "제목입니다.",
       t_limit: 1000,
       m_limit: 512,
       samplein: [""],
       sampleout: [""],
       description: ["본문입니다.", "입력입니다.", "출력입니다."],
       accept: 0,
       attempt: 0,
     }),
     async created() {
       this.prob_no = Number(this.$route.params.num)
       await this.$axios.get(`/api/problem/detail?prob_no=${this.prob_no}&id=${this.$f.userId}`,
         this.$f.makeHeaderObject())
       .then(res => {
         this.t_limit = res.data.t_limit
         this.m_limit = res.data.m_limit
         this.title = res.data.title
         this.owner = res.data.owner
         this.stat = res.data.stat
         this.result = res.data.result
         this.accept = res.data.accept
         this.attempt = res.data.attempt

         this.description = res.data.description

         if (res.data.samplein) this.samplein = res.data.samplein
         if (res.data.sampleout) this.sampleout = res.data.sampleout
       })
       .catch(err => {
         if (err.response.status == 404) {
           this.$router.push('/wrongaccess')
           return
        }
         this.$f.malert()})
     },
  }
</script>
