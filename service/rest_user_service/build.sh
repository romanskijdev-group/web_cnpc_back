#!/bin/bash

SERVICE_PATH=$(dirname "$(realpath "$0")")
SERVICE_NAME=${1:-"rest_user_service_app"}  # Используем первый аргумент командной строки, если он есть, иначе "rest_user_service"

set -e
trap 'echo "🔴 Error build $SERVICE_NAME"; exit 1' ERR

sudo ls
echo "🟡 Start Build:  $SERVICE_NAME"
echo "🔵SERVICE_PATH:  $SERVICE_PATH"

go mod tidy
go vet ./...
go build -o "${SERVICE_NAME}" "$SERVICE_PATH/cmd/app.go"
sudo chmod +x "${SERVICE_NAME}"

echo "🟢 Finish build $SERVICE_NAME generate"