/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50723
Source Host           : localhost:3306
Source Database       : tianqi

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2018-12-28 13:40:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_npc
-- ----------------------------
DROP TABLE IF EXISTS `tb_npc`;
CREATE TABLE `tb_npc` (
  `lId` bigint(20) NOT NULL AUTO_INCREMENT,
  `strName` varchar(255) DEFAULT '' COMMENT '角色名称',
  `strTitle` varchar(255) DEFAULT '' COMMENT '当前称号',
  `nSex` int(1) DEFAULT '0' COMMENT '0 男 1女',
  `nLevel` int(10) unsigned DEFAULT '1' COMMENT '等级 根据阅历提升等级',
  `nExp` bigint(20) DEFAULT '0' COMMENT '阅历',
  `nHP` int(10) DEFAULT '0' COMMENT '体力气血值',
  `nMP` int(10) DEFAULT '0' COMMENT '法力值',
  `nMinAP` int(10) DEFAULT '0' COMMENT 'Attack Power 法术伤害',
  `nMinAD` int(10) DEFAULT '0' COMMENT 'Attack Damage 物理伤害',
  `nMaxAP` int(10) DEFAULT '0' COMMENT 'Attack Power 法术伤害',
  `nMaxAD` int(10) DEFAULT '0' COMMENT 'Attack Damage 物理伤害',
  `nPhyDef` int(10) DEFAULT '0' COMMENT '物防',
  `nMagDef` int(10) DEFAULT '0' COMMENT '法防',
  `nCrit` int(10) DEFAULT '0' COMMENT '会心',
  `nCon` int(10) DEFAULT '0' COMMENT '体质 影响hp+=5*nt 影响物防nDefence+=3*nt',
  `nDex` int(10) DEFAULT '0' COMMENT '敏捷 影响会心 和速度',
  `nStr` int(10) DEFAULT '0' COMMENT '力量',
  `nDod` int(10) DEFAULT '0' COMMENT '躲避',
  `nSup` int(10) DEFAULT '0' COMMENT '法力',
  `fPosX` double(10,2) DEFAULT NULL,
  `fPosY` double(10,2) DEFAULT NULL,
  `fPosZ` double(10,2) DEFAULT NULL,
  `fDirX` double(10,2) DEFAULT NULL,
  `fDirY` double(10,2) DEFAULT NULL,
  `fDirZ` double(10,2) DEFAULT NULL,
  `strMapName` varchar(30) DEFAULT '' COMMENT '角色所在地图名称',
  `DeleteDate` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '角色删除时间',
  `nDeleted` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  `dtUpdateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `dtCreateTime` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`lId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='npc列表';

-- ----------------------------
-- Records of tb_npc
-- ----------------------------

-- ----------------------------
-- Table structure for tb_role
-- ----------------------------
DROP TABLE IF EXISTS `tb_role`;
CREATE TABLE `tb_role` (
  `lId` bigint(20) NOT NULL AUTO_INCREMENT,
  `strName` varchar(255) DEFAULT '' COMMENT '角色名称',
  `strTitle` varchar(255) DEFAULT '' COMMENT '当前称号',
  `nSex` int(1) DEFAULT '0' COMMENT '0 男 1女',
  `nLevel` int(10) unsigned DEFAULT '1' COMMENT '等级 根据阅历提升等级',
  `nExp` bigint(20) DEFAULT '0' COMMENT '阅历',
  `nHP` int(10) DEFAULT '0' COMMENT '体力气血值',
  `nMP` int(10) DEFAULT '0' COMMENT '法力值',
  `nMinAP` int(10) DEFAULT '0' COMMENT 'Attack Power 法术伤害',
  `nMinAD` int(10) DEFAULT '0' COMMENT 'Attack Damage 物理伤害',
  `nMaxAP` int(10) DEFAULT '0' COMMENT 'Attack Power 法术伤害',
  `nMaxAD` int(10) DEFAULT '0' COMMENT 'Attack Damage 物理伤害',
  `nPhyDef` int(10) DEFAULT '0' COMMENT '物防',
  `nMagDef` int(10) DEFAULT '0' COMMENT '法防',
  `nCrit` int(10) DEFAULT '0' COMMENT '会心',
  `nCon` int(10) DEFAULT '0' COMMENT '体质 影响hp+=5*nt 影响物防nDefence+=3*nt',
  `nDex` int(10) DEFAULT '0' COMMENT '敏捷 影响会心 和速度',
  `nStr` int(10) DEFAULT '0' COMMENT '力量',
  `nDod` int(10) DEFAULT '0' COMMENT '躲避',
  `nSup` int(10) DEFAULT '0' COMMENT '法力',
  `fPosX` double(10,2) DEFAULT NULL,
  `fPosY` double(10,2) DEFAULT NULL,
  `fPosZ` double(10,2) DEFAULT NULL,
  `fDirX` double(10,2) DEFAULT NULL,
  `fDirY` double(10,2) DEFAULT NULL,
  `fDirZ` double(10,2) DEFAULT NULL,
  `strMapName` varchar(30) DEFAULT '' COMMENT '角色所在地图名称',
  `DeleteDate` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '角色删除时间',
  `nDeleted` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  `dtUpdateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `dtCreateTime` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`lId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Records of tb_role
-- ----------------------------

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `lId` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键 自增',
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES ('1', 'wuyz', '1', 'wka1', '1111', '0', '', '', '2018-12-27 17:31:17', '2018-12-27 17:31:17', '0');

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
INSERT INTO `tb_version` VALUES ('1', 'weq', '1.0.1', '/weqwe', '3434234234523', '1', '2018-12-20 10:51:47', '2018-12-20 10:51:47', '0');
INSERT INTO `tb_version` VALUES ('2', 'fgh', '1.0.2', '/rtert', '6456576575', '2', '2018-12-25 10:51:52', '2018-12-25 10:51:52', '0');
INSERT INTO `tb_version` VALUES ('3', 'weq', '1.0.1', '/weqwe', '3434234234523', '1', '2018-12-27 10:51:56', '2018-12-27 10:51:56', '0');
