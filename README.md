A Neocities client written in Go
================================

Upload files to you [Neocities](https://neocities.org/) website from the
comfort of your own terminal.

[![GoDoc](https://godoc.org/github.com/peterhellberg/neocities?status.png)](https://godoc.org/github.com/peterhellberg/neocities)

## Installation

You can download precompiled binaries for
[OS X](https://github.com/peterhellberg/neocities/releases/download/0.0.1/neocities-0.0.1-darwin-amd64.zip),
[Linux](https://github.com/peterhellberg/neocities/releases/download/0.0.1/neocities-0.0.1-linux-amd64.zip),
[Windows (32 bit)](https://github.com/peterhellberg/neocities/releases/download/0.0.1/neocities-0.0.1-windows-386.zip) and
[Windows (64 bit)](https://github.com/peterhellberg/neocities/releases/download/0.0.1/neocities-0.0.1-windows-amd64.zip)

Or, if you have [Go](http://golang.org/) installed:

    go get -u github.com/peterhellberg/neocities

## Usage

[![Cat!](https://neocities.org/assets/img/cat.png)](https://neocities.org/)

First you need to set two environment variables:

    export NEOCITIES_USER=<user>
    export NEOCITIES_PASS=<pass>

Then you should be able to upload files to your website:

    neocities upload foo.html bar.js baz.jpg

You can also delete files:

    neocities delete foo.html

You get a list of available commands by default:

```bash
$ neocities
usage: neocities <command> [<args>]

Commands:
   upload    Upload files to Neocities
   delete    Delete files from Neocities
   info      Info about Neocities websites
   version   Show neocities client version

Help for a specific command:
   help [command]
```

## Donate

[![Support Neocities](https://neocities.org/assets/img/support-us.png)](https://neocities.org/donate)

NeoCities is funded by donations. If you’d like to contribute, you can help to pay for server costs using Bitcoin or PayPal.

## License (MIT)

Copyright (c) 2014 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
