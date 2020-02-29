#!/usr/bin/env python3
#
# I reserved 0 to 64**3-1 for myself.

import database

def main():
    cnx = database.connect_db()
    cursor = cnx.cursor()
    i = 62**3-1
    cursor.execute("""
        INSERT INTO oqs_records VALUES ("{}", "00", "reserved")
    """.format(i))
    cnx.commit()
    cursor.close()
    cnx.close()

if __name__ == '__main__':
    main()
