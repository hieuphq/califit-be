package log

import (
	"context"

	"github.com/goadesign/goa"
)

func LogError(ctx context.Context, msg string, keyvals ...interface{}) {
	goa.LogError(ctx, msg, keyvals...)
}
