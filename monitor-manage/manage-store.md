### manager模块与store模块之间的交互

manager主要涉及对**判定服务配置和告警服务配置的增删改查**，主要涉及对应配置的**增删改查**操作，主要基于**mysql**进行存储，故需要stroe模块提供以下接口服务：

#### 需求描述

接口较为简单，只要实现对数据库的增删改查即可，管理模块要啥就给数据库里对应的字段列表即可

1、对判定服务的增删改查

- 新增配置
- 根据id更新配置
- 根据id删除配置
- 根据id获取配置详情
- 获取所有的配置，返回一个列表
- 根据agentID（或ip和local）查看配置，返回一个列表
- 根据agentID（或ip和local）和判定类型（checkType）查看配置，返回一个列表

2、对告警服务配置的增删改查

- 新增配置
- 根据id更新配置
- 根据id删除配置
- 根据id获取配置详情
- 获取所有的配置列表，返回一个列表
- 根据agentID（或ip和local）查看配置，返回一个列表

#### 具体的数据库表设计可参考：

**判定配置表**

```sql
CREATE TABLE `check` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `ip` varchar(50) NOT NULL COMMENT 'agent的ip地址',
  `local` varchar(50) NOT NULL COMMENT 'agent的区域，如北京',
  `metric` tinyint(4) NOT NULL COMMENT '指标类型，1为cpu，2为内存等',
  `checkType` tinyint(4) NOT NULL COMMENT '判定类型，1为平均，2为求和等',
  `period` varchar(255) NOT NULL COMMENT '判定周期',
  `threshold` varchar(255) NOT NULL COMMENT '判定的阈值规则，字符串类型',
  `agentID` varchar(255) NOT NULL COMMENT 'agentID',
  PRIMARY KEY (`id`),
  KEY `select2` (`ip`,`local`),
  KEY `select1` (`ip`,`local`,`checkType`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

**告警配置表**

```sql
CREATE TABLE `alert` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `ip` varchar(50) NOT NULL COMMENT 'ip地址',
  `local` varchar(50) NOT NULL COMMENT '区域',
  `sendType` tinyint(4) NOT NULL COMMENT '告警目标的类型',
  `level` tinyint(4) NOT NULL COMMENT '告警等级',
  `config` varchar(255) NOT NULL COMMENT '告警类型对应的配置，字符串类型',
  `agentID` varchar(255) NOT NULL COMMENT 'agentID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

