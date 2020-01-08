<template>
  <div class="py-4">
    <h3 style="color:rgb(0, 98, 184);" class="ml-2 mb-1">{{title}}</h3>
    <v-card elevation="4">
      <vue-editor
        :placeholder="place"
        v-model="content"
        :editorToolbar="customToolbar">
      </vue-editor>
    </v-card>
  </div>
</template>

<script>
import { VueEditor } from "vue2-editor"

export default {
  props: ['place', 'title'],
  components: {
    VueEditor
  },

  data: () => ({
    content: "",
    customToolbar: [
      ["bold", "italic", "underline"],
      [{ 'script': 'sub'}, { 'script': 'super' }, 'formula'],
      [{'align': ''}, {'align': 'center'}, {'align': 'right'}, {'align': 'justify'}],
      [{ list: "ordered" }, { list: "bullet" }], ['link', "image", "code-block"],
    ],
  }),
};
</script>
