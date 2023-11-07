import { post } from './method'

// 获取验证码
export const getCode =
    (phone: string): Promise<{ code: number, message: string, data: { verification_code: string } }> => post('/user/sendcode', { phone }) as any;
// 手机号加验证码登录
export const loginByCode =
    (phone: string, verification_code: string): Promise<{
        code: number,
        message: string,
        data: { user_id : number, access_token: string, refresh_token: string }
    }> => post('/user/login/phone', { phone, verification_code }) as any
// 手机号加密码登录
export const loginByPass =
    (phone: string, password: string): Promise<{ code: number, message: string, data: { user_id: string, access_token: string, refresh_token: string } }> => post('/user/login/password', { phone, password }) as any
// 忘记密码
export const setPasswordByCode =
    (phone: string, new_password: string): Promise<{ code: number, message: string }> => post('/user/password/forget', { phone, new_password }) as any
// 设置密码
export const setPassord =
    (new_password: string): Promise<{ code: number, message: string }> => post('/user/password/change', { new_password }) as any
// 刷新token
export const refreshToken =
    (): Promise<{
        code: number,
        message: string,
        "data": {
            "access_token": string,
            "refresh_token": string
        }
    }> => post('/user/refreshToken') as any;
