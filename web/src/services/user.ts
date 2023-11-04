import request from '@/utils/request';

/**
 * @param user_id
 * @description 用户信息
 */
export const getUserInfo = (user_id: string) =>{
  return request({
    url:`/user/userinfo?user_id=${user_id}`,
    method: "get",
  })
}

/**
 * @description 刷新token
 */
export const refreshToken = () =>{
  return request({
    url:`/user/refreshToken}`,
    method: "post",
  })
}

/**
 * @param types 修改类型 1:昵称 2:个性签名 3:头像 4:背景图
 * @param value 值
 * @headers Authorization
 * @description 用户信息
 */
export const modifyUserInfo = () =>{
  return request({
    url:`/user/userinfo/modify`,
    method: "post",
  })
}