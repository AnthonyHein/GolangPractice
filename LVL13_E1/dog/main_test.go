// Implements an easy way to go between human years and dog years.
package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}
}

func TestYears(t *testing.T) {
	v := Years(3)
	if v != 21 {
		t.Error("Expected 21, Got ", v)
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Output:
	// 21
}
