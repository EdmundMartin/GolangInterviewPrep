package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetNthFibonacci(t *testing.T) {
	var results []int
	for i := 1; i < 10; i++ {
		results = append(results, GetNthFibonacci(i))
	}
	fmt.Println(results)
	if !reflect.DeepEqual(results, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}) {
		t.Error("Do you even fibonacci")
	}
}
