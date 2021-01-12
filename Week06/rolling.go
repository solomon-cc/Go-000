package rolling

import (
	"sync"
	"time"
)

const WINDOWSIZE=5

type Number struct {
	Buckets map[int64]*bucket
	Mu      *sync.RWMutex
}

type bucket struct {
	Value int64
}

func NewNumber() *Number {
	rn := &Number{
		Buckets: make(map[int64]*bucket),
		Mu:      &sync.RWMutex{},
	}

	return rn
}

func (rn *Number) getCurrentBucket() *bucket {
	now := time.Now().Unix()
	var b *bucket
	var ok bool
	if b, ok = rn.Buckets[now]; !ok {
		b = &bucket{}
		rn.Buckets[now] = b
	}
	return b
}

func (rn *Number) removeOldBuckets() {
	expired := time.Now().Unix() - WINDOWSIZE
	for timestamp := range rn.Buckets {
		if timestamp <= expired {
			delete(rn.Buckets, timestamp)
		}
	}
}

// Increment 累加最新桶的计数器
func (rn *Number) Increment(i int64) {
	rn.Mu.Lock()
	b := rn.getCurrentBucket()
	b.Value += i
	rn.removeOldBuckets()

	rn.Mu.Unlock()
}

// UpdateMax 将最新桶的计数器置为某个最大值
func (rn *Number) UpdateMax(n int64) {
	rn.Mu.Lock()
	b := rn.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	rn.Mu.Unlock()
	rn.removeOldBuckets()
}

// Sum 计算最新 5 个桶内计数器的和
func (rn *Number) Sum(now time.Time) int64 {
	sum := int64(0)

	rn.Mu.RLock()
	defer rn.Mu.RUnlock()

	for timestamp, bucket := range rn.Buckets {
		if timestamp >= now.Unix()-WINDOWSIZE {
			sum += bucket.Value
		}
	}
	return sum
}

// Max 获取最新 5 个桶内计数器的最大值
func (rn *Number) Max(now time.Time) int64 {
	var max int64

	rn.Mu.RLock()
	defer rn.Mu.RUnlock()

	for timestamp, bucket := range rn.Buckets {
		if timestamp >= now.Unix()-WINDOWSIZE {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}
	return max
}

// Avg 计算最新 5 个桶内计数器的平均值
func (rn *Number) Avg(now time.Time) int64 {
	return rn.Sum(now) / WINDOWSIZE
}
