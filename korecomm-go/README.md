# korecomm-go

Golang implementation of the [korecomm-ruby](../korecomm-ruby) POC demo illustrating
a number of key pieces of `Kore::Comm` architecture. `Kore::Comm` is the server
responsible for IO between various platforms and the plugins that implement
behavioral features.

Available as a prebuilt docker image: `quay.io/hegemone/korecomm-go-poc:latest`

## Usage

`make run-image` will pull and run the container from quay.

Expects commands as std input in the form of `<adapter> !<plugin> content`, example:

```
{> hegemone/kore-poc/korecomm-go|korecomm-go+  <}
# make run-image
docker run -it quay.io/hegemone/korecomm-go-poc:latest
INFO[0000] ============================================================
INFO[0000]                  Kore::Comm Golang POC
INFO[0000] ============================================================
INFO[0000] Performing startup validations
INFO[0000] Loading extensions
INFO[0000] Loading plugins from: /usr/lib/kore
INFO[0000] -> bacon.plugins.kore.nsk.io
INFO[0000] Started the reg req func...
INFO[0000] Successfully loaded plugins:
INFO[0000] -> bacon.plugins.kore.nsk.io
INFO[0000] Loading adapters from: /usr/lib/kore
INFO[0000] -> ex-discord.adapters.kore.nsk.io
INFO[0000] file: /usr/lib/kore/ex-discord.adapters.kore.nsk.io.so
INFO[0000] ex-discord.adapters::Init
INFO[0000] -> ex-irc.adapters.kore.nsk.io
INFO[0000] file: /usr/lib/kore/ex-irc.adapters.kore.nsk.io.so
INFO[0000] ex-irc.adapters::Init
INFO[0000] Successfully loaded adapters:
INFO[0000] -> ex-discord.adapters.kore.nsk.io
INFO[0000] -> ex-irc.adapters.kore.nsk.io
INFO[0000] Starting engine
test
ERRO[0011] Must send stdin message in format of '<adapter_name> <content>'
discord !bacon
INFO[0017] Demux -> discord
INFO[0017] bacon.plugins::CmdBacon, IngressMessage: {Content:bacon Originator:{Identity:discord-user AdapterName:ex-discord.adapters.kore.nsk.io}}
INFO[0017] discord client got a message from the adapter! [ gives discord-user a strip of delicious bacon. ]
irc !bacon nelsk
INFO[0025] Demux -> irc
INFO[0025] bacon.plugins::CmdBaconGift, IngressMessage: {Content:bacon nelsk Originator:{Identity:irc-user AdapterName:ex-irc.adapters.kore.nsk.io}}
INFO[0025] irc client got a message from the adapter! [ gives nelsk a strip of delicious bacon as a gift from irc-user ]
nulladapter !bacon
INFO[0036] Client nulladapter not registered with stdin demux, skipping...
```

## Building and running locally

**NOTE**: Demo relies on beta plugin feature introduced with go 1.8 that only supports
linux and Darwin. Additionally, I recently tried a build with 1.9 and had a number
of errors thrown related to plugin builds; I was unable to get it to work.

It would seem the known good Golang version is 1.8.x >= and < 1.9.

Setup go development environment as usual, making sure your $GOPATH is correctly
set to a desired location and you've checked out `hegemone/kore-poc` to
`$GOPATH/src/github.com/hegemone/kore-poc`. Dependencies are managed by glide
and checked in as `vendor/`, so there is no need for any dependency prep prior
to a build.

**IMPORTANT**: When building plugins, go will attempt to install a number of
utilities to your $GOROOT location. It must be writable by your user to allow
for those tools to be installed on a first run.

Build and run is managed by the `Makefile`:

* `make plugins` builds the example bacon plugin as a `.so` libs in `build/`
* `make adapters` builds the example adapters as `.so` libs in `build/`
* `make korecomm` builds the executable as `korecomm` in `build/`, and depends on
`plugins` and `adapters` targets.
* `make run` sets up extension load paths via env vars and runs the executable.
* `make clean` cleans the `build/` directory.
* `make image` will build a Docker image from source in your local registry.
* `make` by default will run `make build`, which is an alias for `korecomm`.

To run from a source build, simply `make run`. It will execute dependency targets
and start the executable.

## Comments

### Points of note
* Pervasive event based concurrent/parallel handling of Ingress, Egress, and
plugin execution. Primary buffered channes are found in the `Engine`
* Uses real golang dynamic plugins, loaded from `.so` libs at runtime.
* Enjoys a lot of the Golang benefits over Ruby implementation. Simple portable
binaries, and I'm willing to bet this is significantly more capable under load.
Would be interesting to setup some performance comparisons.

### Faux pieces
* The `mock` package fakes out API clients that you work normally choose off the
shelf for talking to platforms like Discord or IRC.
* The `StdinDemux` is a reimplementation of a single stdin listener that demuxes
messages to the correct client, making it appear to adapters that they're listening
to their own chat channels.
* `config` isn't entirely fake; I'm expecting to have a similar abstraction layer
in the full implementation, but we faked out reading config files or configuration
vars like secret credentials (API keys). Also no validation or error handling.

### Next Steps
* Overall project next steps outlined in the project issues. Most are higher
level than any one lang's POC.
* Again, this is not production ready at all. Needs a healthy amount of thought around
error handling and hardening. Espectially with extension loading and validation.
A real system needs to be able to validate plugins early and gracefully handle
bad extensions (missing implementations, incorrect signatures, deadlocked cmds).
* I'm not 100% on the extension pattern. My thoughts are in the comments, but
the container struct & loading process feels pretty inelegant and brittle.
