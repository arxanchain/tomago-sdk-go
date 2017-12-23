#!/bin/bash
#
# Copyright Greg Haskins All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e

export GO15VENDOREXPERIMENT=1
echo -n "Obtaining list of tests to run.."
PKGS=`go list github.com/arxanchain/tomago-sdk-go/... | grep -v /vendor/`
echo "DONE!"

echo "Running tests..."
go test -cover -p 1 -timeout=20m $PKGS
