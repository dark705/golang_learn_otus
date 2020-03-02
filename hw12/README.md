# Календарь

## Работа с БД
SQL файл для инициализации БД:
    ./build/sql/init.sql
    
### Тесты
Для эмуляции взаимодействия с БД необходима реальная БД Postgres 11.
В комплекте идёт docker-compose файл разворачивающий ей:

    ./build/docker.up.sh

После этого становятся доступны интеграционный тесты
    
    go test -v github.com/dark705/otus/hw12/internal/storage

## GRPC
### Запуск сервера
    go run ./cmd/server.go
  
### Запуск тестового клиента
    go run ./cmd/client.go
    
### API cпецификация Protobuf
    ./api/protobuf.proto
    
### Скрипт по генерации grpc под Linux (Windows)
    ./build/protobuf.grpc.sh (bat)
Выходная дирректория по умолчанию для сгенерированного кода:
_-/pkg/calendar/protobuf_

### Тесты
    go test -v github.com/dark705/otus/hw12/internal/grpc
Тест поднимает grpc  сервер и эмулирует работу клиента

