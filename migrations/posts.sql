create table if not exists posts
(
    id          varchar(100)                        not null primary key,
    title       varchar(100)                        not null,
    description varchar(100)                        not null,
    created_at  timestamp default current_timestamp not null,
    updated_at  timestamp default current_timestamp,
    canceled_at timestamp default current_timestamp
);