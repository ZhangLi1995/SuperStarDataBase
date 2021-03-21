CREATE TABLE `star_info` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `name_zh` varchar(50) NOT NULL DEFAULT '' COMMENT '中文名',
     `name_en` varchar(50) NOT NULL DEFAULT '' COMMENT '英文名',
     `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
     `birthday` varchar(50) NOT NULL DEFAULT '' COMMENT '出生日期',
     `height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '身高，单位cm',
     `weight` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '体重，单位g',
     `club` varchar(50) NOT NULL DEFAULT '' COMMENT '俱乐部',
     `jersy` varchar(50) NOT NULL DEFAULT '' COMMENT '球衣号码以及主打位置',
     `country` varchar(50) NOT NULL DEFAULT '' COMMENT '国籍',
     `birthaddress` varchar(255) NOT NULL DEFAULT '' COMMENT '出生地',
     `feature` varchar(255) NOT NULL DEFAULT '' COMMENT '个人特点',
     `moreinfo` text COMMENT '更多介绍',
     `sys_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，默认值 1 正常，0 删除',
     `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
     `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `star_info` (`id`, `name_zh`, `name_en`, `avatar`, `birthday`, `height`, `weight`, `club`, `jersy`, `country`, `birthaddress`, `feature`, `moreinfo`, `sys_status`, `sys_created`, `sys_updated`)
VALUES
(1, 'C罗', 'Cristiano Ronaldo', 'http://www.sinaimg.cn/lf/sports/wc_2018/player/14937.jpg', '1985-02-05', 187, 83, '皇家马德里', '前锋7号', '葡萄牙', '', '', '', 1, 0, 1529144171),
(2, '梅西', 'Lionel Messi', 'http://www.sinaimg.cn/lf/sports/wc_2018/player/19054.jpg', '1987-06-24', 170, 72, '巴塞罗那', '前锋10号', '阿根廷', '阿根廷圣菲省罗萨里奥', '', '绰号“新马拉多纳”', 1, 0, 1529135214),
(3, '穆勒', 'Thomas Müller', 'http://www.sinaimg.cn/lf/sports/wc_2018/player/55634.jpg', '1989-09-13', 186, 75, '拜仁慕尼黑', '前锋13号', '德国', '未知', '未知', '更多介绍', 1, 0, 1616323457);