package simplecgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyCFunction(t *testing.T) {
	t.Run("TestMyCFunction - success", func(t *testing.T) {
		result := MyCFunction(2)
		require.Equal(t, 34, result)
	})
}
