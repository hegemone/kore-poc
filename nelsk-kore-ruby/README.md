# nelsk-kore

POC demo illustrating a number of key pieces of `kore-comm`, the server responsible
for IO between various platforms and the plugins that implement features.

## Usage

`bundle install` and make sure you've got the deps installed. Then just `./run`.

Expects commands as input in the form of `<adapter> !<plugin> content`, example:

```
{> dev/kore-poc/nelsk-kore|kore-comm+  <}
# ./run
[2017-08-27T22:28:39] DEBUG Engine#initialize
[2017-08-27T22:28:39] DEBUG Engine#start
[2017-08-27T22:28:39] DEBUG DiscordAdapter#listen
[2017-08-27T22:28:39] DEBUG IRCAdapter#listen
irc !bacon buddy
[2017-08-27T22:29:11] DEBUG IRCAdapter#message_handler
[2017-08-27T22:29:11] DEBUG PlatformAdapter#message_received
[2017-08-27T22:29:11] DEBUG   identity: irc_duder
[2017-08-27T22:29:11] DEBUG   raw: !bacon buddy
[2017-08-27T22:29:11] DEBUG Command matched, parsing...
[2017-08-27T22:29:11] INFO  Engine#route_ingress
PlatformClient[irc] KoreBot: gives buddy a strip of delicious bacon as a gift from irc_duder.
```

### Points of note
* EM based reactor at the heart of the Engine. Enables parallelized work and evented
triggers.
* `Kore::Comm::PlatformAdapter` - implement to define how the engine can speak to
arbirary text communcation platforms like irc, discord, telegram, etc.
* `Kore::Comm::Plugin` - implement to extend JBot behaviors
* Dependency injection thanks to `dry-container` + `dry-auto_inject`. Makes for
very testable components that have a lot of external dependencies like adapters.
* `client` injection into the adapters. Adapters will almost surely be using
some kind of service-specific client to talk to them. Adapters should expect for
them to be injected.
* Master dependencies like logger and config managers can be inherited for free
by subclassing `Kore::Machinery::Base`.

### Faux pieces
* Pretty much everything in `Kore::Mock`.
* `PlatformClient` - is a simple, fake client for talking with an "external"
service. In reality, these would be completely distinct clients for things like
discord that the adapters use to translate between `$PLATFORM` and `Kore`.
* `InputDemux` - Demultiplexes STDIN and sends messages to the appropriate clients.
Not at all relevant in a real environment.

### Next steps
* Integration with `kore-data`, which doesn't exist yet. Good demo would be
to build a basic OpenAPI definition and generate server/client libs. Fire up the
data server with a sqlite db backing it. Demonstrate CRUD against the `kore-data`
server using the client, say `Kore::Data::User`. Example:

```ruby
match /^suggest\s+(\S+)/i, :method => :command_foo

def command_foo(suggestion)
  Kore::Data::Suggestion.new({
    user_id: Kore::Data::User.find_id_by_platform(platform, originator[:identity]),
    suggestion: suggestion,
  }).save
end
```

* Authentication and authorization between `kore-comm` to `kore-data` remains
a big question, and one that absolutely has to be solved if users can take
privileged actions from platforms.. The comm server needs to
*act on the behalf* of a user. Somehow the comm server needs to delegate to
the adapters and ask `"Do we really trust this person is who thye appear to be?"`
If the answer is yes, how do we masquerade as that user when performing CRUD
against the backend?

* Simply not production ready. Needs a healthy amount of error handling and
general polish, but the core architectural ideas are there.
