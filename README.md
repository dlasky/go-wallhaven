# go-wallhaven
golang wrapper for the wallhaven.cc api

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
