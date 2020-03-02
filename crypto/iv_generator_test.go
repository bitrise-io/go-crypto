package crypto_test

import (
	"testing"

	"github.com/bitrise-io/go-crypto/crypto"
	"github.com/stretchr/testify/require"
)

func Test_GenerateIV(t *testing.T) {
	t.Log("test")
	{
		generatedIV, err := crypto.GenerateIV()
		require.NoError(t, err)
		require.Equal(t, int(12), len(generatedIV))
	}
}
