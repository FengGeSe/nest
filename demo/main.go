package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

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
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2005-01-02 15:04:05",
	})

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
		log.WithField("http-addr", conf.HttpAddr).Info("http server is running...")
		go http.Run(conf.HttpAddr, errc)
	}

	// grpc server
	{
		log.WithField("grpc-addr", conf.GrpcAddr).Info("grpc server is running...")
		go grpc.Run(conf.GrpcAddr, errc)
	}

	log.WithField("error", <-errc).Info("Exit")
}
