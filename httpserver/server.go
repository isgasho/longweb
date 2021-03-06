package httpserver

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/longweb/config"
	"github.com/devfeel/longweb/framework/log"
	"github.com/devfeel/longweb/message"
	"strconv"
)

func StartServer() error {

	//初始化DotServer
	app := dotweb.New()

	//设置dotserver日志目录
	app.SetLogPath(config.CurrentConfig.Log.FilePath)

	app.SetEnabledLog(true)
	app.UseRequestLog()

	//设置路由
	InitRoute(app)

	innerLogger := logger.GetInnerLogger()

	//启动监控服务
	pprofport := config.CurrentConfig.HttpServer.PProfPort
	app.SetPProfConfig(true, pprofport)

	if config.CurrentConfig.HttpServer.IsTLS {
		//设置TLS
		app.HttpServer.SetEnabledTLS(true, config.CurrentConfig.HttpServer.TLSCertFile, config.CurrentConfig.HttpServer.TLSKeyFile)
	}

	// 开始服务
	port := config.CurrentConfig.HttpServer.HttpPort
	if config.CurrentConfig.HttpServer.IsTLS{
		innerLogger.Debug("dotweb.StartServer[Tls] => " + strconv.Itoa(port))
	}else {
		innerLogger.Debug("dotweb.StartServer => " + strconv.Itoa(port))
	}
	err := app.StartServer(port)
	return err
}

func ReSetServer() {
	//初始化应用信息
	message.InitAppInfo()
}
