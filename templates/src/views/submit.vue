<template>
  <v-container>
    <codeEditor submit='true'></codeEditor>
  </v-container>
</template>


<script>
import codeEditor from '../semiViews/codeEditor.vue'
export default{
  components: {
    codeEditor,
  },
}
</script>
