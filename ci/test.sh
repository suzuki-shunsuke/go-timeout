#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.."

export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v0.7.0/aqua-installer | bash
aqua i -l
bash scripts/test-code-climate.sh
