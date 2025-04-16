package ports

import (
	"context"
)

type ILoginUser interface {
	LoginUser(ctx context.Context, email, password string) (string, error)
}
