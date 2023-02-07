package benchstring

import (
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Plus(1000, "strings")
	}

}

func BenchmarkStrBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StrBuilder(1000, "strings")
	}
}
