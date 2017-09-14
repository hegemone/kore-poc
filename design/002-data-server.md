- Feature Name: Data server design
- Start Date: 2017-08-25
- RFC PR:
- Showbot Issue:

# Summary
[summary]: #summary

As part of RFC 001, Showbot will move to a client/server based architecture.
This RFC addresses the design of this data server.

# Motivation
[motivation]: #motivation

The data server addresses the need for multiple clients to access the same data
to ensure strong consistency across all of the aforementioned clients.

# Guide-level explanation
[guide-level-explanation]: #guide-level-explanation

The data server will consist of two parts an API to allow data access and 
manipulation and business logic to manipulate the data and store it in a 
database.

Interacting with non-chat service, third-party APIs, i.e. Digital Ocean, should
also be done in the data server. Using the data server to call these APIs cuts 
down on the amount of work necessary to update the entire bot platform should 
the third-party API change.

For simplicity's sake, the initially supported database will be
sqlite3. More databases can be added in the future if the performance of sqlite3
does not meet our requirements.

The API will be implemented using the HTTP protocol to allow for easy client
interaction and allow for (mostly) stateless client interaction. The data server
will also need to ensure certain HTTP requests are from people who are properly
authenticated.

# Reference-level explanation
[reference-level-explanation]: #reference-level-explanation

This section will largely focus on the API and database schema design as they
are the two toughest things to design correctly in any server.

The API will be used for the following interactions:

- Spinning up and down Digital Ocean droplets
- Link suggestions
- Show title suggestions and vote counts
- Storing and regurgitating quotes
- Calling the Twitter API (to be integrated at a later time)
- Calling the necessary APIs for Bitcoin stats (to be integrated at a later time)

The API endpoints should be as follows:

```
GET /api

Returns list of API endpoints
```
```
GET /api/droplets
{
    "droplets": ["DROP1", "DROP2",...,"DROPn"]
}
```
```
POST /api/droplets/{start,stop}
{
    "droplet": "name-of-droplet"
}
```
```
POST /api/link
{
    "link": "url"
}
```
```
POST /api/show/suggest
{
    "nick": "relevant-user-nick"
    "suggestion": "suggested-title"
}
```
```
POST /api/show/start (authed)

{
    "name": "show-name"
}
```
```
GET /api/show/stop

Stops currently live show
```
```
GET /api/quotes?name=<name>

Without query parameter, return all the quotes.
With the query param, return quote for <name>
```
```
POST /api/quotes

{
    "name": "name"
    "quotes": ["quote1", "quote2"]
}

Without query param, set quotes to only specified
With query param, set quote list for <name>
```
```
POST /api/quotes/add

{
    "name": "name-to-add-to",
    "quote": "quote-to-add"
}
```
```
POST /api/quotes/delete

{
    "name": "user-nick",
    "quote": "quote-to-delete"
}
```
```
POST /api/vote

{
    "name": "user-nick"
    "title": "title-of-show"
}
```

As an aside, the programming language to be used to build the server should also
be decided. At this point in time, I see three viable languages for developing
the server with my vote going to Go (not that this is a democracy): Go, Python, 
and Ruby.

#### Ruby considerations
Pros:
- Already used for the existing IRC bot
- Development language of choice for Rikai
- ??? (I don't know enough about Ruby to say what other pros it has)

Cons:
- Can be difficult to ship the application to a new server due to dependencies
(see my complaints about a failing bundle install in [#94](https://github.com/rikai/Showbot/issues/94))
- Many people are not very familiar with Ruby as a language (me included)

#### Python considerations
Pros:
- Well-established language
- Many already familiar with it
- Plenty of packages already available for almost anything

Cons:
- Suffers the same application shipment problem as ruby (mentioned above)
- Not exactly fast

#### Go considerations
Pros:
- Fast, lightweight, and statically-typed
- Simple language to learn and use
- Get a [highly performant HTTP server stack](https://golang.org/pkg/net/http/) 
for free
- Code compiles down into a single binary for easy deployment

Cons:
- Young language
- Not widely used yet in the OSS world

# Drawbacks
[drawbacks]: #drawbacks

Using a backend server makes each of the bots less capable of being run on its
own, possibly reducing the likelihood of someone else in the OSS community 
deciding to use it for their own purposes.

The above drawback can be worked around by allowing the bot to save data
directly to its own sqlite database; however, doing so makes the bot "heavier"
and requires more code that must be maintained.

# Rationale and Alternatives
[alternatives]: #alternatives

Given the alternatives mentioned in RFC 001, using a data server seems to be the
best way to scale out the bot platform for its immediate use in the JB studio.

# Unresolved questions
[unresolved]: #unresolved-questions

This question may be better suited for the IRC bot design RFC but is relevant to
the design of the data server: should all plugins be provided with a way to be 
run statelessly? The droplet plugin seems to make simple calls to the DO api as
well as be extremely specific to JB's livestream broadcasting architecture.