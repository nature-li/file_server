-- 列出表名: .tables
-- 插入数据: INSERT INTO table_name (col1, col2) VALUES(val1, val2);
-- 查询数据: SELECT * FROM table_name;
-- 01 下载者
-- 02 上传者
-- 04 管理员

drop table file_list;
CREATE TABLE file_list (
  id integer primary key autoincrement,
  file_name varchar(255) not null,
  file_size integer not null,
  url_name varchar(255) not null,
  version varchar null,
  md5_value varchar(255) not null,
  user_name varchar(255) not null,
  desc blob null,
  create_time integer not null,
  update_time integer not null
);
CREATE UNIQUE INDEX url_index on file_list (url_name);

drop table user_list;
CREATE TABLE user_list (
  id integer primary key autoincrement,
  user_email varchar(255) not null,
  user_name varchar(255) null,
  user_right integer DEFAULT 1,
  passwd varchar(255) null,
  create_time long null
);
CREATE UNIQUE INDEX user_index on user_list (user_email);
-- 下载者
INSERT INTO user_list(user_email, user_name, user_right, passwd, create_time) VALUES ('download@meitu.com', '下载者', 1, 'e10adc3949ba59abbe56e057f20f883e', datetime(1092941466, 'unixepoch'));
-- 上传者
INSERT INTO user_list(user_email, user_name, user_right, passwd, create_time) VALUES ('upload@meitu.com', '上传者', 3, 'e10adc3949ba59abbe56e057f20f883e', datetime(1092941466, 'unixepoch'));
-- 管理员
INSERT INTO user_list(user_email, user_name, user_right, passwd, create_time) VALUES ('manager@meitu.com', '管理员', 7, 'e10adc3949ba59abbe56e057f20f883e', datetime(1092941466, 'unixepoch'));