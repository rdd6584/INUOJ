<template>
  <v-container>
    <v-textarea
      v-if="compile_msg!=0"
      class="mb-2"
      background-color="rgb(233, 233, 245)"
      readonly
      auto-grow
      hide-details
      solo
      flat
      style="border-radius:0;"
      v-model="compile_msg"
    ></v-textarea>

    <codeEditor ref="codeView" readOnly="2"/>

    <v-data-table
      :headers="headers"
      :items="datas"
      :items-per-page="1"
      hide-default-footer
      class="elevation-2">

      <template v-slot:item.result="{ item }">
        <font v-if="item.result==0" color="black">{{$store.state.result[item.result]}}</font>
        <font v-else-if="item.result==1" color="#00C853">{{$store.state.result[item.result]}}</font>
        <font v-else color="red">{{$store.state.result[item.result]}}</font>
      </template>

      <template v-slot:item.lang="{ item }">
        {{$store.state.lang[item.lang]}}
      </template>

      <template v-slot:item.memory="{ item }">
        <font v-if="item.result==1">{{item.memory}}</font>
        <font v-else></font>
      </template>

      <template v-slot:item.run_time="{ item }">
        <font v-if="item.result==1">{{item.run_time}}</font>
        <font v-else></font>
      </template>

    </v-data-table>
  </v-container>
</template>

<script>
import codeEditor from '../semiViews/codeEditor.vue'
  export default{
    components: {
      codeEditor,
    },
    created() {
      this.$axios.get('/api/source/' + this.$route.params.subm_no, this.$f.makeHeaderObject())
      .then(res => {
        this.$refs.codeView.code = res.data.code
        this.compile_msg = res.data.compile_msg
        this.$refs.codeView.choice = this.$store.state.lang[res.data.submit.lang]
        this.datas.push(res.data.submit)
      }).catch(err => {this.$f.malert()})
    },
    data: () => ({
      compile_msg: "",
      datas: [],
      headers: [
        { text: '제출 번호', value: 'subm_no', sortable: false,  divider: true},
        { text: '아이디', value: 'id', sortable: false,  divider: true},
        { text: '문제 번호', value: 'prob_no', sortable: false,  divider: true },
        { text: '결과', value: 'result', sortable: false,  divider: true },
        { text: '언어', value: 'lang', sortable: false,  divider: true },
        { text: '메모리 (kb)', value: 'memory', sortable: false,  divider: true },
        { text: '시간 (ms)', value: 'run_time', sortable: false,  divider: true },
        { text: '코드 길이 (byte)', value: 'codelen', sortable: false,  divider: true },
        { text: '제출한 시간', value: 'subm_time', sortable: false,  divider: true },
      ],
    }),
  }
</script>
