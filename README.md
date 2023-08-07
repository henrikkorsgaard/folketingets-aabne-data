# folketingets-aabne-data

This projects explore different ways for interacting with the Danish Parliament Open data. 

## Developer notes (will be some sort of documentation at some point)

### Getting the database
The fastest way to get the database for local exploration is by downloading it following the guide from [ft.dk, see page 9 (in Danish)](https://www.ft.dk/-/media/sites/ft/pdf/dokumenter/aabne-data/oda-browser_brugervejledning.ashx) 

To install Microsoft SQL server and tools follow [this guide](https://blog.devart.com/how-to-install-sql-server-on-linux-ubuntu.html).

Very importantly, one need to change the default directories. Follow [this guide](https://www.mssqltips.com/sqlservertip/4652/how-to-change-default-data-and-log-file-directory-for-sql-server-running-on-linux/), but note that the mssql-conf keys have changed, see [MS guide](https://learn.microsoft.com/en-us/sql/linux/sql-server-linux-configure-mssql-conf?view=sql-server-ver16#datadir).

Then you need to move the downloaded database backup, oda.bak, to the data directory. Ensure that the user mssql is the owner and has read/write access. Then run the following command:

`sqlcmd -S localhost -U sa -P [password here!] -Q "RESTORE DATABASE ODA FROM DISK = '/opt/mssql/data/oda.bak' WITH MOVE 'ODA' TO '/opt/mssql/data/ODA.mdf', MOVE 'ODA_log' TO '/opt/mssql/data/ODA_log.ldf'"`

Replace the [password here!] with the database admin password. 

### Update with REST service
You can also use the REST interface to update the database with the following query ([see the guide page 10](https://www.ft.dk/-/media/sites/ft/pdf/dokumenter/aabne-data/oda-browser_brugervejledning.ashx)):

`https://oda.ft.dk/api/Sag?$inlinecount=allpages&$filter=opdateringsdato%20gt%20datetime%272013-05-29T02:00:00%27`

It seems like the 2013 date will get you all the records for each entity. 

## Query the MSSQL database

Connect with: `sqlcmd -S localhost -U sa -P [password here!] -d ODA`

select name from sys.databases
go

remember the GO keyword to run the query.

## Notes on the data
The datamodel is [quite complex](https://oda.ft.dk/Home/OdaModel) in the way it is described by the documentation. 

The key is not to look at the model, but to look at the actual model in the database. 

Easiest way (on linux) to inspect the database is by dumping into textfiles: `sqlcmd -S localhost -U sa -P [password here!] -d ODA -Q "select column_name from information_schema.columns where table_name = 'Debat'" -o "/home/au227822/code/folketingets-aabne-data/debat.txt"`

Get the columns of a table with nullable, datatype etc:

`sqlcmd -S localhost -U sa -P [Password here!] -d ODA -Q "select column_name, data_type, is_nullable from information_schema.columns where table_name = 'Sag'" -o "/home/au227822/code/folketingets-aabne-data/sagcols.txt"`