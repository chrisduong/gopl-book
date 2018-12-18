package counter

import (
	"fmt"
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	writer, count := CountingWriter(os.Stdout)
	// Because *ByteCounter satisfies the io.Writer contract, we can pass it to Fprint
	// Then when Fprint write, it will use the Write method of the type ByteCounter,
	// then it will write the count into the `count field`
	expected, _ := fmt.Fprint(writer, "foo")

	if *count != int64(expected) {
		t.Logf("%d != %d", *count, expected)
		t.Fail()
	}
}

