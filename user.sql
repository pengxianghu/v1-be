create table `user`(
	`id` varchar(64) not null comment 'user_id',
	`name` varchar(64) not null unique comment 'user_name',
	`pwd` varchar(64) not null comment 'user_pwd',
	primary key (`id`)
) comment '用户表';