#!/bin/bash
set -e
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PARENT=`dirname ${DIR}`
NAME=alertmanager-webhook-example

VERSION=`cat ${PARENT}/VERSION`
IMAGE=quay.io/bakins/${NAME}:${VERSION}
