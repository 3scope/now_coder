package main

func LFU(operators [][]int, k int) []int {
	result := make([]int, 0)
	cache := NewLFUCache(k)
	for _, v := range operators {
		if v[0] == 1 {
			cache.Set(v[1], v[2])
		} else if v[0] == 2 {
			result = append(result, cache.Get(v[1]))
		}
	}
	return result
}

type LFUCache struct {
	Data map[int]*struct {
		Times int
		Value int
	}
	Cap int
	// 为了不必每次都遍历两边Map寻找最小的Key
	// MinTimes int
	// KeyStore map[int]*ListNode
}

func NewLFUCache(cap int) *LFUCache {
	return &LFUCache{
		Data: make(map[int]*struct {
			Times int
			Value int
		}),
		Cap: cap,
		// MinTimes: 0,
		// KeyStore: make(map[int]int),
	}
}

func (lc *LFUCache) Get(key int) int {
	data, ok := lc.Data[key]
	if !ok {
		return -1
	}
	data.Times++
	return data.Value
}

func (lc *LFUCache) Set(key, value int) {
	data, ok := lc.Data[key]
	if !ok {
		if lc.Cap > 0 {
			lc.Data[key] = &struct {
				Times int
				Value int
			}{
				Times: 1,
				Value: value,
			}
			lc.Cap--
		} else {
			minKey := 0
			for k, _ := range lc.Data {
				minKey = k
				break
			}
			for k, v := range lc.Data {
				if v.Times < lc.Data[minKey].Times {
					minKey = k
				}
			}
			delete(lc.Data, minKey)
			lc.Data[key] = &struct {
				Times int
				Value int
			}{
				Times: 1,
				Value: value,
			}
		}
	} else {
		data.Times++
		data.Value = value
	}
}
