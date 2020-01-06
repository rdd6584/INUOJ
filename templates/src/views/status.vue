<template>
  <v-app>
    <v-container>
      <v-card dense>
        <v-row justify="center" dense>
          <v-col cols="2">
            <v-text-field
              class="px-3"
              label="문제 번호"
              maxlength="4"
              v-model="prob_no"
            ></v-text-field>
          </v-col>
          <v-text-field
            class="px-3"
            label="아이디"
            maxlength="20"
            v-model="id"
          ></v-text-field>
          <v-col cols="3">
            <v-select
              :items="$store.state.lang"
              class="px-3 pt-2"
              label="언어"
              v-model="lang"
              dense
              outlined
            ></v-select>
          </v-col>
          <v-col cols="3">
            <v-select
              :items="$store.state.result"
              class="pr-3 pt-2"
              label="결과"
              v-model="result"
              dense
              outlined
            ></v-select>
          </v-col>
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
    page: "1",
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
      }}).then(res => {})
  }
}
</script>
