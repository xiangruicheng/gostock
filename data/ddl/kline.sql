CREATE TABLE IF NOT EXISTS `kline` (
     `id` int unsigned NOT NULL AUTO_INCREMENT,
     `code` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '代码',
     `date` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '日期',
     `volume` decimal(20,0) NOT NULL DEFAULT '0' COMMENT '成交量-手',
     `amount` decimal(20,0) NOT NULL DEFAULT '0' COMMENT '成交额-金额',
     `open` decimal(20,4) NOT NULL DEFAULT '0.0000' COMMENT '开盘价',
     `high` decimal(20,4) NOT NULL DEFAULT '0.0000' COMMENT '最高',
     `low` decimal(20,4) NOT NULL DEFAULT '0.0000' COMMENT '最低',
     `close` decimal(20,4) NOT NULL DEFAULT '0.0000' COMMENT '收盘价',
     `chg` decimal(20,4) NOT NULL DEFAULT '0.0000' COMMENT '涨跌幅度',
     `percent` decimal(20,4) NOT NULL DEFAULT '0.0000' COMMENT '涨跌百分比',
     `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
     `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
     PRIMARY KEY (`id`) USING BTREE,
     UNIQUE KEY `uniq_key` (`code`,`date`) USING BTREE,
     KEY `idx_code` (`code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='k线';