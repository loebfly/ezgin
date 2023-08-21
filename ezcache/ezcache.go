package ezcache

import (
	"github.com/loebfly/ezgin/internal/cache"
	"github.com/muesli/cache2go"
	"time"
)

const (
	defaultTable    = "DEFAULT"
	defaultLifeSpan = 5 * time.Minute
)

func Table(name string) cache.Memory {
	return cache.Memory{
		Table: cache2go.Cache(name),
	}
}

func Add(key string, value any, lifeSpan ...time.Duration) {
	if len(lifeSpan) == 0 {
		lifeSpan = []time.Duration{defaultLifeSpan}
	}
	Table(defaultTable).Add(key, value, lifeSpan[0])
}

func Get(key string) (value any, isExist bool) {
	return Table(defaultTable).Get(key)
}

func Delete(key string) {
	Table(defaultTable).Delete(key)
}

func IsExist(key string) bool {
	return Table(defaultTable).IsExist(key)
}

func Clear() {
	Table(defaultTable).Clear()
}

func Size() int {
	return Table(defaultTable).Size()
}
