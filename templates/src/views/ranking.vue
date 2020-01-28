<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="desserts"
      :items-per-page="15"
      hide-default-footer
      class="elevation-1"
    >
      <template v-slot:item.id="{ item }">
        <a style="color:black;" @click="$router.push({path: '/profile/' + item.id})">{{item.id}}</a>
      </template>

      <template v-slot:no-data>등록된 유저가 없습니다</template>
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
      page: 1,
      headers: [
        { text: '랭킹', value: 'rank', divider: true, width: "10%" },
        { text: '아이디', value: 'id', divider: true, width: "15%" },
        { text: '상태 메시지', value: 'pr', divider: true, width: "65%" },
        { text: '맞은 문제', value: 'ac_count', divider: true, width: "10%" },
      ],
      desserts: [],
      data_num: -1,
    }),
    created () {
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
          var req = { page: this.page }
          return req
        })
      },
      async search() {
        await this.makeQuery()
        .then(re => {
          if (re === null) return
          this.$router.push({path:"/ranking", query: re})
        })
      },
      async sendQuery() {
        await this.makeQuery()
        .then(re => {
          if (re === null) return

          var req = this.$f.makeHeaderObject()
          req['params'] = re
          this.$axios.get('/api/ranking', req)
          .then(res => {
            if (res.data.ranklist) this.desserts = res.data.ranklist
            this.data_num = res.data.data_num
          })
        })
      },
    },
  }
</script>
