#!/bin/bash

set -e
trap 'echo "🔴 Скрипт остановлен"; exit' SIGINT

# Должен ли порт быть доступен снаружи
OPEN_PORT_EXTERNALLY="true"
# Порт хоста
HOST_PORT=11001
# Наименование сервиса
SERVICE_NAME="main_app"

# Имя сети Docker
DOCKER_NETWORK_NAME="networklocal1"
# Список дополнительных сетей Docker, к которым нужно подключить контейнер
ADDITIONAL_DOCKER_NETWORKS=""

# IP-адрес для добавления в hosts контейнера
DATABASE_HOST_IP=""

W_BUILD="false"
W_START="false"

# Разбор именованных аргументов
while [ $# -gt 0 ]; do
  case "$1" in
    OPEN_PORT_EXTERNALLY=*)
      OPEN_PORT_EXTERNALLY="${1#*=}"
      ;;
    HOST_PORT=*)
      HOST_PORT="${1#*=}"
      ;;
    SERVICE_NAME=*)
      SERVICE_NAME="${1#*=}"
      ;;
    DOCKER_NETWORK_NAME=*)
      DOCKER_NETWORK_NAME="${1#*=}"
      ;;
    DATABASE_HOST_IP=*)
      DATABASE_HOST_IP="${1#*=}"
      ;;
    W_BUILD=*)
      W_BUILD="${1#*=}"
      ;;
    W_START=*)
      W_START="${1#*=}"
      ;;
    *)
      echo "Ошибка: Неизвестный параметр $1" >&2
      exit 1
      ;;
  esac
  shift
done


# Наименование образа
IMAGE_NAME="${SERVICE_NAME}_i"
# Наименование контейнера
CONTAINER_NAME="${SERVICE_NAME}_c"
# Имя приложения
APP_NAME="${SERVICE_NAME}_app"

if [ "$W_BUILD" = "true" ]; then
  # Билд бинарника
  ./build.sh $APP_NAME

  # Собираем образ
  docker build  --no-cache --build-arg APP_NAME=$APP_NAME -t $IMAGE_NAME .
fi


# Проверка и выполнение запуска контейнера
if [ "$W_START" = "true" ]; then
  # Останавливаем и удаляем существующий контейнер, если он существует
  # Проверяем существование контейнера перед его остановкой и удалением
  if docker ps -a --format "{{.Names}}" | grep -w "^$CONTAINER_NAME$"; then
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
  fi

  # Определяем, нужно ли открывать порт для внешнего мира
  PORT_MAPPING=""
  if [ "$OPEN_PORT_EXTERNALLY" = "true" ]; then
    PORT_MAPPING="-p $HOST_PORT:$HOST_PORT"
  fi

  if [ -n "$DATABASE_HOST_IP" ]; then
    ADD_HOST_OPTION="--add-host=database_host:$DATABASE_HOST_IP"
  else
    ADD_HOST_OPTION=""
  fi

  # Проверяем существование сети и создаем, если не существует
  docker network ls | grep $DOCKER_NETWORK_NAME || docker network create $DOCKER_NETWORK_NAME

  # Запускаем контейнер с новым образом в сетевом режиме host и подключаем к сети по имени
  docker run -d \
  $PORT_MAPPING \
  $ADD_HOST_OPTION \
  --network $DOCKER_NETWORK_NAME \
  --restart always \
  --name $CONTAINER_NAME \
  -e APP_NAME=$APP_NAME \
  $IMAGE_NAME

  # Подключаем контейнер к дополнительным сетям Docker, если список не пустой
  if [ -n "$ADDITIONAL_DOCKER_NETWORKS" ]; then
    for network in $ADDITIONAL_DOCKER_NETWORKS; do
        docker network connect $network $CONTAINER_NAME
        echo "🟢 Container $CONTAINER_NAME is connected to additional network $network"
    done
  fi

  echo "🟢 Container $CONTAINER_NAME is running on port $HOST_PORT"
fi

echo "🟢 Скрипт завершен $APP_NAME"
