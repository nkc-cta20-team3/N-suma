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
('2023-11-01','2023-12-08 09:00:00',1,'2023-12-08 14:00:00',6,'名古屋',2,true,'資格試験のため、公欠を申請します','認可しました',1,100),
('2023-11-01','2023-12-15 10:00:00',null,'2023-12-15 11:00:00',null,'岐阜市XX株式会社',1,true,'移動に時間がかかるので3限まで欲しい',null,1,200),
('2023-11-03','2023-12-15 09:00:00',null,'2023-12-15 10:00:00',null,'学校',-1,true,'オンラインで面接をします','面接先企業を場所欄に()を使って記載してください',2,200),
('2023-11-03','2023-12-12 14:00:00',1,'2023-12-12 15:00:00',6,'東京都YY株式会社',2,true,'移動もあるので、全日公欠を希望します','許可します',3,201)
('2023-11-03','2023-10-30 09:00:00',1,'2023-10-31 21:00:00',6,'実家',2,false,'祖父の葬儀に参加します','許可しました',5,300),
('2023-11-10','2023-11-27 10:00:00',null,'2023-11-27 11:00:00',null,'名古屋市',1,true,'ISO株式会社に面接に行きます',null,7,200),
('2023-11-10','2023-11-28 12:00:00',null,'2023-11-28 15:00:00',null,'名古屋市',1,true,'工学院システム株式会社に面接に行きます',null,7,200),
('2023-11-11','2023-11-29 12:00:00',null,'2023-11-29 13:00:00',null,'名古屋市',1,true,'会社説明会に行きます',null,9,300),
('2023-11-11','2023-12-01 12:00:00',null,'2023-12-01 13:00:00',null,'名古屋市南区',1,true,'252株式会社で面接','テスト',10,201),
('2023-11-13','2023-12-01 9:00:00',null,'2023-12-01 12:00:00',null,'名古屋駅近くの試験センター',1,true,'FEの午前試験があります',null,8,100);
