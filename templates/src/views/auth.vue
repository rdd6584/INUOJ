<template>
  <div>
  </div>
</template>

<script>
export default {
  created() {
    this.doAuth()
  },
  methods: {
    doAuth() {
      this.$axios.post('/api/auth', {token : $route.query.token, email : $route.query.email})
      .then(res => {
        if(res.data.status) localStorage.setItem('auth', true)
        else alert("토큰 정보가 유효하지 않습니다.")
      })
      .catch(err => {
        alert("예기치 못한 오류가 발생했습니다.")
      })
    }
  },
}
</script>
