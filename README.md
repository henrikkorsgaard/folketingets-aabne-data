# folketingets-aabne-data

This projects explore different ways for interacting with the Danish Parliament Open data. 

The goal is to implement a GraphQL microservice for querying the Danish Parliament Open Data.

## Use cases
The dataset should make it possible to do some more visual analysis about the legislative process and/or create dashboard-like interfaces for parliament members, legislation, subject categories etc. 

- Query data about specific members of parliament, including their voting and other parliament activities. 
- Query data about specific legislation and the process behind it.
- Query data about the legislative process (votes, stages, time series etc.)
- (Ranked) Search parliament members by name (and other activities)
- (Ranked) Search legislation (and proces) by subject categories, time etc. 

## Try it!

I have added a small http server that will serve the GraphQL endpoint (/graphql) and a simple GraphIQL (/) front-end for exploration. 

Pull and run and explore the data on localhost:8080/

## Notes
### TODO
- [x] Integrate Sag 
- [x] Design and implement the relationship between Afstemning and Sag
- [x] Design and implement the relationship between Sag and Afstemning
- [ ] Integrate Emneord with Sag
- [ ] Integrate search_sag by Emneord
- [ ] Set "" to NULL on strings in database ingestion
- [ ] Implement a Docker deployment option that integrates:
    - [ ] Sqlite database
    - [ ] Daily updates via the Python update script
    - [ ] Logging
- [ ] Revisit the paper to see if we can realise the common use-cases (if no, expand todo)

.
.
.

- [ ] Do a feasibility study of integrating meeting minutes (XML) and search these through a Solr search api.

### Notes on schema design
One of the critical findings in a recent study by my students and in reviewing the ERD behind the data, is that the ERD introduces too many relations for people to be able to formulate a query. When translating the data to a GraphQL API, I will make an opinionated schema. Here are some notes

**Stemme (vote)**: This table contain the votes cast at a particular parliament vote sesstion (table: Afstemning) and then members of parliament (table: Aktør). It does not make sense to have Stemme on the root query. An atomic vote does not make sense without the relations and querying on votes will lead to subsequent queries to establish the relationship ({stemme {afstemning {id} vote {afstemningid, aktør {name}}}}). Hence, I have made it so that the entry point is either afstemning (vote session) or aktør (voter). 

### Tests
I use tests as a write test first (or in parallel) with the code on the database layer (ftoda), resolvers and schema (resolvers) and (at some point) integration testing. I write these based on intuition and when an issue emerges. I do not do ECP analysis/testing. 
I do not verify the structs and data coverage in the resolver queries. That will be caught with the schema contract and/or graphiql test.

### Data ingestion
The ingest directory contains a README.md and a series of utilities for ingesting the data from the MS SQLServer backup provided by the [Danish Parliament Open Data initiative](https://www.ft.dk/-/media/sites/ft/pdf/dokumenter/aabne-data/oda-browser_brugervejledning.ashx). You can:

- Use the generated PSQL and Sqlite SQL scripts to create the neccesary tables for the data
- Ingest the data from the MS SQL Server into a SQLite database file (around 750mb)
- Ingest the data from the MS SQL Server into a Postgres DB (around 960mb), see config_dev.env for hints about configuration (or change to match local psql setup)

#### Use GIT-LFS to get a sqlite database as a zip
I'm using [GIT-LFS](https://docs.github.com/en/repositories/working-with-files/managing-large-files/installing-git-large-file-storage) to link to a zip of the sqlite database file as a zip.

### Alternative approches [Why not?]

There are several ways I could generate GraphQL schemas and/or use a service like PostGraphile to make a GraphQl service on to of the database. 

First, I'm making this service as a deep dive into GraphQL with the goal of understanding the technology, schema and query relation, and how to make a GraphQL microservice. 
I find that developing this from the database and up is a great way to learn.

Second, to get the most out of GraphQL, generated code (db schema/GQL schema first) and external services rely on FOREIGN KEY relations to establish relationships for the GraphQL schema. The MSSQL database does not contain any foreign keys, so generating the schemas and/or using a service will only get me 50% there. The relations are described in the [ODA oData metadata](https://oda.ft.dk/api/$metadata). I just don't feel like parsing XML at the moment. I also suspect that the relations (NavigationProperty) is a messy source of relations. Once I have understood the dataset, I might add foreign keys to the database where it makes sense.

Third, when dealing with more complex relational databases, a lot of database-relations-clutter is introduced (this particular dataset have a lot of relations that are make reading and exploring the data more difficault). One of the benefits of GraphQL is that you can effectively re-design the query model. However, doing that requires a lot of analysis up against common use-cases. An DB schema or greedy schema first approach would just inherit all the relational clutter and force the user to traverse the relation graph in to the deep end. 

### Analysis ideas emerging from building this service

- The Afstemning.kommentar field seem to contain info on voting errors, with the name of the member who made the mistake. Maybe count who makes the most mistakes when voting.