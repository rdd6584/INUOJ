(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-1c5c29d4"],{1681:function(t,e,r){},"1f4f":function(t,e,r){"use strict";r("a4d3"),r("4de4"),r("a9e3"),r("e439"),r("dbb4"),r("b64b"),r("159b");var a=r("ade3"),s=(r("8b37"),r("80d2")),n=r("7560"),o=r("58df");function i(t,e){var r=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),r.push.apply(r,a)}return r}function c(t){for(var e=1;e<arguments.length;e++){var r=null!=arguments[e]?arguments[e]:{};e%2?i(Object(r),!0).forEach((function(e){Object(a["a"])(t,e,r[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(r,e))}))}return t}e["a"]=Object(o["a"])(n["a"]).extend({name:"v-simple-table",props:{dense:Boolean,fixedHeader:Boolean,height:[Number,String]},computed:{classes:function(){return c({"v-data-table--dense":this.dense,"v-data-table--fixed-height":!!this.height&&!this.fixedHeader,"v-data-table--fixed-header":this.fixedHeader},this.themeClasses)}},methods:{genWrapper:function(){return this.$slots.wrapper||this.$createElement("div",{staticClass:"v-data-table__wrapper",style:{height:Object(s["g"])(this.height)}},[this.$createElement("table",this.$slots.default)])}},render:function(t){return t("div",{staticClass:"v-data-table",class:this.classes},[this.$slots.top,this.genWrapper(),this.$slots.bottom])}})},"2fa4":function(t,e,r){"use strict";r("20f6");var a=r("80d2");e["a"]=Object(a["h"])("spacer","div","v-spacer")},"615b":function(t,e,r){},"62ad":function(t,e,r){"use strict";r("a4d3"),r("4de4"),r("caad"),r("4ec9"),r("a9e3"),r("e439"),r("dbb4"),r("b64b"),r("d3b7"),r("ac1f"),r("3ca3"),r("5319"),r("2ca0"),r("159b"),r("ddb0");var a=r("ade3"),s=(r("4b85"),r("2b0e")),n=r("d9f7"),o=r("80d2");function i(t,e){var r=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),r.push.apply(r,a)}return r}function c(t){for(var e=1;e<arguments.length;e++){var r=null!=arguments[e]?arguments[e]:{};e%2?i(Object(r),!0).forEach((function(e){Object(a["a"])(t,e,r[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(r,e))}))}return t}var l=["sm","md","lg","xl"],u=function(){return l.reduce((function(t,e){return t[e]={type:[Boolean,String,Number],default:!1},t}),{})}(),p=function(){return l.reduce((function(t,e){return t["offset"+Object(o["C"])(e)]={type:[String,Number],default:null},t}),{})}(),d=function(){return l.reduce((function(t,e){return t["order"+Object(o["C"])(e)]={type:[String,Number],default:null},t}),{})}(),f={col:Object.keys(u),offset:Object.keys(p),order:Object.keys(d)};function b(t,e,r){var a=t;if(null!=r&&!1!==r){if(e){var s=e.replace(t,"");a+="-".concat(s)}return"col"!==t||""!==r&&!0!==r?(a+="-".concat(r),a.toLowerCase()):a.toLowerCase()}}var h=new Map;e["a"]=s["a"].extend({name:"v-col",functional:!0,props:c({cols:{type:[Boolean,String,Number],default:!1}},u,{offset:{type:[String,Number],default:null}},p,{order:{type:[String,Number],default:null}},d,{alignSelf:{type:String,default:null,validator:function(t){return["auto","start","end","center","baseline","stretch"].includes(t)}},tag:{type:String,default:"div"}}),render:function(t,e){var r=e.props,s=e.data,o=e.children,i=(e.parent,"");for(var c in r)i+=String(r[c]);var l=h.get(i);return l||function(){var t,e;for(e in l=[],f)f[e].forEach((function(t){var a=r[t],s=b(e,t,a);s&&l.push(s)}));var s=l.some((function(t){return t.startsWith("col-")}));l.push((t={col:!s||!r.cols},Object(a["a"])(t,"col-".concat(r.cols),r.cols),Object(a["a"])(t,"offset-".concat(r.offset),r.offset),Object(a["a"])(t,"order-".concat(r.order),r.order),Object(a["a"])(t,"align-self-".concat(r.alignSelf),r.alignSelf),t)),h.set(i,l)}(),t(r.tag,Object(n["a"])(s,{class:l}),o)}})},"8b37":function(t,e,r){},"8ce9":function(t,e,r){},a844:function(t,e,r){"use strict";r("a4d3"),r("4de4"),r("a9e3"),r("e439"),r("dbb4"),r("b64b"),r("acd8"),r("e25e"),r("159b");var a=r("ade3"),s=(r("1681"),r("8654")),n=r("58df");function o(t,e){var r=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),r.push.apply(r,a)}return r}function i(t){for(var e=1;e<arguments.length;e++){var r=null!=arguments[e]?arguments[e]:{};e%2?o(Object(r),!0).forEach((function(e){Object(a["a"])(t,e,r[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(r,e))}))}return t}var c=Object(n["a"])(s["a"]);e["a"]=c.extend({name:"v-textarea",props:{autoGrow:Boolean,noResize:Boolean,rowHeight:{type:[Number,String],default:24,validator:function(t){return!isNaN(parseFloat(t))}},rows:{type:[Number,String],default:5,validator:function(t){return!isNaN(parseInt(t,10))}}},computed:{classes:function(){return i({"v-textarea":!0,"v-textarea--auto-grow":this.autoGrow,"v-textarea--no-resize":this.noResizeHandle},s["a"].options.computed.classes.call(this))},noResizeHandle:function(){return this.noResize||this.autoGrow}},watch:{lazyValue:function(){this.autoGrow&&this.$nextTick(this.calculateInputHeight)},rowHeight:function(){this.autoGrow&&this.$nextTick(this.calculateInputHeight)}},mounted:function(){var t=this;setTimeout((function(){t.autoGrow&&t.calculateInputHeight()}),0)},methods:{calculateInputHeight:function(){var t=this.$refs.input;if(t){t.style.height="0";var e=t.scrollHeight,r=parseInt(this.rows,10)*parseFloat(this.rowHeight);t.style.height=Math.max(r,e)+"px"}},genInput:function(){var t=s["a"].options.methods.genInput.call(this);return t.tag="textarea",delete t.data.attrs.type,t.data.attrs.rows=this.rows,t},onInput:function(t){s["a"].options.methods.onInput.call(this,t),this.autoGrow&&this.calculateInputHeight()},onKeyDown:function(t){this.isFocused&&13===t.keyCode&&t.stopPropagation(),this.$emit("keydown",t)}}})},b0af:function(t,e,r){"use strict";r("a4d3"),r("4de4"),r("0481"),r("4069"),r("a9e3"),r("e439"),r("dbb4"),r("b64b"),r("159b");var a=r("ade3"),s=(r("615b"),r("10d2")),n=r("297c"),o=r("1c87"),i=r("58df");function c(t,e){var r=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),r.push.apply(r,a)}return r}function l(t){for(var e=1;e<arguments.length;e++){var r=null!=arguments[e]?arguments[e]:{};e%2?c(Object(r),!0).forEach((function(e){Object(a["a"])(t,e,r[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(r)):c(Object(r)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(r,e))}))}return t}e["a"]=Object(i["a"])(n["a"],o["a"],s["a"]).extend({name:"v-card",props:{flat:Boolean,hover:Boolean,img:String,link:Boolean,loaderHeight:{type:[Number,String],default:4},outlined:Boolean,raised:Boolean,shaped:Boolean},computed:{classes:function(){return l({"v-card":!0},o["a"].options.computed.classes.call(this),{"v-card--flat":this.flat,"v-card--hover":this.hover,"v-card--link":this.isClickable,"v-card--loading":this.loading,"v-card--disabled":this.disabled,"v-card--outlined":this.outlined,"v-card--raised":this.raised,"v-card--shaped":this.shaped},s["a"].options.computed.classes.call(this))},styles:function(){var t=l({},s["a"].options.computed.styles.call(this));return this.img&&(t.background='url("'.concat(this.img,'") center center / cover no-repeat')),t}},methods:{genProgress:function(){var t=n["a"].options.methods.genProgress.call(this);return t?this.$createElement("div",{staticClass:"v-card__progress"},[t]):null}},render:function(t){var e=this.generateRouteLink(),r=e.tag,a=e.data;return a.style=this.styles,this.isClickable&&(a.attrs=a.attrs||{},a.attrs.tabindex=0),t(r,this.setBackgroundColor(this.color,a),[this.genProgress(),this.$slots.default])}})},ce7e:function(t,e,r){"use strict";r("a4d3"),r("4de4"),r("e439"),r("dbb4"),r("b64b"),r("159b");var a=r("ade3"),s=(r("8ce9"),r("7560"));function n(t,e){var r=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),r.push.apply(r,a)}return r}function o(t){for(var e=1;e<arguments.length;e++){var r=null!=arguments[e]?arguments[e]:{};e%2?n(Object(r),!0).forEach((function(e){Object(a["a"])(t,e,r[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(r)):n(Object(r)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(r,e))}))}return t}e["a"]=s["a"].extend({name:"v-divider",props:{inset:Boolean,vertical:Boolean},render:function(t){var e;return this.$attrs.role&&"separator"!==this.$attrs.role||(e=this.vertical?"vertical":"horizontal"),t("hr",{class:o({"v-divider":!0,"v-divider--inset":this.inset,"v-divider--vertical":this.vertical},this.themeClasses),attrs:o({role:"separator","aria-orientation":e},this.$attrs),on:this.$listeners})}})},dd74:function(t,e,r){},ecff:function(t,e,r){"use strict";r.r(e);var a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-container",[r("v-card",{staticClass:"px-5",attrs:{elevation:"4"}},[r("v-row",{staticClass:"py-3"},[r("v-btn",{staticClass:"mx-3",attrs:{color:"success",router:"",to:{path:"/submit/"+t.prob_no}}},[t._v("제출")]),r("v-spacer"),r("v-btn",{staticClass:"mx-3",attrs:{dark:"",router:"",to:{path:"/status?prob_no="+t.prob_no}}},[t._v("채점현황")]),r("v-btn",{staticClass:"mx-3",attrs:{dark:"",router:"",to:{path:"/board?&prob_no="+t.prob_no}}},[t._v("게시판")])],1),r("v-row",[r("h2",[t._v(t._s(t.title))]),1==t.result?r("i",{staticClass:"mx-2 mt-3 fas fa-circle",staticStyle:{color:"#00C853"}}):0!=t.result?r("i",{staticClass:"mx-2 mt-3 fas fa-circle",staticStyle:{color:"red"}}):t._e(),r("v-spacer"),""!=t.owner?r("span",{staticClass:"pr-1 pt-3"},[t._v(" 출제자:"),r("a",{staticClass:"line",staticStyle:{color:"black"},on:{click:function(e){return t.$router.push({path:"/profile/"+t.owner})}}},[t._v(t._s(t.owner))])]):t._e()],1),r("v-simple-table",{attrs:{dense:""},scopedSlots:t._u([{key:"default",fn:function(){return[r("thead",[r("tr",[r("th",{staticClass:"text-left"},[t._v("시간 제한")]),r("th",{staticClass:"text-left"},[t._v("메모리 제한")]),r("th",{staticClass:"text-left"},[t._v("맞은 사람")]),r("th",{staticClass:"text-left"},[t._v("시도한 사람")])])]),r("tbody",[r("tr",[r("td",[t._v(t._s(t.t_limit)+" ms")]),r("td",[t._v(t._s(t.m_limit)+" MB")]),r("td",[t._v(t._s(t.accept)+" 명")]),r("td",[t._v(t._s(t.attempt)+" 명")])])])]},proxy:!0}])}),r("v-row",{staticClass:"pt-3"},[r("span",{staticClass:"viewHtml",domProps:{innerHTML:t._s(t.description[0])}})]),r("v-row",[r("h3",{staticClass:"mt-1",staticStyle:{color:"rgb(0, 98, 184)"}},[t._v("입력")])]),r("v-row",[r("span",{staticClass:"viewHtml",domProps:{innerHTML:t._s(t.description[1])}})]),r("v-row",[r("h3",{staticClass:"mt-1",staticStyle:{color:"rgb(0, 98, 184)"}},[t._v("출력")])]),r("v-row",[r("span",{staticClass:"viewHtml",domProps:{innerHTML:t._s(t.description[2])}})]),r("v-divider",{staticClass:"mt-6"}),t._l(t.samplein,(function(e,a){return r("v-row",{staticClass:"px-3"},[r("v-col",{staticClass:"pl-0 ml-0",attrs:{cols:"6"}},[r("v-row",{staticClass:"px-4 pb-1"},[r("h4",{staticClass:"mt-1",staticStyle:{color:"rgb(0, 98, 184)"}},[t._v("예제"+t._s(a+1)+" 입력")]),r("v-spacer"),r("v-btn",{directives:[{name:"clipboard",rawName:"v-clipboard",value:t.samplein[a],expression:"samplein[i]"}],attrs:{dark:"",small:""}},[t._v("복사")])],1),r("v-card",{staticClass:"py-0 my-0",attrs:{color:"rgb(245,245,245)",elevation:"0"}},[r("v-textarea",{attrs:{outlined:"",dense:"","auto-grow":"","hide-details":"",readonly:""},model:{value:t.samplein[a],callback:function(e){t.$set(t.samplein,a,e)},expression:"samplein[i]"}})],1)],1),r("v-col",{staticClass:"mr-0 pr-0",attrs:{cols:"6"}},[r("v-row",{staticClass:"px-4 pb-1"},[r("h4",{staticClass:"mt-1",staticStyle:{color:"rgb(0, 98, 184)"}},[t._v("예제"+t._s(a+1)+" 출력")]),r("v-spacer"),r("v-btn",{directives:[{name:"clipboard",rawName:"v-clipboard",value:t.sampleout[a],expression:"sampleout[i]"}],attrs:{dark:"",small:""}},[t._v("복사")])],1),r("v-card",{attrs:{color:"rgb(245,245,245)",elevation:"0"}},[r("v-textarea",{attrs:{outlined:"",dense:"","auto-grow":"","hide-details":"",readonly:""},model:{value:t.sampleout[a],callback:function(e){t.$set(t.sampleout,a,e)},expression:"sampleout[i]"}})],1)],1)],1)})),r("div",{staticClass:"py-2"})],2)],1)},s=[],n=(r("a4d3"),r("e01a"),r("99af"),r("a9e3"),r("96cf"),r("1da1")),o={data:function(){return{prob_no:0,stat:0,result:0,owner:"",title:"제목입니다.",t_limit:1e3,m_limit:512,samplein:[""],sampleout:[""],description:["본문입니다.","입력입니다.","출력입니다."],accept:0,attempt:0}},created:function(){var t=Object(n["a"])(regeneratorRuntime.mark((function t(){var e=this;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return this.prob_no=Number(this.$route.params.num),t.next=3,this.$axios.get("/api/problem/detail/".concat(this.prob_no,"/").concat(this.$f.userId),this.$f.makeHeaderObject()).then((function(t){e.t_limit=t.data.t_limit,e.m_limit=t.data.m_limit,e.title=t.data.title,e.owner=t.data.owner,e.stat=t.data.stat,e.result=t.data.result,e.accept=t.data.accept,e.attempt=t.data.attempt,e.description=t.data.description,t.data.samplein&&(e.samplein=t.data.samplein),t.data.sampleout&&(e.sampleout=t.data.sampleout)})).catch((function(t){404!=t.response.status?e.$f.malert():e.$router.push("/wrongaccess")}));case 3:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}()},i=o,c=(r("f817"),r("2877")),l=r("6544"),u=r.n(l),p=r("8336"),d=r("b0af"),f=r("62ad"),b=r("a523"),h=r("ce7e"),v=r("0fd9"),m=r("1f4f"),g=r("2fa4"),O=r("a844"),y=Object(c["a"])(i,a,s,!1,null,"1b37d61e",null);e["default"]=y.exports;u()(y,{VBtn:p["a"],VCard:d["a"],VCol:f["a"],VContainer:b["a"],VDivider:h["a"],VRow:v["a"],VSimpleTable:m["a"],VSpacer:g["a"],VTextarea:O["a"]})},f817:function(t,e,r){"use strict";var a=r("dd74"),s=r.n(a);s.a}}]);
//# sourceMappingURL=chunk-1c5c29d4.f5d5ead5.js.map