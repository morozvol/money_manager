create table currency
(
    id   serial
        constraint currency_pk
            primary key,
    name text not null,
    code text not null
);

create unique index currency_code_uindex
    on currency (code);
