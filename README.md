# hack.fm
A little community radio station written in Go


## how to build
The Go project in goproj should be built by setting the GOPATH to it and then building "metalab.at/hack.fm". It is important that the res and goproj (or any directory containing the binary) directories are in the same directory so the keepalive sound can be used. 
## Pi support
set GOARCH to "arm" and the Go toolchain should automatically compile to a Pi-compatible binary.

## Dependencies

You should have a Go toolchain installed on the building computer and youtube-dl and mpv on the computer running it.
