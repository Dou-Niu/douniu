import { defineStore } from "pinia";
import type{ User } from '@/services/user';
import { user as userApi } from '@/services'

interface user {
    refresh_token: string
    access_token: string
    user_id: bigint | string
    user_info: User
}

export const user = defineStore('user', {
    state: (): user => {
        return {
            refresh_token: localStorage.getItem('REFRESH_TOKEN') || '',
            access_token: localStorage.getItem('ACCESS_TOKEN') || '',
            user_id: localStorage.getItem('USER_ID') || '',
            user_info: localStorage.getItem('USER_INFO') ? JSON.parse(localStorage.getItem('USER_INFO') as string) as User : {
                id: BigInt(0),
                phone: '',
                nickname: '',
                follow_count: 0,
                follower_count: 0,
                is_follow: false,
                avatar: 'https://www.kecat.top/avatar.webp',
                background_image: '',
                signature: '',
                total_favorited: '',
                work_count: 0,
                favorite_count: 0,
                collection_count: 0
            }
        }
    },
    actions: {
        getToken(key: 'access_token' | 'refresh_token') {
            return this[key];
        },
        setToken(key: 'refresh_token' | 'access_token', value: string) {
            this[key] = value;
            localStorage.setItem(key.toUpperCase(), value);
        },
        setUserId(id: string) {
            this.user_id = id;
            localStorage.setItem('USER_ID', id);
        },
        setUserInfo(value: User) {
            this.user_info = value;
        },
        async getUserInfo() {
            const user_info = await userApi.getUserInfo(BigInt(this.user_id));
            this.setUserInfo(user_info.data.userinfo);
            localStorage.setItem("USER_INFO", JSON.stringify(this.user_info));
            return user_info.data.userinfo;
        }
    }
})
