<script>
import axios from 'axios'

  export default{
    data: () => ({
      userId: "",
      admin: "",
    }),
    methods: {
      isOnlyNum(st) {
        for (var i of st)
          if (i < '0' || i > '9') return false
        return true
      },
      makeHeaderObject() {
        var token = {'headers' : {'Authorization' : 'INUser ' + this.getToken()}}
        return token
      },
      async getUserValid() {
        var token = this.getToken()
        if (token == null) {
          this.userId = ""
          this.admin = ""
          return null
        }

        var parse = this.decodeToken()
        if (this.getUnixTime() > parse.exp - 5) {
          return axios.get('/api/refresh', this.makeHeaderObject()).then(
            res => {
            localStorage.setItem('token', res.data.token)

            var parse = this.decodeToken()
            this.userId = parse.id
            this.admin = parse.admin

            return parse
          }).catch(err => {
            this.userId = ""
            this.admin = ""
            return null}
          )
        }
        else {
          this.userId = parse.id
          this.admin = parse.admin
          return parse
        }
      },
      getToken() {
        return localStorage.getItem('token')
      },
      getUnixTime() {
        return new Date().getTime() / 1000
      },
      decodeToken() {
        var token = this.getToken()
        if (token == null) return null

        var flag = 0
        var ret = ""
        for (var i in token) {
          if (token[i] == '.') flag++
          else if (flag === 1) ret += token[i]
        }

        return JSON.parse(atob(ret))
      },
      malert() {
        alert("예기치 못한 오류가 발생했습니다.")
      },
    },
  }
</script>
