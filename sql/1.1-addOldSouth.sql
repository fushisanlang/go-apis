-- DROP TABLE IF EXISTS base_user;
-- CREATE TABLE base_user (
--     id INTEGER PRIMARY KEY AUTOINCREMENT,
--     username TEXT ',
--     create_time TEXT  ,
--     last_login_time TEXT ',
--     source INTEGER ,
--     token TEXT  ,
--     admin INTEGER 
-- );               
-- 
-- 
-- DROP TABLE IF EXISTS base_source_type;
-- CREATE TABLE base_source_type(
-- id INTEGER PRIMARY KEY AUTOINCREMENT,
-- source_id INTEGER,
-- source_info TEXT
-- );
-- 
-- INSERT INTO base_source_type VALUES(1,1,"windsmeller"),(2,2,"khronos"),(3,3,"oldsouth");
-- 

DROP TABLE IF EXISTS oldsouth_record;
CREATE TABLE oldsouth_record (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    record_time INT,
    userid INTEGER,
    record_type TEXT,
    record_value REAL

);


DROP TABLE IF EXISTS oldsouth_type;
CREATE TABLE oldsouth_type (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    record_type TEXT,
    record_name TEXT,
    record_name_cn TEXT,
    record_info TEXT
); 


INSERT INTO oldsouth_type (record_type, record_name,record_name_cn, record_info)
VALUES ('FBG', 'fasting blood glucose','空腹血糖', ''),
  ('PBG','postprandial blood glucose','餐后血糖',''),
  ('W','weight','体重',''),
  ('BFR','Body fat rate','体脂率',''),
  ('E3','exercise energy expenditure','运动能量消耗',''),
  ('DoE','duration of exercise','运动时长',''),
  ('RHR','resting heart rate','静息心率',''),
  ('EHR','exercise heart rate','运动心率',''),
  ('ST','sleep time','睡眠时长',''),
  ('FR','food rate','食物比例','200g粗粮主食+200g肉菜+200g素菜'),
  ('SEX','sex','性行为','0 无措施，1 byt，2 紧急，3 短效'),
  ('DP','diastolic (blood) pressure','舒张压',''),
  ('SP','systolic (blood) pressure','收缩压','');
