CREATE TABLE `data_property`
(
    `id`     bigint PRIMARY KEY AUTO_INCREMENT COMMENT '数据ID',
    `type`   VARCHAR(255) NOT NULL COMMENT '数据类型（专利/论文/软件著作权）',
    `title`  VARCHAR(255) NOT NULL COMMENT '标题',
    `source` VARCHAR(255) NOT NULL COMMENT '来源',
    `time`   DATE         NOT NULL COMMENT '采集时间',
    `status` TINYINT      NOT NULL DEFAULT 1 COMMENT '状态（1: 进行中, 2: 已完成）'
) COMMENT '知识产权信息表';

INSERT INTO `data_property` VALUES (1, '专利', '人工智能云诊断平台', '国家知识产权局', '2025-04-16', 2);
INSERT INTO `data_property` VALUES (2, '论文', 'A Survey of Deep Learning Techniques for Autonomous Driving Systems', 'IEEE Transactions', '2025-04-16', 2);
INSERT INTO `data_property` VALUES (3, '软件著作权', '人工智能辅助编程工具', 'OpenKG', '2025-04-16', 2);
INSERT INTO `data_property` VALUES (4, '专利', '细胞病理智能数据集', '国家知识产权局', '2025-04-15', 2);
INSERT INTO `data_property` VALUES (5, '专利', '4000万数字技术专利数据库', '国家知识产权局', '2025-04-15', 2);
INSERT INTO `data_property` VALUES (6, '论文', 'Transformer: A Novel Architecture for neural Machine Translation', 'arXiv', '2025-04-15', 2);
INSERT INTO `data_property` VALUES (7, '论文', 'Generative Adversarial Networks (GANs) for Image Generation', 'arXiv', '2025-04-15', 2);
INSERT INTO `data_property` VALUES (8, '软件著作权', '智能语音识别系统', 'OpenKG', '2025-04-15', 2);
INSERT INTO `data_property` VALUES (9, '软件著作权', '自动驾驶仿真平台', 'OpenKG', '2025-04-15', 2);
INSERT INTO `data_property` VALUES (10, '专利', '专家混合机制和多头潜在注意力机制的大模型架构', '国家知识产权局', '2025-04-15', 2);

CREATE TABLE `data_task`
(
    `id`      bigint PRIMARY KEY AUTO_INCREMENT COMMENT '任务ID',
    `name`    VARCHAR(255) NOT NULL COMMENT '任务名称（如专利数据清洗）',
    `type`    VARCHAR(255) NOT NULL COMMENT '数据类型（专利/论文）',
    `manager` VARCHAR(255) NOT NULL COMMENT '负责人',
    `status`  TINYINT      NOT NULL DEFAULT 1 COMMENT '状态（1: 进行中, 2: 已完成）'
) COMMENT '数据任务表';

INSERT INTO `data_task` VALUES (1, '专利数据清洗', '专利', '张三', 1);
INSERT INTO `data_task` VALUES (2, '论文数据分析类', '论文', '李四', 2);

CREATE TABLE `data_report`
(
    `id`      bigint PRIMARY KEY AUTO_INCREMENT COMMENT '报告ID',
    `name`    VARCHAR(255) NOT NULL COMMENT '报告名称',
    `type`    VARCHAR(255) NOT NULL COMMENT '报告类型（专利分析/技术分析/市场分析）',
    `status`  TINYINT      NOT NULL DEFAULT 1 COMMENT '状态（1: 进行中, 2: 已完成）',
    `time`    DATE         NOT NULL COMMENT '生成时间',
    `content` TEXT         NOT NULL COMMENT '报告内容'
) COMMENT '数据报告表';

INSERT INTO `data_report` VALUES (1, '人工智能领域专利分析报告', '专利分析', 2, '2024-04-16', '这是一份关于人工智能领域专利分析的报告，包含了专利趋势分析、主要申请人分析、技术领域分布等内容。');
INSERT INTO `data_report` VALUES (2, '计算机视觉技术发展报告', '技术分析', 1, '2024-04-15', '这是一份关于计算机视觉技术发展的报告，包含了技术演进路线、关键算法分析、应用场景等内容。');


CREATE TABLE data_patent
(
    `id`        bigint PRIMARY KEY AUTO_INCREMENT COMMENT '专利ID',
    `title`     VARCHAR(255) NOT NULL COMMENT '专利名称',
    `applicant` VARCHAR(255) NOT NULL COMMENT '申请人',
    `date`      DATE         NOT NULL COMMENT '申请日期',
    `field`     VARCHAR(255) NOT NULL COMMENT '技术领域'
) COMMENT='专利数据表';

INSERT INTO `data_patent` VALUES (1, '一种基于深度学习的图像识别方法', '某科技公司', '2024-01-01', '计算机视觉');
INSERT INTO `data_patent` VALUES (2, '自然语言处理中的文本分类方法', '某人工智能研究所', '2024-01-05', '自然语言处理');
INSERT INTO `data_patent` VALUES (3, '基于深度强化学习的决策系统', '某大学', '2024-01-10', '机器学习');
INSERT INTO `data_patent` VALUES (4, '一种音频特征提取方法', '某声学研究所', '2024-01-15', '语音识别');

CREATE TABLE team_member
(
    id   bigint PRIMARY KEY AUTO_INCREMENT COMMENT '成员ID',
    name VARCHAR(50) NOT NULL COMMENT '姓名',
    role TINYINT     NOT NULL DEFAULT 1 COMMENT '角色（1: 查看, 2: 编辑, 3: 管理）'
) COMMENT='团队成员表';

INSERT INTO `team_member` VALUES (1, '张三', 3);
INSERT INTO `team_member` VALUES (2, '李四', 2);