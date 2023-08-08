import os
import pyodbc

from dotenv import load_dotenv
from pathlib import Path
import sqlite3

dotenv_path = Path("../config_dev.env")
load_dotenv(dotenv_path=dotenv_path)

conn =  pyodbc.connect(driver='{ODBC Driver 17 for SQL Server}', host=os.getenv("MSSQL_DATABASE_HOST"), database=os.getenv("MSSQL_DATABASE_NAME"), user=os.getenv("MSSQL_DATABASE_USER"), password=os.getenv("MSSQL_DATABASE_PASS"))
cursor = conn.cursor()

sqlite_conn = sqlite3.connect("data/odatest.sqlite.db")
sqlite_cursor = sqlite_conn.cursor()

tables = ['Afstemning', 'Afstemningstype', 'Aktør', 'AktørAktør', 'AktørAktørRolle', 'Aktørtype', 'Dagsordenspunkt', 'DagsordenspunktDokument', 'DagsordenspunktSag', 'Dokument', 'DokumentAktør', 'DokumentAktørRolle', 'Dokumentkategori', 'Dokumentstatus', 'Dokumenttype', 'Emneord', 'EmneordDokument', 'EmneordSag', 'Emneordstype', 'EntitetBeskrivelse', 'Fil', 'KolloneBeskrivelse', 'Møde', 'MødeAktør', 'Mødestatus', 'Mødetype', 'Omtryk','Periode', 'Sag', 'SagAktør', 'SagAktørRolle', 'SagDokument', 'SagDokumentRolle', 'Sagskategori', 'Sagsstatus', 'Sagstrin', 'SagstrinAktør', 'SagstrinAktørRolle', 'SagstrinDokument', 'Sagstrinsstatus','Sagstrinstype', 'Sagstype', 'Sambehandlinger', 'Stemme', 'Stemmetype', ]
pkeys = {}
# First we need to get the primary key information for each table
pksql = '''SELECT 
    kcu.*
FROM 
    INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu
    INNER JOIN 
    INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
        ON  tc.TABLE_NAME = kcu.TABLE_NAME AND 
            tc.CONSTRAINT_NAME = kcu.CONSTRAINT_NAME
ORDER BY 
    tc.TABLE_NAME,
    tc.CONSTRAINT_NAME,
    kcu.ORDINAL_POSITION
'''

cursor.execute(pksql)
row = cursor.fetchone() 
# See https://learn.microsoft.com/en-us/sql/relational-databases/system-information-schema-views/key-column-usage-transact-sql?view=sql-server-ver16
while row: 
    if "PK_" in row[2] and row[2].index("PK_") == 0 and row[5] in tables:
        pkeys[row[5]] = row[6]
       
    row = cursor.fetchone()

# Create the database
with open("sql/oda.sqlite.sql", "r") as sqlfile:
    sqlite_cursor.executescript(sqlfile.read())

for table in tables:
    print("Ingesting 1000 rows from MSSQL oda." + table + " into SQLite oda."+ table)
    mssql = "SELECT TOP(1000) * FROM " + table + " ORDER BY '" + pkeys[table] + "' DESC" # All the pkeys are id, but you never know
    cursor.execute(mssql)
    all = cursor.fetchall()
   
    vals = "?," * len(all[0])
    vals = vals[0:-1]
    
    sql = "INSERT OR IGNORE INTO " + table + " VALUES (" + vals + ")"
    sqlite_cursor.executemany(sql, all)
    sqlite_conn.commit()

sqlite_conn.close()
    
    