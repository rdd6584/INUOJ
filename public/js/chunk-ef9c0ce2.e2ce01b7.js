(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-ef9c0ce2"],{"4bd4":function(t,e,n){"use strict";n("a4d3"),n("4de4"),n("7db0"),n("caad"),n("e439"),n("dbb4"),n("b64b"),n("07ac"),n("2532"),n("159b");var i=n("ade3"),r=n("58df"),a=n("7e2b"),s=n("3206");function u(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(t);e&&(i=i.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,i)}return n}function o(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?u(Object(n),!0).forEach((function(e){Object(i["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):u(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}e["a"]=Object(r["a"])(a["a"],Object(s["b"])("form")).extend({name:"v-form",inheritAttrs:!1,props:{lazyValidation:Boolean,value:Boolean},data:function(){return{inputs:[],watchers:[],errorBag:{}}},watch:{errorBag:{handler:function(t){var e=Object.values(t).includes(!0);this.$emit("input",!e)},deep:!0,immediate:!0}},methods:{watchInput:function(t){var e=this,n=function(t){return t.$watch("hasError",(function(n){e.$set(e.errorBag,t._uid,n)}),{immediate:!0})},i={_uid:t._uid,valid:function(){},shouldValidate:function(){}};return this.lazyValidation?i.shouldValidate=t.$watch("shouldValidate",(function(r){r&&(e.errorBag.hasOwnProperty(t._uid)||(i.valid=n(t)))})):i.valid=n(t),i},validate:function(){return 0===this.inputs.filter((function(t){return!t.validate(!0)})).length},reset:function(){this.inputs.forEach((function(t){return t.reset()})),this.resetErrorBag()},resetErrorBag:function(){var t=this;this.lazyValidation&&setTimeout((function(){t.errorBag={}}),0)},resetValidation:function(){this.inputs.forEach((function(t){return t.resetValidation()})),this.resetErrorBag()},register:function(t){this.inputs.push(t),this.watchers.push(this.watchInput(t))},unregister:function(t){var e=this.inputs.find((function(e){return e._uid===t._uid}));if(e){var n=this.watchers.find((function(t){return t._uid===e._uid}));n&&(n.valid(),n.shouldValidate()),this.watchers=this.watchers.filter((function(t){return t._uid!==e._uid})),this.inputs=this.inputs.filter((function(t){return t._uid!==e._uid})),this.$delete(this.errorBag,e._uid)}}},render:function(t){var e=this;return t("form",{staticClass:"v-form",attrs:o({novalidate:!0},this.attrs$),on:{submit:function(t){return e.$emit("submit",t)}}},this.$slots.default)}})},7803:function(t,e,n){"use strict";n.r(e);var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-container",[n("v-row",{attrs:{justify:"center",align:"center"}},[n("v-form",{ref:"form",staticStyle:{"min-width":"43%"},model:{value:t.valid,callback:function(e){t.valid=e},expression:"valid"}},[n("v-text-field",{attrs:{rules:t.idRules,counter:20,label:"아이디",maxlength:"20",required:""},on:{keyup:[function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.register()},function(e){return t.valCheck(1)}]},model:{value:t.id,callback:function(e){t.id=e},expression:"id"}}),n("v-text-field",{attrs:{rules:t.passRules,type:"password",label:"비밀번호",maxlength:"20",required:""},on:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.register()}},model:{value:t.pass,callback:function(e){t.pass=e},expression:"pass"}}),n("v-text-field",{attrs:{rules:t.passConfRules,type:"password",label:"비밀번호 확인",maxlength:"20",required:""},on:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.register()}},model:{value:t.passConf,callback:function(e){t.passConf=e},expression:"passConf"}}),n("v-text-field",{attrs:{rules:t.emailRules,label:"E-mail",maxlength:"50",required:""},on:{keyup:[function(e){return t.valCheck(2)},function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.register()}]},model:{value:t.email,callback:function(e){t.email=e},expression:"email"}}),n("v-btn",{staticClass:"mr-4",attrs:{disabled:!t.valid,color:"success"},on:{click:function(e){return t.register()}}},[t._v(" 회원가입 ")])],1)],1)],1)},r=[],a={data:function(){return{valid:!0,id:"",pass:"",passRules:[function(t){return!!t||"비밀번호를 입력하세요"},function(t){return/^(?=.*[a-zA-Z])(?=.*\d).{8,20}$/.test(t)&&t.length>=8&&t.length<=20||"비밀번호 8~20글자이며 알파벳과 숫자를 반드시 포함해야 됩니다."}],passConf:"",email:"@inu.ac.kr",idCheck:!1,emailCheck:!1}},methods:{register:function(){var t=this;this.$axios.post("/api/regi-done",{id:this.id,pass:this.$sha256(this.pass),email:this.email}).then((function(e){"ok"==e.data.status?t.$router.push({path:"/auth?id="+t.id}):alert("회원가입 도중 오류가 발생했습니다.")})).catch((function(t){alert("회원가입 도중 오류가 발생했습니다.")}))},valCheck:function(t){var e=this,n="";1==t?n="id=".concat(this.id):2==t&&(n="email=".concat(this.email)),this.$axios.get("/api/register/valid?"+n).then((function(n){1==t?e.idCheck=n.data.status:2==t&&(e.emailCheck=n.data.status)}))}},computed:{idRules:function(){var t=[function(t){return!!t||"ID를 입력하세요"}];return t.push((function(t){return/^[a-zA-Z0-9]+/.test(t)&&t.length>=2&&t.length<=20||"아이디는 2~20글자이며 알파벳과 숫자로만 구성됩니다."})),t.push(this.idCheck||"중복되는 ID가 있습니다."),t},emailRules:function(){var t=[function(t){return!!t||"이메일을 입력하세요"}];return t.push((function(t){return/^[a-zA-Z0-9\_]+@inu\.ac\.kr/.test(t)||"이메일은 @inu.ac.kr로 끝나야 합니다."})),t.push(this.emailCheck||"중복되는 이메일이 있습니다."),t},passConfRules:function(){var t=[function(t){return!!t||"비밀번호를 확인하세요"}];return t.push(this.pass===this.passConf||"비밀번호를 확인하세요"),t}}},s=a,u=n("2877"),o=n("6544"),c=n.n(o),l=n("8336"),d=n("a523"),f=n("4bd4"),h=n("0fd9"),p=n("8654"),v=Object(u["a"])(s,i,r,!1,null,null,null);e["default"]=v.exports;c()(v,{VBtn:l["a"],VContainer:d["a"],VForm:f["a"],VRow:h["a"],VTextField:p["a"]})}}]);
//# sourceMappingURL=chunk-ef9c0ce2.e2ce01b7.js.map