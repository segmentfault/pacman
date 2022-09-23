# Pacman - Yet Another Toolkit to Build Golang Application Quickly

![logo](./docs/pacman-logo.svg)

[![LICENSE](https://img.shields.io/badge/License-MIT-green)](https://github.com/segmentfault/pacmam/blob/master/LICENSE)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)

<!-- TOC depthfrom:2 orderedlist:false -->

- [Why call it Pacman?](#why-call-it-pacman)
- [Changelog](#changelog)
- [Design Philosophy](#design-philosophy)
- [Requirements and install](#requirements-and-install)
- [How to Use It](#how-to-use-it)
- [Contributing](#contributing)
- [License](#license)

<!-- /TOC -->

Notice: This document also has [Simplified Chinese 中文版本](README_CN.md) Edition.

## Why call it Pacman?

We inspired by two things:

1. [Pacman](https://wiki.archlinux.org/title/Pacman) is the package manager of Arch Linux, which is a very popular Linux distribution.
We want to make it easy to use third-party packages in Go projects, just like what pacman does in Arch Linux.
2. Pacman is also a video game character, which is very popular in the 1980s.
We want to make it easy to build a Golang application quickly, just like Pacman eats all the dots in the maze.

## Changelog

`20220919` initial release

## Design Philosophy

Golang is very stable and robust for server-side application development. But some general issues are confusing and blocking the progress of development, such as:

1. We need many third-party libraries to build the application. We need to choose a suitable one or multiple solutions at the same time, but do not want to be too dependent on them;
2. The upgrade and replacement of third-party dependencies will cause failures, eg. some third-party libraries do not support major version updates.
3. We're building microservices that may require common features, so we don't need to develop them repeatedly.

Based on above the issues, we began to think about how to manage third-party libraries. Far as we are aware of the Spring Boot which famous framework from Java world managing the third-party libraries or dependencies by Bean. This is the appropriate implementation for Ioc - The injection of dependencies.

## Requirements and install

This project requires golang v1.18 or higher version with new features such as general generic and so on.

## How to Use It

To use this toolkit is very simple and easy. For example, we wanna including the general cache module which has implemented:

First, we import the interface from the definition of the cache module.

```go
import "github.com/segmentfault/pacman/cache/v2"
```

Second, we import the implementation from the other path.

```go
import "github.com/segmentfault/pacman/contrib/cache/v2"
```

Finally, we use it!

```go
cache := v2.NewCache[int, string]()
cache.Set(1, "hello")

got, ok := cache.Get(1)
assert.Equal(t, ok, true)
assert.Equal(t, got, "hello")
```

The cache module should be running correctly.

## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.

## License

[MIT](https://github.com/segmentfault/pacmam/blob/master/LICENSE)
