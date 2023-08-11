import os
import pyodbc
import psycopg2

from dotenv import load_dotenv
from pathlib import Path

dotenv_path = Path("../config_dev.env")
load_dotenv(dotenv_path=dotenv_path)

conn =  pyodbc.connect(driver='mssqæ', host=os.getenv("MSSQL_DATABASE_HOST"), database=os.getenv("MSSQL_DATABASE_NAME"), user=os.getenv("MSSQL_DATABASE_USER"), password=os.getenv("MSSQL_DATABASE_PASS"))

cursor = conn.cursor()
#"postgres://%s:%s@%s:%s/%s", user,pass,host,port,name
psql_conn = psycopg2.connect(database=os.getenv("PSQL_DATABASE_NAME"),user=os.getenv("PSQL_DATABASE_USER"),password=os.getenv("PSQL_DATABASE_PASS"), host=os.getenv("PSQL_DATABASE_HOST"), port=os.getenv("PSQL_DATABASE_PORT"))
psql_cursor = psql_conn.cursor()
with open("sql/oda.psql.sql", "r") as sqlfile:
    psql_cursor.execute(sqlfile.read())
    psql_conn.commit()


tables = ['Afstemning', 'Afstemningstype', 'Aktør', 'AktørAktør', 'AktørAktørRolle', 'Aktørtype', 'Dagsordenspunkt', 'DagsordenspunktDokument', 'DagsordenspunktSag', 'Dokument', 'DokumentAktør', 'DokumentAktørRolle', 'Dokumentkategori', 'Dokumentstatus', 'Dokumenttype', 'Emneord', 'EmneordDokument', 'EmneordSag', 'Emneordstype', 'EntitetBeskrivelse', 'Fil', 'KolloneBeskrivelse', 'Møde', 'MødeAktør', 'Mødestatus', 'Mødetype', 'Omtryk','Periode', 'Sag', 'SagAktør', 'SagAktørRolle', 'SagDokument', 'SagDokumentRolle', 'Sagskategori', 'Sagsstatus', 'Sagstrin', 'SagstrinAktør', 'SagstrinAktørRolle', 'SagstrinDokument', 'Sagstrinsstatus','Sagstrinstype', 'Sagstype', 'Sambehandlinger', 'Stemme', 'Stemmetype', ]

for table in tables:
    print("Ingesting all rows from MSSQL oda." + table + " into PSQL oda."+ table)
    mssql = "SELECT * FROM " + table
    cursor.execute(mssql)
    all = cursor.fetchall()
   
    vals = "%s," * len(all[0])
    vals = vals[0:-1]
    
    sql = "INSERT INTO " + table + " VALUES (" + vals + ")"
    psql_cursor.executemany(sql, all)
    psql_conn.commit()
    break

psql_conn.close()
    