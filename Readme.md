###### 一、sqlite建表：
```
drop table file_list if exists;
CREATE TABLE file_list (
  id integer primary key autoincrement,
  file_name varchar(255) not null,
  file_size integer not null,
  url_name varchar(255) not null,
  version varchar(255) null,
  md5_value varchar(255) not null,
  user_name varchar(255) not null,
  desc blob null,
  create_time integer not null,
  update_time integer not null
);
CREATE UNIQUE INDEX url_index on file_list (url_name);

drop table user_list if exists;
CREATE TABLE user_list (
  id integer primary key autoincrement,
  user_email varchar(255) not null,
  user_name varchar(255) default '',
  user_right long DEFAULT 1,
  passwd varchar(255) default '',
  create_time long null
);
CREATE UNIQUE INDEX user_index on user_list (user_email);
-- 下载者
INSERT INTO user_list(user_email, user_name, user_right, passwd, create_time) VALUES ('download', '下载者', 1, 'e10adc3949ba59abbe56e057f20f883e', 1092941466);
-- 上传者
INSERT INTO user_list(user_email, user_name, user_right, passwd, create_time) VALUES ('upload', '上传者', 3, 'e10adc3949ba59abbe56e057f20f883e', null);
-- 管理员
INSERT INTO user_list(user_email, user_name, user_right, passwd, create_time) VALUES ('admin', '管理员', 7, 'e10adc3949ba59abbe56e057f20f883e', 1092941466);
```

###### 二、运行：
修改conf.yaml， 主要配置
logs
templates
db
data
listen_port
cookie_secret


###### 三、安装第三主插件
sh prepare.sh

###### 四、运行
nohup http_server --conf=conf.yaml &