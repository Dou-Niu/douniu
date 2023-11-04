package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/errorx"
	"douniu/server/video/model"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	GetVideoListInfo *GetVideoListInfoLogic
}

func NewSearchVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchVideoLogic {
	return &SearchVideoLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		GetVideoListInfo: NewGetVideoListInfoLogic(ctx, svcCtx),
	}
}

func (l *SearchVideoLogic) SearchVideo(in *pb.SearchVideoReq) (*pb.FeedResp, error) {
	esclient := l.svcCtx.ESClient
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("title", in.KeyWords))
	// 设置分页参数
	searchSource.From((int(in.Page) - 1) * consts.DefaultSizeOfPage) // 从第0条记录开始
	searchSource.Size(consts.DefaultSizeOfPage)                      // 每页返回10条记录

	// 设置排序参数
	searchSource.Sort("create_time", false) // 逆序排序

	/* this block will basically print out the ES query */
	queryStr, err1 := searchSource.Source()
	_, err2 := json.Marshal(queryStr)
	if err1 != nil || err2 != nil {
		logx.Error("[esclient][GetResponse]err during query marshal=", err1, err2)
		return nil, errors.Wrapf(errorx.NewDefaultError("es序列化错误"), "es序列化错误 Req：%v", in)

	}
	/* until this block */

	searchService := esclient.Search().Index(consts.EsVideoIndex).SearchSource(searchSource)
	searchResult, err := searchService.Do(l.ctx)
	if err != nil {
		logc.Error(l.ctx, "[ProductsES][GetPIds]Error=", err)
		return nil, errors.Wrapf(errorx.NewDefaultError("[ProductsES][GetPIds]Error="+err.Error()), "es查询错误 Req：%v", in)

	}
	videoIds := make([]int64, 0)
	for _, hit := range searchResult.Hits.Hits {
		sourceBytes, err := hit.Source.MarshalJSON()
		if err != nil {
			fmt.Println("[Getting Students][MarshalJSON] Err=", err)
			continue
		}
		video := model.Video{}
		err = json.Unmarshal(sourceBytes, &video)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		videoIds = append(videoIds, video.Id)
	}
	if len(videoIds) == 0 {
		return &pb.FeedResp{
			IsFinal:      true,
			NextMaxValue: -1,
		}, nil
	}
	videoList, err := l.GetVideoListInfo.GetVideoListInfo(&pb.GetVideoListInfoReq{
		MeUserId:    in.MeUserId,
		VideoIdList: videoIds,
	})
	if err != nil {
		return nil, err
	}
	list := videoList.VideoList

	return &pb.FeedResp{
		VideoList:    list,
		NextMaxValue: in.Page + 1,
	}, nil

}
