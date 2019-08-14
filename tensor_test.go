package tensor

import (
	"fmt"
	"testing"
)

func TestFloat64_Type(t *testing.T) {
	ints := []int{1, 233, 45, 454, 654, 6,}
	fmt.Println(ints)
	test(ints)
	fmt.Println(ints)

}
func test(a []int) {
	for i, v := range a {
		a[i] = -v
	}
	a =append(a,2222222222222222222)
}
