package consts

const (
	UserId = "userId"

	MsgTypeRecv = 0
	MsgTypeSend = 1
	CommentAdd  = 1
	CommentDel  = 2
)

const (
	UserMachineId = iota
	ChatMachineId
	VideoMachineId
	FavoriteMachineId
)

const (
	SingleHotScore = 8640
)

const (
	Token = "token"
)

// StoreType 存储类型(表示文件存到哪里)
type StoreType int

const (
	_ StoreType = iota
	// StoreLocal : 节点本地
	StoreLocal
	// StoreMinio : Minio集群
	StoreMinio
	// StoreCOS : 腾讯云COS
	StoreCOS
)

const (
	MinioBucket = "gophertok-video"
)

const (
	CoverTemp = "/home/project/gophertok/temp/"
)

// 用户信息
const (
	DefaultAvatar          = "http://s36rnw3k2.hn-bkt.clouddn.com/douniu/%E6%96%97%E7%89%9B%E9%BB%98%E8%AE%A4%E7%94%A8%E6%88%B7%E5%A4%B4%E5%83%8F.jpg"
	DefaultBackgroundImage = "http://s36rnw3k2.hn-bkt.clouddn.com/douniu/%E6%96%97%E7%89%9B%E9%BB%98%E8%AE%A4%E7%94%A8%E6%88%B7%E8%83%8C%E6%99%AF%E5%9B%BE.jpg"
	DefaultSignature       = "这个家伙很懒什么都没留下"
)

// 点赞
const (
	FavoriteAdd = 1
	FavoriteDel = 2
)
