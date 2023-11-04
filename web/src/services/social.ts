import request from "@/utils/request";

/**
 * @param to_user_id 关注的人的id
 * @param action_type 操作类型 1为关注，2为取消关注
 */
export const followPeople = (to_user_id:any,action_type:any)=>{
  return request({
    url: "/relation/action",
    method: "post",
    data:{
      to_user_id,
      action_type
    }
  })
}

/**
 * @param user_id 用户id
 * @param page_num 页数
 * @description 关注列表
 */
export const getFollowList = (user_id:any,page:any)=>{
  return request({
    url: `/relation/follow/list?user_id=${user_id}&page_num=${page}`,
    method: "get",
  })
}

/**
 * @param user_id 用户id
 * @param page_num 页数
 * @description 粉丝列表
 */
export const getFollowerList = (user_id:string,page:number)=>{
  return request({
    url: `/relation/follower/list?user_id=${user_id}&page_num=${page}`,
    method: "get",
  })
}

/**
 * @param user_id 用户id
 * @description 获取好友列表
 */
export const getFriendList = (user_id:string)=>{
  return request({
    url: `/relation/friend/list?user_id=${user_id}`,
    method: "get",
  })
}

