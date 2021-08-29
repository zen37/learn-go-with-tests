package iteration

import "testing"

func Test_Repeat(t *testing.T) {

	actual := repeat("a", 4)
	expect := "aaaa"

	if actual != expect {
		t.Errorf("expected %q, but got %q:", expect, actual)
	}
}

//func Benchmark_Test(b *testing.B) {
func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 11)
	}
}
