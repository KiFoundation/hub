package types

import (
	"crypto/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestDepositKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(DepositKeyPrefix, address.MustLengthPrefix(addr)...),
				DepositKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			DepositKey(addr)
		})
	}
}
