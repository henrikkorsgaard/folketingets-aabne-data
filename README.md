# Folketingets Åbne Data

Just checking commits -- this is not a good way to do this

## User stories

"As a user, I want to see all the regulation passes within a given time period"

"As a user, I want to zoom in on a specific regulation, politician etc."

"As a user, I want to be able to select a politician and see their votes across multiple regulations and proposals"

## Implementation notes

### Repository pattern with interfaces

- Wrap oData API
- Wrap sqlite DB API

## stories

### Controversial legislation

As a citizen, I would like to be able to look up controversial legislation. This is defined by being a close votes, frequently debated, and/or if the voting pattern breaks the party positions.

#### Task 1: Map Sag to legislation
A Sag is a law draft when it has the Sagstype 'Lovforslag'.

Example to find the lovforslag about US troops in Denmark

https://oda.ft.dk/api/Sag?$filter=typeid%20eq%203%20and%20substringof(%27forsvarssamarbejde%27,%20titel) gives "Om forsvarssamarbejde mellem Danmark og Amerikas Forenede Stater m.v."


- Task 2: Map Sag to Vote


## Creating the local database should properly focus on getting taxonomies out, e.g. emneord etc.


### Controversial legislation

- L 87: Smykkeloven
- L 188: Om forsvarssamarbejde mellem Danmark og Amerikas Forenede Stater m.v
- Minister pension
- Buskørselsloven
- 

# Design considerations

I can either let the API pattern follow through /lovforslag/{id} and try to let the HTML pages solve this OR let the HTMl pattern follow through, e.g. /lovforslag?id={id} and adopt the server.

I would rather do the former without having to do significant dynamic routing on the pure static frontend.

If I use HTMX boost, then I have a very thin frontend application and everything is served by the backend. That creates tighter coupling:

- This backend can only serve this particular frontend
- I cannot mix and match the frontend components.