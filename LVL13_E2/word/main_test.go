package word

import (
    "fmt"
    "testing"
)

func TestCount (t *testing.T) {
    v := Count("I am Anthony.")
    if v != 3 {
        t.Error("Expected 3 got ", v)
    }

}

func Benchmark (b *testing.B) {
    for i := 0; i < b.N; i++ {
        Count("I am Anthony.")
    }
}

func Example () {
    fmt.Println(Count("I am Anthony."))
    // Output:
    // 3
}
