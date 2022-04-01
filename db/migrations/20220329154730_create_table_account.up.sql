create table account
(
    id          serial
        constraint account_pk
            primary key,
    balance     float default 0 not null,
    id_currency int             not null
        constraint account_currency_id_fk
            references currency
);

