package infra

import (
	"context"
	"fmt"
	"time"
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
		return nil, fmt.Errorf("connect database: %w", err)
	}

	// init redis
	rdsCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err = rdsCli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("connect redis: %w", err)
	}
	rdsCache := rcache.New(rdsCli)

	// init notification service
	notifSvc := notifsvc.New(conf.SMTP)

	// init jwt session manager
	jwtManager := session.NewJWTManager(conf.JWT.Secret, rdsCli)

	return &Resources{
		Database:            psqlx,
		Redis:               rdsCache,
		SessionManager:      jwtManager,
		NotificationService: notifSvc,
	}, nil
}
