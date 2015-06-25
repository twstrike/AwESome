#!/usr/bin/env bash

# This file SHOULD NOT be executed, it should be sourced.
CURRENT_SOURCE=$_
if [[ "$CURRENT_SOURCE" != *init.sh ]]; then
    echo "For this to work correctly, you need to source this file, not execute it:"
    echo "  source init.sh"
else
    CURRENT_DIR="$(cd "$(dirname "$CURRENT_SOURCE")" && pwd -P)"
    export GOPATH=$CURRENT_DIR/.gopkg
    export PATH=$GOPATH/bin:$PATH

    mkdir -p $GOPATH
    mkdir -p $GOPATH/src
    mkdir -p $GOPATH/pkg
    mkdir -p $GOPATH/bin
    mkdir -p $GOPATH/src/github.com/twstrike

    if [ ! -e $GOPATH/src/github.com/twstrike/AwESome ]; then
        ln -s $CURRENT_DIR $GOPATH/src/github.com/twstrike/AwESome
    fi
fi
