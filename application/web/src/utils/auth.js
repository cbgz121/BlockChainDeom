import Cookies from 'js-cookie'

const TokenKey = 'account_id_token'
const UserName = 'username'
const PassWord = 'password'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function getTokenUserName() {
  return Cookies.get(UserName)
}

export function getTokenPassWord() {
  return Cookies.get(PassWord)
}

export function getTokenPassword() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}
export function setTokenUsername(token) {
  return Cookies.set(UserName, token)
}
export function setTokenPassword(token) {
  return Cookies.set(PassWord, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}
