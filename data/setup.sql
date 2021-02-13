drop table works;
drop table sessions;
drop table users;

create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);

create table works (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  date       varchar(64) not null unique,
  title      varchar(64) not null unique,
  money      integer not null unique,
  station    varchar(64) not null unique,
  job_id     integer not null unique,
  working_time_start varchar(64) not null unique,
  working_time_finish varchar(64) not null unique,
  business_detail text not null unique,
  Evaluation int,
  created_at timestamp not null
);
