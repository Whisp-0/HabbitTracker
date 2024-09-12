#!/bin/bash

# Запуск контейнера с MongoDB
echo "Запуск контейнера с MongoDB..."
docker start mongoToDoRPG

# Проверка, запустился ли контейнер
if [ $? -eq 0 ]; then
    echo "Контейнер с MongoDB успешно запущен."
else
    echo "Ошибка при запуске контейнера с MongoDB."
    exit 1
fi

# Запуск вашего Go-приложения
echo "Запуск Go-приложения..."
go run cmd/main.go

# Проверка, запустилось ли приложение
if [ $? -eq 0 ]; then
    echo "Go-приложение успешно запущено."
else
    echo "Ошибка при запуске Go-приложения."
    exit 1
fi
