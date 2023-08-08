# Utils and information

This folder contains a few ingest utilities for setting up and ingesting the data into PostgreSQL or SQLite. Python is used for ingestion. I use Golang to test, because then they are part of a the global tests.


## Getting the data from ft.dk
The easiest way to obtain the full data is to download the Microsoft SQLServer backup from the [link in the guide, see page 9 (in Danish)](https://www.ft.dk/-/media/sites/ft/pdf/dokumenter/aabne-data/oda-browser_brugervejledning.ashx).

For linux, to install Microsoft SQL server and tools follow [this guide](https://blog.devart.com/how-to-install-sql-server-on-linux-ubuntu.html).

Very importantly, one need to change the default directories. Follow [this guide](https://www.mssqltips.com/sqlservertip/4652/how-to-change-default-data-and-log-file-directory-for-sql-server-running-on-linux/), but note that the mssql-conf keys have changed, see [MS guide](https://learn.microsoft.com/en-us/sql/linux/sql-server-linux-configure-mssql-conf?view=sql-server-ver16#datadir).

Then you need to move the downloaded database backup, oda.bak, to the data directory. Ensure that the user mssql is the owner and has read/write access. Then run the following command:

`sqlcmd -S localhost -U sa -P [password here!] -Q "RESTORE DATABASE ODA FROM DISK = '/opt/mssql/data/oda.bak' WITH MOVE 'ODA' TO '/opt/mssql/data/ODA.mdf', MOVE 'ODA_log' TO '/opt/mssql/data/ODA_log.ldf'"`

Replace the [password here!] with the database admin password. 

## Generate-oda-sql-setup-py
This python script will connect to a Microsoft Sql Server with the oda.bak backup from above. 

I had multiple driver and OpenSSL (?) issues when using the ODBC Driver 18 for SQL Server. Following [this guide helped]( https://www.cdata.com/kb/tech/sql-odbc-python-linux.rst) and also rolling back to the ODBC Driver 17. 

For the script to work out of the box, you need to create the user 'odadev' with the password 'odadev1234', see [this guide.](https://www.sqlservertutorial.net/sql-server-administration/sql-server-create-user/)

The script is used to generate the SQL script to create the tables for Postgres and Sqlite. The generated files are in the sql folder. 

## Database_sql_test.go
This file test  the connection to the MS SQL Server and the setup of Postgresql and Sqlite. Run it with `go test -v` to make sure that everything works before using the ingest file. 