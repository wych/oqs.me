#!/bin/sh

RANDOM_NAME=`openssl rand -hex 6`
SCRIPT_DIR=`dirname $0`
ROOT_DIR=`cd $SCRIPT_DIR/.. > /dev/null 2>&1 && pwd`

docker build -t $RANDOM_NAME $ROOT_DIR/app && \
(docker run --rm -v $ROOT_DIR/bin:/cmd $RANDOM_NAME ; \
docker image rm $RANDOM_NAME )