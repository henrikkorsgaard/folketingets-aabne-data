import pyodbc
import os
from dotenv import load_dotenv
from pathlib import Path

dotenv_path = Path("../config_dev.env")
load_dotenv(dotenv_path=dotenv_path)

cnxn =  pyodbc.connect(driver='{ODBC Driver 17 for SQL Server}', host=os.getenv("MSSQL_DATABASE_HOST"), database=os.getenv("MSSQL_DATABASE_NAME"), user=os.getenv("MSSQL_DATABASE_USER"), password=os.getenv("MSSQL_DATABASE_PASS"))
cursor = cnxn.cursor()


# From https://oda.ft.dk/Home/OdaModel
# and sqlcmd -S localhost -U sa -P [password here!] -d ODA -Q "select *^Crom information_schema.tables" -o "/home/au227822/code/folketingets-aabne-data/tables.txt"
# Some of these are views on Sag, based on typeid: Aktstykke, Almdel, Debat, EUSag, Forslag
# I have excluded these for now. 
tables = ['Afstemning', 'Afstemningstype', 'Aktør', 'AktørAktør', 'AktørAktørRolle', 'Aktørtype', 'Dagsordenspunkt', 'DagsordenspunktDokument', 'DagsordenspunktSag', 'Dokument', 'DokumentAktør', 'DokumentAktørRolle', 'Dokumentkategori', 'Dokumentstatus', 'Dokumenttype', 'Emneord', 'EmneordDokument', 'EmneordSag', 'Emneordstype', 'EntitetBeskrivelse', 'Fil', 'KolloneBeskrivelse', 'Møde', 'MødeAktør', 'Mødestatus', 'Mødetype', 'Omtryk','Periode', 'Sag', 'SagAktør', 'SagAktørRolle', 'SagDokument', 'SagDokumentRolle', 'Sagskategori', 'Sagsstatus', 'Sagstrin', 'SagstrinAktør', 'SagstrinAktørRolle', 'SagstrinDokument', 'Sagstrinsstatus','Sagstrinstype', 'Sagstype', 'Sambehandlinger', 'Stemme', 'Stemmetype', ]
pkeys = {}

mssql_to_psql_types = {
    "int": "INTEGER",
    "smallint": "SMALLINT",
    # Currently, I do not have a good way of determing the length of the datatype in the information_schema
    "nvarchar": "TEXT",
    "varchar": "TEXT",
    "char": "TEXT",
    "datetime": "TIMESTAMP",
    "bit":"BOOLEAN"
}

mssql_to_sqlite_types = {
    "int": "INTEGER",
    "smallint": "INTEGER",
    "nvarchar": "TEXT",
    "varchar": "TEXT",
    "char": "TEXT",
    "datetime": "TEXT",
    "bit":"INTEGER"
}

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

# very crude way of cleaning the file for each run
sqlitefile = open("sql/oda.sqlite.sql", "w")
sqlitefile.write("")
sqlitefile.close()
sqlitefile = open("sql/oda.sqlite.sql", "a")

psqlfile = open("sql/oda.psql.sql", "w")
psqlfile.write("")
psqlfile.close()
psqlfile = open("sql/oda.psql.sql", "a")


for table in tables:
    print("Creating SQL files for " + table)
    sql = "SELECT column_name, data_type, is_nullable FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = '" + table + "'"
    cursor.execute(sql)
    row = cursor.fetchone()
    
    sqlitecreate = "DROP TABLE IF EXISTS " + table + ";\n\n"
    sqlitecreate += "CREATE TABLE IF NOT EXISTS " + table + "  (\n"

    psqlcreate = "DROP TABLE IF EXISTS " + table + ";\n\n"
    psqlcreate += "CREATE TABLE IF NOT EXISTS " + table + "  (\n"


    while row:
        
        if row[0] == pkeys[table]:
            sqlitecreate += "\t" + row[0] + " " + mssql_to_sqlite_types[row[1]] + " PRIMARY KEY,\n"
            psqlcreate += "\t" + row[0] + " " + mssql_to_psql_types[row[1]] + " PRIMARY KEY,\n"
        else:
            notnull = " NOT NULL,\n" if row[2] == "NO" else ",\n"
            sqlitecreate += "\t" + row[0] + " " + mssql_to_sqlite_types[row[1]] + notnull
            psqlcreate += "\t" + row[0] + " " + mssql_to_psql_types[row[1]] + notnull

        row = cursor.fetchone()
    
    sqlitecreate = sqlitecreate[0:-2]
    sqlitecreate += "\n);\n\n"
    sqlitefile.write(sqlitecreate)

    psqlcreate = psqlcreate[0:-2]
    psqlcreate += "\n);\n\n"
    psqlfile.write(psqlcreate)
    

psqlfile.close()
