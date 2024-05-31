CREATE TABLE IF NOT EXISTS `stock_block` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `type` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '类型 1东财行业板块 2东财概念板块',
    `code` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '代码' COLLATE 'utf8_general_ci',
    `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '名称' COLLATE 'utf8_general_ci',
    `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `code` (`code`) USING BTREE
) COMMENT='板块信息' COLLATE='utf8_general_ci' ENGINE=InnoDB;