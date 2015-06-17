#!/usr/bin/env bash

# This file SHOULD NOT be executed, it should be sourced.

CURRENT_DIR="$(cd $(dirname $0) && pwd)"

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
