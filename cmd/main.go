package main

import (
    "fmt"
    "avito_task_segments/pkg/segments"
    "avito_task_segments/pkg/users_segments"
    "avito_task_segments/pkg/common/db"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    // add env variables as needed
    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    fmt.Println(port, dbUrl)

    router := gin.Default()
    dbHandler := db.Init(dbUrl)

    segments.RegisterRoutes(router, dbHandler)
    users_segments.RegisterRoutes(router, dbHandler)

    router.GET("/", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "port":  port,
            "dbUrl": dbUrl,
        })
    })

    router.Run(port)

}