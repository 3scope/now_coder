package designpatterns

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var (
	m sync.Mutex
	o sync.Once
	s *singleton
)

func GetInstance() *singleton {
	o.Do(func() {
		s = new(singleton)
	})
	return s
}

func GetInstanceMutex() *singleton {
	// 双重校验。
	if s == nil {
		m.Lock()
		if s == nil {
			s = new(singleton)
		}
		m.Unlock()
	}
	return s
}

// 以上为一个单例模式包里面的内容。
func main() {
	obj := GetInstance()
	fmt.Println(obj)
}
