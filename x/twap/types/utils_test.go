package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/osmosis-labs/osmosis/v11/app/apptesting/osmoassert"
)

func TestGetAllUniqueDenomPairs(t *testing.T) {
	tests := map[string]struct {
		denoms       []string
		wantedPairGT []string
		wantedPairLT []string
		panics       bool
	}{
		"basic":    {[]string{"A", "B"}, []string{"B"}, []string{"A"}, false},
		"basicRev": {[]string{"B", "A"}, []string{"B"}, []string{"A"}, false},
		// AB > A
		"prefixed": {[]string{"A", "AB"}, []string{"AB"}, []string{"A"}, false},
		"basic-3":  {[]string{"A", "B", "C"}, []string{"C", "C", "B"}, []string{"B", "A", "A"}, false},
		"panics":   {[]string{"A", "A"}, []string{}, []string{}, true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			osmoassert.ConditionalPanic(t, tt.panics, func() {
				pairGT, pairLT := GetAllUniqueDenomPairs(tt.denoms)
				require.Equal(t, pairGT, tt.wantedPairGT)
				require.Equal(t, pairLT, tt.wantedPairLT)
			})
		})
	}
}

func TestLexicographicalOrderDenoms(t *testing.T) {
	tests := map[string]struct {
		firstDenom     string
		secondDenom    string
		expectedDenomA string
		expectedDenomB string
		expectErr      bool
	}{
		"basic":    {"A", "B", "A", "B", false},
		"basicRev": {"B", "A", "A", "B", false},
		"realDenoms": {"uosmo", "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
			"ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", "uosmo", false},
		"sameDenom": {"A", "A", "", "", true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// system under test
			denomA, denomB, err := LexicographicalOrderDenoms(tt.firstDenom, tt.secondDenom)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.Equal(t, denomA, tt.expectedDenomA)
				require.Equal(t, denomB, tt.expectedDenomB)
			}
		})
	}
}
