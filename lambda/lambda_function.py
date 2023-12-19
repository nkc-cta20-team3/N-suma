import sys
import logging
import pymysql
import json

#20230706変更　secret追加
# RDSの情報登録
#更新した後
rds_host  = "ctagr3-rds.cmedpe62g1hv.us-east-1.rds.amazonaws.com"
user_name = "admin"
password = "admin0pass"
db_name = "cta20gr3"

logger = logging.getLogger()
logger.setLevel(logging.INFO)

# 接続
try:
    conn = pymysql.connect(host=rds_host, user=user_name, passwd=password, db=db_name, connect_timeout=5)
except pymysql.MySQLError as e:
    logger.error("ERROR: Unexpected error: Could not connect to MySQL instance.")
    logger.error(e)
    sys.exit()

logger.info("SUCCESS: Connection to RDS MySQL instance succeeded")

def lambda_handler(event, context):
    #logger.info(event)
    #print(json.dumps(event, indent=4))
    #data = event
    data = json.loads(event['body'])
    value1 = data['value1']
    value2 = data['value2']
    value3 = data['value3']
    value4 = data['value4']
    value5 = data['value5']
    value6 = data['value6']
    value7 = data['value7']
    value8 = data['value8']
    value9 = data['value9']
    value10 = data['value10']

    item_count = 0
    sql_string = f"insert into absence_document (student_id, company_id,reason_id, request_date, absence_start_date, absence_start_flame, absence_end_date, absence_end_flame, location, student_comment) values({value1}, {value2},{value3},'{value4}','{value5}',{value6},'{value7}',{value8},'{value9}','{value10}')"

    with conn.cursor() as cur:
        cur.execute("create table if not exists Customer ( CustID  int NOT NULL, Name varchar(255) NOT NULL, PRIMARY KEY (CustID))")
        cur.execute(sql_string)
        conn.commit()
        cur.execute("select * from Customer")
        logger.info("The following items have been added to the database:")
        for row in cur:
            item_count += 1
            logger.info(row)
    conn.commit()

    return "Added %d items to RDS MySQL table" %(item_count)