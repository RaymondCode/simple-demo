package video

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tikstart/common/utils"
	"tikstart/http/internal/svc"
	"tikstart/http/internal/types"
	"tikstart/http/schema"
	"tikstart/rpc/user/userClient"
	"tikstart/rpc/video/videoClient"
	"time"
)

type GetVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoListLogic) GetVideoList(req *types.GetVideoListRequest) (resp *types.GetVideoListResponse, err error) {
	userClaims, err := utils.ParseToken(req.Token, l.svcCtx.Config.JwtAuth.Secret)
	var userId int64 = 0
	if err == nil {
		userId = userClaims.UserId
	}
	var LatestTime int64 = 0
	if req.LatestTime == 0 {
		LatestTime = time.Now().Unix()
	} else {
		LatestTime = req.LatestTime / 1000
	}
	GetVideoListResponse, err := l.svcCtx.VideoRpc.GetVideoList(l.ctx, &videoClient.GetVideoListRequest{
		Num:        20,
		LatestTime: LatestTime,
	})
	if err != nil {
		return nil, schema.ServerError{
			ApiError: schema.ApiError{
				StatusCode: 500,
				Code:       50000,
				Message:    "Internal Server Error",
			},
			Detail: err,
		}
	}
	resp = &types.GetVideoListResponse{}
	resp.BasicResponse = types.BasicResponse{
		StatusCode: 0,
		StatusMsg:  "Success",
	}
	if len(GetVideoListResponse.VideoList) != 0 {
		resp.Next_time = GetVideoListResponse.VideoList[len(GetVideoListResponse.VideoList)-1].CreateTime
	}
	for _, v := range GetVideoListResponse.VideoList {
		//获取作者信息和关注情况
		GetUserInfoResponse, err := l.svcCtx.UserRpc.QueryById(l.ctx, &userClient.QueryByIdRequest{
			UserId: v.AuthorId,
		})
		if err != nil {
			return nil, schema.ServerError{
				ApiError: schema.ApiError{
					StatusCode: 500,
					Code:       50000,
					Message:    "Internal Server Error",
				},
				Detail: err,
			}
		}
		//获取视频收藏状态
		isFavorite := false
		if userId != 0 {
			IsFavoriteVideoReply, err := l.svcCtx.VideoRpc.IsFavoriteVideo(l.ctx, &videoClient.IsFavoriteVideoRequest{
				UserId:  userId,
				VideoId: v.Id,
			})
			if err != nil {
				return nil, schema.ServerError{
					ApiError: schema.ApiError{
						StatusCode: 500,
						Code:       50000,
						Message:    "Internal Server Error",
					},
					Detail: err,
				}
			}
			isFavorite = IsFavoriteVideoReply.IsFavorite
		}
		resp.VideoList = append(resp.VideoList, types.Video{
			Id:    v.Id,
			Title: v.Title,
			Author: types.User{
				Id:   v.AuthorId,
				Name: GetUserInfoResponse.Username,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isFavorite,
		})
	}

	return resp, nil
}
