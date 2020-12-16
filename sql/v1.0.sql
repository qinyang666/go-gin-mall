create table pms_product_category
(
   id                   int(10) not null auto_increment,
   parent_id            int(10)  unsigned not null comment '上级分类的编号：0表示一级分类',
   name                 varchar(64)  not null default '' comment '名称',
   level                int(1) not null default '0' comment '分类级别：0->1级；1->2级',
   product_count        int unsigned not null default '0' comment '商品数量',
   product_unit         varchar(64) not null default '' comment '商品单位',
   nav_status           int(1) not null default '0' comment '是否显示在导航栏：0->不显示; 1->显示',
   show_status          int(1) not null default '0' comment '显示状态：0->不显示；1->显示',
   sort                 int unsigned not null default '0' comment '排序',
   icon                 varchar(255) not null default '' comment '图标',
   keywords             varchar(255) not null default '' comment '关键字',
   description          varchar(1000) not null default '' comment '描述',
   created              datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated              timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   primary key (id)
)engine=innodb default charset=utf8mb4;