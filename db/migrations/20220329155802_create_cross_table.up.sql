create table cross_account_currency
(
    id          serial
        constraint cross_account_currency_pk
            primary key,
    id_account  int not null,
    id_currency int not null
);

