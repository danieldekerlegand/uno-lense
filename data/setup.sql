create table users (
  id         integer primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null
);

create table remote_connections (
  id         integer primary key,
  uuid       varchar(64) not null unique,
  ip         varchar(64) not null,
  name       varchar(255),
  connected  integer
);

create table sessions (
  id         integer primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);

create table lessons (
  id         integer primary key,
  uuid       varchar(64) not null unique,
	topic 		 text,
  details    text,
	base_image varchar(255),
	published  integer,
  user_id    integer references users(id),
  created_at timestamp not null
);

create table components (
  id         integer primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  lesson_id  integer references lessons(id),
  created_at timestamp not null
);
