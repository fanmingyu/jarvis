-- 短信记录表
CREATE TABLE `sms_record` (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `app` varchar(50) NOT NULL COMMENT '应用名称',
    `mobile` varchar(20) NOT NULL COMMENT '短信接收的手机号',
    `content` varchar(1024) NOT NULL COMMENT '短信内容',
    `channel` varchar(30) NOT NULL COMMENT '短信所走的通道',
    `msgid` varchar(100) NOT NULL DEFAULT '' COMMENT'第三方返回的短信id',
    `create_time` int(10) UNSIGNED NOT NULL COMMENT '短信创建时间',
    `request_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '请求第三方的时间',
    `request_status` tinyint(1) UNSIGNED NOT NULL COMMENT '短信向第三方请求的状态，1——已保存，2——请求失败，3——请求成功',
    `report_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '短信回执中的第三方返回的时间',
    `report_recv_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收到短信回执的时间',
    `report_status` varchar(20) NOT NULL DEFAULT '' COMMENT '短信回执返回的短信状态',

    PRIMARY KEY (`id`),
    KEY `idx_mobile` (`mobile`),
    KEY `idx_msgid` (`msgid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='短信记录';

ALTER TABLE `sms_record` ADD INDEX idx_create_time(`create_time`);
ALTER TABLE `sms_record` ADD COLUMN num int(3) NOT NULL DEFAULT 0;
ALTER TABLE `sms_record` ADD COLUMN tpl_id int(10) NOT NULL DEFAULT 0 COMMENT '短信发送所使用的模板id';

-- 应用配置表
CREATE TABLE `sms_app` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL COMMENT '应用名',
    `secret` varchar(50) NOT NULL COMMENT '秘钥',
    `prefix` varchar(50) NOT NULL COMMENT '签名，短信前缀',
    `channel` varchar(30) NOT NULL COMMENT '短信所走通道',
    `worker` varchar(20) NOT NULL COMMENT '短信所用的worker',
    `create_time` int(10) UNSIGNED NOT NULL COMMENT '配置创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',

    PRIMARY KEY(`id`),
    UNIQUE KEY `uni_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='短信配置——应用配置';

-- 模板配置表
CREATE TABLE `sms_tpl` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL COMMENT '模板名称',
    `content` varchar(1024) NOT NULL COMMENT '模板内容',
    `create_time` int(10) UNSIGNED NOT NULL COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL COMMENT '更新时间',

    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='短信配置——模板配置';

-- 黑名单列表
CREATE TABLE `sms_blacklist` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `mobile` varchar(20) NOT NULL COMMENT '黑名单手机号',
    `create_time` int(10) UNSIGNED NOT NULL COMMENT '创建时间',

    PRIMARY KEY(`id`),
    UNIQUE KEY `uni_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='短信配置——黑名单列表';

-- 上行短信记录表
CREATE TABLE `sms_reply` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `msgid` varchar(100) NOT NULL COMMENT '信息编号，第三方给短信的编号',
    `sp_code` varchar(30) NOT NULL COMMENT '上行短信的目标地址',
    `mobile` varchar(30) NOT NULL COMMENT '用户的手机号',
    `content` varchar(1024) NOT NULL COMMENT '上行短信内容',
    `recv_time` int(10) NOT NULL COMMENT '上行短信的时间',
    `create_time` int(10) NOT NULL COMMENT '接收上行短信的时间',
    `account` varchar(100) NOT NULL COMMENT '第三方的账号',

    PRIMARY KEY (`id`),
    KEY `idx_mobile` (`mobile`),
    KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='上行短信记录';
