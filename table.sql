CREATE TABLE file_list (
  id integer primary key autoincrement,
  file_name varchar(255) not null,
  nick_name varchar(255) not null,
  version varchar null,
  md5_value varchar(255) not null,
  user_name varchar(255) not null,
  desc varchar(1024) null,
  create_time long not null,
  update_time long not null
);

CREATE TABLE user_list (
  id integer primary key autoincrement,
  user_name varchar(255) not null,
  passwd varchar(255) not null,
  create_time long not null
);