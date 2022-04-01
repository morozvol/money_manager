drop table cross_account_operation;

alter table operation
    add id_account int not null;
