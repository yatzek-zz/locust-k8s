#!/bin/bash

#bail on fail
set -eo pipefail

LOCUST=( "/usr/local/bin/locust" )

LOCUST+=( -f ${LOCUST_SCRIPT} )
LOCUST+=( --host=$TARGET_HOST )

LOCUST_MODE=${LOCUST_MODE:-standalone}
if [[ "$LOCUST_MODE" = "master" ]]; then
    LOCUST+=( --master)
elif [[ "$LOCUST_MODE" = "worker" ]]; then
    LOCUST+=( --slave --master-host=$LOCUST_MASTER)
    # wait for master
    while ! wget -sqT5 $LOCUST_MASTER:$LOCUST_MASTER_WEB >/dev/null 2>&1; do
        echo "Waiting for master"
        sleep 5
    done
fi

echo "${LOCUST[@]}"

#replace bash, let locust handle signals
exec ${LOCUST[@]}