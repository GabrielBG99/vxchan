#!/bin/sh

set -e

filebeat setup
service filebeat start

exec "$@"
