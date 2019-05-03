package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		exptected := "aaaaa"

		if repeated != exptected {
			t.Errorf("exptected '%s' got '%s'", exptected, repeated)
		}
	})

	t.Run("repeat character 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		exptected := "aaaaaaaaaa"

		if repeated != exptected {
			t.Errorf("exptected '%s' got '%s'", exptected, repeated)
		}
	})

}

func ExampleRepeat() {

	fmt.Println(Repeat("a", 5))
	//Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
