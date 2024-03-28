#!/bin/bash

set -ex
gitBranch=`git branch | grep \* | cut -d ' ' -f2`
labelVersion=$gitBranch.`date -u +%Y.%m%d.%H%M`

echo "Tagging with $labelVersion"
git tag "$labelVersion" --force
git push && \
    git push --tags --force
