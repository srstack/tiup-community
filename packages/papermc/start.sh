#!/bin/sh

if [[ "$1" != "-p" ]]; then
    WORKDIR=${TIUP_COMPONENT_DATA_DIR:-$HOME/.tiup}/storage/papermc
    mkdir -p $WORKDIR
    cd $WORKDIR
fi

exec "$TIUP_COMPONENT_INSTALL_DIR/papermc" "$@"
