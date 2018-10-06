package main

import (
	"apiserver/demo/config"
	"apiserver/demo/model"
	"apiserver/demo/router"
	"apiserver/demo/router/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	// 解析命令行参数
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	//for {
	//	fmt.Println(viper.GetString("runmode"))
	//	time.Sleep(4 * time.Second)
	//}

	// Create the Gin engine.
	g := gin.New()

	//设置运行模式 debug release test；不同模式对应不同的日志打印
	gin.SetMode(viper.GetString("runmode"))

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	/*
		有时候 API 进程起来不代表 API 服务器正常，
		问题：API 进程存在，但是服务器却不能对外提供服务。
		因此在启动 API 服务器时，如果能够最后做一个自检会更好些。
		添加了自检程序，在启动 HTTP 端口前 go 一个 pingServer 协程，启动 HTTP 端口后，该协程不断地 ping /sd/health 路径，如果失败次数超过一定次数，
		则终止 HTTP 服务器进程。通过自检可以最大程度地保证启动后的 API 服务器处于健康状态。
	*/
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Infof("The router has been deployed successfully.")
	}()

	// Start to listening the incoming requests.
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Infof("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
