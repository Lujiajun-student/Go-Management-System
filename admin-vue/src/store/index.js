// vuex状态管理
import Vue from 'vue'
import Vuex from 'vuex'
import mutations from './mutations'

Vue.use(Vuex)

const state = new Vuex.Store({
    // todo
    mutations
})

export default state