# Gilded Rose

This is a refactoring kata created by Terry Hughes (http://twitter.com/TerryHughes). This version in Go is based on
Emily Bache's [translation](https://github.com/emilybache/GildedRose-Refactoring-Kata/tree/master/go).

The [Gilded Rose Requirements](https://github.com/jamesjoshuahill/gildedrose/blob/master/requirements.txt) explain what
the code is for.

The [Golden Master](https://github.com/jamesjoshuahill/gildedrose/blob/master/golden_master.txt) records the correct
output for 30 days. You can confirm there is no diff by comparing the output or running the test provided:

```bash
$ diff <(go run main.go 30) golden_master.txt
```

## Get

```bash
$ go get github.com/jamesjoshuahill/gildedrose
```

## Test

The test suite uses [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) and can be run with `go test`, or using the `ginkgo` CLI:

```bash
$ cd $GOPATH/src/github.com/jamesjoshuahill/gildedrose
$ go get github.com/onsi/ginkgo/ginkgo
$ ginkgo
```
