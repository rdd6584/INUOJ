<template>
  <v-container>
    <codeEditor ref="code" readOnly="2" language="1">

    </codeEditor>
  </v-container>
</template>

<script>
import codeEditor from '../semiViews/codeEditor.vue'
  export default{
    components: {
      codeEditor,
    },
    data: () => ({
      code: "",
    }),
  }
</script>
