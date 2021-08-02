package integers

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	got := add(2, 2)
	want := 4

	if got != want {
		t.Errorf("expected %d, but got %d", want, got)
	}
}

func ExampleAdd() {
	sum := add(1, 5)
	fmt.Println(sum)
	//Output: 6
}
