from dotenv import load_dotenv
from pathlib import Path
import sqlite3

import requests
import time

dotenv_path = Path("../config_dev.env")
load_dotenv(dotenv_path=dotenv_path)

conn = sqlite3.connect("data/oda.api.sqlite.db")
cursor = conn.cursor()

tables = ['Afstemning', 'Afstemningstype', 'Aktør', 'AktørAktør', 'AktørAktørRolle', 'Aktørtype', 'Dagsordenspunkt', 'DagsordenspunktDokument', 'DagsordenspunktSag', 'Dokument', 'DokumentAktør', 'DokumentAktørRolle', 'Dokumentkategori', 'Dokumentstatus', 'Dokumenttype', 'Emneord', 'EmneordDokument', 'EmneordSag', 'Emneordstype', 'EntitetBeskrivelse', 'Fil', 'KolloneBeskrivelse', 'Møde', 'MødeAktør', 'Mødestatus', 'Mødetype', 'Omtryk','Periode', 'Sag', 'SagAktør', 'SagAktørRolle', 'SagDokument', 'SagDokumentRolle', 'Sagskategori', 'Sagsstatus', 'Sagstrin', 'SagstrinAktør', 'SagstrinAktørRolle', 'SagstrinDokument', 'Sagstrinsstatus','Sagstrinstype', 'Sagstype', 'Sambehandlinger', 'Stemme', 'Stemmetype' ]

# We need to check if the database file exist and if the database is sound? 
# Check for schema consistency?
# Or remove drop table from the sqlite text?
# Create the database
with open("sql/oda.sqlite.sql", "r") as sqlfile:
    cursor.executescript(sqlfile.read())

def insert_data(rows, table):
    data = []

    for row in rows:
        rowdata = ()
        for v in row:
            rowdata = rowdata + (row[v],)
        
        data.append(rowdata)

    vals = "?," * len(rows[0].keys())
    vals = vals[0:-1]

    sql = "INSERT OR IGNORE INTO " + table + " VALUES (" + vals + ")"
    cursor.executemany(sql, data)
    conn.commit()

def request_handler(query, table, count):
    response = requests.get(query)
    if response.status_code != 200:
        print("Response from " +  query + " returned status code: " + response.status_code)
        exit(1)
    
    data = response.json()
    if len(data["value"]) == 0:
        print("Nothing to update from {}".format(table))
        return

    update_count = data["odata.count"]
    status = "Updating table {}: Records {} - {} out of {} new records".format(table, str(count), str(count+100), update_count)
    print(status)
    insert_data(data["value"], table)
    time.sleep(1)
    if "odata.nextLink" in data:
        return data["odata.nextLink"]
    
    return


for table in tables:
    print("Checking for updates on {}".format(table))
    id = 0
    sql = "SELECT id FROM " + table + " ORDER BY id DESC LIMIT 1;"
    cursor.execute(sql)
    result = cursor.fetchall()
    if len(result) > 0:
        id = result[0][0]
   
    # While the API recommends updating by date, I am not sure that is the right move. 
    # Some of the early ids have a never update date, so we will have date skips
    # This could be avoided by ordering the SELECT opdateringsdata FROM <table> ORDER BY opdateringsdato DESC LIMIT 1
    # But I feel safer doing it by the id as it is the primary key in the MSSQL database behind the API
    query = 'https://oda.ft.dk/api/{}?$inlinecount=allpages&$filter=id%20gt%20{}'.format(table, str(id))
    
    count = 0
    # While is better than recursion, because Python has a recursion depth limit of a 1000
    while True:
        query = request_handler(query, table, count)
        if query == None:
            break
        count+=100




    
    