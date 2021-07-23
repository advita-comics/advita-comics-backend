-- +goose Up
create table if not exists company
(
    id                        integer primary key auto_increment,
    name                      varchar(255)            not null unique,
    termination_amount        integer                 not null,
    expiration_date           timestamp               not null,

    status                    SMALLINT  default '10'  not null,
    created_at                TIMESTAMP default NOW() not null,
    updated_at                TIMESTAMP default NOW() not null
)CHARACTER SET utf8 COLLATE utf8_unicode_ci ENGINE=INNODB;

create table if not exists user
(
    id                        integer primary key auto_increment,
    name                      varchar(255),
    email                     varchar(255)            not null unique,
    get_report                bool                    not null default false,
    follow_process            bool                    not null default false,
    role                      varchar(255)            not null default 'USER',

    status                    SMALLINT  default '10'  not null,
    created_at                TIMESTAMP default NOW() not null,
    updated_at                TIMESTAMP default NOW() not null
)CHARACTER SET utf8 COLLATE utf8_unicode_ci ENGINE=INNODB;

create table if not exists donation_type
(
    id                        integer primary key auto_increment,
    name                      varchar(255)            not null unique,
    min_amount                float                   not null,
    description               text,

    status                    SMALLINT  default '10'  not null,
    created_at                TIMESTAMP default NOW() not null,
    updated_at                TIMESTAMP default NOW() not null
)CHARACTER SET utf8 COLLATE utf8_unicode_ci ENGINE=INNODB;

create table if not exists comics
(
    id                        integer primary key auto_increment,
    name                      varchar(255)            not null unique,
    description               text                    not null,
    path                      varchar(255)            not null unique,

    status                    SMALLINT  default '10'  not null,
    created_at                TIMESTAMP default NOW() not null,
    updated_at                TIMESTAMP default NOW() not null
)CHARACTER SET utf8 COLLATE utf8_unicode_ci ENGINE=INNODB;

create table if not exists donation
(
    id                        integer primary key auto_increment,
    amount                    float                   not null,
    company_id                integer                 not null,
    donation_type_id          integer                 not null,
    user_id                   integer                 not null,
    comics_id                 integer                 not null,
    personalisation           json                    not null,

    status                    SMALLINT  default '10'  not null,
    created_at                TIMESTAMP default NOW() not null,
    updated_at                TIMESTAMP default NOW() not null,

    INDEX company_id_idx (company_id),
    FOREIGN KEY (company_id)
        REFERENCES company(id)
        ON DELETE CASCADE,

    INDEX donation_type_id_idx (donation_type_id),
    FOREIGN KEY (donation_type_id)
        REFERENCES donation_type(id)
        ON DELETE CASCADE,

    INDEX user_id_idx (user_id),
    FOREIGN KEY (user_id)
        REFERENCES user(id)
        ON DELETE CASCADE,

    INDEX comics_id_idx (comics_id),
    FOREIGN KEY (comics_id)
        REFERENCES comics(id)
        ON DELETE CASCADE
)CHARACTER SET utf8 COLLATE utf8_unicode_ci ENGINE=INNODB;

create table if not exists transaction
(
    id                       integer primary key auto_increment,
    amount                    float                   not null,
    donation_id               integer                 not null,
    external_id               integer                 not null,
    request                   text,
    response                  text,

    status                    SMALLINT  default '10'  not null,
    created_at                TIMESTAMP default NOW() not null,
    updated_at                TIMESTAMP default NOW() not null,

    INDEX donation_id_idx (donation_id),
    FOREIGN KEY (donation_id)
        REFERENCES donation(id)
        ON DELETE CASCADE
) CHARACTER SET utf8 COLLATE utf8_unicode_ci ENGINE=INNODB;

-- +goose Down
drop table if exists company;
drop table if exists user;
drop table if exists donation_type;
drop table if exists comics;
drop table if exists donation;
drop table if exists transaction;

