package data

import (
	"context"
	"log"

	"github.com/adrian-lin-1-0-0/go-study/clean_architecture/biz"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

type contextTxKey struct{}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

func NewData(db *gorm.DB) (*Data, func(), error) {
	d := &Data{
		db: db,
	}
	return d, func() {
	}, nil
}

// NewDB gorm Connecting to a Database
func NewDB() *gorm.DB {

	db, err := gorm.Open(mysql.Open("mysql-connection"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&User{}, &Card{}); err != nil {
		log.Fatal(err)
	}
	return db
}
