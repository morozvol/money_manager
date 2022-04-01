alter table cross_account_currency
    rename column id_currency to id_operation;

alter table cross_account_currency
    rename to cross_account_operation;

alter table cross_account_operation
    add constraint cross_account_operation_account_id_fk
        foreign key (id_account) references account;

alter table cross_account_operation
    add constraint cross_account_operation_operation_id_fk
        foreign key (id_operation) references operation;