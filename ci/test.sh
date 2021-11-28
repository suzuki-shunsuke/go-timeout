#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.."

export PATH=$HOME/.aqua/bin:$PATH
curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v0.3.0/aqua-installer | bash
aqua i -l
bash scripts/test-code-climate.sh
