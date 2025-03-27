# Meebin backend

## Setup

Для поднятия Backend-сервера нужно:

1. Склонировать главный репозиторий
2. Перейти в директорию backend
3. Поднять контейнеры с backend:
    1. Для запуска готового приложения с миграциями: ввести команду docker compose -f docker-compose.bin.yml up
    2. Для запуска приложения с hot-reload и без миграций: ввести команду docker compose -f docker-compose.air.yml

На порту, который указан в .env-файле, будет поднят backend-сервер.

Для накатки миграций в случае работы с hot-reload сервером следует использовать Makefile