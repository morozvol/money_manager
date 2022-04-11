create table account_type
(
    id          serial
        constraint account_type_pk
            primary key,
    name        text not null,
    symbol      text not null,
    description text
);

