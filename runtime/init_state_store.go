package runtime

import (
	redisStateStore "github.com/gmtek/goflow/core/redis-statestore"
	"github.com/gmtek/goflow/core/sdk"
)

func initStateStore(redisURI string) (stateStore sdk.StateStore, err error) {
	stateStore, err = redisStateStore.GetRedisStateStore(redisURI)
	return stateStore, err
}
