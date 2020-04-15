import Cookies from 'js-cookie'
import CookieParser from 'cookie'
export const strict = false

export const state = () => ({
  user: null
})

export const mutations = {
  setUser(state, payload) {
    state.user = payload
  }
}

export const actions = {
  nuxtServerInit({ commit }, { req }) {
    if(req.headers.cookie) {
      let data = CookieParser.parse(req.headers.cookie)
      commit("setUser", data.user)
    }
  },
  setUser({ commit }, payload) {
    if(payload === null) {
      commit('setUser', null)
      Cookies.remove("user");
    } else {
      commit('setUser', payload)
      Cookies.set("user", payload)
    }
  }
}

export const getters = {
  isAuthenticated(state) {
    return !!state.user
  }
}