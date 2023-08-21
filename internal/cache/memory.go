package cache

import (
	"github.com/loebfly/ezgin/ezlogs"
	"github.com/muesli/cache2go"
	"time"
)

type Memory struct {
	Table *cache2go.CacheTable
}

func (receiver Memory) Add(key string, value any, lifeSpan time.Duration) {
	receiver.Table.Add(key, lifeSpan, value)
}

func (receiver Memory) Get(key string) (value any, isExist bool) {
	item, err := receiver.Table.Value(key)
	if err != nil {
		return nil, false
	}
	return item.Data(), true
}

func (receiver Memory) Delete(key string) {
	_, err := receiver.Table.Delete(key)
	if err != nil {
		ezlogs.CError("cache", "delete cache error: {}", err)
	}
}

func (receiver Memory) IsExist(key string) bool {
	return receiver.Table.Exists(key)
}

func (receiver Memory) Clear() {
	receiver.Table.Flush()
}

func (receiver Memory) Size() int {
	return receiver.Table.Count()
}
