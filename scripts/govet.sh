#!/bin/bash
#
# Copyright Greg Haskins All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e

echo "LINT: Running code checks.."
echo "Running go vet"

cd $GOPATH/src/github.com/arxanchain/tomago-sdk-go/

for i in `ls -d */|grep -v gotools |grep -v scripts`
do
    OUTPUT=`go vet ./$i... && golint ./$i... && misspell ./$i/*`
    if [[ $OUTPUT ]]; then
        echo "The following files contain go vet errors"
        echo $OUTPUT
        exit 1
    fi
done
