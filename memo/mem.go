package memo

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var mc *memcache.Client

func SetMemcache(k string, v string, host string) {
	if mc == nil {
		mc = memcache.New(host)
	}
	mc.Set(&memcache.Item{Key: k, Value: []byte(v)})

}

func GetMemcache(k string, host string) string {
	if mc == nil {
		mc = memcache.New(host)

	}
	it, err := mc.Get(k)
	if err != nil {
		return ""
	}
	return string(it.Value)
}
