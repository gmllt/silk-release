#!/bin/bash

set -eu
set -o pipefail

configure_db "${DB}"
# shellcheck disable=SC2068
# Double-quoting array expansion here causes ginkgo to fail
args=${@} 
go run github.com/onsi/ginkgo/v2/ginkgo  --skip-package cni $args
# run in serial
go run github.com/onsi/ginkgo/v2/ginkgo $args --nodes=1 ./cni
