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
    datas: [
    {
      'subm_no':123,
      'id':'rdd6584',
      'prob_no':14,
      'result':1,
      'lang':0,
      'run_time':110,
      'memory':344,
      'codelen':957,
      'subm_time':'2020-01-07 03:35:20',
    },
    {
      'subm_no':12333,
      'id':'rdd653384',
      'prob_no':1224,
      'result':2,
      'lang':4,
      'run_time':2110,
      'memory':34244,
      'codelen':57,
      'subm_time':'2020-01-07 03:35:20',
    },
   ],
  }),
  created() {
    if (typeof this.$route.query.prob_no !== 'undefined') this.prob_no = this.$route.query.prob_no
    if (typeof this.$route.query.id !== 'undefined') this.id = this.$route.query.id
    if (typeof this.$route.query.lang !== 'undefined') this.lang = this.$store.state.lang[this.$route.query.lang]
    if (typeof this.$route.query.result !== 'undefined') this.result = this.$store.state.result[this.$route.query.result]
    if (typeof this.$route.query.page !== 'undefined') this.page = this.$route.query.page

    this.$axios.get('/api/status', {
      params: {
        prob_no: this.prob_no,
        id: this.id,
        lang: this.$store.state.langOrd[this.lang],
        result: this.$store.state.resultOrd[this.result],
        page: this.page,
      }}).then(res => {
        data_num = res.data.data_num,
        datas = res.data.datas
      }).catch(err => {
        // alert("문제 불러오기 오류")
      })
  }
}
</script>
