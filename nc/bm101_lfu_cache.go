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
	// 存放键值对。
	KeyValue map[int]int
	// 通过“key”可以得到访问次数。
	KeyTimes map[int]int
	// 通过访问次数可以得到相同次数的“key”。
	TimesKeySet map[int][]int
	// 存放最少的访问次数。
	MinTimes int
}

func LFUFactory(cap int) *LFUCache {
	return &LFUCache{
		Capacity: cap,
		KeyValue: make(map[int]int),
		KeyTimes: make(map[int]int),
		// 有很多相同访问次数的“key”，它们按照时间顺序排列。
		TimesKeySet: map[int][]int{},
		MinTimes:    0,
	}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.KeyValue[key]
	if ok {
		// 需要增加相应的访问次数。
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

	// 判断需要增加访问次数的页面，是否是第一次访问的新页面，如果是的话，就不会在“TimesKeySet”中找到对应的下标。
	index := -1
	for i := 0; i < len(this.TimesKeySet[oldTimes]); i++ {
		if this.TimesKeySet[oldTimes][i] == key {
			index = i
			break
		}
	}
	// 从原来的集合中删除。
	if index != -1 {
		this.TimesKeySet[oldTimes] = append(this.TimesKeySet[oldTimes][:index], this.TimesKeySet[oldTimes][index+1:]...)
	}
	// 加入到新的集合。
	this.TimesKeySet[newTimes] = append(this.TimesKeySet[newTimes], key)
	// 如果原有的集合空了，并且“key”是最小的“times”，那么需要更新“times”的值。
	if len(this.TimesKeySet[oldTimes]) == 0 {
		if oldTimes == this.MinTimes {
			this.MinTimes++
		}
	}
}

func (this *LFUCache) Set(key, value int) {
	// 如果“key”存在，那么修改值，并且访问次数加一。
	if _, ok := this.KeyValue[key]; ok {
		this.KeyValue[key] = value
		this.IncreaseTimes(key)
		return
	}
	// 如果不存在，那么根据容量判断是否需要淘汰页面。
	if this.Capacity > 0 {
		this.KeyValue[key] = value
		this.IncreaseTimes(key)
		this.Capacity--
	} else {
		// 此时还没有更新“MinTimes”，因此会淘汰原有的页面。
		this.RemoveMinTimes()
		this.KeyValue[key] = value
		this.IncreaseTimes(key)
	}
	// 因为不存在，所以最少访问次数一定是1。
	this.MinTimes = 1
}

func (this *LFUCache) RemoveMinTimes() {
	key := this.TimesKeySet[this.MinTimes][0]
	this.TimesKeySet[this.MinTimes] = this.TimesKeySet[this.MinTimes][1:]
	delete(this.KeyTimes, key)
	delete(this.KeyValue, key)
}
