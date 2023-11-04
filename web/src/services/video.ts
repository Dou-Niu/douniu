import request from "@/utils/request";

/**
 * @param video_url 视频链接
 * @param cover_url 封面链接
 * @param title 视频标题
 * @param partition 分区
 * @headers Authorization
 * @description 上传视频
 */
export const uploadVideo = (data:any) => {
  return request({
    url: "/video/publish",
    method: "post",
    data: data,
  })
}

/**
 * @param key_words 视频链接
 * @param page 页码
 * @headers Authorization
 * @description 上传视频
 */
export const searchVideo = (params:any) => {
  return request({
    url: "/video/search",
    method: "get",
    params
  })
}

/**
 * @param max_hot
 * @headers Authorization
 * @description 热门视频
 */
export const getHotVideos = (max_hot:number) => {
  return request({
    url: `/video/feed/hot?max_hot=${max_hot}`,
    method: "get",
  })
}

/**
 * @param latest_time
 * @headers Authorization
 * @description 关注视频
 */
export const getFollowVideos = (latest_time:number) => {
  return request({
    url: `/video/feed/hot?latest_time=${latest_time}`,
    method: "get",
  })
}

/**
 * @param max_value
 * @parameter user_id 用户Id
 * @param sort 排序方法
 * @headers Authorization
 * @description 某个人的视频
 */
export const getFeedVideos = (params:any) => {
  const {max_value,user_id,sort} = params
  return request({
    url: `/video/feed/user?user_id=${user_id}&max_hot=${max_value}&sort=${sort}`,
    method: "get",
  })
}

/**
 * @param max_value
 * @parameter partition 分区
 * @param sort 排序方法
 * @headers Authorization
 * @description 分区的视频
 */
export const getPartVideos = (params:any) => {
  const {max_value,partition,sort} = params
  return request({
    url: `/video/feed/partition?max_value=${max_value}&sort=${sort}&partition=${partition}`,
    method: "get",
  })
}

/**
 * @param video_id
 * @headers Authorization
 * @description 删除视频
 */
export const deleteVideo = (video_id:any,partition:number) => {
  return request({
    url: `/video/delete?video_id=${video_id}&partition=${partition}}`,
    method: "delete",
  })
}

/**
 * @param video_id
 * @headers Authorization
 * @description 查看单个视频
 */
export const lookVideo = (video_id:any) => {
  return request({
    url: `/video?video_id=${video_id}`,
    method: "delete",
  })
}