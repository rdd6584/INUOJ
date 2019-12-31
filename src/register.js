var input_list = new Vue({
  el: '#input_list',
  data:{
    id:"myID",
    pass:"",
    pass_confirm:"",
    email:""
  },
  methods:{
    check_length: function(le, ri, t) {
      return t.length >= le && t.length <= ri;
    },
    check_format: function() {
      for (var ch of this.id) {
        if (ch >= '0' && ch <= '9' || ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z');
        else return false;
      }
      return true;
    },
    check_email: function() {
      var regExp = /^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]{2,3}$/i;
      return this.email.match(regExp) != null;
    },
    check_all: function() {
      flag = 0;
      if (check_length(3, 20, this.id) == true && check_length(8, 20, this.pass))
        if (check_email() && check_format() && pass == pass_confirm)
          return true;
        return false;
    },
    /*
    req: function() {
      http.interceptors.push(function(request) {
        request.method = 'POST';
        request.headers.set('X-CSRF-TOKEN', 'TOKEN');
        request.headers.set('Authorization', 'Bearer TOKEN');

      });
    }
    */
  }
});
