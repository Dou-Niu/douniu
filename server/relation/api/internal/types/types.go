// Code generated by goctl. DO NOT EDIT.
package types

type RelationActionRequest struct {
	ToUserId   int64 `json:"to_user_id" vd:"$>0;msg:'to_user_id error'"`
	ActionType int64 `json:"action_type" vd:"$==1||$==2;msg:'action_type error'"`
}

type RelationActionResponse struct {
}

type RelationFollowListRequest struct {
	UserId  int64 `form:"user_id" vd:"$>0;msg:'user_id error'"`
	PageNum int64 `form:"page_num"`
}

type RelationFollowListResponse struct {
	UserList []*User `json:"user_list"`
}

type RelationFollowerListRequest struct {
	UserId  int64 `form:"user_id" vd:"$>0;msg:'user_id error'"`
	PageNum int64 `form:"page_num"`
}

type RelationFollowerListResponse struct {
	UserList []*User `json:"user_list"`
}

type FriendListRequest struct {
	UserId int64 `form:"user_id" vd:"$>0;msg:'user_id error'"`
}

type FriendListResponse struct {
	FriendList []*FriendUser `json:"user_list"`
}

type User struct {
	Id              int64  `json:"id"`
	Username        string `json:"name"`
	Avatar          string `json:"avatar"`
	FollowCount     int64  `json:"follow_count"`
	TotalFavorited  int64  `json:"total_favorited"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
	FollowerCount   int64  `json:"follower_count"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
	IsFollow        bool   `json:"is_follow"`
}

type FriendUser struct {
	User
	MsgType int64  `json:"msg_type"  vd:"$==1||$==2;msg:'MsgType error'"`
	Message string `json:"message,optional"`
}
