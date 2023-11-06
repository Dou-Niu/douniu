import { post, get } from './method';
import { User } from '@/types/user';

export type Comment = {
    id: bigint,
    content: string,
    create_time: string,
    sub_count:number
    user: User
}

export const sendComment = (video_id : bigint, action_type: number,  parent_id : bigint,content?: string, comment_id ?: bigint): Promise<{
    "code": number
    "message": string
    comment: Comment
}> => post('/comment/action', { video_id, action_type,  parent_id, content, comment_id })

/**
 * 评论区展示采用游标分页
 * @param video_id 
 * @param last_comment_id last_comment_id是上一页的最后一条评论id，第一次打开评论，那就是0
 * @returns 
 */
export const getVideoComment = (video_id : bigint, last_comment_id : bigint): Promise<{
    "code": number
    "message": string
    "data":{
        comment_list: Comment[]
    }
}> => get('/comment/list', { video_id, last_comment_id })

/**
 * 获取评论的评论
 * @param comment_id 
 * @returns 
 */
export const getComment = (comment_id : bigint): Promise<{
    "code": number
    "message": string
    comment: Comment
}> => get('/comment/detail', { comment_id });