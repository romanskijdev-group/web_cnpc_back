
#!/bin/bash

set -e
trap 'echo "🔴 Error build "; exit 1' ERR

THIS_PATH=$(dirname "$(realpath "$0")")
CORE_PATH="$THIS_PATH/core"  # Название ядра

echo "CORE_PATH:  $CORE_PATH"

sudo ls
cd "$CORE_PATH"
./build_proto.sh
# go mod tidy
# go list -u -m all | awk '/\[.*\]$/ {print}'
go vet ./...
go mod tidy
