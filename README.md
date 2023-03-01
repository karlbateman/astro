# Astro

> A simple echo server, which responds to [Maelstrom] messages.

## Introduction

This project is a node which is written in Go and compiled to a binary which
receives JSON messages from STDIN and sends JSON messages to STDOUT. See the full
Maelstrom [protocol specification] for more information.

[protocol specification]: https://github.com/jepsen-io/maelstrom/blob/main/doc/protocol.md

We're using a [Maelstrom Go Library] which provides `maelstrom.Node` that handles the aforementioned boilerplate for us.
Essentially, it'll let you register handler functions for each message type—similar to how [http.Handler] works in the
standard library.

[maelstrom go library]: https://pkg.go.dev/github.com/jepsen-io/maelstrom/demo/go
[http.handler]: https://pkg.go.dev/net/http#Handler

## Specification

Our node will receive an "echo" message from Maelstrom which will look like this:

```json
{
    "src": "c1",
    "dest": "n1",
    "body": {
        "type": "echo",
        "msg_id": 1,
        "echo": "Please echo 35"
    }
}
```

Nodes and clients are sequentially numbered (e.g. n1, n2, etc). Nodes are prefixed with "n" and external clients are
prefixed with "c". Message IDs are unique per source node but this is handled by the Go library.

Astro sends a message back to the client with the same body, but adds a message type of "echo_ok". It also associates
itself to the original message by setting the "in_reply_to" field to the original message ID. This reply field is
automatically handled by using the `Node.Reply()` method.

The response message looks like this:

```json
{
    "src": "n1",
    "dest": "c1",
    "body": {
        "type": "echo_ok",
        "msg_id": 1,
        "in_reply_to": 1,
        "echo": "Please echo 35"
    }
}
```

## Setup

Before you begin, you will need [Maelstrom] and [Go] installed on your system. This section will walk you through
setting up a Mac computer to run this project, you will need to adjust these commands if you're using Linux or Windows.

[go]: https://www.go.dev

### Installing Go

To install the Go programming language, it's standard library and associated toolchain head over to
<https://go.dev/dl>. Mac users can simply download the relevant `.pkg` file based on your arch Apple Silicon (ARM64) vs
Intel (x64). Once your download is complete, run the `.pkg` file and follow the instructions to complete the
installation process.

### Installing Maelstrom

Maelstrom is built using [Closure] so we'll need to install [OpenJDK]. It also provides some plotting and graphing
utilities which rely on [Graphviz] and [gnuplot]. Assuming you're using [Homebrew], these can be installed by running
the following:

[closure]: https://clojure.org/
[openjdk]: https://openjdk.org/
[graphviz]: https://graphviz.org/
[gnuplot]: http://www.gnuplot.info/
[homebrew]: https://brew.sh/

```bash
brew install openjdk graphviz gnuplot
```

> You can find more details on the [Prerequisites] section of the Maelstrom docs.

[prerequisites]: https://github.com/jepsen-io/maelstrom/blob/main/doc/01-getting-ready/index.md#prerequisites

Next, we'll need to download Maelstrom itself. Head over to the [releases page] on the Maelstrom Github page and
download the latest release and unpack it. You can run the maelstrom binary from inside the newly created directory.

[releases page]: https://github.com/jepsen-io/maelstrom/releases

## Running the node

Clone this project into your `$GOPATH` using the following command.

```bash
mkdir -p "${GOPATH}/src/github.com/karlbateman" \
    && git clone --branch develop \
        https://github.com/karlbateman/astro.git \
        "${GOPATH}/src/github.com/karlbateman/astro"
```

Change into the newly created directory and run the following to produce a binary which can be run by Maelstrom.

```bash
cd "${GOPATH}/src/github.com/karlbateman/astro"
    \ make build
```

Once this has completed, change into the directory where you unpacked Maelstrom previously and run the following
command to launch the node and run the tests.

```bash
./maelstrom test \
    --workload=echo \
    --bin="${GOPATH}/src/github.com/karlbateman/astro/bin/astro" \
    --node-count=1 \
    --time-limit=10
```

> If `maelstrom` is available on your path, you can run `make test` from within the Astro directory to run the node.

This command instructs `maelstrom` to run the "echo" workload against Astro. It runs a single node and it will send
"echo" commands for 10 seconds.

Maelstrom will only inject network failures and it will not intentionally crash Astro so we don't need to worry about
persistence.

If everything ran correctly, you should see a bunch of log messages, stats and then finally a pleasant message from
Maelstrom:

```text
Everything looks good! ヽ(‘ー`)ノ
```

## Debugging Maelstrom

If the test fails, you can run the Maelstrom web server to view the test results in more depth.

```bash
./maelstrom serve
```

You can then open a browser and visit <http://localhost:8080> to view the results. Consult the Maelstrom [docs] for
further details.

[docs]: https://github.com/jepsen-io/maelstrom/blob/main/doc/results.md

## Copyright

Copyright © 2023 Karl Bateman. All Rights Reserved.

[maelstrom]: https://github.com/jepsen-io/maelstrom
