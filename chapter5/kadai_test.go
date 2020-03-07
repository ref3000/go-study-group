package chapter5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		err := Validate([]string{"dummyPath"}, 1)
		assert.NoError(t, err)
	})

	t.Run("異常系: パス未指定", func(t *testing.T) {
		t.Parallel()

		err := Validate([]string{}, 1)
		assert.Error(t, err)
	})

	t.Run("異常系: -f が 1 未満", func(t *testing.T) {
		t.Parallel()

		err := Validate([]string{"dummyPath"}, 0)
		assert.Error(t, err)
	})
}

func TestCut(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		// setup
		src := bytes.NewBufferString("11,12,13\n21,22,2\n31,32,33")
		dst := new(bytes.Buffer)

		// exercise
		err := Cut(src, dst, ",", 2)

		// verify
		assert.NoError(t, err)
		expected := []byte("12\n22\n32\n")
		assert.Equal(t, expected, dst.Bytes())
	})

	t.Run("異常系: -f 範囲外", func(t *testing.T) {
		t.Parallel()
		// setup
		src := bytes.NewBufferString("11,12,13\n21,22,2\n31,32,33")
		dst := new(bytes.Buffer)

		// exercise
		err := Cut(src, dst, ",", 4)

		// verify
		assert.Error(t, err)
	})
}
