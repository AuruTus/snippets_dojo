package test_snipets

import (
	"context"
	"fmt"
	"math/rand"
	snipets "snipets_dojo/src"
	"snipets_dojo/utils/cfmt"

	"golang.org/x/exp/constraints"
)

type ordered interface {
	constraints.Integer | constraints.Float
}

type ArrPartitionTstr struct{}

var _ snipets.Tstr = (*ArrPartitionTstr)(nil)

func (t *ArrPartitionTstr) Test(ctx context.Context) error {
	type targ float32
	arr := randArr[targ](0, 10, 10)
	cfmt.Printf(ctx, "Test arr %v\n", arr)
	cfmt.Printf(ctx, "testSwap\n")
	testPartition(arrCopy(arr), partitionWithSwap[targ])
	cfmt.Printf(ctx, "testAssign\n")
	testPartition(arrCopy(arr), partitionWithAssign[targ])
	return nil
}

func randArr[T ordered](min, max T, len int) []T {
	arr := make([]T, len)
	for i := 0; i < len; i++ {
		arr[i] = (T)((float32)(rand.Intn((int)(max-min))) + (rand.Float32()))
	}
	return arr
}

func arrCopy[T ordered](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

func testPartition[T ordered](arr []T, partition func([]T, int, int) int) {
	printArr := func(arr []T, turn, head, rear, prtn_key int) {
		fmt.Printf("turn %d, head %d, rear %d, key %d: %v\n", turn, head, rear, prtn_key, arr)
	}

	// recursive closure
	var sort func(arr []T, head, rear, turn int)
	sort = func(arr []T, head, rear, turn int) {
		if head < 0 || rear < 0 || head >= rear {
			return
		}
		prtn_key := partition(arr, head, rear)
		printArr(arr, turn, head, rear, prtn_key)
		sort(arr, head, prtn_key-1, turn+1)
		sort(arr, prtn_key+1, rear, turn+1)
	}

	sort(arr, 0, len(arr)-1, 0)
}

func swap[T ordered](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partitionWithSwap[T ordered](arr []T, head, rear int) int {
	tmp := arr[head]
	for head < rear {
		for head < rear && arr[rear] >= tmp {
			rear--
		}
		swap(arr, head, rear)
		for head < rear && arr[head] <= tmp {
			head++
		}
		swap(arr, head, rear)
	}
	return head
}

func partitionWithAssign[T ordered](arr []T, head, rear int) int {
	tmp := arr[head]
	for head < rear {
		for head < rear && arr[rear] >= tmp {
			rear--
		}
		arr[head] = arr[rear]
		for head < rear && arr[head] <= tmp {
			head++
		}
		arr[rear] = arr[head]
	}
	arr[head] = tmp
	return head
}
