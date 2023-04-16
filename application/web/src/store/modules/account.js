import {
  login,
  authentication,
  queryAccountList,
  adminlogin
} from '@/api/account'
import {
  getToken,
  setToken,
  removeToken,
  setTokenUsername,
  setTokenPassword
} from '@/utils/auth'
import {
  resetRouter
} from '@/router'

const getDefaultState = () => {
  return {
    token: getToken(),
    accountId: '',
    userName: '',
    balance: 0,
    roles: []
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_ACCOUNTID: (state, accountId) => {
    state.accountId = accountId
  },
  SET_USERNAME: (state, userName) => {
    state.userName = userName
  },
  SET_BALANCE: (state, balance) => {
    state.balance = balance
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_PASSWORD: (state, passWord) => {
    state.passWord = passWord
  }
}

const actions = {
  login({
    commit
  }, accountIdargs) {
    return new Promise((resolve, reject) => {
      login(
        accountIdargs
      ).then(response => {
        commit('SET_TOKEN', response[0].accountId)
//        commit('SET_ROLES', ['editor'])
        commit('SET_USERNAME',response[0].userName)
        // commit('SET_ACCOUNTID', response[0].accountId)
        commit('SET_PASSWORD',response[0].passWord)
        setToken(response[0].userName)
        setTokenUsername(response[0].userName)
        setTokenPassword(response[0].passWord)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  adminlogin({
    commit
  }, accountIdargs) {
    return new Promise((resolve, reject) => {
      adminlogin(
        accountIdargs
      ).then(response => {
        commit('SET_TOKEN', response[0].accountId)
        commit('SET_USERNAME',response[0].userName)
        // commit('SET_ACCOUNTID', response[0].accountId)
        commit('SET_PASSWORD',response[0].passWord)
        setToken(response[0].userName)
        setTokenUsername(response[0].userName)
        setTokenPassword(response[0].passWord)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  authentication({
    commit
  }, accountIdargs) {
    return new Promise((resolve, reject) => {
      authentication(
        accountIdargs
      ).then(response => {
        commit('SET_TOKEN', response.accountId)
        commit('SET_USERNAME',response.userName)
        commit('SET_ACCOUNTID', response.accountId)
        commit('SET_PASSWORD',response.passWord)
        commit('SET_BALANCE', response.balance)
        setToken(response.accountId)
        setTokenUsername(response.userName)
        setTokenPassword(response.passWord)
       // alert(state.accountId)
        resolve(response.accountId)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // get user info
  getInfo({
    commit
  }, tokenpass) {
    return new Promise((resolve, reject) => {
      //alert(tokenpass.username)
      if (tokenpass.username === 'admin') {
      adminlogin({
          username:tokenpass.username,
          password:tokenpass.password,
        })
        .then(response => {
          const roles = ['admin']
          commit('SET_ROLES', roles)
          commit('SET_USERNAME', response[0].userName)
          commit('SET_ACCOUNTID', response[0].accountId)
          resolve(roles)
        })
        .catch(error => {
          alert(JSON.stringify(error))
          reject(error)
        })
    } else {
      login({
        username: tokenpass.username,
        password: tokenpass.password
      }).then(response => {
       // alert(tokenpass.username)
        const roles = ['editor']
        commit('SET_ROLES', roles)
        commit('SET_ACCOUNTID', response[0].accountId)
        commit('SET_USERNAME', response[0].userName)
        resolve(roles)
      }).catch(error => {
        reject(error)
      })
    }
    })
  },
  logout({
    commit
  }) {
    return new Promise(resolve => {
      removeToken()
      resetRouter()
      commit('RESET_STATE')
      resolve()
    })
  },

  resetToken({
    commit
  }) {
    return new Promise(resolve => {
      removeToken()
      commit('RESET_STATE')
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
