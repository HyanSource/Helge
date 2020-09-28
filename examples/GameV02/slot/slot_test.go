package slot

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	s := NewSlot()

	fmt.Println(s.GetTable())
}
