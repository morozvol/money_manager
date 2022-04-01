alter table operation
    add constraint operation_account_id_fk
        foreign key (id_account) references account;
