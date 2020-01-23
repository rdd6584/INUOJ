import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    lang: ['모든 언어', 'C++', 'C', 'Go', 'Java', 'pypy3','Python3', 'Kotlin', ],
    langOrd: {'모든 언어': 0, 'C++': 1, 'C': 2, 'Go': 3, 'Java': 4, 'pypy3': 5, 'Python3': 6, 'Kotlin': 7, },

    slang: ['C++', 'C', 'Go', 'Java', 'pypy3','Python3', 'Kotlin', ],
    langCodeMirror: ['text/x-c++src', 'text/x-csrc', 'text/x-c++src', 'text/x-java', 'text/x-python', 'text/x-python', 'text/x-kotlin'],

    result: ['모든 결과', 'Accept', 'Wrong answer', 'Time limit exceeded', 'Memory limit exceeded', 'Runtime error', 'Compile error', '채점 실패'],
    resultOrd: {'채점중': 0, '모든 결과': 0, 'Accept': 1, 'Wrong answer': 2, 'Time limit exceeded': 3, 'Memory limit exceeded': 4, 'Runtime error': 5, 'Compile error': 6, '채점 실패': 7},

    stat: ['미등록', '채점불가', '공개', '비공개'],
    statOrd: {"미등록" : 0, "채점불가" : 1, "공개" : 2, "비공개" : 3},
    statNedg: [1, 0, 0, 0],

    category2: ['공지', '질문', '자유'],
    category: ['전체', '공지', '질문', '자유'],
    category0: ['질문', '자유'],
    categoryOrd: {"전체" : 0, "공지" : 1, "질문" : 2, "자유" : 3},
  },
  getters: {

  },
  mutations: {
  },
  actions: {
  },
  modules: {
  }
})
