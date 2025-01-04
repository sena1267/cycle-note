package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/sena1267/cycle-note/config"
	"github.com/sena1267/cycle-note/util/appctx"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

// TODO: user_idをミドルウェア等で指定する(全クエリにwhere句を設定する)

type DBClient struct {
	db *bun.DB
}

func NewDBClient(cfg config.DB) (*DBClient, error) {
	bunDB, err := NewBunDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init bunDB. %s", err)
	}

	return &DBClient{db: bunDB}, nil
}

func NewBunDB(cfg config.DB) (*bun.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("failed to load location. %s", err)
	}
	mysqlConfig := mysql.Config{
		DBName: cfg.Database,
		User:   cfg.UserName,
		Passwd: cfg.Password,
		// ↓コンテナ内のアプリケーションから見たときのホスト名は "db"
		Addr:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	sqlDB, err := sql.Open("mysql", mysqlConfig.FormatDSN())

	if err != nil {
		return nil, fmt.Errorf("failed to open a database. %w", err)
	}

	db := bun.NewDB(sqlDB, mysqldialect.New())

	return db, nil
}

func (client *DBClient) DB() *bun.DB {
	return client.db
}

func ApplyFixedWhere[T interface {
	*bun.SelectQuery | *bun.InsertQuery
}](ctx context.Context, query T) T {
	userID := appctx.User(ctx).ID
	switch q := any(query).(type) {
	case *bun.SelectQuery:
		return any(q.Where("user_id = ?", userID)).(T)
	case *bun.InsertQuery:
		return any(q.Where("user_id = ?", userID)).(T)
	default:
		return query
	}
}
