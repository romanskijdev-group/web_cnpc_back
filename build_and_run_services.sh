#!/bin/bash

set -e
trap 'echo "🔴 Error build $SERVICE_NAME"; exit 1' ERR

THIS_PATH=$(dirname "$(realpath "$0")")
SERVICE_LIST=(
  "rest_admin_service" \
  "rest_user_service" \
  "chats_service" \
  "openai_service" \
  "payment_service" \
  "user_service" \
  "system_service" \
  )

echo "THIS_PATH:  $THIS_PATH"


for service in "${SERVICE_LIST[@]}"; do
  SERVICE_PATH="$THIS_PATH/service/$service"
  echo "_"
  echo "🦄🦄 SERVICE_PATH: $SERVICE_PATH"
  cd "$SERVICE_PATH"
  # go list -u -m all | awk '/\[.*\]$/ {print}'
  # go get -u ./... && go mod tidy
    gofmt -w .
  ./build.sh
done

# for service in "${SERVICE_LIST[@]}"; do
#   echo "🦄 Запуск сервиса: $service"
#   SERVICE_PATH="$THIS_PATH/service/$service"
#   cd "$SERVICE_PATH"
#     gnome-terminal -- bash -c "cd '$SERVICE_PATH' && go run cmd/app.go; echo 'Нажмите Enter для выхода...'; read"
# done