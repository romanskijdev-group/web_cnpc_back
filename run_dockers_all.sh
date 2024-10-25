#!/bin/bash

set -e
trap 'echo "🔴 Error build $SERVICE_NAME"; exit 1' ERR

./build_core.sh

# Инициализация переменных параметров
DOCKER_NETWORK_NAME=""
DATABASE_HOST_IP=""

# Обработка именованных аргументов
while getopts ":n:a:t:d:" opt; do
  case $opt in
    n) DOCKER_NETWORK_NAME="$OPTARG"
    ;;
    d) DATABASE_HOST_IP="$OPTARG"
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    ;;
  esac
done

# Проверка, что все параметры были указаны
if [ -z "$DOCKER_NETWORK_NAME" ] || [ -z "$DATABASE_HOST_IP" ]; then
  echo "🔴 All parameters (DOCKER_NETWORK_NAME, DATABASE_HOST_IP) must be specified and non-empty."
  exit 1
fi

THIS_PATH=$(dirname "$(realpath "$0")")

SERVICE_LIST=(
  "rest_admin_service::15000::true" \
  "rest_user_service::15004::true" \
  "chats_service::15001::true" \
  "payment_service::15003::false" \
  "user_service::15005::false" \
  "system_service::15015::false" \
  "openai_service::15002::false" \
  )

echo "THIS_PATH:  $THIS_PATH"

for item in "${SERVICE_LIST[@]}"; do
  (
    service=$(echo "$item" | awk -F'::' '{print $1}')
    port=$(echo "$item" | awk -F'::' '{print $2}')
    open_port_externally=$(echo "$item" | awk -F'::' '{print $3}')
    SERVICE_PATH="$THIS_PATH/service/$service"
    echo "_"
    echo "🦄🦄 service: $service Port: $port"
    cd "$SERVICE_PATH"
    ./docker_run.sh \
    OPEN_PORT_EXTERNALLY=$open_port_externally HOST_PORT=$port \
    SERVICE_NAME=$service \
    DOCKER_NETWORK_NAME=$DOCKER_NETWORK_NAME \
    DATABASE_HOST_IP=$DATABASE_HOST_IP \
    W_BUILD="true" W_START="true"
  ) &
done

wait