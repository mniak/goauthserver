#!/bin/bash

export CONFORMANCE_SERVER="https://localhost:8443/"
export CONFORMANCE_DEV_MODE="true"

python run-test-plan.py \
"oidcc-config-certification-test-plan" plans/core_config.json