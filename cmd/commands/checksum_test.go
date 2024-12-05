/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright Contributors to the cpackget project. */

package commands_test

import (
	"errors"
	"os"
	"testing"

	errs "github.com/open-cmsis-pack/cpackget/cmd/errors"
)

// TODO: Compare actual ErrFileNotFound output
var checksumCreateCmdTests = []TestCase{
	{
		name:        "test different number of parameters",
		args:        []string{"checksum-create"},
		expectedErr: errors.New("accepts 1 arg(s), received 0"),
	},
	{
		name:        "test help command",
		args:        []string{"help", "checksum-create"},
		expectedErr: nil,
	},
	{
		name:        "test creating checksum of nonexisting pack",
		args:        []string{"checksum-create", "DoesNotExist.Pack.1.2.3.pack"},
		expectedErr: errs.ErrFileNotFound,
	},
	{
		name:        "test using nonexisting hash function",
		args:        []string{"checksum-create", "Vendor.Pack.1.2.3.pack", "-a", "sha1"},
		expectedErr: errs.ErrHashNotSupported,
		setUpFunc: func(t *TestCase) {
			f, _ := os.Create("Vendor.Pack.1.2.3.pack.sha256.checksum")
			f.Close()
		},
		tearDownFunc: func() {
			os.Remove("Vendor.Pack.1.2.3.pack.sha256.checksum")
		},
	},
}

var checksumVerifyCmdTests = []TestCase{
	{
		name:        "test different number of parameters",
		args:        []string{"checksum-verify"},
		expectedErr: errors.New("accepts 1 arg(s), received 0"),
	},
	{
		name:        "test help command",
		args:        []string{"help", "checksum-verify"},
		expectedErr: nil,
	},
	{
		name:        "test verifying checksum of nonexisting pack",
		args:        []string{"checksum-verify", "DoesNotExist.Pack.1.2.3.pack"},
		expectedErr: errs.ErrFileNotFound,
		setUpFunc: func(t *TestCase) {
			f, _ := os.Create("DoesNotExist.Pack.1.2.3.pack.sha256.checksum")
			f.Close()
		},
		tearDownFunc: func() {
			os.Remove("DoesNotExist.Pack.1.2.3.pack.sha256.checksum")
		},
	},
	{
		name:        "test verifying checksum of nonexisting checksum file",
		args:        []string{"checksum-verify", "Vendor.Pack.1.2.3.pack"},
		expectedErr: errs.ErrFileNotFound,
		tearDownFunc: func() {
			os.Remove("Vendor.Pack.1.2.3.pack.sha256.checksum")
		},
	},
}

func TestChecksumCreateCmd(t *testing.T) {
	runTests(t, checksumCreateCmdTests)
}

func TestChecksumVerifyCmd(t *testing.T) {
	runTests(t, checksumVerifyCmdTests)
}
