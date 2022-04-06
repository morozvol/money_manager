alter table "user"
    rename column  id_default_currency to id_status;

alter table "user"
    drop constraint user_currency_id_fk;

