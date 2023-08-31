package db

import (
    "log"
    "time"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{
        NowFunc: func() time.Time {
            ti, _ := time.LoadLocation("Europe/Moscow")
            return time.Now().In(ti)
        },
    })

    if err != nil {
        log.Fatalln(err)
    }

    return db
}