package cache

import (
	"github.com/loebfly/klogs"
	"github.com/muesli/cache2go"
	"time"
)

type Memory struct {
	table *cache2go.CacheTable
}

func (receiver Memory) Add(key string, value interface{}, lifeSpan time.Duration) {
	receiver.table.Add(key, lifeSpan, value)
}

func (receiver Memory) Get(key string) (value interface{}, isExist bool) {
	item, err := receiver.table.Value(key)
	if err != nil {
		return nil, false
	}
	return item.Data(), true
}

func (receiver Memory) Delete(key string) {
	_, err := receiver.table.Delete(key)
	if err != nil {
		klogs.C("cache").Error("delete cache error", err)
	}
}

func (receiver Memory) IsExist(key string) bool {
	return receiver.table.Exists(key)
}

func (receiver Memory) Clear() {
	receiver.table.Flush()
}

func (receiver Memory) Size() int {
	return receiver.table.Count()
}
