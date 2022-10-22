#!/bin/sh

if [[ "$1" != "-p" ]]; then
    WORKDIR=${TIUP_COMPONENT_DATA_DIR:-$HOME/.tiup}/storage/hmcl
    mkdir -p $WORKDIR
    cd $WORKDIR
fi

exec java -jar "$TIUP_COMPONENT_INSTALL_DIR/usr/share/java/hmcl/hmcl.jar" "$@"
