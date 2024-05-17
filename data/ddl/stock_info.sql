CREATE TABLE IF NOT EXISTS `stock_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '代码',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `market` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '市场 SH SZ',
  `cyb` tinyint(1) NOT NULL DEFAULT '0' COMMENT '创业板标识 0不是 1是',
  `hs300` tinyint(1) NOT NULL DEFAULT '0' COMMENT '沪深300标识 0不是 1是',
  `kcb` tinyint(1) NOT NULL DEFAULT '0' COMMENT '科创板标识 0不是 1是',
  `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_code` (`code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='股票信息';
