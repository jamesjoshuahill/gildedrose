# Gilded Rose

[![Build Status](https://travis-ci.org/jamesjoshuahill/gildedrose.svg?branch=master)](https://travis-ci.org/jamesjoshuahill/gildedrose)

This is a refactoring kata created in C# by [Terry Hughes](http://twitter.com/TerryHughes).

This version is based on Emily Bache's
[translation to Go](https://github.com/emilybache/GildedRose-Refactoring-Kata/tree/master/go). It provides a test suite
and functioning [main.go](https://github.com/jamesjoshuahill/gildedrose/blob/master/cmd/gildedrose/main.go).

## Requirements

The [requirements](https://github.com/jamesjoshuahill/gildedrose/blob/master/REQUIREMENTS.md) explain what
the code is for.

The [golden master](https://github.com/jamesjoshuahill/gildedrose/blob/master/cmd/gildedrose/testdata/golden_master.txt)
records the correct output for 30 days.

You can check the program meets the requirements by comparing its output with the golden master:

```bash
$ diff <(go run cmd/gildedrose/main.go) cmd/gildedrose/testdata/golden_master.txt
```

## Get

```bash
$ go get github.com/jamesjoshuahill/gildedrose
$ cd $GOPATH/src/github.com/jamesjoshuahill/gildedrose
$ brew install dep
$ dep ensure
```

## Test

The test suite uses [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) and can be run
with `go test`, or using the `ginkgo` CLI:

```bash
$ cd $GOPATH/src/github.com/jamesjoshuahill/gildedrose
$ go get github.com/onsi/ginkgo/ginkgo
$ ginkgo -r -p
```

The integration test expects the program's output to match the
[golden master](https://github.com/jamesjoshuahill/gildedrose/blob/master/cmd/gildedrose/testdata/golden_master.txt).

The unit tests check that every item updates correctly.

For conjured items there are
[pending unit tests](https://github.com/jamesjoshuahill/gildedrose/blob/master/app_test.go#L95-L109) and a
[golden master with conjured item](https://github.com/jamesjoshuahill/gildedrose/blob/master/cmd/gildedrose/testdata/golden_master_with_conjured_item.txt)
ready to use in the integration test.

## Refactor

There are many ways to use this kata. Here are some suggestions:

1. Implement conjured items with the fewest changes.
1. Refactor the code so that it is easy to extend, then implement conjured items.
1. Delete all the tests, then write enough tests to allow you to implement conjured items without breaking the other items.
1. Delete all the code, then use the tests provided to drive out a new implementation.
1. New requirement: any item can be conjured and their quality changes twice as fast, e.g. Conjured Aged Brie.
