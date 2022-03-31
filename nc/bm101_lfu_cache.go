package main

func LFU(operators [][]int, k int) []int {
	// write code here
	result := make([]int, 0)
	lfu := LFUFactory(k)
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			lfu.Set(operators[i][1], operators[i][2])
		} else {
			result = append(result, lfu.Get(operators[i][1]))
		}
	}

	return result
}

type LFUCache struct {
	Capacity int
	// Use key to get value.
	KeyValue map[int]int
	// Use key to get times.
	KeyTimes map[int]int
	// Use times to get key set.
	TimesKeySet map[int][]int
	MinTimes    int
}

func LFUFactory(cap int) *LFUCache {
	return &LFUCache{
		Capacity: cap,
		KeyValue: make(map[int]int),
		KeyTimes: make(map[int]int),
		// Sort by time when put it in.
		TimesKeySet: map[int][]int{},
		MinTimes:    0,
	}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.KeyValue[key]
	if ok {
		// Increase one time of the key.
		this.IncreaseTimes(key)
		return value
	} else {
		return -1
	}
}

func (this *LFUCache) IncreaseTimes(key int) {
	this.KeyTimes[key]++
	newTimes := this.KeyTimes[key]
	oldTimes := newTimes - 1

	index := -1
	for i := 0; i < len(this.TimesKeySet[oldTimes]); i++ {
		if this.TimesKeySet[oldTimes][i] == key {
			index = i
			break
		}
	}
	// Delete key in the original set.
	if index != -1 {
		this.TimesKeySet[oldTimes] = append(this.TimesKeySet[oldTimes][:index], this.TimesKeySet[oldTimes][index+1:]...)
	}
	// Add key to new one.
	this.TimesKeySet[newTimes] = append(this.TimesKeySet[newTimes], key)
	// When the old one is empty.
	if len(this.TimesKeySet[oldTimes]) == 0 {
		if oldTimes == this.MinTimes {
			this.MinTimes++
		}
	}
}

func (this *LFUCache) Set(key, value int) {
	// If key is existed.
	if _, ok := this.KeyValue[key]; ok {
		this.KeyValue[key] = value
		this.IncreaseTimes(key)
		return
	}

	if this.Capacity > 0 {
		this.KeyValue[key] = value
		this.IncreaseTimes(key)
		this.Capacity--
	} else {
		this.RemoveMinTimes()
		this.KeyValue[key] = value
		this.IncreaseTimes(key)
	}
	// No key in any set.
	this.MinTimes = 1
}

func (this *LFUCache) RemoveMinTimes() {
	key := this.TimesKeySet[this.MinTimes][0]
	this.TimesKeySet[this.MinTimes] = this.TimesKeySet[this.MinTimes][1:]
	delete(this.KeyTimes, key)
	delete(this.KeyValue, key)
}
