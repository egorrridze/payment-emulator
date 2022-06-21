CREATE TABLE IF NOT EXISTS payments (
    id              serial          not null unique,
    user_id         integer         not null,
    user_email      varchar(255)    not null,
    summ            integer         not null,
    currency        varchar(255)    not null,
    creation_time   timestamp       not null,
    update_time     timestamp       not null,
    status          varchar(255)    not null
);