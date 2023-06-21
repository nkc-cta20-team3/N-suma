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
/* 
    学生の所属クラス兼教員の担当するクラスを管理する

    class_idは3桁の数字
    3桁目(100の位)で分野の種別を判別する
    2桁目で学科を判別する
    1桁目でクラスを判別する

    例: 100, 200, 300, 400, 500, 600
    100: コンピューター・IT分野
    200: ゲーム・CG分野
    300: 映像・音響分野
    400: 電気分野
    500: 情報通信分野
    600: 機械・ロボット・CAD分野

    例: 110, 120, 130, 140, 150
    110: 情報総合学科
    120: 情報システム科
    130: AIイノベーション学科/AIシステムコース
    140: AIイノベーション学科/AIシステムコース
    150: 情報処理学科
    160: 高度情報学科

    例: 111, 112, 113, 114, 115, 116, 117, 118
    111: 情報総合学科/1年A組
    112: 情報総合学科/1年B組
    113: 情報総合学科/2年A組
    114: 情報総合学科/2年B組
    115: 情報総合学科/3年A組
    116: 情報総合学科/3年B組
    117: 情報総合学科/4年A組
    118: 情報総合学科/4年B組

*/


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
/* 
    IDの2桁目が学科の種別ごとに違うなどのIDを振り分けたい(予定) 
*/


DROP TABLE IF EXISTS `division`;
CREATE TABLE `division`(
    `division_id` int AUTO_INCREMENT,
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
/* 
    IDの2桁目が学科の種別ごとに違うなどのIDを振り分けたい(予定) 
*/


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
