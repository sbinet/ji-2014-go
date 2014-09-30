ji-2014-go
==========

`ji-2014-go` is a simple repository holding sources for the `Go`
hands-on session of the [JI-2014](http://ji.in2p3.fr)

## Bootstrapping the work environment

### Installing the `Go` toolchain

The `Go` hands-on session obviously needs for you to install the `Go`
toolchain.

There are 3 ways to do so:
- install `Go` via your favorite package manager (`yum`, `apt-get`,
  `fink`, ...)
- install `Go` via `docker`
- install `Go` manually.

While all 3 methods are valid ones, to reduce the complexity of the
debugging/configuration/OS matrix, we'll only recommend the last one.

**WARNING:** Do note that installing `Go` via `homebrew` on `MacOSX`
is known to have some deficiencies. On `Macs`, it is really
recommended you use the manual method.

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

$ which godoc
/somewhere/go/bin/godoc
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

(on other platforms/architectures, the output might differ
slightly. The important env.vars. are `GOPATH` and `GOROOT`.)

### Testing `go get`

Now that the `go` tool is correctly setup, let's try to fetch some
code (you'll need `git` to fetch the code):

```sh
$ go get github.com/sbinet/ji-2014-go/cmd/ji-hello
```

`go get` downloaded (cloned, in `git` speak) the whole
`github.com/sbinet/ji-2014-go` repository (under `$GOPATH/src`) and
compiled the `ji-hello` command.
As the compilation was successful, it also installed the `ji-hello`
command under `$GOPATH/bin`.

The `ji-hello` command is now available from your shell:

```sh
$ ji-hello
Hello JI-2014!

$ ji-hello you
Hello you!
```

## Documentation

The `Go` programming language is quite new (released in 2009) but
ships already with quite a fair amount of documentation.
Here are a few pointers:

- http://golang.org/doc/code.html
- http://tour.golang.org
- http://golang.org/doc/effective_go.html
- http://dave.cheney.net/resources-for-new-go-programmers
- http://gobyexample.com

For more advanced topics:

- http://talks.golang.org
- http://blog.golang.org
- https://groups.google.com/d/forum/golang-nuts (`Go` community forum)
- the internets. typical queries are `go something` or `golang something`
