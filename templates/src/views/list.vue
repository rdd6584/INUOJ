<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="desserts"
      :items-per-page="15"
      hide-default-footer
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat dark color="blue">
          <v-spacer></v-spacer>
          <v-col class="pr-0 pl-5" cols="2">
            <v-text-field
              v-model="title"
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

      <template v-slot:item.title="{ item }">
        <a v-if="item.result==0" style="color:black;" @click="$router.push({path:'/problem/' + item.prob_no})">{{item.title}}</a>
        <a v-else-if="item.result==1" style="color:#00C853;" @click="$router.push({path:'/problem/' + item.prob_no})">{{item.title}}</a>
        <a v-else style="color:red;" @click="$router.push({path:'/problem/' + item.prob_no})">{{item.title}}</a>
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
  export default {
    data: () => ({
      title: "",
      page: 1,
      headers: [
        { text: '문제 번호', value: 'prob_no', divider: true, width: "10%" },
        { text: '제목', value: 'title', divider: true, width: "70%" },
        { text: '맞은 사람', value: 'accept', divider: true, width: "10%" },
        { text: '시도한 사람', value: 'attempt', divider: true, width: "10%" },
      ],
      desserts: [],
      data_num: -1,
    }),
    created () {
      if (typeof this.$route.query.title !== 'undefined') this.title = this.$route.query.title
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
      async makeQuery() {
        return await this.$f.getUserValid()
        .then(re => {
          if (re === null) {
            this.$router.push('/login')
            return null
          }

          var req = {
            title: this.title,
            page: this.page,
            id: this.$f.userId
          }
          return req
        })
      },
      async search() {
        await this.makeQuery()
        .then(re => {
          if (re === null) return
          this.$router.push({path:"/list", query: re})
        })
      },
      async sendQuery() {
        await this.makeQuery()
        .then(re => {
          if (re === null) return

          var req = this.$f.makeHeaderObject()
          req['params'] = re
          this.$axios.get('/api/list', req)
          .then(res => {
            if (res.data.problems) this.desserts = res.data.problems
            else this.desserts = []
          })
        })
      },
    },
  }
</script>
