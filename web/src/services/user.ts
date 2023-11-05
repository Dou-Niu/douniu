import { post, get } from './method'

export type User = {
    id: bigint,
    phone: string,
    nickname: string,
    follow_count: number,
    follower_count: number,
    is_follow: boolean,
    avatar: string,
    background_image: string,
    signature: string,
    total_favorited: string,
    work_count: number,
    favorite_count: number,
    collection_count: number
}

// 修改用户信息
export const changeUserInfo = (types: number, value: string): Promise<{ code: number, message: string }> => post('/user/userinfo/modify', { types, value }) as any
// 获取用户信息
export const getUserInfo =
    (user_id: bigint): Promise<{
        code: number, message: string,
        data: {
            userinfo: User
        }
    }> => get('/user/userinfo', { user_id }) as any;