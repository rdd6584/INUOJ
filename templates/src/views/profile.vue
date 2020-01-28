<template>
  <v-container>
    <v-card class="pb-1 mb-2" elevation="1">
      <v-row class="mb-0 my-2 px-6">
        <h1>{{userId}}</h1>
        <v-btn
          v-if="$f.userId==userId"
          class="mx-1 mt-2"
          router :to="{path:'/modify/' + $f.userId}"
          icon>
          <i class="fas fa-cog" style="color:grey; font-size:22px;"></i>
        </v-btn>
        <h4 class="px-6 pt-4">Rank : {{rank}}</h4>
      </v-row>
      <v-row v-if="changePr==true" class="mb-3">
        <v-text-field
          ref="onPr"
          class="pl-3 mt-0"
          maxlength="50"
          style="max-width: 850px"
          full-width
          flat
          dense
          @keyup.enter="savePr()"
          counter="50"
          v-model="pr">
        </v-text-field>
        <v-btn class="ml-0 mt-2" @click="savePr()" icon>저장</v-btn>
      </v-row>
      <v-row v-else class="pt-1 mb-3">
        <div class="pt-1 pl-6">{{pr}}</div>
        <v-btn
          v-if="$f.userId==userId"
          class="mt-0 ml-0"
          @click="changePr=true;"
          icon>
            <i class="fas fa-pencil-alt"></i>
          </v-btn>
      </v-row>
    </v-card>

      <v-hover v-slot:default="{ hover }">
        <v-card :elevation="hover ? 16 : 1" @click="getList(1)">
          <v-card-title>맞은 문제 {{ac_count}}개</v-card-title>
          <v-card-subtitle>이 유저가 맞춘 문제를 확인하세요!</v-card-subtitle>
        </v-card>
      </v-hover>

      <v-dialog v-model="acOn" inset>
        <v-sheet min-height="400px" class="text-center">
          <v-btn
            class="mt-0 mb-0"
            text
            color="error"
            @click="acOn=false"
          >close</v-btn>
          <v-card elevation="0">
            <v-card-text v-if="ac_count != 0" class="py-0 px-3">
              <a class="line px-3" v-for="(val, i) in numbers" @click="$router.push({path:'/problem/' + val})">
                <font color="black">{{val}}:</font><font color="#00C853">{{titles[i]}}</font>
              </a>
            </v-card-text>

            <v-card-text v-else class="py-0 px-3">
              문제가 존재하지 않습니다.
            </v-card-text>
          </v-card>
        </v-sheet>
      </v-dialog>

      <v-hover v-slot:default="{ hover }">
        <v-card class="mt-3" :elevation="hover ? 16 : 1" @click="getList(2)">
          <v-card-title>틀린 문제 {{wa_count}}개</v-card-title>
          <v-card-subtitle>이 유저가 시도했으나 맞추지 못한 문제를 확인하세요!</v-card-subtitle>
        </v-card>
      </v-hover>

      <v-dialog v-model="waOn" inset>
        <v-sheet min-height="400px" class="text-center">
          <v-btn
            class="mt-0 mb-0"
            text
            color="error"
            @click="waOn=false"
          >close</v-btn>
          <v-card elevation="0">
            <v-card-text v-if="wa_count != 0" class="py-0 px-3">
              <a class="line px-3" v-for="(val, i) in numbers" @click="$router.push({path:'/problem/' + val})">
                <font color="black">{{val}}:</font><font color="red">{{titles[i]}}</font>
              </a>
            </v-card-text>

            <v-card-text v-else class="py-0 px-3">
              문제가 존재하지 않습니다.
            </v-card-text>
          </v-card>
        </v-sheet>
      </v-dialog>

      <v-hover v-slot:default="{ hover }">
        <v-card router :to="{path:'/status?id=' + userId}" class="mt-3" :elevation="hover ? 16 : 1" @click="">
          <v-card-title>제출 횟수 {{all_count}}회</v-card-title>
          <v-card-subtitle>이 유저가 제출한 답안을 확인하세요!</v-card-subtitle>
        </v-card>
      </v-hover>

  </v-container>
</template>

<style scoped>
  a.line {text-decoration:none;}
  a.line:hover {text-decoration:underline;}
</style>

<script>
export default{
  data: () => ({
    userId: "",
    pr: "",
    rank: 0,
    ac_count: 0,
    wa_count: 0,
    all_count: 0,
    acOn: false,
    waOn: false,
    numbers: [],
    titles: [],
    changePr: false,
  }),
  created() {
    this.userId = this.$route.params.id
    this.$axios.get("/api/profile/" + this.userId, this.$f.makeHeaderObject())
    .then(res => {
      this.pr = res.data.pr
      this.ac_count = res.data.ac_count
      this.wa_count = res.data.wa_count
      this.all_count = res.data.all_count
      this.rank = res.data.rank
    }).catch(err => {
       if (err.response.status == 404) {
         this.$router.push('/wrongaccess')
         return
      }
       this.$f.malert()})
  },
  methods: {
    savePr() {
      this.$f.getUserValid()
      .then(re => {
        if (re == null) {
          this.$router.push("/login")
          return
        }
        this.$axios.post('/api/edit/pr', {
          id:this.userId,
          pr:this.pr,
        }, this.$f.makeHeaderObject())
        .then(res => { this.changePr=false })
        .catch(err => { this.$f.malert() })
      })
    },
    getList(result) {
      this.$f.getUserValid().then(
        re => {
          if (re == null) {
            this.$router.push({path: '/login'})
            return
          }
          this.$axios.get(`/api/user/problem?result=${result}&id=${this.userId}`, this.$f.makeHeaderObject())
          .then(res => {
            this.numbers = res.data.numbers
            this.titles = res.data.titles
          }).catch(err => {this.$f.malert()})

          if (result == 1) this.acOn = true
          else if (result == 2) this.waOn = true
      })
    },
  },
}
</script>
