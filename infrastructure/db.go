package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/sena1267/cycle-note/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type DBClient struct {
	DB *bun.DB
}

func NewDBClient(cfg config.DB) (*DBClient, error) {
	bunDB, err := NewBunDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init bunDB. %s", err)
	}

	return &DBClient{DB: bunDB}, nil
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
