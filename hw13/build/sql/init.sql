create table if not exists events
(
    id           serial                   not null constraint events_pk primary key,
    start_time   timestamp with time zone not null,
    end_time     timestamp with time zone not null,
    title        varchar(256)             not null,
    description  text,
    is_scheduled boolean default false not null
);
