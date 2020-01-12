<template>
  <div class="ma-0 pa-0">
    <v-row class="ma-0 pa-0">
      <v-btn large class="mt-4 ml-4" color="success" v-if="submit">제출</v-btn>
      <v-spacer></v-spacer>
      <v-col class="pa-2" cols="2">
        <v-card elevation="5">
          <v-select
            outlined
            v-model="choice"
            hide-details
            :items="$store.state.lang"
        ></v-select>
      </v-card>
      </v-col>
    </v-row>
    <v-card elevation="2">
      <codemirror
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
  props: ['submit'],
  components: {
     codemirror
   },
   data: () => ({
     code: "",
     myOption: {
       mode: "text/x-c++src",
       lineNumbers: true,
       indentUnit: 4,
       indentWithTabs: true,
       lineWrapping: true,
     },
     choice: "C++",
   }),
}
</script>
