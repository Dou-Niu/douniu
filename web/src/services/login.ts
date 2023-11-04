import request from "@/utils/request";

/**
 * @description 发送验证码
 */
export const sendCode = (phone: string) =>{
  return request({
    url: "/login/sendcode",
    method: "post",
    data: {
      phone,
    },
  });
}

/**
 * @param phone
 * @param code
 * @description 通过验证码登录
 */
export const loginByCode = (phone: string, code: string) =>{
  return request({
    url: "/login/code",
    method: "post",
    data: {
      phone,
      code
    },
  });
}

/**
 * @param phone
 * @param password
 * @description 通过密码登录
 */
export const loginByPassword = (phone: string, password: string) =>{
  return request({
    url: "/login/password",
    method: "post",
    data: {
      phone,
      password
    },
  });
}

/**
 * @param user_id
 * @param new_password
 * @description 修改密码
 */
export const forgetPassword = (user_id: string, new_password: string) =>{
  return request({
    url:"/password/forget",
    method: "post",
    data: {
      user_id,
      new_password
    }
  })
}

/**
 * 
 * @param user_id 
 * @param new_password 
 * @headers Authorization
 * @description 修改密码
 */
export const changePassword = (user_id: string, new_password: string) =>{
  return request({
    url:"/password/change",
    method: "post",
    data: {
      user_id,
      new_password
    }
  })
}

