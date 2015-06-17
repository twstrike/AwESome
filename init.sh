#!/usr/bin/env bash

# This file SHOULD NOT be executed, it should be sourced.

if command -v realpath>/dev/null 2>&1; then
    CURRENT_FILE=`realpath "$0"`
else
    CURRENT_FILE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/$(basename $0)"
fi
CURRENT_DIR=`dirname "$CURRENT_FILE"`

export GOPATH=$CURRENT_DIR/gopkg
export PATH=$GOPATH/bin:$PATH

mkdir -p $GOPATH
mkdir -p $GOPATH/src
mkdir -p $GOPATH/pkg
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/src/github.com/twstrike

if [ ! -e $GOPATH/src/github.com/twstrike/AwESome ]; then
    ln -s $CURRENT_DIR $GOPATH/src/github.com/twstrike/AwESome
fi
