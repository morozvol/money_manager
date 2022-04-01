alter table cross_account_operation
    rename column id_operation to id_currency;

alter table cross_account_operation
    rename to cross_account_currency;

alter table cross_account_currency
    drop constraint cross_account_operation_account_id_fk;

alter table cross_account_currency
    drop constraint cross_account_operation_operation_id_fk;