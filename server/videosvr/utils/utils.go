package utils

import (
	"bytes"
	"fmt"
	"github.com/Keqing-win/camp_tiktok/pkg/pb"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os/exec"
	"path/filepath"
	"strings"
	"videosvr/config"
	"videosvr/log"
	"videosvr/utils/otgrpc"
)

var (
	FavoriteSvrClient pb.FavoriteServiceClient
	RelationSvrClient pb.RelationServiceClient
	CommentSvrClient  pb.CommentServiceClient
)

func NewSvrConn(svrName string) (*grpc.ClientConn, error) {
	consulInfo := config.GetGlobalConfig().ConsulConfig
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, svrName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		log.Errorf("NewSvrConn with svrname %s err:%v", svrName, err)
		return nil, err
	}
	return conn, nil
}

// func NewUserSvrClient(svrName string) pb.UserServiceClient {
// 	conn, err := NewSvrConn(svrName)
// 	if err != nil {
// 		return nil
// 	}
// 	return pb.NewUserServiceClient(conn)
// }

func NewCommentSvrClient(svrName string) pb.CommentServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewCommentServiceClient(conn)
}

func NewRelationSvrClient(svrName string) pb.RelationServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewRelationServiceClient(conn)
}

func NewFavoriteSvrClient(svrName string) pb.FavoriteServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewFavoriteServiceClient(conn)
}

func GetRelationSvrClient() pb.RelationServiceClient {
	return RelationSvrClient
}

func GetFavoriteSvrClient() pb.FavoriteServiceClient {
	return FavoriteSvrClient
}

func GetCommentSvrClient() pb.CommentServiceClient {
	return CommentSvrClient
}

func InitSvrConn() {
	RelationSvrClient = NewRelationSvrClient(config.GetGlobalConfig().SvrConfig.RelationSvrName)
	FavoriteSvrClient = NewFavoriteSvrClient(config.GetGlobalConfig().SvrConfig.FavoriteSvrName)
	CommentSvrClient = NewCommentSvrClient(config.GetGlobalConfig().SvrConfig.CommentSvrName)
}

func GetImageFile(videoPath string) (string, error) {
	temp := strings.Split(videoPath, "/")
	videoName := temp[len(temp)-1]
	b := []byte(videoName)
	videoName = string(b[:len(b)-3]) + "jpg"
	picPath := config.GetGlobalConfig().MinioConfig.PicPath
	picName := filepath.Join(picPath, videoName)
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "1", "-f", "image2", "-t", "0.01", "-y", picName)
	log.Info(cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Errorf("cmd.Run() failed with %s\n", err.Error(), stderr.String())
		return "", err
	} else {
		log.Info(out.String())
	}

	//err := cmd.Run()
	//if err != nil {
	//	log.Errorf("cmd.Run() failed with %s\n", err)
	//	return "", err
	//}
	return picName, nil
}
