package main

import (
	"fmt"
	"time"

	memcached "github.com/reckcn/cache/memcached"
	serialization "github.com/reckcn/cache/serialization"
)

func main() {
	address := []string{"192.168.145.6:11211", "192.168.145.5:11211"}
	memcached.Instance().SetClient(new(serialization.JsonProvider), 100*time.Millisecond, address...)
	//fmt.Println(instance.Client)
	key := fmt.Sprintf("testkey_%v", 11)
	memcached.Instance().SetValue(key, "my value 1", 1000)
	key2 := fmt.Sprintf("testkey_%v", 12)
	memcached.Instance().SetValue(key2, "my value 2", 1000)
	memcached.Instance().SetValue("hkh", nil, 1000)
	for index := 0; index < 1; index++ {
		var str = ""
		//err := memcached.Instance().GetValue(key, &str)
		//mapd := make(map[string]string)
		maps, err := memcached.Instance().GetMultiValue([]string{key, key2, "hkh"})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
			fmt.Println(maps)
			for _, item := range maps {
				fmt.Println(string(item.Value))
			}
			fmt.Println("====================================")
		}
	}
}
