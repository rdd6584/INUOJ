<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="desserts"
      hide-default-footer
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat dark color="blue">
          <v-spacer></v-spacer>
          <v-col class="pr-0 pl-5" cols="2">
            <v-text-field
              v-model="searchM"
              background-color="rgb(250, 252, 255)"
              light
              class="pt-6 mt-1"
              outlined
              dense
              placeholder="검색">
            </v-text-field>
          </v-col>
          <v-btn @click="search()" icon>
            <i class="fas fa-search"></i>
          </v-btn>
        </v-toolbar>
      </template>

      <template v-slot:item="{ index }">
        <div v-if="index < notice_num" style="background-color:green;"></div>
      </template>

      <template v-slot:item.title="{ item }">
        <a style="color:black;" @click="$router.push({path:'/post/' + item.post_no})">{{item.title}}</a>
      </template>

      <template v-slot:item.category="{ item }">
        <div v-if="item.prob_no == 0">
          {{item.category}}
        </div>
        <div v-else>
          <a style="color:black;" @click="$router.push({path:'/problem/' + item.prob_no})">{{item.prob_no}}번 {{item.category}}</a>
        </div>
      </template>

      <template v-slot:item.id="{ item }">
        <a style="color:black;" @click="$router.push({path:'/profile/' + item.id})">{{item.id}}</a>
      </template>

      <template v-slot:no-data>등록된 문제가 없습니다</template>
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
// 마우스 올리면 문제 나오는 거 해야되고

export default {
  data: () => ({
    page: 1,
    select: "제목", // 검색을 뭐로 할건지?
    searchM: "",
    category: "전체",
    headers: [
      { text: '제목', value: 'title', divider: true, width: "50%" },  // 옆에 댓글 수
      { text: '분류', value: 'category', divider: true, width: "15%" },
      { text: '작성자', value: 'id', divider: true, width: "15%" },
      { text: '작성일', value: 'post_time', divider: true, width: "20%" },
    ],
    desserts: [],
    notice: [],
    normal: [],
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
          if (res.data.normal) this.desserts.concat(res.data.notice)

          for (var i of this.desserts)
            i.category = this.$store.state.category2[i.category]
        })
      })
    },
  },
}
</script>
