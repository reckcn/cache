package memcached

import (
	"fmt"
	"testing"
)

type ColorGroup struct {
	Id     int
	Name   string
	Colors []string
}

var group = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

func set(b *testing.B) {
	// address := []string{"192.168.145.6:11211", "192.168.145.5:11211"}
	// instance := memcached.Instance()
	// instance.SetClient(address...)
	// instance.Client.Timeout = 100 * time.Millisecond
	for i := 0; i < b.N; i++ {
		fmt.Println(1)
		// key := fmt.Sprintf("testkey_%v", i)
		// memcached.Instance().Client.Set(&memcache.Item{Key: key, Value: []byte("my value")})
	}

}
