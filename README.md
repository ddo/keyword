# keyword [![Build Status][travis-img]][travis-url] [![Doc][godoc-img]][godoc-url]
> :mag: word matching in text 

[travis-img]: https://img.shields.io/travis/ddo/go-keyword.svg?style=flat-square
[travis-url]: https://travis-ci.org/ddo/go-keyword

[godoc-img]: https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat-square
[godoc-url]: https://godoc.org/github.com/ddo/go-keyword

##Example

```go
var wordChecker = keyword.New("love, happy", "hate, sad, die", false)

if wordChecker.Check("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore LOVE magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.") {
    println("positive sentence")
}
```