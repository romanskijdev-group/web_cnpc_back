FROM ubuntu:22.04

WORKDIR /app

# Определяем аргумент сборки
ARG APP_NAME

COPY ./${APP_NAME} /app/${APP_NAME}
COPY ./.env /app/.env

# Обновляем хранилище доверенных сертификатов
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*
    
RUN chmod +x /app/${APP_NAME}

RUN echo "APP_NAME=${APP_NAME}"

CMD ["sh", "-c", "/app/${APP_NAME}"]