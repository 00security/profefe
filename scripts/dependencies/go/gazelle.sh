#!/bin/bash

set -e
set -x

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
REPO_DIR="$(cd "$(dirname ${SCRIPT_DIR}/../../../../)"; pwd -P)"

pushd ${REPO_DIR}
# -build_extra_args="-exclude=**/testdata","-exclude=go/packages/packagestest" are used to stop the warning messages in gazelle for things like wire
# https://github.com/bazelbuild/bazel-gazelle/issues/610

# use prune true on the first one only so we keep a clean set of deps
echo "--- Updating 3rd Party Repositories"
bazel run //:gazelle -- update-repos -repo_root=${REPO_DIR} -prune=true -from_file=go.mod -to_macro=thirdparty/go/repositories.bzl%go_repositories -build_extra_args="-exclude=**/testdata","-exclude=go/packages/packagestest"
echo "--- Running Gazelle"
bazel run //:gazelle -- update -index=false  -exclude="go/packages/packagestest" 
popd