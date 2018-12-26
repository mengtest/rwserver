/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50723
Source Host           : localhost:3306
Source Database       : tianqi

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2018-12-26 21:03:33
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `lId` bigint(20) NOT NULL,
  `strName` varchar(50) DEFAULT NULL COMMENT '用户名',
  `strPwd` varchar(100) DEFAULT NULL COMMENT '密码',
  `strRealName` varchar(50) DEFAULT '' COMMENT '真实姓名',
  `strIdCardNo` varchar(30) DEFAULT '' COMMENT '身份证号',
  `nAuthStatus` tinyint(2) DEFAULT '0' COMMENT '认证状态：0未认证 1认证中 2认证失败 3认证成功',
  `strMobile` varchar(20) DEFAULT '' COMMENT '手机号',
  `strEmail` varchar(20) DEFAULT '' COMMENT '邮箱',
  `dtUpdateTime` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dtCreateTime` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `nDeleted` tinyint(1) DEFAULT '0' COMMENT '删除状态 0未删除 1删除',
  PRIMARY KEY (`lId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_user
-- ----------------------------

-- ----------------------------
-- Table structure for tb_version
-- ----------------------------
DROP TABLE IF EXISTS `tb_version`;
CREATE TABLE `tb_version` (
  `lId` bigint(20) NOT NULL AUTO_INCREMENT,
  `strName` varchar(50) DEFAULT '' COMMENT '包名',
  `strVersion` varchar(50) DEFAULT '' COMMENT '版本号',
  `strPath` varchar(255) DEFAULT '' COMMENT '文件地址',
  `strMd5` varchar(255) DEFAULT '' COMMENT 'md5',
  `nAppType` tinyint(2) DEFAULT '1' COMMENT '类型 1:登录器rdc 2客户端rbc',
  `dtUpdateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `dtCreateTime` datetime DEFAULT NULL COMMENT '创建时间',
  `nDeleted` tinyint(1) DEFAULT '0' COMMENT '删除状态 0 正常 1已删除',
  PRIMARY KEY (`lId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_version
-- ----------------------------
INSERT INTO `tb_version` VALUES ('1', 'weq', '1.0.1', '/weqwe', '3434234234523', '1', '2018-12-25 17:17:44', '2018-12-25 17:17:42', '0');
INSERT INTO `tb_version` VALUES ('2', 'fgh', '1.0.2', '/rtert', '6456576575', '2', '2018-12-26 17:18:11', '2018-12-26 17:18:16', '0');
INSERT INTO `tb_version` VALUES ('3', 'weq', '1.0.1', '/weqwe', '3434234234523', '1', '2018-12-24 17:17:44', '2018-12-31 17:17:42', '0');
