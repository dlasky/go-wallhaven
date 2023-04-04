[![GoDoc](https://godoc.org/github.com/dlasky/go-wallhaven?status.svg)](https://godoc.org/github.com/dlasky/go-wallhaven)

# go-wallhaven

golang wrapper for the wallhaven.cc api version 1

## Installation

`go get -u github.com/dlasky/go-wallhaven`

## Basic Usage

```golang
results, err := wallhaven.SearchWallpapers(nil)
if err != nil {
  log.Fatal(err)
}
log.Printf("%v", results)
```
