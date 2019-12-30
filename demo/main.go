package main

import (
	"flag"
	"fmt"

	log "demo/log"

	conf "demo/conf"
	grpc "demo/server/grpc"
	http "demo/server/http"

	_ "demo/router/grpc"
	_ "demo/router/http"
)

var (
	_version   = "unknown"
	_gitCommit = "unknown"
	_goVersion = "unknown"
	_buildTime = "unknown"
	_osArch    = "unknown"
)

func init() {
	// flags
	flag.StringVar(&conf.HttpAddr, "http-addr", conf.GetEnv("HttpAddr", "0.0.0.0:8080"), "http服务地址")
	flag.StringVar(&conf.GrpcAddr, "grpc-addr", conf.GetEnv("GrpcAddr", "0.0.0.0:5000"), "grpc服务地址")

	// log
	fmt.Printf(`Application Info:   
   Version:          %s
   Go version:       %s
   Git commit:       %s
   Built:            %s
   OS/Arch:          %s`, _version, _goVersion, _gitCommit, _buildTime, _osArch)
	fmt.Println()
}

func main() {
	flag.Parse()

	errc := make(chan error)

	// http server
	{
		log.Info("http server is running...", log.Field("http-addr", conf.HttpAddr))
		go http.Run(conf.HttpAddr, errc)
	}

	// grpc server
	{
		log.Info("grpc server is running...", log.Field("grpc-addr", conf.GrpcAddr))
		go grpc.Run(conf.GrpcAddr, errc)
	}

	log.Error("Exit", log.Field("error", <-errc))
}
