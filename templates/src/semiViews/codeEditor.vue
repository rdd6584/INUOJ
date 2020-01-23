<template>
  <div class="ma-0">
    <v-row v-if="readOnly!=2 && readOnly!=3" class="ma-0 pa-0">
      <v-btn @click="$emit('send')" large class="mt-4 ml-4" color="success" v-if="readOnly!=1">제출</v-btn>
      <v-spacer></v-spacer>
      <v-col class="pa-2" cols="2">
        <v-card elevation="5">
          <v-select
            outlined
            v-model="choice"
            hide-details
            :items="$store.state.slang"
        ></v-select>
      </v-card>
      </v-col>
    </v-row>
    <v-card elevation="2">
      <codemirror
        ref="mirror"
        v-model="code"
        :options="myOption"
      ></codemirror>
    </v-card>
  </div>
</template>

<style>
.CodeMirror {
  border: 1px solid #eee;
  height: auto;
  min-height: 500px;
}
.CodeMirror-scroll {
  min-height: 500px;
}
</style>


<script>
import { codemirror } from 'vue-codemirror-lite'
require('codemirror/mode/go/go')
require('codemirror/mode/python/python')
require('codemirror/mode/clike/clike')

export default{
  props: {
    readOnly: { default:1 },
  },
  created() {
    if (this.readOnly == 2) this.myOption.readOnly = true
  },
  components: {
     codemirror
   },
   data: () => ({
     code: "",
     myOption: {
       readOnly: false,
       mode: "text/x-c++src",
       lineNumbers: true,
       indentUnit: 4,
       indentWithTabs: true,
       lineWrapping: true,
     },
     choice: "C++",
   }),
   watch: {
     choice(val) {
       this.myOption.mode = this.$store.state.langCodeMirror[this.$store.state.slang.indexOf(val)];
       this.$refs.mirror.editor.setOption('mode', this.myOption.mode)
     },
   },
   methods: {

   },
}
</script>
