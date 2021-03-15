drop table messages;
drop table apply_users;
drop table sessions;
drop table works;
drop table users;


create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  avatar_url varchar(255),
  avatar_id  varchar(255),
  created_at timestamp not null
);

create table works (
  id         serial primary key,
  date       varchar(255),
  title      varchar(255),
  money      varchar(255),
  job_id     varchar(255),
  evaluation varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);

create table apply_users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  user_id    integer references users(id),
  work_id    integer references works(id),
  created_at timestamp not null
);

create table messages (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  text       varchar(255),
  user_id    integer references users(id),
  user_name  varchar(255),
  work_id    integer references works(id),
  chat_uuid  varchar(64) not null unique,
  created_at timestamp not null
);
