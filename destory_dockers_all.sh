#!/bin/bash

set -e
trap 'echo "🔴 Error build $SERVICE_NAME"; exit 1' ERR

THIS_PATH=$(dirname "$(realpath "$0")")
SERVICE_LIST=(
  "rest_admin_service" \
  "rest_user_service" \
  "chats_service" \
  "payment_service" \
  "user_service" \
  "system_service" \
  "openai_service" \
)


for service in "${SERVICE_LIST[@]}"; do
  SERVICE_PATH="$THIS_PATH/service/$service"

  CONTAINER_NAME="${service}_zodi_c"
  IMAGE_NAME="${service}_zodi_i"

  # Check if the container exists
  if docker ps -a --format '{{.Names}}' | grep -Eq "^${CONTAINER_NAME}\$"; then
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
  else
    echo "Container $CONTAINER_NAME not found, skipping stop and remove."
  fi

  # Check if the image exists
  if docker images --format '{{.Repository}}:{{.Tag}}' | grep -Eq "^${IMAGE_NAME}:latest\$"; then
    docker rmi $IMAGE_NAME
  else
    echo "Image $IMAGE_NAME not found, skipping remove."
  fi
done

docker system prune --all --volumes --force