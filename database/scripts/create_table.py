#!/usr/bin/env python3
#
# Initialize tables

import database

sql_file_table = [
    '../sql/oqs_record.sql',
    '../sql/user_record.sql',
]

def main():
    cnx = database.connect_db()
    cursor = cnx.cursor()
    for i in sql_file_table:
        with  open(i) as f:
            sql_words = f.read()
            cursor.execute(sql_words)
    cursor.close()
    cnx.close()

if __name__ == '__main__':
    main()