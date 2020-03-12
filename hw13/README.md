# Календарь

## GRPC сервер
### Запуск сервера
    go run ./cmd/server.go
    либо    
    make run_server

Следует учитывать что сервер использует реальную БД, к которой должен иметь доступ. 
В комплекте идёт Docker с настроенной БД
  
### Запуск тестового клиента
    go run ./cmd/client.go
    либо
    make run_client
      
### API cпецификация Protobuf
    ./api/protobuf.proto
    
### Скрипт по генерации grpc под Linux (Windows)
    ./build/protobuf.grpc.sh (bat)
    либо
    make build_grpc

Выходная дирректория по умолчанию для сгенерированного кода:
_-/pkg/calendar/protobuf_

### Тесты
    go test -v github.com/dark705/otus/hw12/internal/grpc
Тест поднимает grpc сервер (используется inmemory хранилище), и эмулирует работу клиента

## Работа с БД Postgres
SQL файл для инициализации БД:
    ./build/sql/init.sql

Пакет реализующий интерфейс хранилища:
    ./internal/storage/postgres.go

### Тесты
Для эмуляции взаимодействия с БД необходима реальная БД Postgres 11.
В комплекте идут docker - скрипты:

    ./build/docker.up.sh #(создание контейнера с тестовой бд)
    либо
    make docker_up
    
    ./build/docker.down.sh #(для удаления контейнеров и образов)
    либо
    make docker_down 

В процессе развёртывания docker контейнера, начальное состояние автоматически загружается из _./build/sql/init.sql_.
После этого становятся доступны интеграционный тесты с БД:
    
    go test -v github.com/dark705/otus/hw12/internal/storage
    
Запустить все тесты можно также с подготовленным окружением:

    make test

При этом развёртывается контейнер, прогоняются тесты, после чего контейнер удаляется.
