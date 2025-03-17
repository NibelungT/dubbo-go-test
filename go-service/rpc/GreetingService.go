package rpc

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/client"
	dubbo_go_hessian2 "github.com/apache/dubbo-go-hessian2"
)

/**
 * @Author: 51130
 * @Date: 2025-03-14 15:56
 * @Description: GreetingService.go - 描述该文件的功能
 * @Version: 1.0
 */

type GreetRequest struct {
	Name string
}

func (x *GreetRequest) JavaClassName() string {
	return "org.example.api.GreetRequest"
}

type GreetResponse struct {
	Greeting string
}

func (x *GreetResponse) JavaClassName() string {
	return "org.example.api.GreetResponse"
}

type Result struct {
	Code int

	Message string

	Data interface{}
}

func (r *Result) JavaClassName() string {

	return "org.example.api.RpcResult"
}

func init() {
	dubbo_go_hessian2.RegisterPOJO(new(GreetResponse))
	dubbo_go_hessian2.RegisterPOJO(new(Result))
}

// 调用不同服务的方法
func SayHi(param string) (string, error) {

	d, err := GetDubboClient()
	if err != nil {
		return "", err
	}

	req, err := d.cli.Dial(
		"org.example.api.GreetingsService",
		client.WithVersion("1.0.0"),
		client.WithGroup("dubbo"),
	)
	if err != nil {
		return "", err
	}

	var response string
	err = req.CallUnary(context.Background(), []interface{}{param}, &response, "sayHi")
	if err != nil {
		return "", err
	}

	return response, nil
}

func Greet(param *GreetRequest) (*GreetResponse, error) {
	d, err := GetDubboClient()
	if err != nil {
		return nil, err
	}

	req, err := d.cli.Dial(
		"org.example.api.GreetingsService",
		client.WithVersion("1.0.0"),
		client.WithGroup("dubbo"),
	)
	if err != nil {
		return nil, err
	}

	response := new(GreetResponse)
	err = req.CallUnary(context.Background(), []interface{}{param}, &response, "greet")
	if err != nil {
		return nil, err
	}

	return response, nil
}

func SayHiGeneric(param *GreetRequest) (*GreetResponse, error) {
	d, err := GetDubboClient()
	if err != nil {
		return nil, err
	}

	req, err := d.cli.Dial(
		"org.example.api.GreetingsService",
		client.WithVersion("1.0.0"),
		client.WithGroup("dubbo"),
	)
	if err != nil {
		return nil, err
	}

	response := new(Result)
	err = req.CallUnary(context.Background(), []interface{}{param}, &response, "sayHiGeneric")
	if err != nil {
		return nil, err
	}

	if Result, ok := response.Data.(GreetResponse); ok {

		return &Result, nil
	}

	return nil, nil
}
