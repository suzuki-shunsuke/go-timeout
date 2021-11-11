#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.."

export PATH=$HOME/.aqua/bin:$PATH
curl -sSfL https://raw.githubusercontent.com/suzuki-shunsuke/aqua-installer/v0.2.0/aqua-installer | bash
aqua i -l
bash scripts/test-code-climate.sh
