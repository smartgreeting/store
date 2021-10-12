create table `hc_user`(
    `id` int unsigned not null AUTO_INCREMENT,
    `username` varchar(32) not null default "" comment "用户名",
    `password` varchar(32) not null comment "密码",
    `avatar` varchar(128) not null default "" comment "头像地址",
    `sex` tinyint(1) not null default 0 comment "性别：0.男 1.女",
    `phone` varchar(11) not null default "" comment "手机号",
    `email` varchar(64) comment "邮箱",
    `address` varchar(64) comment "地址",
    `hobby` varchar(128) comment "爱好",
    `deleted` tinyint(1) not null default 0 comment "是否注销",
    `created_at` int(16) comment "创建时间",
    `updated_at` int(16) comment "更新时间",
    primary key (`id`) using BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商城用户表';