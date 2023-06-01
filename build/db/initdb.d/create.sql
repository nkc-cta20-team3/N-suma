DROP SCHEMA IF EXISTS database;
CREATE SCHEMA database;
USE database;

DROP TABLE IF EXISTS `class`;
CREATE TABLE `class`(
    `class_name` varchar(6),
    `formal_name` varchar(32),
    PRIMARY KEY (`class_name`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='クラス表';
# show create table `class` ;


DROP TABLE IF EXISTS `position`;
CREATE TABLE `position`(
    `position_id` int,
    `position_name` varchar(6) NOT NULL,
    PRIMARY KEY (`position_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='役職表';
# show create table `position` ;


DROP TABLE IF EXISTS `company`;
CREATE TABLE `company`(
    company_id int AUTO_INCREMENT,
    company_name varchar(64) NOT NULL,
    offer_number int,
    PRIMARY KEY (`company_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='企業表';
# show create table `company` ;


DROP TABLE IF EXISTS `absence_reason`;
CREATE TABLE `absence_reason`(
    reason_id int,
    absence_category varchar(16) NOT NULL,
    absence_reason varchar(32)  NOT NULL,
    detail text,
    PRIMARY KEY (`reason_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='公欠理由表';
# show create table `absence_reason` ;


DROP TABLE IF EXISTS `students`;
CREATE TABLE `students`(
    student_id int,
    student_name varchar(20) NOT NULL,
    mail_address varchar(64) NOT NULL,
    password varchar(20) NOT NULL,
    class_name varchar(6) NOT NULL,
    PRIMARY KEY (`student_id`),
    FOREIGN KEY (class_name) REFERENCES `class`(class_name)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='学生表';
# show create table `students` ;


DROP TABLE IF EXISTS `teachers`;
CREATE TABLE `teachers`(
    teacher_id int AUTO_INCREMENT,
    teacher_name varchar(20) NOT NULL,
    mail_address varchar(64) NOT NULL,
    password varchar(20) NOT NULL,
    class_name varchar(6),
    position_id int,
    PRIMARY KEY (`teacher_id`),
    FOREIGN KEY (class_name) REFERENCES `class`(class_name),
    FOREIGN KEY (position_id) REFERENCES `position`(position_id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='教員表';
# show create table `teachers` ;


DROP TABLE IF EXISTS `absence_document`;
CREATE TABLE `absence_document`(
    document_id int AUTO_INCREMENT,
    student_id int NOT NULL,
    company_id int,
    reason_id int NOT NULL,
    request_date date NOT NULL,
    absence_start_date date NOT NULL,
    absence_start_flame int NOT NULL,
    absence_end_date date NOT NULL,
    absence_end_flame int NOT NULL,
    location varchar(16) NOT NULL,
    read_flag boolean NOT NULL,
    status int NOT NULL,
    student_comment text NOT NULL,
    teacher_comment text,
    PRIMARY KEY (`document_id`),
    FOREIGN KEY(student_id) REFERENCES `students`(student_id),
    FOREIGN KEY(company_id) REFERENCES `company`(company_id),
    FOREIGN KEY(reason_id) REFERENCES `absence_reason`(reason_id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_as_cs
COMMENT='公欠届表';
# show create table `absence_document` ;