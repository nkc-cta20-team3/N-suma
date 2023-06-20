DROP SCHEMA IF EXISTS cta20gr3;
CREATE SCHEMA cta20gr3;
USE cta20gr3;


DROP TABLE IF EXISTS `classification`;
CREATE TABLE `classification`(
    `class_id` int,
    `class_abbr` varchar(16) NOT NULL,
    `class_name` varchar(32) NOT NULL,
    PRIMARY KEY (`class_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='クラス表/マスターデータ';
# show create table `class` ;


DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`(
    `post_id` int,
    `post_name` varchar(16) NOT NULL,
    PRIMARY KEY (`post_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='役職表/マスターデータ';
# show create table `post` ;


DROP TABLE IF EXISTS `division`;
CREATE TABLE `division`(
    `division_id` int,
    `division_name` varchar(16) NOT NULL,
    `division_detail` varchar(32) NOT NULL,
    `division_remark` varchar(8000),
    PRIMARY KEY (`division_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='区分表/マスターデータ';
# show create table `division` ;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
    `user_uid` varchar(255),
    `user_name` varchar(80),
    `user_number` int,
    `post_id` int,
    `class_id` int,
    PRIMARY KEY (`user_uid`),
    FOREIGN KEY (`post_id`) REFERENCES `post`(`post_id`),
    FOREIGN KEY (`class_id`) REFERENCES `classification`(`class_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='ユーザー表';
# show create table `user` ;


DROP TABLE IF EXISTS `oa`;
CREATE TABLE `oa`(
    `document_id` int AUTO_INCREMENT,
    `request_at` date NOT NULL,
    `start_date` date NOT NULL,
    `start_flame` int NOT NULL,
    `end_date` date NOT NULL,
    `end_flame` int NOT NULL,
    `location` varchar(16) NOT NULL,
    `status` int NOT NULL,
    `read_flag` boolean NOT NULL,
    `student_comment` varchar(8000) NOT NULL,
    `teacher_comment` varchar(8000),
    `user_uid` varchar(255) NOT NULL,
    `division_id` int NOT NULL,
    PRIMARY KEY (`document_id`),
    FOREIGN KEY (`user_uid`) REFERENCES `user`(`user_uid`),
    FOREIGN KEY (`division_id`) REFERENCES `division`(`division_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='公欠表';
# show create table `oa` ;
