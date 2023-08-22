// Code generated
// This file is a generated precompile config test with the skeleton of test functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package md5

import (
	"math/big"
	"testing"

	"github.com/ava-labs/subnet-evm/precompile/precompileconfig"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
)

// TestVerify tests the verification of Config.
func TestVerify(t *testing.T) {
	tests := map[string]testutils.ConfigVerifyTest{
		"valid config": {
			Config:        NewConfig(big.NewInt(3)),
			ExpectedError: "",
		},
		// CUSTOM CODE STARTS HERE
		// Add your own Verify tests here, e.g.:
		// "your custom test name": {
		// 	Config: NewConfig(big.NewInt(3),),
		// 	ExpectedError: ErrYourCustomError.Error(),
		// },
	}
	// Run verify tests.
	testutils.RunVerifyTests(t, tests)
}

// TestEqual tests the equality of Config with other precompile configs.
func TestEqual(t *testing.T) {
	tests := map[string]testutils.ConfigEqualTest{
		"non-nil config and nil other": {
			Config:   NewConfig(big.NewInt(3)),
			Other:    nil,
			Expected: false,
		},
		"different type": {
			Config:   NewConfig(big.NewInt(3)),
			Other:    precompileconfig.NewNoopStatefulPrecompileConfig(),
			Expected: false,
		},
		"different timestamp": {
			Config:   NewConfig(big.NewInt(3)),
			Other:    NewConfig(big.NewInt(4)),
			Expected: false,
		},
		"same config": {
			Config:   NewConfig(big.NewInt(3)),
			Other:    NewConfig(big.NewInt(3)),
			Expected: true,
		},
		// CUSTOM CODE STARTS HERE
		// Add your own Equal tests here
	}
	// Run equal tests.
	testutils.RunEqualTests(t, tests)
}
