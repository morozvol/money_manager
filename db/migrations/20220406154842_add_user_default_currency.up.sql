alter table "user"
    rename column id_status to id_default_currency;

alter table "user"
    add constraint user_currency_id_fk
        foreign key (id_default_currency) references currency;

