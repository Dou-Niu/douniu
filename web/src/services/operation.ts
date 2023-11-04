import request from "@/utils/request";

/**
 * @param video_id 视频Id
 * @param action_type 操作类型 1为点赞，2为取消点赞
 * @description 点赞
 */
export const favorite = (video_id:any,action_type:any) =>{
  return request({
    url:"/favorite/action",
    method: "post",
    data:{
      video_id,
      action_type,
    }
  })
}

/**
 * @param user_id 用户id
 * @param page 页数
 * @description 点赞列表
 */
export const getFavorList = (user_id:any,page:any) =>{
  return request({
    url:`/favorite/list?user_id=${user_id}&page_num=${page}`,
    method: "get",
  })
}

/**
 * @param video_id 视频Id
 * @param action_type 操作类型 1为收藏，2为取消收藏
 * @description 收藏
 */
export const collection = (video_id:any,action_type:any,partition:any) =>{
  return request({
    url:"/collection/action",
    method: "post",
    data:{
      video_id,
      action_type,
      partition
    }
  })
}

/**
 * @param user_id 用户id
 * @param page 页数
 * @description 收藏列表
 */
export const getCollectionList = (user_id:any,page:any) =>{
  return request({
    url:`/collection/list?user_id=${user_id}&page_num=${page}`,
    method: "get",
  })
}

/**
 * @param video_id 视频id
 * @param action_type 操作类型 1为评论，2为删除评论
 * @param content 可选，用户填写的评论内容，在action_type=1的时候使用
 * @param parent_id 评论id，0是评论视频
 * @param comment_id 可选，要删除的评论id，在action_type=2的时候使用
 */
export const comment = (data:any) =>{
  return request({
    url:"/collection/action",
    method: "post",
    data
  })
}

/**
 * @param video_id 视频id
 * @param last_comment_id 是上一页的最后一条评论id，第一次打开评论，那就是0
 * @description 视频评论
 */
export const getCommentsList = (video_id:any,last_comment_id:any) =>{
  return request({
    url:`/comment/list?video_id=${video_id}&last_comment_id=${last_comment_id}`,
    method: "get",
  })
}

/**
 * @param last_comment_id 是上一页的最后一条评论id，第一次打开评论，那就是0
 * @description 查看回复评论
*/
export const getReplyList = (comment_id:any) =>{
  return request({
    url:`/comment/detail?comment_id=${comment_id}`,
    method: "get",
  })
}

/**
 * @param video_id
 * @headers Authorization
 * @description 分享视频
 */
export const shareVideo = (video_id:any) => {
  return request({
    url: `/video/share?video_id=${video_id}}`,
    method: "get",
  })
}
