USE cta20gr3;

INSERT INTO `class`(class_name,formal_name)
VALUES
('CTA20','情報総合学科4年A組'),
('CTB20','情報総合学科4年B組'),
('CTA21','情報総合学科3年A組'),
('CTB21','情報総合学科3年B組');


INSERT INTO `position`(position_id,position_name)
VALUES
(1,'担任'),
(2,'主任'),
(3,'学科長'),
(4,'学部長'),
(5,'所属長');

INSERT INTO `company`(company_name,offer_number)
VALUES
('株式会社AAA',1001),
('株式会社BBB',1002),
('株式会社CCC',1003),
('株式会社DDD',1004);

INSERT INTO `absence_reason`(reason_id,absence_category,absence_reason,detail)
VALUES
(1,'国家試験','FE','詳細'),
(2,'国家試験','AP','詳細'),
(3,'国家試験','SG','詳細');

INSERT INTO `students`(student_id,student_name,mail_address,password,class_name)
VALUES
(20206001,'斎藤勇','nkc20206001@denpa.jp','20206001','CTA20'),
(20206100,'小澤直樹','nkc20206100@denpa.jp','20206100','CTA20'),
(20206200,'大橋健一','nkc20206200@denpa.jp','20206200','CTA20');

INSERT INTO `teachers`(teacher_name,mail_address,password,class_name,position_id)
VALUES
('田中太郎','00000001@denpa.jp','00000001','CTA20',1),
('高橋啓太','00000002@denpa.jp','00000002',NULL,2),
('藤井雄太','00000003@denpa.jp','00000003',NULL,3),
('櫻井健一','00000004@denpa.jp','00000004',NULL,4),
('鈴木秀昭','00000005@denpa.jp','00000005',NULL,5);

INSERT INTO `absence_document`(student_id,company_id,reason_id,request_date,absence_start_date,absence_start_flame,absence_end_date,absence_end_flame,location,read_flag,status,student_comment,teacher_comment)
VALUES
(20206001,null,1,'2020-10-01','2020-10-02',1,'2020-10-02',2,'A棟',false,0,'学生コメント',null),
(20206001,null,2,'2020-10-01','2020-10-03',1,'2020-10-03',2,'A棟',false,0,'学生コメント',null),
(20206001,null,3,'2020-10-01','2020-10-04',1,'2020-10-04',2,'A棟',false,0,'学生コメント',null);

