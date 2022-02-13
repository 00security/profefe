#!/usr/bin/env bash

# This file is used during bazel build time to set the Version variable in each go binary

echo GITSHA $(git rev-parse --short HEAD 2>/dev/null)
echo BUILDTIME $(date -u +"%Y-%m-%dT%H:%M:%SZ")
echo STABLE_BUILD_VERSION $STABLE_BUILD_VERSION

if [ -z $VERSION ]; then
  echo VERSION
else
  echo VERSION git-$(GITSHA)
fi