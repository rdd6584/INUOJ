import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    lang: ['모든 언어', 'C++', 'C', 'Go', 'Java', 'pypy3','Python3', 'Kotlin', ],
    langOrd: {'모든 언어': 0, 'C++': 1, 'C': 2, 'Go': 3, 'Java': 4, 'pypy3': 5, 'Python3': 6, 'Kotlin': 7, },
    result: ['모든 결과', 'Accept', 'Wrong answer', 'Time limit exceeded', 'Memory limit exceeded', 'Runtime error', 'Compile error'],
    resultOrd: { '모든 결과': 0, 'Accept': 1, 'Wrong answer': 2, 'Time limit exceeded': 3, 'Memory limit exceeded': 4, 'Runtime error': 5, 'Compile error': 6, },
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
