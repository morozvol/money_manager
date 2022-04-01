alter table account
    alter column id_user drop not null;

alter table account
    drop constraint account_user_id_fk;
