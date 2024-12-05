/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright Contributors to the cpackget project. */

package commands_test

import (
	"testing"

	errs "github.com/open-cmsis-pack/cpackget/cmd/errors"
	"github.com/open-cmsis-pack/cpackget/cmd/installer"
)

var (
	urlPath      string = "https://www.keil.com"
	wrongURLPath string = "https://www.keilll.com"
)

var connectionCmdTests = []TestCase{
	{
		name:        "test help command",
		args:        []string{"help", "connection"},
		expectedErr: nil,
	},
	{
		name:        "test checking connection",
		args:        []string{"connection", urlPath},
		expectedErr: nil,
	},
	{
		name:          "test checking invalid url",
		args:          []string{"connection", wrongURLPath},
		expectedErr:   errs.ErrOffline,
		expErrUnwwrap: true,
	},

	{ // set up environment for next test
		name:      "test checking default connection",
		args:      []string{"init"},
		noCleanup: true,
		setUpFunc: func(t *TestCase) {
			server := NewServer()
			t.args = append(t.args, server.URL()+installer.PublicIndex)
			server.AddRoute(installer.PublicIndex, []byte(`<?xml version="1.0" encoding="UTF-8" ?>
<index schemaVersion="1.1.0" xs:noNamespaceSchemaLocation="PackIndex.xsd" xmlns:xs="https://www.w3.org/2001/XMLSchema-instance">
<vendor>TheVendor</vendor>
<url>https://www.keil.com/</url>
<timestamp>2021-10-17T12:21:59.1747971+00:00</timestamp>
<pindex>
  <pdsc url="https://www.keil.com" vendor="Keil" name="PackName" version="1.2.3" />
</pindex>
</index>`))
		},
	},
	{
		name:        "test checking default connection",
		args:        []string{"connection"},
		expectedErr: nil,
	},
}

func TestConnectionCmd(t *testing.T) {
	runTests(t, connectionCmdTests)
}
