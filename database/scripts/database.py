#!/usr/bin/env python3
#
# Database initialize utils 

import mysql.connector
import os

USER = 'root'
PASSWORD = os.getenv("DB_PASSWORD")
HOST = '127.0.0.1'
PORT = 3309
DATABASE = 'oqs'

def connect_db():
    cnx = mysql.connector.connect(user=USER, password=PASSWORD, host=HOST, port=PORT, database=DATABASE)
    return cnx
