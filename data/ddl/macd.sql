CREATE TABLE IF NOT EXISTS `macd` (
    `id` int NOT NULL AUTO_INCREMENT,
    `code` varchar(20) NOT NULL DEFAULT '' COMMENT '代码',
    `date` varchar(8) NOT NULL DEFAULT '' COMMENT '日期',
    `close` float(20,6) NOT NULL DEFAULT '0' COMMENT 'close',
    `ema12` float(20,6) NOT NULL DEFAULT '0' COMMENT 'ema12',
    `ema26` float(20,6) NOT NULL DEFAULT '0' COMMENT 'ema26',
    `diff` float(20,6) NOT NULL DEFAULT '0' COMMENT 'diff',
    `dea` float(20,6) NOT NULL DEFAULT '0' COMMENT 'dea',
    `macd` float(20,6) NOT NULL DEFAULT '0' COMMENT 'macd',
    `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_code_date` (`code`,`date`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='macd表';