package main

import (
	"fmt"
	"time"

	"github.com/opusdv/lru_cache/pkg/cache"
)

var testData = map[string]string{
	"TestKey1": "TestValue1",
	"TestKey2": "TestValue2",
	"TestKey3": "TestValue3",
	"TestKey4": "TestValue4",
	"TestKey5": "TestValue5",
}

func GetTestData(key string) string {
	return testData[key]
}

func main() {
	c := cache.NewLRUCache(3)

	for i := 0; i < 2; i++ {
		for k, v := range testData {
			if v, ok := c.Get(k); ok {
				fmt.Printf("Value from cache %s\n", v)
				time.Sleep(time.Second * 4)
				continue
			}
			s := GetTestData(k)
			fmt.Printf("Value %s\n", s)
			c.Add(k, v)
			time.Sleep(time.Second * 4)
		}
	}
}
