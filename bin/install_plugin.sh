#!/usr/bin/env bash
CURRENTDIR=`pwd`
PLUGIN_PATH="$CURRENTDIR/plugin-cf-manifest-generator"

go build
cf uninstall-plugin manifest-generator
cf install-plugin "$PLUGIN_PATH" -f