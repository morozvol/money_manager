alter table account
    drop column id_account_type;

alter table account
    drop constraint account_account_type_id_fk;

