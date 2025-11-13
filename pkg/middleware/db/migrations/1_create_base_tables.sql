create table if not exists users (
    id bigserial primary key,
    username varchar not null unique,
    email varchar not null unique,
    password_encrypted varchar not null,
    birth_date date,
    created_at timestamp,
    updated_at timestamp
);
create unique index if not exists users_email on users using btree (email);
create unique index if not exists users_username on users using btree (username);

create table if not exists roles (
    id bigserial primary key,
    name varchar not null unique,
    permissions jsonb
);

create table if not exists user_roles (
    id bigserial primary key,
    user_id bigint references users not null,
    role_id bigint references roles not null
);
create index if not exists user_roles_user on user_roles using btree (user_id);
create index if not exists user_roles_role on user_roles using btree (role_id);

create table if not exists categories (
    id bigserial primary key,
    name varchar not null,
    description varchar
);

create table if not exists questions (
    id bigserial primary key,
    user_id bigint references users not null,
    category_id bigint references categories not null,
    header varchar not null,
    description varchar,
    is_closed bool,
    created_at timestamp,
    updated_at timestamp
);
create index if not exists question_user on questions using btree (user_id);
create index if not exists question_category on questions using btree (category_id);

create table if not exists answers (
    id bigserial primary key,
    user_id bigint references users not null,
    question_id bigint references questions not null,
    text varchar not null,
    is_best bool default false,
    created_at timestamp,
    updated_at timestamp
);
create index if not exists answer_user on answers using btree (user_id);
create index if not exists answer_question on answers using btree (question_id);

create table if not exists rates (
    id bigserial primary key,
    rankable_type varchar not null,
    rankable_id bigint not null,
    user_id bigint references users not null,
    rank int not null default 1 check (rank between 1 and 5),
    created_at timestamp,
    updated_at timestamp
);
create index if not exists rate_user on rates using btree (user_id);
create index if not exists rate_ratable on rates using btree (rankable_id, rankable_type)