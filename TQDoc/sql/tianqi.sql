/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50723
Source Host           : localhost:3306
Source Database       : tianqi

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2018-12-25 20:41:00
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_version
-- ----------------------------
DROP TABLE IF EXISTS `tb_version`;
CREATE TABLE `tb_version` (
  `lId` bigint(20) NOT NULL AUTO_INCREMENT,
  `strName` varchar(50) DEFAULT '' COMMENT '文件名字',
  `strVersion` varchar(50) DEFAULT '' COMMENT '版本号',
  `strPath` varchar(255) DEFAULT '' COMMENT '文件地址',
  `strMd5` varchar(255) DEFAULT '' COMMENT 'md5',
  `nType` tinyint(2) DEFAULT '1' COMMENT '类型 1:登录器 2客户端',
  `nDelete` tinyint(1) DEFAULT '0' COMMENT '删除状态 0 正常 1已删除',
  `dtUpdateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `dtCreateTime` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`lId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_version
-- ----------------------------
INSERT INTO `tb_version` VALUES ('1', 'weq', '1.0.1', '/weqwe', '3434234234523', '1', '0', '2018-12-25 17:17:44', '2018-12-25 17:17:42');
INSERT INTO `tb_version` VALUES ('2', 'fgh', '1.0.2', '/rtert', '6456576575', '2', '0', '2018-12-26 17:18:11', '2018-12-26 17:18:16');
INSERT INTO `tb_version` VALUES ('3', 'weq', '1.0.1', '/weqwe', '3434234234523', '1', '0', '2018-12-24 17:17:44', '2018-12-31 17:17:42');
