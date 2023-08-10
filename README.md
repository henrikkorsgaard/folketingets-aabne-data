# folketingets-aabne-data

This projects explore different ways for interacting with the Danish Parliament Open data. 

The goal is to implement a GraphQL microservice for querying the Danish Parliament Open Data.

## Try it!
I have added a small http server that will serve the GraphQL endpoint (/graphql) and a simple GraphIQL (/graphiql) front end for exploration. 

Pull and run and explore the data on localhost:8080/graphiql!

## Notes

### Tests
I mainly run tests to make sure that I cover the contract implied in the GraphQL schema definition and making sure that I can call the database with the individual entities and resolvers. I do not verify the structs and data coverage in the resolver queries. That will be caught with the schema contract and/or graphIql test.

### Data ingestion
The ingest directory contains a README.md and a series of utilities for ingesting the data from the MS SQLServer backup provided by the [Danish Parliament Open Data initiative](https://www.ft.dk/-/media/sites/ft/pdf/dokumenter/aabne-data/oda-browser_brugervejledning.ashx). You can:

- Use the generated PSQL and Sqlite SQL scripts to create the neccesary tables for the data
- Ingest the data from the MS SQL Server into a SQLite database file (around 750mb)
- Ingest the data from the MS SQL Server into a Postgres DB (around 960mb), see config_dev.env for hints about configuration (or change to match local psql setup)

### Docker deployment?

- I should pack this in a docker instance
- See how much Docker will compress the full database
- Make a CRON job for updating the database. 

### Analysis ideas emerging from building this service

- The Afstemning.kommentar field seem to contain info on voting errors, with the name of the member who made the mistake. Maybe count who makes the most mistakes when voting.