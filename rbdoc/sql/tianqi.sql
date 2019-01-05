/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50717
 Source Host           : localhost
 Source Database       : tianqi

 Target Server Type    : MySQL
 Target Server Version : 50717
 File Encoding         : utf-8

 Date: 01/05/2019 14:55:59 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `tb_npc`
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
--  Table structure for `tb_role`
-- ----------------------------
DROP TABLE IF EXISTS `tb_role`;
CREATE TABLE `tb_role` (
  `lId` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `lUserId` bigint(20) DEFAULT NULL COMMENT '用户ID',
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
  `nDodge` int(10) DEFAULT NULL COMMENT '回避值',
  `nCrit` int(10) DEFAULT '0' COMMENT '会心',
  `nHit` int(10) DEFAULT '0' COMMENT '命中',
  `nCon` int(10) DEFAULT '0' COMMENT '体 影响nHP+=5*nCon 影响物防nPhyDef+=3*nCon',
  `nDex` int(10) DEFAULT '0' COMMENT '敏 影响会心和施法速度',
  `nStr` int(10) DEFAULT '0' COMMENT '力 影响物理伤害和命中 nMinAD+=2*nStr; nMaxAD=3*nStr  nHit+=1*nStr',
  `nAvoid` int(10) DEFAULT '0' COMMENT '避',
  `nSp` int(10) DEFAULT '0' COMMENT '法',
  `fPosX` double(10,2) DEFAULT NULL,
  `fPosY` double(10,2) DEFAULT NULL,
  `fPosZ` double(10,2) DEFAULT NULL,
  `fDirX` double(10,2) DEFAULT NULL,
  `fDirY` double(10,2) DEFAULT NULL,
  `fDirZ` double(10,2) DEFAULT NULL,
  `strMapName` varchar(30) DEFAULT '' COMMENT '角色所在地图名称',
  `nChunkX` int(10) DEFAULT '0',
  `nChunkY` int(10) DEFAULT '0',
  `nDeleted` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  `dtDeleteDate` datetime DEFAULT NULL COMMENT '角色删除时间',
  `dtUpdateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `dtCreateTime` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`lId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
--  Records of `tb_role`
-- ----------------------------
BEGIN;
INSERT INTO `tb_role` VALUES ('1', '1', '慕临风', '一剑霜寒十四州', '0', '50', '45897228', '5849', '1209', '100', '257', '324', '547', '588', '432', '878', '112', '210', '100', '60', '200', '70', '30', '100.00', '100.00', '100.00', '0.00', '0.00', '0.00', 'tzy', '11', '12', '0', null, null, '2019-01-05 12:16:58');
COMMIT;

-- ----------------------------
--  Table structure for `tb_user`
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
--  Records of `tb_user`
-- ----------------------------
BEGIN;
INSERT INTO `tb_user` VALUES ('1', 'wuyz', 'e19f69ca630aef5591f2107b56228dfc', 'wka1', '1111', '0', '', '', '2019-01-05 11:42:50', '2019-01-05 11:42:50', '0');
COMMIT;

-- ----------------------------
--  Table structure for `tb_version`
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
--  Records of `tb_version`
-- ----------------------------
BEGIN;
INSERT INTO `tb_version` VALUES ('1', 'weq', '1.0.1', '/weqwe', '3434234234523', '1', '2018-12-20 10:51:47', '2018-12-20 10:51:47', '0'), ('2', 'fgh', '1.0.2', '/rtert', '6456576575', '2', '2018-12-25 10:51:52', '2018-12-25 10:51:52', '0'), ('3', 'weq', '1.0.1', '/weqwe', '3434234234523', '1', '2018-12-27 10:51:56', '2018-12-27 10:51:56', '0');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
