//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcomputeschedule_test

import (
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/v3/testutil"
)

const (
	pathToPackage = "sdk/resourcemanager/computeschedule/armcomputeschedule/testdata"
)

func TestMain(m *testing.M) {
	code := run(m)
	os.Exit(code)
}

func run(m *testing.M) int {
	f := testutil.StartProxy(pathToPackage)
	defer f()
	return m.Run()
}
