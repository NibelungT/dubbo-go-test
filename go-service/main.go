package main

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"go-service/rpc"
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	hi, err := rpc.SayHi("Nibelung")
	if err == nil {
		logger.Info("sayhi接口的出参为：" + hi)
	}

	greet, err := rpc.Greet(&rpc.GreetRequest{Name: "Nibelung"})
	if err == nil && greet != nil {
		logger.Info("greet接口的出参为：" + greet.Greeting)
	}

	generic, err := rpc.SayHiGeneric(&rpc.GreetRequest{Name: "Nibelung"})
	if err == nil && generic != nil {
		logger.Info("sayHiGeneric接口的出参为：" + generic.Greeting)
	}

}
