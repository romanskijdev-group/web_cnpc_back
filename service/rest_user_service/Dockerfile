FROM ubuntu:22.04

WORKDIR /app

# Определяем аргументы сборки для переменных среды Vault
ARG VAULT_ADDR
ARG VAULT_TOKEN

# Устанавливаем переменные среды из аргументов сборки
ENV VAULT_ADDR=${VAULT_ADDR}
ENV VAULT_TOKEN=${VAULT_TOKEN}

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