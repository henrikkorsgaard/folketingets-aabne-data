import os
import pyodbc

from dotenv import load_dotenv
from pathlib import Path
import sqlite3

dotenv_path = Path("../config_dev.env")
load_dotenv(dotenv_path=dotenv_path)

conn =  pyodbc.connect(driver='{ODBC Driver 17 for SQL Server}', host=os.getenv("MSSQL_DATABASE_HOST"), database=os.getenv("MSSQL_DATABASE_NAME"), user=os.getenv("MSSQL_DATABASE_USER"), password=os.getenv("MSSQL_DATABASE_PASS"))
cursor = conn.cursor()

sqlite_conn = sqlite3.connect("data/oda.sqlite.db")
sqlite_cursor = sqlite_conn.cursor()

# Create the database
with open("sql/oda.sqlite.sql", "r") as sqlfile:
    sqlite_cursor.executescript(sqlfile.read())

tables = ['Afstemning', 'Afstemningstype', 'Aktør', 'AktørAktør', 'AktørAktørRolle', 'Aktørtype', 'Dagsordenspunkt', 'DagsordenspunktDokument', 'DagsordenspunktSag', 'Dokument', 'DokumentAktør', 'DokumentAktørRolle', 'Dokumentkategori', 'Dokumentstatus', 'Dokumenttype', 'Emneord', 'EmneordDokument', 'EmneordSag', 'Emneordstype', 'EntitetBeskrivelse', 'Fil', 'KolloneBeskrivelse', 'Møde', 'MødeAktør', 'Mødestatus', 'Mødetype', 'Omtryk','Periode', 'Sag', 'SagAktør', 'SagAktørRolle', 'SagDokument', 'SagDokumentRolle', 'Sagskategori', 'Sagsstatus', 'Sagstrin', 'SagstrinAktør', 'SagstrinAktørRolle', 'SagstrinDokument', 'Sagstrinsstatus','Sagstrinstype', 'Sagstype', 'Sambehandlinger', 'Stemme', 'Stemmetype', ]

for table in tables:
    print("Ingesting all rows from MSSQL oda." + table + " into SQLite oda."+ table)
    mssql = "SELECT * FROM " + table
    cursor.execute(mssql)
    all = cursor.fetchall()
   
    vals = "?," * len(all[0])
    vals = vals[0:-1]
    
    sql = "INSERT OR IGNORE INTO " + table + " VALUES (" + vals + ")"
    sqlite_cursor.executemany(sql, all)
    sqlite_conn.commit()

sqlite_conn.close()
    
    