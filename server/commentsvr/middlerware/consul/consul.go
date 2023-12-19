package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	Registry(address string, port int, name string, tags []string, id string) error
	DeRegister(servicedId string) error
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry) DeRegister(servicedId string) error {
	//consul cfg配置
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	//获取consul client
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceDeregister(servicedId)
	return err
}

func (r *Registry) Registry(address string, port int, name string, tags []string, id string) error {
	//consul cfg配置
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	//获取consul client
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", address, port),
		Timeout:                        "30s",
		Interval:                       "6s",
		DeregisterCriticalServiceAfter: "15s", //超过此时间自动取消注册
	}
	//生成注册对象
	registration := &api.AgentServiceRegistration{
		Name:    name,
		ID:      id,
		Port:    port,
		Address: address,
		Tags:    tags,
		Check:   check,
	}
	//将对象注册到consul中
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil

}
