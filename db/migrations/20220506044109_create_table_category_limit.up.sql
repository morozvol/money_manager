create table category_limit
(
    id          serial
        constraint category_limit_pk
            primary key,
    id_user     int   not null,
    id_category int   not null,
    id_currency int   not null,
    sum         float not null
);