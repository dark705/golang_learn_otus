create sequence events_id_seq as integer;
alter sequence events_id_seq owner to postgres;

create table events
(
    id          serial                   not null constraint events_pk primary key,
    start_time  timestamp with time zone not null,
    end_time    timestamp with time zone not null,
    title       varchar(256)             not null,
    description text
);

alter table events
    owner to tmp;