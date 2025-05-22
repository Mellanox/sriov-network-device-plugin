#!/bin/sh


# Copyright 2025 sriov-network-device-plugin Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

SRIOV_DP_SYS_BINARY_DIR="/usr/bin/"
LOG_DIR=""
LOG_LEVEL=10
RESOURCE_PREFIX=""
CONFIG_FILE=""
CLI_PARAMS=""
USE_CDI=false

usage()
{
    /bin/echo -e "This is an entrypoint script for SR-IOV Network Device Plugin"
    /bin/echo -e ""
    /bin/echo -e "./entrypoint.sh"
    /bin/echo -e "\t-h --help"
    /bin/echo -e "\t--log-dir=$LOG_DIR"
    /bin/echo -e "\t--log-level=$LOG_LEVEL"
    /bin/echo -e "\t--resource-prefix=$RESOURCE_PREFIX"
    /bin/echo -e "\t--config-file=$CONFIG_FILE"
    /bin/echo -e "\t--use-cdi"
}

while [ "$1" != "" ]; do
    PARAM="$(echo "$1" | awk -F= '{print $1}')"
    VALUE="$(echo "$1" | awk -F= '{print $2}')"
    case $PARAM in
        -h | --help)
            usage
            exit
            ;;
        --log-dir)
            LOG_DIR=$VALUE
            ;;
        --log-level)
            LOG_LEVEL=$VALUE
            ;;
        --resource-prefix)
            RESOURCE_PREFIX=$VALUE
            ;;
        --config-file)
            CONFIG_FILE=$VALUE
            ;;
        --use-cdi)
            USE_CDI=true
            ;;
        *)
            echo "ERROR: unknown parameter \"$PARAM\""
            usage
            exit 1
            ;;
    esac
    shift
done

CLI_PARAMS="-v $LOG_LEVEL"

if [ "$LOG_DIR" != "" ]; then
    mkdir -p "/var/log/$LOG_DIR"
    CLI_PARAMS="$CLI_PARAMS --log_dir /var/log/$LOG_DIR --alsologtostderr"
else
    CLI_PARAMS="$CLI_PARAMS --logtostderr"
fi

if [ "$RESOURCE_PREFIX" != "" ]; then
    CLI_PARAMS="$CLI_PARAMS --resource-prefix $RESOURCE_PREFIX"
fi

if [ "$CONFIG_FILE" != "" ]; then
    CLI_PARAMS="$CLI_PARAMS --config-file $CONFIG_FILE"
fi

if [ "$USE_CDI" = true ]; then
    CLI_PARAMS="$CLI_PARAMS --use-cdi"
fi
set -f
# shellcheck disable=SC2086
exec $SRIOV_DP_SYS_BINARY_DIR/sriovdp $CLI_PARAMS
