package main

/**
并发设置map
*/
import (
	"fmt"
	"reflect"
	"sync"
)

var (
	wg sync.WaitGroup
	m  = sync.Map{}
)

func get(key int) interface{} {
	val, _ := m.Load(key)
	return val
}
func set(key, value int) {
	m.Store(key, value)
}
func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(key int) {
			set(key, key+100)
			val, _ := m.Load(key)

			var transformVal int64 = 0
			v := reflect.ValueOf(val)
			switch v.Kind() {
			case reflect.String:
				fmt.Println(`string`)
			case reflect.Int:
				transformVal = v.Int()
			}

			fmt.Printf("key: %v ; value: %v \n", key, transformVal)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
