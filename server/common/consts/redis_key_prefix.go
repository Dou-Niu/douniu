package consts

const (
	ProjectNamePrefix = "douniu:"
	LastMessagePrefix = ProjectNamePrefix + "lastMessage:"
	VideoFavorPrefix  = ProjectNamePrefix + "videoFavor:"
)

// User服务
const (
	PhoneCode      = ProjectNamePrefix + "phoneCode:"
	UserIsRegister = ProjectNamePrefix + "isRegister:"
)

const (
	// VideoHotScore 视频热度排行
	VideoHotScore = ProjectNamePrefix + "videoHotScore"
)

const (

	// VideoCommentCountPrefix 视频评论数
	VideoCommentCountPrefix = ProjectNamePrefix + "videoCommentCount:"

	VideoRankKey        = ProjectNamePrefix + "videoRank"
	UserVideoRankPrefix = ProjectNamePrefix + "userVideoRank:"

	// UserFavoriteIdPrefix 用户点赞的视频Id列表
	UserFavoriteIdPrefix = ProjectNamePrefix + "userFavoriteId:"

	// VideoFavoritedIdPrefix 视频被点赞的用户Id列表
	VideoFavoritedIdPrefix = ProjectNamePrefix + "videoFavoritedId:"

	// UserFavoritedCountPrefix 用户获赞数
	UserFavoritedCountPrefix = ProjectNamePrefix + "userFavoritedCount:"

	UserFollowPrefix   = ProjectNamePrefix + "userFollow:"
	UserFollowerPrefix = ProjectNamePrefix + "userFollower:"
)