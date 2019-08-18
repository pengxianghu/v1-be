'''
@Description: 
@Autor: pengxianghu
@Date: 2019-08-17 22:10:49
@LastEditTime: 2019-08-17 22:19:54
'''
#!/usr/bin/python
# -*- coding: utf-8 -*-

import pymysql
import string
import time

def do_schedule():
    db = pymysql.connect(host='119.23.70.24', port=25003, user='root',passwd='990085', db='v1-be', charset='utf8')
    cursor = db.cursor()
    sql = "UPDATE schedule SET `status` = `status` + 1 WHERE `active` = 1 AND `status` < 4"
    try:
        cursor.execute(sql)
        db.commit()
    except:
        db.rollback()
    db.close()
    cursor.close()

def main():
    do_schedule()

if __name__ == "__main__":
    main()
