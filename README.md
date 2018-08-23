# Gilded Rose

[![Build Status](https://travis-ci.org/jamesjoshuahill/gildedrose.svg?branch=master)](https://travis-ci.org/jamesjoshuahill/gildedrose)

This is a refactoring kata created in C# by [Terry Hughes](http://twitter.com/TerryHughes). This version is based on
Emily Bache's [translation to Go](https://github.com/emilybache/GildedRose-Refactoring-Kata/tree/master/go). It provides unit tests
and a functioning binary in `cmd/gildedrose` with integration tests.

## Get

```bash
$ go get github.com/jamesjoshuahill/gildedrose
$ cd $GOPATH/src/github.com/jamesjoshuahill/gildedrose
$ brew install dep
$ dep ensure
```

## Test

The test suite uses [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) and can be run with `go test`, or using the `ginkgo` CLI:

```bash
$ cd $GOPATH/src/github.com/jamesjoshuahill/gildedrose
$ go get github.com/onsi/ginkgo/ginkgo
$ ginkgo -r -p -keepGoing
```

The integration test checks that the output of the program matches the [Golden Master](https://github.com/jamesjoshuahill/gildedrose/blob/master/cmd/gildedrose/testdata/golden_master.txt).

The unit tests check that every item is updated correctly.

There are pending tests for Conjured items.

## Refactor

The [Gilded Rose Requirements](https://github.com/jamesjoshuahill/gildedrose/blob/master/REQUIREMENTS.md) explain what
the code is for.

The [Golden Master](https://github.com/jamesjoshuahill/gildedrose/blob/master/cmd/gildedrose/testdata/golden_master.txt) records the correct
output for 30 days. You can confirm there is no diff by running the integration test provided, or comparing the output:

```bash
$ diff <(go run cmd/gildedrose/main.go) cmd/gildedrose/testdata/golden_master.txt
```

The goal of the refactoring is to make it easy to add support for Conjured items.

## Scenarios

There are many ways to approach this exercise. Here are some suggestions:

1. Refactor the code so that it is easy to extend, then implement Conjured items.
1. Implement Conjured items with the fewest changes.
1. Delete all the tests, then write enough tests to allow you to implement Conjured items without breaking the other items.
1. Delete all the code, then test drive a new implementation.
1. New requirement: any item can be conjured and their quality changes twice as fast, e.g. Conjured Aged Brie.
