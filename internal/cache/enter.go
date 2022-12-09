package cache

import (
	"github.com/muesli/cache2go"
	"time"
)

type enter int

const Enter = enter(0)

func (enter) Table(name string) Memory {
	return Memory{
		table: cache2go.Cache(name),
	}
}

func (receiver enter) Add(key string, value any, lifeSpan time.Duration) {
	receiver.Table("DEFAULT").Add(key, value, lifeSpan)
}

func (receiver enter) Get(key string) (value any, isExist bool) {
	return receiver.Table("DEFAULT").Get(key)
}

func (receiver enter) Delete(key string) {
	receiver.Table("DEFAULT").Delete(key)
}

func (receiver enter) IsExist(key string) bool {
	return receiver.Table("DEFAULT").IsExist(key)
}

func (receiver enter) Clear() {
	receiver.Table("DEFAULT").Clear()
}

func (receiver enter) Size() int {
	return receiver.Table("DEFAULT").Size()
}
