import { post, get, Resp, del } from "./method";
import {Video} from '@/types'

export enum Partition {
    "游戏",
    "生活",
    "影视",
    "动漫",
    "知识"
}

export const publishVideo = (video_url: string, cover_url: string, title: string, partition: number): Promise<{
    "code": number
    "message": string
    "data": 0
}> => post("/video/publish", {
    video_url, cover_url, title, partition
}) as any

// export interface Video {
//     "video_id": bigint
//     "author": {
//         "id": bigint
//         "nickname": string
//         "follow_count": number
//         "follower_count": number
//         "is_follow": boolean
//         "avatar": string
//         "background_image": string
//         "signature": string
//         "total_favorited": number
//         "work_count": number
//         "favorite_count": number
//         "collection_count": number
//     }
//     "play_url": string
//     "cover_url": string
//     "favorite_count": number
//     "collection_count": number
//     "comment_count": number
//     "is_favorite": boolean
//     "title": string
//     "partition": number
//     "create_time": string
// }

export const searchVideo = (key_words: string, page: number): Promise<{
    "code": number
    "message": string
    "data": {
        "next_max_value": 2
        "is_final": false
        "video_list": Video[]
    }
}> => get('/video/search', { key_words, page })

export const getHomeVideo = (latest_time: number): Promise<{
    code: number,
    message: string,
    data: {
        next_time: number
        video_list: Video[]
    }
}> => get('/video/feed/home', { latest_time })

export const getHotVideo = (max_hot: number): Resp<{
    NextMaxHot: number
    video_list: Video[]
}> => get('/video/feed/hot', { max_hot })

export const getFollowFeed = (latest_time: number): Resp<{
    next_max_value: number,
    is_final: number,
    video_list: Video[]
}> => get('/video/feed/follow', { latest_time })


/**
 * 
 * @param user_id 
 * @param max_value 
 * @param sort 1-热点排序 2-最新排序
 * @returns 
 */
export const getUserAllVideo = (user_id: bigint, max_value: number, sort: number): Resp<{
    next_max_value: number,
    is_final: number,
    video_list: Video[]
}> => get('/video/feed/user', { user_id, max_value, sort })

/**
 * 
 * @param max_value 
 * @param sort 1-热点排序，2-最新排序
 * @param partition 1-游戏 2-生活 3-影视 4-动漫 5-知识
 * @returns 
 */
export const getPartitionVideo = (max_value: number, sort: number, partition: number): Resp<{
    next_max_value: number,
    is_final: number,
    video_list: Video[]
}> => get("/video/feed/partition", { max_value, sort, partition })

export const delVideo = (video_id: bigint, partition: number): Resp<0> => del('/video/delete', { video_id, partition })

export const shareVideo = (video_id: bigint): Resp<{
    "share_url": string
}> => get('/video/share', { video_id })

export const getVideoInfo = (video_id: bigint) => get('/video', { video_id });

/**
 * 
 * @param video_id 
 * @param action_type 1是点赞，2是取消点赞，其余报错
 * @param partition 
 * @returns 
 */
export const toLikeVideo = (video_id : bigint, action_type: number, partition: number) => post('/favorite/action', { video_id, action_type, partition })

export const getLikeVideoList = (user_id: bigint, page_num: number) => get('/favorite/list', { user_id, page_num })

/**
 * 
 * @param video_id 
 * @param action_type 1是点赞，2是取消点赞，其余报错
 * @param partition 
 * @returns 
 */
export const toCollectVideo = (video_id : bigint, action_type: number, partition: number) => post('/collection/action', { video_id, action_type, partition })
export const getCollectVideoList = (user_id: bigint, page_num: number) => get('/collection/list', { user_id, page_num })
