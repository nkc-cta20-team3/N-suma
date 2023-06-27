USE cta20gr3;

INSERT INTO `classification`(class_id,class_abbr,class_name)
VALUES
(117,'CTA20','情報総合学科4年A組'),
(118,'CTB20','情報総合学科4年B組'),
(115,'CTA21','情報総合学科3年A組'),
(116,'CTB21','情報総合学科3年B組');

INSERT INTO `post`(post_id,post_name)
VALUES
(0,'管理者'),
(1,'学生'),
(2,'担任'),
(3,'主任'),
(4,'学科長'),
(5,'学部長'),
(6,'所属長');

INSERT INTO `division`(division_id,division_name,division_detail,division_remark)
VALUES
(100,'国家試験','FE','基本情報'),
(101,'国家試験','AP','応用情報'),
(102,'国家試験','SG','情報セキュリティマネジメント');

INSERT INTO `user`(user_uuid,user_name,user_number,post_id,class_id)
VALUES
('1waffa4r','田中太郎',20206000,1,117),
('ff3ru6dg ','高橋啓太',null,2,117),
('ohigw923','藤井雄太',null,3,null),
('foqihnfl','斎藤勇',null,4,null),
('o32tgbkl','小澤直樹',null,5,null),
('4hpivesw','大橋健一',null,6,null),
('9vacugh2','櫻井健一',20206100,1,116),
('36gew305','鈴木秀昭',20206200,1,118);

INSERT INTO `oa`(request_at,start_date,start_flame,end_date,end_flame,location,status,read_flag,student_comment,teacher_comment,user_id,division_id)
VALUES
('2020-10-01','2020-10-01',1,'2020-10-01',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01',1,'2020-10-01',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01',1,'2020-10-01',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01',1,'2020-10-01',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01',1,'2020-10-01',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01',1,'2020-10-01',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01',1,'2020-10-01',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01',1,'2020-10-01',2,'A棟',1,false,'テスト','テスト',1,100);
