package featureflag

import (
	"testing"
)

func TestFeatureFlags(t *testing.T) {
	const (
		testAlphaGate = "TestAlpha"
		testBetaGate  = "TestBeta"
	)

	tests := []struct {
		arg  string
		want map[string]bool
	}{
		{
			arg:  "",
			want: map[string]bool{},
		},
		{
			arg: "TestAlpha=false",
			want: map[string]bool{
				testAlphaGate: false,
			},
		},
		{
			arg: "TestAlpha=true",
			want: map[string]bool{
				testAlphaGate: true,
			},
		},
		{
			arg: "TestAlpha=foo",
			want: map[string]bool{
				testAlphaGate: false,
			},
		},
		{
			arg: "TestAlpha=false,TestBeta=true",
			want: map[string]bool{
				testAlphaGate: false,
				testBetaGate:  true,
			},
		},
		{
			arg: "TestAlpha=true,TestBeta=false",
			want: map[string]bool{
				testAlphaGate: true,
				testBetaGate:  false,
			},
		},
		{
			arg: "TestAlpha=false,TestBeta=false",
			want: map[string]bool{
				testAlphaGate: false,
				testBetaGate:  false,
			},
		},
		{
			arg: "TestAlpha=true,TestBeta=true",
			want: map[string]bool{
				testAlphaGate: true,
				testBetaGate:  true,
			},
		},
		{
			arg: "TestAlpha",
			want: map[string]bool{
				testAlphaGate: true,
			},
		},
		{
			arg: "TestAlpha=false,TestBeta",
			want: map[string]bool{
				testAlphaGate: false,
				testBetaGate:  true,
			},
		},
		{
			arg: "TestAlpha,TestBeta",
			want: map[string]bool{
				testAlphaGate: true,
				testBetaGate:  true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.arg, func(t *testing.T) {
			flags = make(map[string]FeatureFlag) // reset
			Set(test.arg)
			for k, v := range test.want {
				actual, ok := flags[k]
				if !ok {
					t.Fatalf("want: %s=%v, got: nil", k, v)
				}
				if actual.Enabled() != v {
					t.Errorf("want: %s=%v, got: %v", k, v, actual)
				}
			}
		})
	}
}
