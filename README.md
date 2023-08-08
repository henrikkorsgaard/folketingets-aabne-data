# folketingets-aabne-data

This projects explore different ways for interacting with the Danish Parliament Open data. 

The goal is to implement a GraphQL microservice for querying the Danish Parliament Open Data.

## Notes

### Data ingestion
The ingest directory contains a README.md and a series of utilities for ingesting the data from the MS SQLServer backup provided by the [Danish Parliament Open Data initiative](https://www.ft.dk/-/media/sites/ft/pdf/dokumenter/aabne-data/oda-browser_brugervejledning.ashx). You can:

- Use the generated PSQL and Sqlite SQL scripts to create the neccesary tables for the data
- Ingest the data from the MS SQL Server into a SQLite database file (around 750mb)
- Ingest the data from the MS SQL Server into a Postgres DB (around 960mb), see config_dev.env for hints about configuration (or change to match local psql setup)
