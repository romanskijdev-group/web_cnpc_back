#!/bin/bash

SERVICE_PATH=$(dirname "$(realpath "$0")")
SERVICE_NAME="system_service"  # Название сервиса

set -e
trap 'echo "🔴 Error build $SERVICE_NAME"; exit 1' ERR

sudo ls
echo "🟡 Start Build:  $SERVICE_NAME"
echo "🔵SERVICE_PATH:  $SERVICE_PATH"

go mod tidy
go vet ./...
go build -o "${SERVICE_NAME}_app" "$SERVICE_PATH/cmd/app.go"
sudo chmod +x "${SERVICE_NAME}_app"

echo "🟢 Finish build $SERVICE_NAME generate"