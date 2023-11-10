USE cta20gr3;

INSERT INTO `classification`(class_id,class_abbr,class_name)
VALUES
(132,'AI22','AIシステム科2年'),
(123,'CSX21','情報システム科3年'),
(115,'CTA21','情報総合学科3年A組'),
(117,'CTA20','情報総合学科4年A組'),
(118,'CTB20','情報総合学科4年B組');


INSERT INTO `post`(post_id,post_name)
VALUES
(0,'管理者'),
(1,'学生'),
(2,'教員');

INSERT INTO `division`(division_id,division_name,division_detail,division_remark)
VALUES
(100,'国家試験','FE','基本情報処理技術者試験'),
(101,'国家試験','AP','応用情報処理技術者試験'),
(102,'国家試験','SC','情報処理安全確保支援士試験'),
(200,'就職活動','面接試験(1次)',''),
(201,'就職活動','面接試験(最終)',''),
(300,'忌引','葬儀','休暇は3日まで'),
(400,'その他','その他','担任と要相談');

INSERT INTO `user`(user_uuid,user_name,user_number,post_id,class_id,mail_address,user_flag)
VALUES
('1waffa4r','田中太郎',20226000,1,132,'aaa.gmail.com',true),
('ff3ru6dg','田中次郎',20226002,1,132,'bbb.gmail.com',true),
('ohigw923','佐藤一子',20216003,1,115,'ccc.gmail.com',true),
('foqihnfl','佐藤次子',20216004,1,115,'ddd.gmail.com',true),
('o32tgbkl','森田林',20206005,1,117,'eee.gmail.com',true),
('4hpivesw','熱田神男',20206006,1,117,'fff.gmail.com',true),
('9vacugh2','秋田冬美',20206007,1,118,'ggg.gmail.com',true),
('36gew305','上田　下',20206008,1,118,'hhh.gmail.com',true),
('y5r2a6p8',' kebin fake ',20206009,1,118,'iii.gmail.com',true),
('a5p2r9q0','虚無田無心',null,2,null,'jjj.gmail.com',true),
('z3c7d8x1','加藤聡美',null,2,123,'kkk.gmail.com',true),
('k6o4s9f2','遠山千佳',null,2,115,'lll.gmail.com',true),
('e1v0m8y7','啓発登',null,2,118,'mmm.gmail.com',true),
('h2t5l6n3','和井田昇',null,2,118,'nnn.gmail.com',true),
('d6t2w9v8','無所俗助',null,2,null,'ooo.gmail.com',true),
('h4o1z7f3','江藤　近	',null,0,null,'ppp.gmail.com',true),
('q5p0r9a2','辞田博志',null,2,null,'qqq.gmail.com',false);

INSERT INTO `oa`(request_at,start_time,start_flame,end_time,end_flame,location,status,read_flag,student_comment,teacher_comment,user_id,division_id)
VALUES
('2020-10-01','2020-10-01 12:00:00',1,'2020-10-01 12:00:00',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01 12:00:00',1,'2020-10-01 12:00:00',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01 12:00:00',1,'2020-10-01 12:00:00',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01 12:00:00',1,'2020-10-01 12:00:00',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01 12:00:00',1,'2020-10-01 12:00:00',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01 12:00:00',1,'2020-10-01 12:00:00',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01 12:00:00',1,'2020-10-01 12:00:00',2,'A棟',1,false,'テスト','テスト',1,100),
('2020-10-01','2020-10-01 12:00:00',1,'2020-10-01 12:00:00',2,'A棟',1,false,'テスト','テスト',1,100);
