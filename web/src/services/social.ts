import { post, get, Resp } from './method';
/**
 * 关注
 * @param to_user_id 
 * @param action_type 
 * @returns 
 */
export const toFollow = (to_user_id : number, action_type: number) => post('/relation/action', { to_user_id, action_type })
/**
 * 获取某人的关注列表
 * @param user_id 
 * @param page_num 
 * @returns 
 */
export const getSbFollowList = (user_id : number, page_num: number) => get('/relation/follow/list', { user_id, page_num })
/**
 * 获取某人的粉丝列表
 * @param user_id 
 * @param page_num 
 * @returns 
 */
export const getSbFollowerList = (user_id : number, page_num: number) => get('/relation/follower/list', { user_id, page_num })

export const getFriendsList = (user_id : number): Resp<{
    "user_list": []
}> => get('/relation/friend/list', { user_id })