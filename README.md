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
- Klima love?

# Design considerations
In terms of the design, I've hit a forking path: 

I can either design this services as a client agnostic API that coincidently sends HTML instead of JSON. This decouples style and context, but not structure. Relations are not decoupled (links) and there is a snatch around the API design. The service _should_ follow a RESTful design, e.g. /lovforslag returns a list and /lovforslag/:id returns a specific lovforslag. However, this creates an issue for clients. They would need some sort of routing intermediary, e.g. translating /lovforslag/:id to lovforslag?id=:id to avoid mirroring the file structure in the client layer, e.g. `<a href="/lovforslag/:id.html">Link</a>` or a thick client that handle this. If needing an intermediary or thick client, then the API might as well be JSON-based RPC-like API. 

Or I can design the service as an [HATEOAS](https://htmx.org/essays/hateoas/) application. This means that this server serves the application. If needing a decoupled service, I can add JSON return on selected endpoints if needed _or_ refactor the API into its own service. Then this application become the intermediary backend-for-the-frontend. A HTMX BFF cannot be designed to deliver html as a client agnostic service in the same way as a JSON based API. 

In my opinion, JSON-based APIs cannot be in the back-end-for-the-frontend layer, because then a lot of work is pushed to the client. JSON-based APIs either need a BFF server in the middle or result in thick frontend clients. I don't like thick frontend clients for just binding data to HTML and rendering. They may be useful for interactive islands, when the Model-View-Controller is entirely in a frontend interface. 

Going for a tight coupling with a BFF HOATEOAS design have two downsides:

- I cannot do a full static HTML client without a server or HTMX without static site generation (That sounds like a fun challenge).
- I cannot use the BFF API from other services without transclusion, which might be hard in terms of handling relations and links (Another fun challenge). 

