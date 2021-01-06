package rolling

import (
	"testing"
	"time"
)

func TestMax(t *testing.T) {

	n := NewNumber()
	for _, x := range []int64{10, 11, 9} {
		n.UpdateMax(x)
		time.Sleep(1 * time.Second)
	}

	if n.Max(time.Now()) != 11 {
		t.Error("failed")
	}

}

func TestAvg(t *testing.T) {

	n := NewNumber()
	for _, x := range []int64{1, 2, 3, 4, 5} {
		n.Increment(x)
		time.Sleep(1 * time.Second)
	}

	if n.Avg(time.Now()) != 3 {
		t.Error("failed")
	}

}


func TestSum(t *testing.T) {

	n := NewNumber()
	for _, x := range []int64{1, 2, 3, 4, 5} {
		n.Increment(x)
		time.Sleep(1 * time.Second)
	}

	if n.Sum(time.Now()) != 15 {
		t.Error("failed")
	}

}
