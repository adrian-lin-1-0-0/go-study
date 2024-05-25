# Clean Architecture

## Transaction

### 以 gorm 為例
> 參考kratos範例的作法 :
> https://github.com/go-kratos/examples/tree/main/transaction/gorm

可以把orm再封裝一層，透過context傳遞tx;
tx的取得方式,取得tx後的操作方式，都可以透過這個封裝來處理。

可以感受到這個封裝的好處,不只是實作上統一,修改也很方便。

```go
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
```