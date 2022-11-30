package snipets

import (
	"context"
	"fmt"
	"math/rand"
	"snipets_dojo/utils/cfmt"

	"golang.org/x/exp/constraints"
)

type ordered constraints.Ordered

type ArrPartitionTstr struct{}

var _ Tstr = (*ArrPartitionTstr)(nil)

func (t *ArrPartitionTstr) Test(ctx context.Context) error {
	arr := randIntArr(0, 10, 10)
	cfmt.Printf(ctx, "testSwap\n")
	testPartition(arrCopy(arr), partitionWithSwap[int])
	cfmt.Printf(ctx, "testAssign\n")
	testPartition(arrCopy(arr), partitionWithAssign[int])
	return nil
}

func randIntArr(min, max, len int) []int {
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = rand.Intn(max-min) + min
	}
	return arr
}

func arrCopy[T ordered](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

func printArr[T ordered](arr []T) {
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}

func testPartition[T ordered](arr []T, partition func([]T, int, int) int) {
	for range arr {
		printArr(arr)
		partition(arr, 0, len(arr)-1)
	}
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
