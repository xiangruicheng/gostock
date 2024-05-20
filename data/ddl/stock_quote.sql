
CREATE TABLE IF NOT EXISTS `stock_quote` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `code` VARCHAR(8) NULL DEFAULT '' COMMENT '代码' COLLATE 'utf8_general_ci',
    `name` VARCHAR(50) NULL DEFAULT '' COMMENT '名称' COLLATE 'utf8_general_ci',
    `pe_forecast` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '动态市盈率',
    `pe_ttm` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT 'TTM市盈率',
    `pe_lyr` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '静态市盈率',
    `pb` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '市净率PB',
    `total_shares` DECIMAL(20,2) NULL DEFAULT '0.00' COMMENT '总股本',
    `float_shares` DECIMAL(20,6) NULL DEFAULT '0.000000' COMMENT '流通股',
    `float_market_capital` DECIMAL(20,6) NULL DEFAULT '0.000000' COMMENT '流通市值',
    `market_capital` DECIMAL(20,6) NULL DEFAULT '0.000000' COMMENT '市值',
    `amount` DECIMAL(20,6) NULL DEFAULT '0.000000' COMMENT '成交额',
    `volume` DECIMAL(20,6) NULL DEFAULT '0.000000' COMMENT '成交量',
    `turnover_rate` FLOAT NULL DEFAULT '0' COMMENT '换手率',
    `amplitude` FLOAT NULL DEFAULT '0' COMMENT '振幅',
    `navps` FLOAT NULL DEFAULT '0' COMMENT '每股净值产',
    `eps` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '每股收益',
    `volume_ratio` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '量比',
    `pankou_ratio` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '委比',
    `high` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '最高',
    `low` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '最低',
    `open` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '开盘价',
    `current` FLOAT(10,2) NULL DEFAULT '0.00' COMMENT '当前价',
    `dividend` FLOAT(10,6) NULL DEFAULT '0.000000' COMMENT '股息',
    `dividend_yield` FLOAT(10,6) NULL DEFAULT '0.000000' COMMENT '股息率',
    `date` VARCHAR(8) NULL DEFAULT '' COMMENT '日期' COLLATE 'utf8_general_ci',
    `c_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `u_time` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `code` (`code`) USING BTREE
    )
    COMMENT='股票基本信息'
    COLLATE='utf8_czech_ci'
    ENGINE=InnoDB
;