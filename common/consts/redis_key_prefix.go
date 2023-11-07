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
	// VideoTimeScore 视频时序排行
	VideoTimeScore = ProjectNamePrefix + "videoTimeScore"

	// VideoEveryUserTimeScore  每个用户的所有视频时序排行
	VideoEveryUserTimeScore = ProjectNamePrefix + "videoEveryUserTimeScore:"
	// VideoEveryUserHotScore  每个用户的所有视频热度排行
	VideoEveryUserHotScore = ProjectNamePrefix + "videoEveryUserHotScore:"

	// VideoPartitionTimeScore 每个分区的所有视频时序排序
	VideoPartitionTimeScore = ProjectNamePrefix + "videoPartitionTimeScore:"
	// VideoPartitionHotScore 每个分区的所有视频时序排序
	VideoPartitionHotScore = ProjectNamePrefix + "videoPartitionHotScore:"
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

	// UserCollectPrefix 用户收藏列表
	UserCollectPrefix = ProjectNamePrefix + "userCollect:"
	// VideoCollectCountPrefix 视频收藏数量
	VideoCollectCountPrefix = ProjectNamePrefix + "videoCollectCount:"

	// UserFollowPrefix 用户关注列表
	UserFollowPrefix = ProjectNamePrefix + "userFollow:"
	// UserFollowerPrefix 用户粉丝列表
	UserFollowerPrefix = ProjectNamePrefix + "userFollower:"
)
