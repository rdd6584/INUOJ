<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="desserts"
      :items-per-page="15"
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat dark color="gray">
          <v-spacer></v-spacer>
          <v-col cols="2">
            <v-text-field class="pt-6 mt-1" outlined dense placeholder="검색"></v-text-field>
          </v-col>
          <v-btn color="blue">icon</v-btn>
        </v-toolbar>
      </template>

      <template v-slot:item.action="{ item }">
        <v-icon
          small
          class="mr-2"
          @click="toDetail(item.ori_no)"
        >detail
        </v-icon>
        <v-icon
          small
          class="mr-2"
          @click="editItem(item)"
        >status
        </v-icon>
        <v-icon
          small
          @click="del(item)"
        >delete
        </v-icon>
      </template>
      <template v-slot:no-data>등록된 문제가 없습니다</template>
    </v-data-table>
  </v-container>
</template>


<script>
  export default {
    data: () => ({
      dialog: false,
      edial: false,
      headers: [
        { text: '관리번호', value: 'ori_no', sortable: false, divider: true, width: "10%" },
        { text: '등록번호', value: 'prob_no', sortable:false, divider: true, width: "10%" },
        { text: '제목', value: 'title', sortable: false, divider: true, width: "50%" },
        { text: '편집', value: 'action', sortable: false, divider: true, width: "30%" },
      ],
      desserts: [],
      editedIndex: -1,
      defaultItem: {
        ori_no: 0,
        prob_no: 0,
        title: "",
        stat: 0,
      },
      sel: 0,
      tsel: 0,
    }),
    created () {
      this.$axios.get("/api/bdmin/list?id=" + this.$f.decodeToken().id, this.$f.makeHeaderObject())
      .then(re => {
        if (re.data.problems === null) return
        this.desserts = re.data.problems
      })
      .catch(err => { this.$f.malert() })
    },
    methods: {
      del(item) {

      },
      toDetail(item) { this.$router.push("/sup/detail/" + item) },
      editItem (item) {
        this.editedIndex = this.desserts.indexOf(item)
        this.sel = item.stat
        this.tsel = item.stat
        this.edial = true
      },
      async save () {
        var flag = false
        if (this.tsel === 0 || this.sel == 2) {flag = confirm("문제를 공개하기 전에 다시 한번 생각해보세요. "
        + "지문은 오해의 소지가 없나요? 데이터는 정확한가요? "
        + "공식 풀이는 증명했나요? 최소 2인에게 검수를 받으셨나요?")
          if (flag == true) flag = confirm("문제를 공개하면 데이터를 변경할 수 없습니다. "
          + " 문제가 잘못된 경우 사이트와 이용자에게 심각하게"
          + " 부정적인 영향을 끼칠 수 있다는 것에 유의하셨습니까?")
        }
        else { flag = confirm("정말로 문제 공개상태를 변경하시겠습니까?") }
        if (flag == false) return

        this.desserts[this.editedIndex].stat = this.sel

        await this.$f.getUserValid()
        .then(res => {
          if (res == null) {
            this.$router.push("/login")
            return
          }
          this.$axios.post("/api/bdmin/update/stat",
            { ori_no : this.desserts[this.editedIndex].ori_no,
            fromstat : this.tsel,
            tostat : this.sel }
          , this.$f.makeHeaderObject())
          .catch(err => {this.$f.malert()})
        })
        this.edial = false
      },
      async create () {
        await this.$f.getUserValid().then(
          res => {
            if (res === null) {
              this.$router.push({path:"/login"})
              return
            }

          this.$axios.get("/api/bdmin/new?id=" + res.id, this.$f.makeHeaderObject())
            .then(re => {
              this.defaultItem.ori_no = re.data.ori_no
              this.desserts.unshift(Object.assign({}, this.defaultItem))
            })
            .catch(err => {this.$f.malert()})
          })
        this.dialog = false
      },
    },
  }
</script>
