package ezcache

import (
	"github.com/loebfly/ezgin/internal/cache"
	"github.com/muesli/cache2go"
	"time"
)

func Table(name string) cache.Memory {
	return cache.Memory{
		Table: cache2go.Cache(name),
	}
}

func Add(key string, value any, lifeSpan time.Duration) {
	Table("DEFAULT").Add(key, value, lifeSpan)
}

func Get(key string) (value any, isExist bool) {
	return Table("DEFAULT").Get(key)
}

func Delete(key string) {
	Table("DEFAULT").Delete(key)
}

func IsExist(key string) bool {
	return Table("DEFAULT").IsExist(key)
}

func Clear() {
	Table("DEFAULT").Clear()
}

func Size() int {
	return Table("DEFAULT").Size()
}
