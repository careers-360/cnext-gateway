package main

import (
   "flag"
   "log"
   "os"
   "time"
   "github.com/gin-gonic/gin"
   "github.com/gin-contrib/cache"
   "github.com/gin-contrib/cache/persistence"
   "github.com/luraproject/lura/v2/config"
   "github.com/luraproject/lura/v2/logging"
   "github.com/luraproject/lura/v2/proxy"
   "github.com/luraproject/lura/v2/transport/http/server"
   krakendgin "github.com/luraproject/lura/v2/router/gin"
)

func main() {
    port := flag.Int("p", 0, "Port of the service")
    debug := flag.Bool("d", true, "Enable the debug")
    configFile := flag.String("c", "config.json", "Path to the configuration filename")

    flag.Parse()

    parser := config.NewParser()
    serviceConfig, err := parser.Parse(*configFile)
    if err != nil {
        log.Fatal("ERROR:", err.Error())
    }
    serviceConfig.Debug = serviceConfig.Debug || *debug
    if *port != 0 {
        serviceConfig.Port = *port
    }

    store := persistence.NewInMemoryStore(time.Minute * 60)

    logger, _ := logging.NewLogger("DEBUG", os.Stdout, "Cnext")

    routerFactory := krakendgin.NewFactory(krakendgin.Config{
        Engine:       gin.Default(),
        ProxyFactory: customProxyFactory{logger, proxy.DefaultFactory(logger)},
        Logger:       logger,
        HandlerFactory: func(configuration *config.EndpointConfig, proxy proxy.Proxy) gin.HandlerFunc {
            return cache.CachePage(store, configuration.CacheTTL, krakendgin.EndpointHandler(configuration, proxy))
        },
        RunServer:    server.RunServer,
    })

    routerFactory.New().Run(serviceConfig)
}

// customProxyFactory adds a logging middleware wrapping the internal factory
type customProxyFactory struct {
    logger  logging.Logger
    factory proxy.Factory
}

// New implements the Factory interface
func (cf customProxyFactory) New(cfg *config.EndpointConfig) (p proxy.Proxy, err error) {
    p, err = cf.factory.New(cfg)
    if err == nil {
        p = proxy.NewLoggingMiddleware(cf.logger, cfg.Endpoint)(p)
    }
    return
}