create table category
(
    id                 serial
        constraint category_pk
            primary key,
    name               text not null,
    type               int  not null,
    id_owner           int
        constraint category_user_id_fk
            references "user",
    id_parent_category int
        constraint category_category_id_fk
            references category
);