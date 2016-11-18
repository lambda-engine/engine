package metrics

import (
	"golang.org/x/net/context"

	"github.com/zorkian/go-datadog-api"

	"google.golang.org/appengine/datastore"

	"github.com/lambda-engine/engine/util"
	"github.com/lambda-engine/engine/audit"

)

const (
	PREFIX    string = "banga"
	DS_METRIC string = "METRIC"
)

type Metric struct {
	Category string
	Action   string
	Label    string
	Value    int64
	Created  int64
}

var _datadog *datadog.Client

func init() {
	_datadog = datadog.NewClient("api key", "application key")
}

func Count(ctx context.Context, category, action, label string, value int) {

	now := util.Timestamp()

	v := int64(value)

	m := Metric{
		category,
		action,
		label,
		v,
		now,
	}

	key := datastore.NewIncompleteKey(ctx, DS_METRIC, nil)
	_, err := datastore.Put(ctx, key, &m)

	if err != nil {
		audit.Error(ctx, "metrics.count", err)
	}
}
