import pyodbc #pip install pyodbc  

# I had multiple driver and OpenSSL (?) issues when using the ODBC Driver 18 for SQL Server
# See registred drivers with: odbcinst -q -d
# See https://www.cdata.com/kb/tech/sql-odbc-python-linux.rst
# If I ever need to create a user for a mssql db: https://www.sqlservertutorial.net/sql-server-administration/sql-server-create-user/

server = 'localhost' 
database = 'ODA' 
username = 'odadev' 
password = 'odadev1234' ## need a test user with a different password 
cnxn =  pyodbc.connect(driver='{ODBC Driver 17 for SQL Server}', host=server, database=database, user=username, password=password)
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
psqlfile = open("./oda.psql.sql", "w")
psqlfile.write("")
psqlfile.close()
psqlfile = open("./oda.psql.sql", "a")


for table in tables:
    print("Creating SQL files for " + table)
    sql = "SELECT column_name, data_type, is_nullable FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = '" + table + "'"
    cursor.execute(sql)
    row = cursor.fetchone()
    psqlcreate = "DROP TABLE IF EXISTS " + table + ";\n\n"
    psqlcreate += "CREATE TABLE IF NOT EXISTS " + table + "  (\n"
    while row:
        
        if row[0] == pkeys[table]:
            psqlcreate += "\t" + row[0] + " " + mssql_to_psql_types[row[1]] + " PRIMARY KEY,\n"
        else:
            notnull = " NOT NULL,\n" if row[2] == "YES" else ",\n"
            psqlcreate += "\t" + row[0] + " " + mssql_to_psql_types[row[1]] + notnull

        row = cursor.fetchone()
    
    psqlcreate = psqlcreate[0:-2]
    psqlcreate += "\n);\n\n"
    psqlfile.write(psqlcreate)
    

psqlfile.close()
"""
tables = ['Afstemning', 'Afstemningstype', 'Aktør', 'AktørAktør', 'AktørAktørRolle', 'Aktørtype', 'Dagsordenspunkt']






cursor.execute("select column_name, data_type, is_nullable from information_schema.columns where table_name = 'Sag'") 
row = cursor.fetchone() 

sql = "CREATE TABLE IF NOT EXISTS sat ("

while row: 
    for 
    print(row[0], row[1], row[2])
    row = cursor.fetchone()

sql += ");"

"""