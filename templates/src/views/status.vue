<template>
  <v-app>
    <v-container>
      <v-card dense>
        <v-row justify="center" dense>
          <v-col cols="1">
            <v-text-field
              class="px-3"
              label="문제"
              maxlength="4"
              v-model="prob_no"
            ></v-text-field>
          </v-col>
          <v-col cols="2">
            <v-text-field
              class="pr-3"
              label="아이디"
              maxlength="20"
              v-model="id"
            ></v-text-field>
          </v-col>
          <v-col cols="3">
            <v-select
              :items="$store.state.lang"
              class="pr-3 pt-3"
              label="언어"
              v-model="lang"
              dense
              outlined
            ></v-select>
          </v-col>
          <v-col cols="3">
            <v-select
              :items="$store.state.result"
              class="pr-2 pt-3"
              label="결과"
              v-model="result"
              dense
              outlined
            ></v-select>
          </v-col>
          <v-col class="pt-4 pl-6" cols="1">
            <v-btn
            @click="sendQuery()"
            color="success"
            >검색</v-btn>
          </v-col>
        </v-row>
      </v-card>
      <v-card class="mt-3">
        <v-row justify="center" dense>
          <v-data-table
            :headers="headers"
            :items="datas"
            :items-per-page="15"
            hide-default-footer
            class="elevation-1"
          ></v-data-table>
        </v-row>
      </v-card>
    </v-container>
  </v-app>
</template>

<script>
export default{
  data: () => ({
    prob_no: "",
    id: "",
    lang: "모든 언어",
    result: "모든 결과",
    page: 1,
    data_num: 0,
    headers: [
      {
        text: '채점 번호',
        align: 'left',
        divider: true,
        value: 'subm_no',
      },
      { text: '아이디', value: 'id', divider: true, sortable: false },
      { text: '문제 번호', value: 'prob_no', divider: true },
      { text: '결과', value: 'result', divider: true },
      { text: '언어', value: 'lang', divider: true },
      { text: '메모리', value: 'memory', divider: true },
      { text: '시간', value: 'run_time', divider: true },
      { text: '코드 길이', value: 'codelen', divider: true },
      { text: '제출한 시간', value: 'subm_time', divider: true },
    ],
    datas: [],
  }),
  async created() {
    if (typeof this.$route.query.prob_no !== 'undefined') this.prob_no = this.$route.query.prob_no
    if (typeof this.$route.query.id !== 'undefined') this.id = this.$route.query.id
    if (typeof this.$route.query.lang !== 'undefined') this.lang = this.$store.state.lang[this.$route.query.lang]
    if (typeof this.$route.query.result !== 'undefined') this.result = this.$store.state.result[this.$route.query.result]
    if (typeof this.$route.query.page !== 'undefined') this.page = this.$route.query.page

    await this.sendQuery()
  },
  methods: {
    modifyDatas() {
      for (var i of this.datas) {
        if (typeof i.result == 'number') i.result = this.$store.state.result[i.result]
        if (typeof i.lang == 'number') i.lang = this.$store.state.lang[i.lang]
      }
    },
    modifyProbNo() {
      var ret = ""
      for (var i of this.prob_no)
        if (this.prob_no[i] >= '0' && this.prob_no[i] <= '9')
          ret += this.prob_no[i]

      if (ret.length == 0) ret = "0"
      return ret
    },

    async sendQuery() {
      await this.$f.getUserValid().then(
        res => {
          if (res == null)
            this.$router.push({path : '/'})
        })
        .catch(err => {
          this.$f.malert()
          this.$router.push({path : '/'})
        })
      var req = await this.$f.makeHeaderObject()
      req['params'] = {
        prob_no: this.modifyProbNo(),
        id: this.id,
        lang: this.$store.state.langOrd[this.lang],
        result: this.$store.state.resultOrd[this.result],
        page: this.page,
      }
     await this.$axios.get('/api/status', req)
      .then(res => {
          this.data_num = res.data.data_num
          if (res.data.datas) this.datas = res.data.datas
          else this.datas = []
          this.modifyDatas()
        }).catch(err => {
          // alert("문제 불러오기 오류")
        })
    },
  },
}
</script>
