package log

import (
	"context"

	"go.elastic.co/apm"
)

func NewElasticAPM(c context.Context, name, spanType string) {
	span, _ := apm.StartSpan(c, name, spanType)
	defer span.End()
}
