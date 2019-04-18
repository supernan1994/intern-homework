create table `tbl_attribute` (
  `a_id` int unsigned not null auto_increment comment '属性id',
  `attr_name` varchar(32) not null comment '属性名',
  PRIMARY KEY (`a_id`)
);

create table `tbl_user` (
  `u_id` int unsigned not null auto_increment comment '用户id',
  `username` varchar(32) not null comment '用户名',
  PRIMARY KEY (`u_id`)
);

create table `tbl_user_attr` (
  `ua_id` int unsigned not null auto_increment comment '用户属性关系id',
  `u_id` int unsigned not null comment '用户id',
  `a_id` int unsigned not null comment '属性id',
  `attr_value` varchar(128) NOT NULL COMMENT '属性值',
  PRIMARY KEY (`ua_id`)
);

create table `tbl_knowledge` (
  `k_id` int unsigned not null auto_increment comment '知识点id',
  `standard_question` varchar(128) NOT NULL COMMENT '知识点标准问',
  PRIMARY KEY (`k_id`)
);

create table `tbl_group` (
  `g_id` int unsigned not null auto_increment comment '答案组id',
  `group_name` varchar(128) NOT NULL COMMENT '答案组名',
  PRIMARY KEY (`g_id`)
);

create table `tbl_group_attr` (
  `ga_id` int unsigned not null auto_increment comment '答案组属性关系id',
  `g_id` int unsigned not null comment '答案组id',
  `a_id` int unsigned not null comment '属性id',
  `attr_value` varchar(128) NOT NULL COMMENT '属性值',
  PRIMARY KEY (`ga_id`)
);

create table `tbl_answer` (
  `a_id` int unsigned not null auto_increment comment '答案id',
  `k_id` int unsigned not null comment '知识点id',
  `g_id` int unsigned not null comment '答案组id',
  `answer` varchar(1024) not null default '' comment '答案内容',
  PRIMARY KEY (`a_id`)
);

