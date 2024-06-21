CREATE TABLE IF NOT EXISTS `trade_user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(64) NOT NULL DEFAULT '' COMMENT 'name',
    `recharge_money` float(20,6) NOT NULL DEFAULT '0' COMMENT '充值金额',
    `hold_money` float(20,6) NOT NULL DEFAULT '0' COMMENT '持仓金额',
    `usable_money` float(20,6) NOT NULL DEFAULT '0' COMMENT '可用金额',
    `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单表';


CREATE TABLE IF NOT EXISTS `trade_order` (
                                             `id` int NOT NULL AUTO_INCREMENT,
                                             `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'uid',
    `code` varchar(20) NOT NULL DEFAULT '' COMMENT '代码',
    `date` varchar(8) NOT NULL DEFAULT '' COMMENT '日期',
    `price` float(20,6) NOT NULL DEFAULT '0' COMMENT '价格',
    `number` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
    `fee` float(20,6) NOT NULL DEFAULT '0' COMMENT '费用',
    `money` float(20,6) NOT NULL DEFAULT '0' COMMENT '总金额',
    `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_uid` (`uid`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单表';

CREATE TABLE IF NOT EXISTS `trade_hold` (
                                            `id` int NOT NULL AUTO_INCREMENT,
                                            `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'uid',
    `code` varchar(20) NOT NULL DEFAULT '' COMMENT '代码',
    `price` float(20,6) NOT NULL DEFAULT '0' COMMENT '价格',
    `number` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
    `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_uid` (`uid`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='持仓表';