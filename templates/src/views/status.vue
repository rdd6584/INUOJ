<template>
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
          @click="search()"
          color="success"
          >검색</v-btn>
        </v-col>
      </v-row>
    </v-card>

  <v-data-table
    :headers="headers"
    :items="datas"
    :items-per-page="15"
    hide-default-footer
    class="mt-0 pt-0 elevation-2">

    <template v-slot:item.subm_no="{ item }">
      <a @click="$router.push({path:'/source/' + item.subm_no})">{{item.subm_no}}</a>
    </template>

    <template v-slot:item.prob_no="{ item }">
      <v-tooltip right>
        <template v-slot:activator="{ on }">
          <a v-on="on" @click="$router.push({path:'/problem/' + item.prob_no})">{{item.prob_no}}</a>
        </template>
        <span>{{item.prob_title}}</span>
      </v-tooltip>
    </template>

    <template v-slot:item.id="{ item }">
      <a @click="$router.push({path:'/profile/' + item.id})">{{item.id}}</a>
    </template>

    <template v-slot:item.result="{ item }">
      <font v-if="$store.state.resultOrd[item.result]==0" color="black">{{item.result}}</font>
      <font v-else-if="$store.state.resultOrd[item.result]==1" color="#00C853">{{item.result}}</font>
      <font v-else color="red">{{item.result}}</font>
    </template>

    <template v-slot:item.memory="{ item }">
      <font v-if="$store.state.resultOrd[item.result]==1">{{item.memory}}</font>
      <font v-else></font>
    </template>

    <template v-slot:item.run_time="{ item }">
      <font v-if="$store.state.resultOrd[item.result]==1">{{item.run_time}}</font>
      <font v-else></font>
    </template>

    </v-data-table>

    <v-pagination
      v-model="page"
      :length="pageLength"
      :page="page"
      total-visible="10"
      class="pt-1"
    ></v-pagination>
  </v-container>
</template>

<style scoped>
  a {color:black; text-decoration:none;}
  a:hover {color:black; text-decoration:underline;}
</style>

<script>
export default{
  data: () => ({
    prob_no: "",
    id: "",
    lang: "모든 언어",
    result: "모든 결과",
    page: 1,
    data_num: -1,
    headers: [
      { text: '제출 번호', value: 'subm_no', divider: true},
      { text: '아이디', value: 'id', divider: true},
      { text: '문제 번호', value: 'prob_no', divider: true },
      { text: '결과', value: 'result', divider: true },
      { text: '언어', value: 'lang', divider: true },
      { text: '메모리 (kb)', value: 'memory', divider: true },
      { text: '시간 (ms)', value: 'run_time', divider: true },
      { text: '코드 길이 (byte)', value: 'codelen', divider: true },
      { text: '제출한 시간', value: 'subm_time', divider: true },
    ],
    datas: [],
  }),
  created() {
    if (typeof this.$route.query.prob_no !== 'undefined') this.prob_no = this.$route.query.prob_no
    if (typeof this.$route.query.id !== 'undefined') this.id = this.$route.query.id
    if (typeof this.$route.query.lang !== 'undefined') this.lang = this.$store.state.lang[this.$route.query.lang]
    if (typeof this.$route.query.result !== 'undefined') this.result = this.$store.state.result[this.$route.query.result]
    if (typeof this.$route.query.page !== 'undefined') this.page = this.$route.query.page

    this.sendQuery()
  },
  computed: {
    pageLength() {
      var ret = parseInt(this.data_num / 15)
      if (ret === 0 || this.data_num % 15 != 0) ret++
      return ret
    },
  },
  watch: {
    page:function(val) {
      if (this.data_num != -1) this.search()
    },
  },
  methods: {
    modifyDatas() {
      for (var i of this.datas) {
        if (typeof i.result == 'number') {
          if (i.result === 0) i.result = "채점중"
          else i.result = this.$store.state.result[i.result]
        }
        if (typeof i.lang == 'number') i.lang = this.$store.state.lang[i.lang]
      }
    },

    async makeQuery() {
      return await this.$f.getUserValid().
      then(res => {
        if (res === null) {
          this.$router.push('/login')
          return null
        }
        var req = {
          prob_no: Number(this.$f.modifyString(this.prob_no)),
          id: this.id,
          lang: this.$store.state.langOrd[this.lang],
          result: this.$store.state.resultOrd[this.result],
          page: Number(this.$f.modifyString(this.page)),
        }
        return req
      })
    },
    async search() {
      await this.makeQuery()
      .then(re => {
        if (re === null) return
        this.$router.push({path:'/status', query:re})
      })
    },
    async sendQuery() {
      await this.makeQuery()
      .then(re => {
        if (re === null) return

        var req = this.$f.makeHeaderObject()
        req['params'] = re

        this.$axios.get('/api/status', req)
         .then(res => {
             this.data_num = res.data.data_num
             if (res.data.datas) this.datas = res.data.datas
             else this.datas = []
             this.modifyDatas()
           }).catch(err => {this.$f.malert()})
       })
     }
   }
 }
</script>
