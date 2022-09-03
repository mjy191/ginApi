CREATE TABLE `user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `cdt` datetime DEFAULT CURRENT_TIMESTAMP,
    `mdt` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    `name` varchar(255) DEFAULT NULL,
    `age` int(11) DEFAULT NULL,
    `username` varchar(30) DEFAULT NULL,
    `password` varchar(40) DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表';

CREATE TABLE `phone` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `userId` int(11) DEFAULT NULL,
     `phone` varchar(11) DEFAULT NULL,
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='手机号码表';

CREATE TABLE `order` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `cdt` datetime DEFAULT CURRENT_TIMESTAMP,
     `mdt` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
     `userId` int(11) unsigned DEFAULT NULL,
     `status` int(11) unsigned DEFAULT NULL,
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='订单表';