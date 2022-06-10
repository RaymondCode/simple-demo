package service

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/common"
	"github.com/warthecatalyst/douyin/dao"
	"github.com/warthecatalyst/douyin/logx"
	"github.com/warthecatalyst/douyin/model"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

func PublishVideoInfo(data *multipart.FileHeader, userId int64, title string) error {
	return newPublishVideoInfoFlow(data, userId, title).Do()
}

type publishVideoInfoFlow struct {
	data   *multipart.FileHeader
	userId int64
	title  string
}

func newPublishVideoInfoFlow(data *multipart.FileHeader, userId int64, title string) *publishVideoInfoFlow {
	return &publishVideoInfoFlow{
		data:   data,
		userId: userId,
		title:  title,
	}
}

func (p *publishVideoInfoFlow) Do() error {
	fileName := p.data.Filename
	//首先检查video扩展名和大小
	if !common.CheckFileExt(fileName) {
		logx.DyLogger.Error("wrong input video type")
		return errors.New("video Error")
	}
	if !common.CheckFileMaxSize(p.data.Size) {
		logx.DyLogger.Error("file extends the maximum value")
		return errors.New("video Error")
	}

	//然后把文件保存至本地
	err := p.saveFile(common.GlobalSavePath)
	if err != nil {
		logx.DyLogger.Error("Saving goes wrong")
		return errors.New("save Error")
	}
	//截取视频的第一帧作为cover
	saveDir := path.Join(common.GlobalSavePath, strconv.FormatInt(p.userId, 10))
	saveVideo := saveDir + "/" + p.data.Filename
	coverName := common.GetFileNameWithOutExt(fileName) + "_cover" + ".jpeg"
	saveCover := saveDir + "/" + coverName
	err = common.ExtractCoverFromVideo(saveVideo, saveCover)
	if err != nil {
		logx.DyLogger.Error("Saving goes wrong")
		return errors.New("save Error")
	}

	//上传视频和封面
	logx.DyLogger.Info("Saving Complete, in Upload")
	err = p.uploadServer()
	if err != nil {
		logx.DyLogger.Error("An error occurs in upload the file to Aliyun")
		return errors.New("upload Error")
	}
	err = p.uploadCoverToServer(saveCover, coverName)
	if err != nil {
		logx.DyLogger.Error("An error occurs in upload the file to Aliyun")
		return errors.New("upload Error")
	}

	rand.Seed(time.Now().UnixNano())

	//调用dao层函数操作数据库添加对应的Video字段
	video := model.Video{
		VideoID:       rand.Int63n(100000), //随机生成VideoId,之后可以进行调整
		VideoName:     p.title,
		UserID:        p.userId,
		FavoriteCount: 0,
		CommentCount:  0,
		PlayURL:       common.GetUploadURL(p.userId, p.data.Filename),
		CoverURL:      common.GetUploadURL(p.userId, coverName),
	}

	err = dao.NewPublishDaoInstance().AddVideo(&video)
	if err != nil {
		logx.DyLogger.Error("An error occurs in ")
		return errors.New("database Error")
	}
	return nil
}

func (p *publishVideoInfoFlow) saveFile(savePath string) error {
	userSavePath := path.Join(savePath, strconv.FormatInt(p.userId, 10))
	if flag, _ := common.PathExists(userSavePath); !flag {
		err := os.Mkdir(userSavePath, os.ModePerm)
		if err != nil {
			logx.DyLogger.Error("Error in making directory")
			return err
		}
	}
	src, err := p.data.Open()
	if err != nil {
		logx.DyLogger.Error("Error in Saving", err)
		return err
	}
	defer src.Close()

	out, err := os.Create(userSavePath + "/" + p.data.Filename)
	if err != nil {
		logx.DyLogger.Error("Error in Saving", err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (p *publishVideoInfoFlow) uploadServer() error {
	Endpoint := common.UploadServerURL
	AccessKeyID := common.UploadAccessKeyID         // AccessKeyId
	AccessKeySecret := common.UploadAccessKeySecret // AccessKeySecret

	client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		logx.DyLogger.Error("Error In upload1:", err)
		return err
	}

	// 指定bucket
	bucket, err := client.Bucket(common.UploadServerBucket)
	if err != nil {
		logx.DyLogger.Error("Error In upload2:", err)
		return err
	}

	src, err := p.data.Open()
	if err != nil {
		logx.DyLogger.Error("Error in upload3:", err)
		return err
	}
	defer src.Close()

	// 先将文件流上传至douyin_test目录下
	uploadPath := common.UploadBucketDirectory + "/" + strconv.FormatInt(p.userId, 10) + "/" + p.data.Filename
	err = bucket.PutObject(uploadPath, src)
	if err != nil {
		logx.DyLogger.Error("Error in upload4:", err)
		return err
	}

	logx.DyLogger.Info("file upload success")
	return nil
}

//把封面上传到云端
func (p *publishVideoInfoFlow) uploadCoverToServer(filePath, fileName string) error {
	client, err := oss.New(common.UploadServerURL, common.UploadAccessKeyID, common.UploadAccessKeySecret)
	if err != nil {
		logx.DyLogger.Error(err)
		return err
	}

	bucket, err := client.Bucket(common.UploadServerBucket)
	if err != nil {
		logx.DyLogger.Error(err)
		return err
	}

	uploadPath := common.UploadBucketDirectory + "/" + strconv.FormatInt(p.userId, 10) + "/" + fileName
	err = bucket.PutObjectFromFile(uploadPath, filePath)
	if err != nil {
		logx.DyLogger.Error(err)
		return err
	}
	return nil
}

func PublishListInfo(userId int64) (*VideoList, error) {
	return newPublishListInfoFlow(userId).Do()
}

type publishListInfoFlow struct {
	userId int64
}

func newPublishListInfoFlow(userId int64) *publishListInfoFlow {
	return &publishListInfoFlow{userId: userId}
}

func (p *publishListInfoFlow) Do() (*VideoList, error) {
	videoIds, err := dao.NewPublishDaoInstance().GetVideoPublistList(p.userId)
	if err != nil {
		return nil, err
	}
	var videolist VideoList
	for _, videoId := range videoIds {
		user, err := videoService.getUserFromVideoId(videoId)
		if err != nil {
			return nil, err
		}
		video, err := dao.NewVideoDaoInstance().GetVideoFromId(videoId)

		videolist = append(videolist, api.Video{
			Id:            videoId,
			Author:        *user,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    true, //被videolist返回的肯定点赞过了
		})
	}
	return &videolist, nil
}
