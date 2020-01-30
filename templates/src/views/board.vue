<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="desserts"
      :items-per-page="15 + notice_num"
      hide-default-footer
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat dark color="blue">
          <v-btn @click="if(category!='전체') {category='전체'; search();}" text>전체</v-btn>
          <v-btn @click="if(category!='공지') {category='공지'; search();}" text>공지</v-btn>
          <v-btn @click="if(category!='질문') {category='질문'; search();}" text>질문</v-btn>
          <v-btn @click="if(category!='자유') {category='자유'; search();}" text>자유</v-btn>
          <v-btn @click="$router.push({path:'/writepost'})" text large>글 작성</v-btn>

          <v-spacer></v-spacer>
          <v-col class="pr-0" cols="2">
            <v-select
              class="pt-6 mt-1"
              :items="selectList"
              label="분류"
              dense
              solo
              light
              outlined
              v-model="select"
            ></v-select>
          </v-col>

          <v-col class="ml-3 pr-0 pl-0" cols="3">
            <v-text-field
              v-model="searchM"
              background-color="rgb(250, 252, 255)"
              light
              class="pt-6 mt-1"
              outlined
              dense
              @keyup.enter="search()"
              placeholder="검색">
            </v-text-field>
          </v-col>
          <v-btn @click="search()" icon>
            <i class="fas fa-search"></i>
          </v-btn>
        </v-toolbar>
      </template>

      <template v-slot:item.title="{ item }">
        <p class="pa-0 ma-0">
          <i v-if="desserts.indexOf(item) < notice_num" class="pr-2 fas fa-star"></i>
          <a style="color:black;" @click="$router.push({path:'/post/' + item.post_no})">{{item.title}}</a>
          <span class="pl-1" v-if="item.cmt_no">({{item.cmt_no}})</span>
        </p>
      </template>

      <template v-slot:item.category="{ item }">
        <div v-if="item.prob_no==0 || item.prob_title==''">
          {{item.category}}
        </div>
        <v-tooltip right v-else>
          <template v-slot:activator="{ on }">
            <a v-if="item.result==0" v-on="on" style="color:black;" @click="$router.push({path:'/problem/' + item.prob_no})">{{item.prob_no}}번 </a>
            <a v-else-if="item.result==1" v-on="on" style="color:#00C853;" @click="$router.push({path:'/problem/' + item.prob_no})">{{item.prob_no}}번 </a>
            <a v-else v-on="on" style="color:red;" @click="$router.push({path:'/problem/' + item.prob_no})">{{item.prob_no}}번 </a>
            {{item.category}}
            <div class="pa-0 ma-0" v-if="$f.admin==2">
              <a class="pl-2" @click="modifyNotice(item.post_no, 1)">올림</a>
              <a class="pl-2" @click="modifyNotice(item.post_no, 2)">내림</a>
            </div>
          </template>
          <span>{{item.prob_title}}</span>
        </v-tooltip>
      </template>

      <template v-slot:item.id="{ item }">
        <a style="color:black;" @click="$router.push({path:'/profile/' + item.id})">{{item.id}}</a>
      </template>

      <template v-slot:no-data>등록된 글이 없습니다</template>
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
  a {text-decoration:none;}
  a:hover {text-decoration:underline;}
</style>

<script>
// 검색 인터페이스 구축하고.

export default {
  data: () => ({
    page: 1,
    select: "제목",
    selectList: ["제목", "문제 번호", "작성자"],
    searchM: "",
    category: "전체",
    headers: [
      { text: '제목', value: 'title', sortable:false, divider: true, width: "50%" },  // 옆에 댓글 수
      { text: '분류', value: 'category', sortable:false, divider: true, width: "15%" },
      { text: '작성자', value: 'id', sortable:false, divider: true, width: "15%" },
      { text: '작성일', value: 'post_time', sortable:false, divider: true, width: "20%" },
    ],
    desserts: [],
    data_num: -1,
    notice_num: 0,
  }),
  created () {
    if (typeof this.$route.query.page !== 'undefined') this.page = this.$route.query.page
    if (typeof this.$route.query.title !== 'undefined') {
      this.searchM = this.$route.query.title
      this.select = "제목"
    }
    else if (typeof this.$route.query.prob_no !== 'undefined') {
      this.searchM = this.$route.query.prob_no
      this.select = "문제 번호"
    }
    else if (typeof this.$route.query.author !== 'undefined') {
      this.searchM = this.$route.query.author
      this.select = "작성자"
    }
    if (typeof this.$route.query.category !== 'undefined') this.category = this.$store.state.category[this.$route.query.category]
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
    modifyNotice(no, tice) {
      this.$f.getUserValid()
      .then(re => {
        if (re === null) {
          this.$router.push('/login')
          return null
        }
        this.$axios.post('/api/admin/update/notice', {
            post_no : no,
            notice : tice,
          }, this.$f.makeHeaderObject()
        ).then(res => {
          alert("정상적으로 변경되었습니다.")
        }).catch(err => {this.$f.malert()})
      })
    },
    async makeQuery() {
      return await this.$f.getUserValid()
      .then(re => {
        if (re === null) {
          this.$router.push('/login')
          return null
        }

        var req = {
          page: this.page,
          id: this.$f.userId,
          category: this.$store.state.categoryOrd[this.category],
        }
        if (this.select == "제목") req['title'] = this.searchM
        else if (this.select == "작성자") req['author'] = this.searchM
        else if (this.select == "문제 번호") req['prob_no'] = this.searchM
        return req
      })
    },
    async search() {
      await this.makeQuery()
      .then(re => {
        if (re === null) return
        this.$router.push({path:"/board", query: re})
      })
    },
    async sendQuery() {
      await this.makeQuery()
      .then(re => {
        if (re === null) return

        var req = this.$f.makeHeaderObject()
        req['params'] = re
        this.$axios.get('/api/board/list', req)
        .then(res => {
          // 데이터 처리하는 거 하기.
          if (res.data.notice) {
            this.desserts = res.data.notice
            this.notice_num = this.desserts.length
          }

          if (res.data.normal) this.desserts = this.desserts.concat(res.data.normal)
          for (var i of this.desserts)
            i.category = this.$store.state.category[i.category]
          this.data_num = res.data.data_num
        })
      })
    },
  },
}
</script>
