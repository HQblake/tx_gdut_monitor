/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : monitor

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 23/02/2022 14:33:12
*/

-- ----------------------------
-- Database structure for monitor
-- ----------------------------
DROP DATABASE IF EXISTS `monitor`;
CREATE DATABASE `monitor` CHARACTER SET utf8mb4 COLLATE = utf8mb4_general_ci;
USE `monitor`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for agent
-- ----------------------------
DROP TABLE IF EXISTS `agent`;
CREATE TABLE `agent`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Agent IP',
  `local` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Agent 地区',
  `port` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Agent 端口',
  `isLive` tinyint(1) NOT NULL DEFAULT 1 COMMENT 'Agent 存活状态: 1为存活，0为失活',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for agent_metric
-- ----------------------------
DROP TABLE IF EXISTS `agent_metric`;
CREATE TABLE `agent_metric`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `agentId` int NOT NULL,
  `metricId` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `map_agent`(`agentId`) USING BTREE,
  INDEX `map_metric`(`metricId`) USING BTREE,
  CONSTRAINT `map_agent` FOREIGN KEY (`agentId`) REFERENCES `agent` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `map_metric` FOREIGN KEY (`metricId`) REFERENCES `metric` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for alert
-- ----------------------------
DROP TABLE IF EXISTS `alert`;
CREATE TABLE `alert`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `agentId` int NOT NULL,
  `sendType` tinyint NOT NULL COMMENT '告警目标的类型',
  `level` tinyint NOT NULL COMMENT '告警等级',
  `config` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '告警类型对应的配置，JSON字符串类型',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `alert_agent`(`agentId`) USING BTREE,
  CONSTRAINT `alert_agent` FOREIGN KEY (`agentId`) REFERENCES `agent` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for check
-- ----------------------------
DROP TABLE IF EXISTS `check`;
CREATE TABLE `check`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `agentId` int NOT NULL,
  `metricId` int NOT NULL COMMENT '指标类型ID',
  `method` tinyint NOT NULL COMMENT '聚合方式，1为平均，2为求和等',
  `period` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '聚合周期',
  `threshold` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '判定的阈值规则，JSON字符串类型',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `select1`(`metricId`) USING BTREE,
  INDEX `select2`(`agentId`) USING BTREE,
  CONSTRAINT `check_agent` FOREIGN KEY (`agentId`) REFERENCES `agent` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `check_metric` FOREIGN KEY (`metricId`) REFERENCES `metric` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for history
-- ----------------------------
DROP TABLE IF EXISTS `history`;
CREATE TABLE `history`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `agentId` int NOT NULL COMMENT 'Agent ID',
  `metricId` int NOT NULL COMMENT '指标类型ID',
  `value` float(10, 6) NOT NULL COMMENT 'Agent指标值',
  `threshold` float(10, 6) NOT NULL COMMENT '阈值',
  `method` tinyint NOT NULL COMMENT '聚合方式',
  `level` tinyint NOT NULL COMMENT '告警等级',
  `duration` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '告警持续时间',
  `start` int NOT NULL COMMENT '开始时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `history_agent`(`agentId`) USING BTREE,
  INDEX `history_metric`(`metricId`) USING BTREE,
  CONSTRAINT `history_agent` FOREIGN KEY (`agentId`) REFERENCES `agent` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `history_metric` FOREIGN KEY (`metricId`) REFERENCES `metric` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for metric
-- ----------------------------
DROP TABLE IF EXISTS `metric`;
CREATE TABLE `metric`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Procedure structure for AddAgentInfo
-- ----------------------------
DROP PROCEDURE IF EXISTS `AddAgentInfo`;
delimiter ;;
CREATE PROCEDURE `AddAgentInfo`(IN i VARCHAR(50), 
	IN l VARCHAR(50), 
	IN p VARCHAR(50), 
	IN mid INT)
BEGIN
	-- 定义 MetricID 与 AgentID 变量
	DECLARE agentId INT DEFAULT -1;
	DECLARE id INT DEFAULT -1;
	
-- 	-- 获取metricID，若不存在则插入新的metric
-- 	SELECT m.id FROM metric AS m WHERE m.`name`=m INTO metricId;
-- 	IF metricId=-1 THEN
-- 		INSERT INTO metric(`name`) VALUES(m);
-- 		SELECT m.id FROM metric AS m WHERE m.`name`=m INTO metricId;
-- 	END IF;

	-- 获取AgentID，若不存在则插入新的agent
	SELECT a.id FROM agent AS a WHERE a.`ip`=i AND a.`local`=l INTO agentId;
	IF agentId=-1 THEN
		INSERT INTO agent(`ip`,`local`,`port`) VALUES(i,l,p);
		SELECT a.id FROM agent AS a WHERE a.`ip`=i AND a.`local`=l INTO agentId;
	END IF;

	
	-- 检测Agent是否包含该指标，若不存在则插入新的agent-metric关联
	SELECT am.id FROM agent_metric AS am WHERE am.`agentId`=agentId AND am.`metricId`=mid INTO id;
	IF id=-1 THEN
		INSERT INTO agent_metric(`agentId`,`metricId`) VALUES(agentId,mid);
	END IF;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;