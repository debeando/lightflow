#!/usr/bin/env bash

# Apache License Version 2.0, January 2004
# https://github.com/debeando/lightflow/blob/master/LICENSE.md

if ! type "jq" > /dev/null; then
  echo "Require command tool, please install: jq"
  exit 1
fi

if [ -z "$GITHUB_TOKEN" ]
then
  echo "Require environment variable: GITHUB_TOKEN"
  exit 1
fi

TAG=$(grep "Number string =" < cli/version.go | awk -F'"' '{$0=$2}1')

git push --delete origin "v${TAG}"
git tag --delete "v${TAG}"
git tag "v${TAG}"
git push --tags --force

curl --silent --output /dev/null --data "{\"tag_name\": \"v${TAG}\",\"target_commitish\": \"master\",\"name\": \"Pre Release v${TAG}\",\"body\": \"\",\"draft\": true,\"prerelease\": true}" "https://api.github.com/repos/debeando/lightflow/releases?access_token=${GITHUB_TOKEN}"

ID=$(curl -sH "Authorization: token $GITHUB_TOKEN" https://api.github.com/repos/debeando/lightflow/releases | jq -r '.[0].id')

rm -rf pkg/*

BUILD_DATE=$(date +%Y%m%d%H%M)
export BUILD_DATE

go generate ./...
mkdir -p pkg/linux_amd64/
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X github.com/debeando/lightflow/command.BuildTime=${BUILD_DATE}" -o pkg/linux_amd64/lightflow main.go
tar -czf pkg/linux_amd64/lightflow-linux_amd64.tar.gz -C pkg/linux_amd64/ lightflow

curl -# \
     --silent \
     --output /dev/null \
     -XPOST \
     -H "Authorization:token ${GITHUB_TOKEN}" \
     -H "Content-Type:application/octet-stream" \
     --data-binary @pkg/linux_amd64/lightflow-linux_amd64.tar.gz \
     "https://uploads.github.com/repos/debeando/lightflow/releases/${ID}/assets?name=lightflow-linux_amd64.tar.gz"

echo -e "\r"
