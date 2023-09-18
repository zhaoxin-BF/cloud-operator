#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
echo "script_root ${SCRIPT_ROOT}"
CODEGEN_PKG="./hack"
echo $CODEGEN_PKG
echo "------------"
echo "$(dirname "${BASH_SOURCE[0]}")/../../../.."
source "${CODEGEN_PKG}/code-gen.sh"
# exit 1
# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.

kube::codegen::gen_helpers \
    --input-pkg-root github.com/zhaoxin-BF/cloud-operator/apis \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

echo "begin gen client"

kube::codegen::gen_client \
    --with-watch \
    --input-pkg-root github.com/zhaoxin-BF/cloud-operator/apis \
    --output-pkg-root github.com/zhaoxin-BF/cloud-operator/pkg/generated \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"
