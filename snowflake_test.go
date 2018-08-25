package snowflake

import (
	"testing"
)

func BenchmarkSnowflake_UUID(b *testing.B) {

	sf := NewSnowflack()

	for n := 0; n < b.N; n++ {
		sf.UUID()
	}
}
