package runtime

import (
	redisDataStore "github.com/gmtek/goflow/core/redis-datastore"
	"github.com/gmtek/goflow/core/sdk"
)

func initDataStore(redisURI string) (dataStore sdk.DataStore, err error) {
	dataStore, err = redisDataStore.GetRedisDataStore(redisURI)
	return dataStore, err
}
