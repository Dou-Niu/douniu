// Code generated by goctl. DO NOT EDIT.
package types

type CollectionActionRequest struct {
	VideoId    int64 `json:"video_id" vd vd:"$>0;msg:'video_id error'"`
	ActionType int64 `json:"action_type" vd:"$==1||$==2;msg:'action_type error'"`
	Partition  int64 `json:"partition" vd:"$>=0;msg:'partition error'"`
}

type CollectionActionResponse struct {
}

type CollectioneListRequest struct {
	UserId  int64 `form:"user_id" vd:"$>0;msg:'user_id error'"`
	PageNum int64 `form:"page_num" vd:"$>=0;msg:'page_num error'"`
}

type CollectioneListResponse struct {
	VideoList []*Video `json:"video_list"`
}

type FavoriteLikeRequest struct {
	VideoId    int64 `json:"video_id" vd vd:"$>0;msg:'video_id error'"`
	ActionType int64 `json:"action_type" vd:"$==1||$==2;msg:'action_type error'"`
	Partition  int64 `json:"partition" vd:"$>0;msg:'partition error'"`
}

type FavoriteLikeResponse struct {
}

type FavoriteListRequest struct {
	UserId  int64 `form:"user_id" vd:"$>0;msg:'user_id error'"`
	PageNum int64 `form:"page_num" vd:"$>=0;msg:'page_num error'"`
}

type FavoriteListResponse struct {
	VideoList []*Video `json:"video_list"`
}

type Video struct {
	Id              int64  `json:"id"`
	Author          *User  `json:"author" copier:"User"`
	Title           string `json:"title"`
	PlayUrl         string `json:"play_url"`
	CoverUrl        string `json:"cover_url"`
	FavoriteCount   int64  `json:"favorite_count"`
	CommentCount    int64  `json:"comment_count"`
	IsFavorite      bool   `json:"is_favorite"`
	CollectionCount int64  `json:"collection_count"`
	IsCollect       bool   `json:"is_collect"`
}

type User struct {
	Id              int64  `json:"id"`
	Nickname        string `json:"nickname"`
	Avatar          string `json:"avatar"`
	FollowCount     int64  `json:"follow_count"`
	TotalFavorited  int64  `json:"total_favorited"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
	FollowerCount   int64  `json:"follower_count"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
	IsFollow        bool   `json:"is_follow"`
	CollectionCount int64  `json:"collection_count"`
}
