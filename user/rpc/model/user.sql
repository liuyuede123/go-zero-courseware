CREATE TABLE `user`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `login_name`  varchar(255) NOT NULL DEFAULT '' COMMENT '登录名',
    `username`        varchar(255) NOT NULL DEFAULT '' COMMENT '用户姓名',
    `sex`      tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
    `password`    varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
    `is_delete`   tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0-未删除 1-已删除',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `udx_login_name` (`login_name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;