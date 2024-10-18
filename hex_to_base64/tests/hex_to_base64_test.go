package hextobase64_test

import (
	"hextobase64"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHexToBase64(t *testing.T) {
	t.Parallel()

	want := "AdSQ"

	got := hextobase64.HexToBase64("01d490")

	if !cmp.Equal(want, got) {
		t.Errorf("Test failed; wanted %s, got %s", want, got)
	}
}
