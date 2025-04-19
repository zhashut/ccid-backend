CREATE TABLE `data_property`
(
    `id`     bigint PRIMARY KEY AUTO_INCREMENT COMMENT '数据ID',
    `type`   VARCHAR(255) NOT NULL COMMENT '数据类型（专利/论文/软件著作权）',
    `title`  VARCHAR(255) NOT NULL COMMENT '标题',
    `source` VARCHAR(255) NOT NULL COMMENT '来源',
    `time`   DATE         NOT NULL COMMENT '采集时间',
    `status` TINYINT      NOT NULL DEFAULT 1 COMMENT '状态（1: 进行中, 2: 已完成）',
) COMMENT '知识产权信息表';

CREATE TABLE `data_task`
(
    `id`      bigint PRIMARY KEY AUTO_INCREMENT COMMENT '任务ID',
    `name`    VARCHAR(255) NOT NULL COMMENT '任务名称（如专利数据清洗）',
    `type`    VARCHAR(255) NOT NULL COMMENT '数据类型（专利/论文）',
    `manager` VARCHAR(255) NOT NULL COMMENT '负责人',
    `status`  TINYINT      NOT NULL DEFAULT 1 COMMENT '状态（1: 进行中, 2: 已完成）'
) COMMENT '数据任务表';

CREATE TABLE `data_report`
(
    `id`      bigint PRIMARY KEY AUTO_INCREMENT COMMENT '报告ID',
    `name`    VARCHAR(255) NOT NULL COMMENT '报告名称',
    `type`    VARCHAR(255) NOT NULL COMMENT '报告类型（专利分析/技术分析/市场分析）',
    `status`  TINYINT      NOT NULL DEFAULT 1 COMMENT '状态（1: 进行中, 2: 已完成）',
    `time`    DATE         NOT NULL COMMENT '生成时间',
    `content` TEXT         NOT NULL COMMENT '报告内容'
) COMMENT '数据报告表';

CREATE TABLE data_patent
(
    `id`        bigint PRIMARY KEY AUTO_INCREMENT COMMENT '专利ID',
    `title`     VARCHAR(255) NOT NULL COMMENT '专利名称',
    `applicant` VARCHAR(255) NOT NULL COMMENT '申请人',
    `date`      DATE         NOT NULL COMMENT '申请日期',
    `field`     VARCHAR(255) NOT NULL COMMENT '技术领域',
) COMMENT='专利数据表';

CREATE TABLE team_member
(
    id   bigint PRIMARY KEY AUTO_INCREMENT COMMENT '成员ID',
    name VARCHAR(50) NOT NULL COMMENT '姓名',
    role TINYINT     NOT NULL DEFAULT 1 COMMENT '角色（1: 查看, 2: 编辑, 3: 管理）'
) COMMENT='团队成员表';