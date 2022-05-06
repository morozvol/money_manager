create table exchange_rate
(
    id               serial
        constraint exchange_rate_pk
            primary key,
    id_currency_from int             not null,
    id_currency_to   int             not null,
    rate             float default 1 not null,
    date             date            not null
);