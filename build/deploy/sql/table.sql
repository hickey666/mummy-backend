CREATE TABLE product (
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    name VARCHAR(255) NOT NULL DEFAULT '' COMMENT '产品名称',
    price DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '价格',
    description TEXT COMMENT '描述',
    is_deleted TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除 0:未删除 1:已删除',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '产品表';

CREATE Table category (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    name VARCHAR(255) NOT NULL DEFAULT '' COMMENT '分类名称',
    description TEXT COMMENT '描述',
    is_deleted TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除 0:未删除 1:已删除',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '分类表';

create Table product_category (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    pid INT NOT NULL DEFAULT 0 COMMENT '产品ID',
    cid INT NOT NULL DEFAULT 0 COMMENT '分类ID',
    is_deleted TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除 0:未删除 1:已删除',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `product_id_category_id` (`pid`, `cid`)
) engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '产品分类关联表';

CREATE TABLE product_storage (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    pid INT NOT NULL DEFAULT 0 COMMENT '产品ID',
    total INT NOT NULL DEFAULT 0 COMMENT '总库存',
    used INT NOT NULL DEFAULT 0 COMMENT '已用库存',
    residue INT NOT NULL DEFAULT 0 COMMENT '剩余库存',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '产品库存表';

create Table product_storage_log (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    pid INT NOT NULL DEFAULT 0 COMMENT '产品ID',
    before_total INT NOT NULL DEFAULT 0 COMMENT '总库存 before',
    put_in INT NOT NULL DEFAULT 0 COMMENT '新增库存',
    put_out INT NOT NULL DEFAULT 0 COMMENT '减少库存',
    after_total INT NOT NULL DEFAULT 0 COMMENT '总库存 after',
    is_deleted TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除 0:未删除 1:已删除',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '产品库存日志表';