import request from '@/utils/request'

//获取登录界面角色选择列表
export function queryAccountList() {
  return request({
    url: '/queryAccountList',
    method: 'post'
  })
}

// 登录
export function login(logindata) {
  return request({
    url: '/login',
    method: 'post',
    data:logindata
  })
}

// 删除用户
export function deleteid(logindata) {
  return request({
    url: '/deleteid',
    method: 'post',
    data: logindata
  })
}

export function register(registerdata) {
  return request({
    url: '/register',
    method: 'post',
    data: registerdata
  })
}

export function authentication(authendata) {
  return request({
    url: '/authentication',
    method: 'post',
    data: authendata
  })
}

export function captcha() {
  return request({
    url: '/captcha',
    method: 'get'
  })
}
