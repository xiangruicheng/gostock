CREATE TABLE IF NOT EXISTS `stock_block_code` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `bk_code` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '板块代码' COLLATE 'utf8_general_ci',
    `code` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '股票代码' COLLATE 'utf8_general_ci',
    `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `idx_code` (`bk_code`, `code`) USING BTREE
) COMMENT='板块的股票代码' COLLATE='utf8_general_ci' ENGINE=InnoDB;