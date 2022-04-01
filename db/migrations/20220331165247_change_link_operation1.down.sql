create table cross_user_account
(
    id         serial
        constraint cross_user_account_pk
            primary key,
    id_user    integer not null
        constraint cross_user_account_user_id_fk
            references "user",
    id_account integer not null
        constraint cross_user_account_account_id_fk
            references account
);

alter table cross_user_account
    owner to postgres;
