ji-2014-go
==========

`ji-2014-go` is a simple repository holding sources for the `Go`
hands-on session of the [JI-2014](http://ji.in2p3.fr)

## Bootstrapping the work environment

### Installing the `Go` toolchain

The `Go` hands-on session obviously needs for you to install the `Go`
toolchain.

There are 3 ways to do so:
- install `Go` via your favorite package manager (`yum`, `apt-get`, ...)
- install `Go` via `docker`
- install `Go` manually.

The first way is distribution/OS dependent. Please refer to your
favorite forum.
We'll only describe the 2 other ways.

#### Installing `Go` manually

This is best explained on the official page:
http://golang.org/doc/install

On linux-64b, it would perhaps look like:

```sh
$ mkdir /somewhere
$ cd /somewhere
$ curl -O -L https://storage.googleapis.com/golang/go1.3.2.linux-amd64.tar.gz
$ tar zxf go1.3.2.linux-amd64.tar.gz
$ export GOROOT=/somewhere/go
$ export PATH=$GOROOT/bin:$PATH
$ which go
/somewhere/go/bin/go
```

#### Installing `Go` via `docker`

If you have `docker` installed, installing `Go` is as easy as:

```sh
$ docker pull golang:latest
$ docker images | grep golang
golang                 latest              1744b082eb8f        2 days ago          373.4 MB
```

You can now test it like so:

```sh
$ docker run -i -t golang /bin/bash
$ go env
GOARCH="amd64"
GOBIN=""
GOCHAR="6"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/go"
GORACE=""
GOROOT="/usr/src/go"
GOTOOLDIR="/usr/src/go/pkg/tool/linux_amd64"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1"
```

### Setting up the work environment

Like `python` and its `$PYTHONPATH` environment variable, `Go` uses
`$GOPATH` to locate packages' source trees.
You can choose whatever you like (obviously a directory under which
you have read/write access, though.)
In the following, we'll assume you chose `$HOME/ji-go-work`:

```sh
$ mkdir -p $HOME/ji-go-work
$ export GOPATH=$HOME/ji-go-work
$ export PATH=$GOPATH/bin:$PATH
```

Make sure the `go` tool is correctly setup:

```sh
$ go env
GOARCH="amd64"
GOBIN=""
GOCHAR="6"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="$HOME/ji-go-work"
GORACE=""
GOROOT="/somewhere/go"
GOTOOLDIR="/somewhere/go/pkg/tool/linux_amd64"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1"
```

(on other platforms/architectures, the output might slightly
differ. The important env.vars. are `GOPATH` and `GOROOT`.)

### Testing `go get`

Now that the `go` tool is correctly setup, let's try to fetch some
code (you'll need `git` to fetch the code):

```sh
$ go get github.com/sbinet/ji-2014-go/cmd/ji-hello
```
