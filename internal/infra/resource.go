package infra

import (
	"fmt"
	"log"
	"tx-bank/internal/config"
	"tx-bank/internal/infra/cache"
	rcache "tx-bank/internal/infra/cache/redis"
	"tx-bank/internal/infra/notification"
	notifsvc "tx-bank/internal/infra/notification/impl"
	"tx-bank/internal/infra/session"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Resources struct {
	// rdbms
	Database *sqlx.DB

	// cache
	Redis cache.Cache

	// session manager
	SessionManager session.Manager

	// notification
	NotificationService notification.Service
}

func NewResources(conf *config.Config) (*Resources, error) {
	// init db
	psqlsource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.DBName)
	psqlx, err := sqlx.Connect("postgres", psqlsource)
	if err != nil {
		log.Fatalf("Failed create database conn, with err:%+v", err)
		return nil, err
	}

	// init redis
	rdsCli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if rdsCli == nil {
		log.Fatalf("failed to init redis")
	}
	rdsCache := rcache.New(rdsCli)

	// init notification service
	notifSvc := notifsvc.New(conf.SMTP)

	// init jwt session manager
	jwtManager := session.NewJWTManager(string(conf.JWT.Secret), rdsCli)

	return &Resources{
		Database:            psqlx,
		Redis:               rdsCache,
		SessionManager:      jwtManager,
		NotificationService: notifSvc,
	}, nil
}
