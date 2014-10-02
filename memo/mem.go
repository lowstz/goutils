package memo

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func SetMemcache(k string, v string, host string) {
	mc := memcache.New(host)
	mc.Set(&memcache.Item{Key: k, Value: []byte(v)})

}

func GetMemcache(k string, host string) string {
	mc := memcache.New(host)
	it, err := mc.Get(k)
	if err != nil {
		return ""
	}
	return string(it.Value)
}
