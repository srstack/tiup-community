#!/bin/sh

WORKDIR=${TIUP_COMPONENT_DATA_DIR:-$HOME/.tiup}/storage/papermc
mkdir -p $WORKDIR
cd $WORKDIR

exec java -jar "$TIUP_COMPONENT_INSTALL_DIR/papermc.jar" "$@"
