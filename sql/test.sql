# 测试数据 paasword=123456
INSERT INTO `go`.`user`(`name`, `age`, `username`, `password`) VALUES ('张三', 10, 'zhangsan', '7c4a8d09ca3762af61e59520943dc26494f8941b');
INSERT INTO `go`.`user`(`name`, `age`, `username`, `password`) VALUES ('李四', NULL, 'lisi', '4be30d9814c6d4e9800e0d2ea9ec9fb00efa887b');
INSERT INTO `go`.`phone`(`id`, `userId`, `phone`) VALUES (1, 1, '13688889999');
INSERT INTO `go`.`phone`(`id`, `userId`, `phone`) VALUES (2, 2, '1234567');
INSERT INTO `go`.`order`(`id`, `userId`, `status`) VALUES (1, 1, 1);