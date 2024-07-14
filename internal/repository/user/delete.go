package user

import (
	"context"
	"fmt"
)

func (r *repo) Delete(_ context.Context, id int64) error {
	fmt.Println("id:", id)
	return nil
}
