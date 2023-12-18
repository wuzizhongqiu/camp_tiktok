package main

import (
	"fmt"
	"github.com/Keqing-win/camp_tiktok/pkg/pb"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
	"usersvr/config"
	"usersvr/log"
	"usersvr/middlerware/cache"
	"usersvr/middlerware/consul"
	"usersvr/middlerware/db"
	"usersvr/middlerware/lock"
	"usersvr/service"
)

func Init() {
	err := config.Init()
	if err != nil {
		log.Fatalf("init config failed, err:%v\n", err)
	}
	log.InitLog()
	log.Info("log init success")
}

func Run() error {
	//服务器监听
	cfg := config.GetGlobalConfig().SvrConfig
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("Listen: error %v", err)
		return fmt.Errorf("listen: error %v", err)
	}
	//启动grpc server
	server := grpc.NewServer()
	//注册grpc server
	pb.RegisterUserServiceServer(server, &service.UserService{})
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	//注册服务到consul中
	//获取consul
	cfgConsult := config.GetGlobalConfig().ConsulConfig
	consultClient := consul.NewRegistryClient(cfgConsult.Host, cfgConsult.Port)
	//注册服务到consul中
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	err = consultClient.Registry(cfg.Host, cfg.Port, cfg.Name, cfgConsult.Tags, serviceID)
	if err != nil {
		log.Fatalf("consul registry: error %v", err)
		return fmt.Errorf("consul registry: error %v", err)
	}
	log.Info("Init Consul Register success")
	//启动
	log.Info("tiktok usersvr listening on %s:%d", cfg.Host, cfg.Port)
	//实现非阻塞的方式启动 gRPC 服务器。这样，主程序可以继续执行后续的代码，而不需要等待服务器完全启动和开始接受请求。
	go func() {
		err = server.Serve(listen)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()
	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 服务终止，注销 consul 服务
	if err = consultClient.DeRegister(serviceID); err != nil {
		log.Info("注销失败")
		return fmt.Errorf("注销失败")
	} else {
		log.Info("注销成功")
	}
	return nil
}

func main() {
	Init()
	defer db.CloseDb()
	defer cache.CloseRedis()
	defer log.Sync()
	defer lock.CloseRedSync()
	if err := Run(); err != nil {
		log.Errorf("UserSvr run err:%v", err)
	}
}
