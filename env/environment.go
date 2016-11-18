package env

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"

	"github.com/lambda-engine/engine/audit"
)

const (
	DS_ENVIRONMENT string = "ENVIRONMENT"
)

type EnvironmentValue struct {
	Key          string
	Value        string
	DefaultValue string
}

func GetValue(ctx context.Context, key string) string {
	value := ""
	cacheKey := "env." + key

	cache, err := memcache.Get(ctx, cacheKey)
	if err != nil {
		// cache missed
		k := datastore.NewKey(ctx, DS_ENVIRONMENT, key, 0, nil)
		var e EnvironmentValue

		err = datastore.Get(ctx, k, &e)
		if err == nil {
			value = e.Value

			// update the cache
			cache := memcache.Item{}
			cache.Key = cacheKey
			cache.Value = []byte(value)
			memcache.Set(ctx, &cache)
		} else {
			audit.Critical(ctx, "environment.value.missing", key)
		}
	} else {
		value = string(cache.Value)
	}

	return value

}
