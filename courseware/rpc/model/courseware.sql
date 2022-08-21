CREATE TABLE `courseware`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `code`        varchar(255) NOT NULL DEFAULT '' COMMENT '编号',
    `name`        varchar(255) NOT NULL DEFAULT '' COMMENT '用户姓名',
    `type`        tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '课件类型 1-h5 2-scorm 3-多媒体多章节',
    `is_delete`   tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0-未删除 1-已删除',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `udx_code` (`code`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
