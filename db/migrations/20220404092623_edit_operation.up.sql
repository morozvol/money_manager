alter table operation
    rename column type to id_category;

alter table operation
    add constraint operation_category_id_fk
        foreign key (id_category) references category;

