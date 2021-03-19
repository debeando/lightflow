#!/bin/bash

# Apache License Version 2.0, January 2004
# https://github.com/debeando/lightflow/blob/master/LICENSE.md

set -e

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

if ! [[ "${OSTYPE}" == "linux"* ]]; then
  echo "Only works on Linux amd64."
  exit
fi

if ! type "wget" > /dev/null; then
  echo "The program 'wget' is currently not installed, please install it to continue."
  exit
fi

FILE="lightflow-linux_amd64.tar.gz"
TAG=$(wget -qO- "https://api.github.com/repos/debeando/lightflow/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -f /usr/local/bin/lightflow ]; then
  rm -f /usr/local/bin/lightflow
fi

if [ -L /usr/bin/lightflow ]; then
  rm -f /usr/bin/lightflow
fi

if [ -n "${FILE}" ]; then
  wget -qO- "https://github.com/debeando/lightflow/releases/download/${TAG}/${FILE}" | tar xz -C /usr/local/bin/
fi

if [ -f /usr/local/bin/lightflow ]; then
  ln -s /usr/local/bin/lightflow /usr/bin/lightflow
fi