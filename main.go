package main

import (
        "flag"
        "log"

        "github.com/gin-gonic/gin"
)

func main() {

        var err error
        var Config_File string = "/opt/k9/etc/k9-client-logging-proxy.yaml"

        config_file := flag.String("config", "", "Configuration file to use")

        if *config_file != "" {
                Config_File = *config_file
        }

        log.Printf("* Using configuration file - %s\n", Config_File)

        /* Load configuration into global memory */

        LoadConfig(Config_File)

        DropPrivileges(Config.Core.Runas)

        log.Printf("Setting gin to \"%s\" mode.\n", Config.Proxy.HTTP_Mode)
        gin.SetMode(Config.Proxy.HTTP_Mode) /* 'debug', 'release' or 'test' */

        router := gin.Default()

        router.Use(HTTP_Logger())

        router.Use(Authenticate_API())

	router.POST("/CLIENT_LOGGING_NEED_URL", Process_Client_Log)

        /* Non-TLS */

        if Config.Proxy.HTTP_TLS == false {

                log.Printf("Listening for unencrypted traffic on %s.", Config.Proxy.HTTP_Listen)
                err = router.Run(Config.Proxy.HTTP_Listen)

        } else {

                /* TLS */

                log.Printf("Listening for TLS traffic on %s.", Config.Proxy.HTTP_Listen)
                err = router.RunTLS(Config.Proxy.HTTP_Listen, Config.Proxy.HTTP_Cert, Config.Proxy.HTTP_Key)
        }

        if err != nil {

                if Config.Proxy.HTTP_TLS == false {

                        log.Fatalf("Cannot bind to %s", Config.Proxy.HTTP_Listen)

                } else {

                        log.Fatalf("Cannot bind it %s or cannot open %s or %s.\n", Config.Proxy.HTTP_Listen, Config.Proxy.HTTP_Cert, Config.Proxy.HTTP_Key)

                }

        }

}

