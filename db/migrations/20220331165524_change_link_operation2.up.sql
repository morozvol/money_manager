alter table account
    alter column id_user set not null;

alter table account
    add constraint account_user_id_fk
        foreign key (id_user) references "user";