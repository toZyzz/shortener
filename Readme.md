# Сокращение ссылок на golang
---

## Описание
Http сервис, реализующий методы:
- Сохранение оригинального URL в хранилище и возврат сокращённого
- Возвращение оригинального URL через сокращённый

Для выбора типа хранилища (In-memory/PostgreSQL) необходимо изменить параметр MODE в файле .env на in_memory/db соответственно.

## Установка
1. Клонируйте репозиторий:
    ```bash
    git clone
    ```

2. Перейдите в директорию проекта:
    ```bash
    cd shortener
    ```

3. Измените `.env` указав свои настройки:
    ```plaintext
    PORT=8080
    DATABASE_URL="host=localhost user=postgres password=root dbname=urls port=5433 sslmode=disable"
    MODE="in_memory"
    REDIS_URL=172.29.86.99:6379
    REDIS_PASSWORD=
    REDIS_DB=0
    ```

## Usage
Запустите приложение:
   ```bash
    go run main.go
   ```