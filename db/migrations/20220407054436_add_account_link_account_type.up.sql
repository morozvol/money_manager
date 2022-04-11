alter table account
    add id_account_type int default 1 not null;

alter table account
    add constraint account_account_type_id_fk
        foreign key (id_account_type) references account_type;

