# Gilded Rose

This is a refactoring kata created by Terry Hughes (http://twitter.com/TerryHughes). This version in Go is based on
Emily Bache's [translation](https://github.com/emilybache/GildedRose-Refactoring-Kata/tree/master/go).

The [Gilded Rose Requirements](https://github.com/jamesjoshuahill/gildedrose/blob/master/requirements.txt) explain what
the code is for.

The [Golden Master](https://github.com/jamesjoshuahill/gildedrose/blob/master/golden_master.txt) records the correct
output for 30 days. You can confirm there is no diff by running:

```bash
$ diff <(go run gilded_rose.go 30) golden_master.txt
```
