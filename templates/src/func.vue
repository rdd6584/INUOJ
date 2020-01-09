<script>
import axios from 'axios'
  export default{
    data: () => ({

    }),
    methods: {
      isOnlyNum(st) {
        for (i of st)
          if (i < '0' || i > '9') return false
        return true
      },
      makeHeaderObject() {
        var token = {'headers' : {'Authorization' : 'INUser ' + this.getToken()}}
        return token
      },
      async getUserValid() {
        return true
        
        var token = this.getToken()
        if (token == null) return null

        var parse = this.decodeToken()
        if (this.getUnixTime() > parse.exp - 5) {
          return axios.get('/api/refresh', this.makeHeaderObject()).then(
            res => {
            localStorage.setItem('token', res.data.token)
            return this.decodeToken()
          }).catch(err => {return null})
        }
        else return parse
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
