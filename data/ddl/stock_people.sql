CREATE TABLE IF NOT EXISTS  `stock_people` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `code` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '代码',
    `date` VARCHAR(8) NOT NULL DEFAULT '' COMMENT '股东人数统计日期',
    `holder_num` INT(11) NOT NULL DEFAULT '0' COMMENT '户数',
    `avg_market` float(20,6) NOT NULL DEFAULT '0' COMMENT '人均市值',
    `avg_hold_num` float(20,6) NOT NULL DEFAULT '0' COMMENT '人均股数',
    `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `idx_code_date` (`code`, `date`) USING BTREE
    )
    COMMENT='股票股东户数信息'
    COLLATE='utf8_general_ci'
    ENGINE=InnoDB
;