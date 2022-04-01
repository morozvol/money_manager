create table cross_account_operation
(
    id           integer default nextval('cross_account_currency_id_seq'::regclass) not null
        constraint cross_account_currency_pk
            primary key,
    id_account   integer                                                            not null
        constraint cross_account_operation_account_id_fk
            references account,
    id_operation integer                                                            not null
        constraint cross_account_operation_operation_id_fk
            references operation
);

alter table cross_account_operation
    owner to postgres;

alter table operation
    drop column id_account;
