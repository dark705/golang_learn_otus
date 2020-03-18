# Календарь
Параметры подключения к RabbitMQ, Postgres, периодичность проверки на необходимость отпраки уведомлений, число отправщиков и другие параметры, находятся в конфигурационном файле:

_./config/config.yaml_

## Scheduler сервер
### Запуск сервера
    go run ./cmd/sheduler.go
    либо    
    make run_sheduler

### Запуск отправщиков
    go run ./cmd/sender.go
    либо    
    make run_sender

### Компиляция сервера и отправщика
    make build_scheduler
    make build_sender

Следует учитывать что сервер и отправщик используют реальную БД Postgres, а также реальный брокер RabbitMQ к которым должен иметь доступ. 
В комплекте идёт Docker с настроенными компонентами, для их запуска и остановки достатовно выполнить:

    ./build/docker.up.sh #(создание контейнера с тестовой бд)
    либо
    make docker_up
    
    ./build/docker.down.sh #(для удаления контейнеров и образов)
    либо
    make docker_down 

## GRPC сервер
### Запуск сервера
    go run ./cmd/api.go
    либо    
    make run_api

Следует учитывать что сервер использует реальную БД, к которой должен иметь доступ. 
В комплекте идёт Docker с настроенной БД
  
### Запуск тестового клиента
    go run ./cmd/api.client.go
    либо
    make run_api_client
      
### API cпецификация Protobuf
    ./api/protobuf.proto
    
### Скрипт по генерации grpc под Linux (Windows)
    ./build/protobuf.grpc.sh (bat)
    либо
    make build_grpc

Выходная дирректория по умолчанию для сгенерированного кода: 

_/pkg/calendar/protobuf_

### Тесты
    go test -v github.com/dark705/otus/hw14/internal/grpc
Тест поднимает grpc сервер (используется inmemory хранилище), и эмулирует работу клиента

## Работа с БД Postgres
SQL файл для инициализации БД:

_./build/sql/init.sql_

Пакет реализующий интерфейс хранилища:
    
_./internal/storage/postgres.go_

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
    
    go test -v github.com/dark705/otus/hw14/internal/storage
    
Запустить все тесты можно также с подготовленным окружением:

    make test

При этом развёртывается контейнер, прогоняются тесты, после чего контейнер удаляется.
