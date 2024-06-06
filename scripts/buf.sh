#!/bin/sh

cd proto
buf lint || exit $?

buf generate --template buf.gen.yaml
