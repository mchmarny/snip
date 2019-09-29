#!/usr/bin/env bash

# WIP: DON'T USE IT, still not working
RELEASE_VERSION=$1

FILE_PATH="bin/snip-windows-${RELEASE_VERSION}"
FILE_NAME=$(basename $FILE_PATH)

GH_API="https://api.github.com"
GH_REPO="$GH_API/repos/mchmarny/snip"
GH_RELS="$GH_REPO/releases/latest"
GH_AUTH="Authorization: token $GITHUB_ACCESS_TOKEN"

GH_INFO=$(curl -sH "$GH_AUTH" $GH_RELS)
GH_RELS_ID=$(echo $GH_INFO | jq ".id")
GH_ASSET="https://uploads.github.com/repos/mchmarny/snip/releases/${GH_RELS_ID}/assets?name=${FILE_NAME}access_token=${GITHUB_ACCESS_TOKEN}"
# echo $GH_ASSET

curl -v -u mchmarny --data-binary @"$FILE_PATH" \
  -H "Authorization: token $GITHUB_ACCESS_TOKEN" \
  -H "Content-Type: application/octet-stream" "$GH_ASSET"


echo "DONE"
exit 0