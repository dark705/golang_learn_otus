# Календарь
Параметры подключения к RabbitMQ, Postgres, периодичность проверки на необходимость отпраки уведомлений, число отправщиков и другие параметры, находятся в конфигурационном файле:

_./config/config.yaml_

## Docker
Все файлы относящиеся к Docker находятся в папке _./build/docker_

Подготовка окружения:

    docker-compose -f ./build/docker/docker-compose-environment.yml up -d
    
Собрать и запустить приложение:
    
    docker-compose -f ./build/docker/docker-compose.yml up
    
Последовательно запустить окружение, микрсервисы и выполнить интеграционные тесты:

    make test

Последовательно запустить окружение и микрсервисы:

    ./build/run.sh
        
## Компиляция
    make build
        
При этом компилируются соответствующие файлы в ./bin
      
### API cпецификация Protobuf
    ./api/protobuf.proto
    
### Скрипт по генерации grpc под Linux (Windows)
    ./build/protobuf.grpc.sh (bat)
    либо
    make build_grpc

Выходная дирректория по умолчанию для сгенерированного кода: 

_/pkg/calendar/protobuf_

### Тесты
    go test -v github.com/dark705/otus/hw15/internal/grpc
Тест поднимает grpc сервер (используется inmemory хранилище), и эмулирует работу клиента

## Работа с БД Postgres
SQL файл для инициализации БД:

_./migrations/init.sql_

### Тесты
Для эмуляции взаимодействия с БД необходима реальная БД Postgres 11.
В комплекте идёт настроенное окружение:

    make docker_up_environment
    
В процессе развёртывания docker контейнера, начальное состояние автоматически загружается из _./migrations/init.sql_.
После этого становятся доступны интеграционный тесты с БД:
    
    go test -v github.com/dark705/otus/hw15/internal/storage
    
Запустить все тесты можно также с подготовленным окружением:

    make test

При этом развёртывается контейнеры, прогоняются тесты, после чего контейнеры удаляется.
