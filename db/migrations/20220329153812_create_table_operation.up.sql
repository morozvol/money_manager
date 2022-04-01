create table operation
(
    id   serial
        constraint operation_pk
            primary key,
    time timestamp not null,
    sum  float     not null,
    type int       not null
);