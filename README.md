

# bredit


bredit stands for: **B**encode **R**eplacer and **EDIT**or  - which is written in the Golang language.

It's released under GPLv2 license and therefor is free software.

bredit is used to edit and replace data that utilizes the bencode format such as torrent & rtorrent files.
It's designed to edit files in bulk but you can manually edit a file if needed.

Currently bredit only supports CLI but is planned to support editing via GUI.


# Installation

## From Binary
There is no installation required since the software is compiled with static.

## Compile source
⚠Some if not all of this might change in the future with a simple Makefile.⚠

Make sure to install [zeebo's encode library](https://github.com/zeebo/bencode):
`go get github.com/zeebo/bencode`

If you're going to compiling statically, make sure to install libgo-static and glibc-static is installed (name may vary)

To compile it yourself you'll need to make sure that at least `$GOOS` and `$GOARCH` is set:
- `$GOOS`
	* Underlying OS (Linux, Mac OS X, Windows, BSD, ...)
- `$GOARCH`
	* Which computer architecture you're using (arm, x64, i386)
- `$GOPATH` - Not needed but recommended
	* The path to the go project

---
you can either use gc (golang default) or gccgo. 

With gc it's simply:
If you have set `$GOPATH `:

`go build bredit` 

Without `$GOPATH`:

`cd /bredit/; go build *`

With gccgo it's (static):

` go build -compiler gccgo -gccgoflags '-static -O3' main.go fileIO.go bencode.go`


# Usage
Please refer to the [Wiki](https://github.com/piperun/bredit/wiki/)
      
---
There are some limitations currently when using bredit in CLI mode.
