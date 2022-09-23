<center>
  <img src="docs/logo.png" alt="drawing" width="120" align="center" />
</center>

# Pacman

Pacman 是一个 Golang 库，它抽象了 Golang 应用程序开发中常用的库和模式。
它提供了一组标准的接口和实现，可以用来快速构建 Golang 应用程序。

<!-- TOC depthfrom:2 orderedlist:false -->

- [更新记录](#%E6%9B%B4%E6%96%B0%E8%AE%B0%E5%BD%95)
- [设计原则](#%E8%AE%BE%E8%AE%A1%E5%8E%9F%E5%88%99)
- [如何使用](#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8)
- [参考资源](#%E5%8F%82%E8%80%83%E8%B5%84%E6%BA%90)

<!-- /TOC -->

## 为什么叫 Pacman?

我们从两个方面得到灵感：

1. [Pacman](https://wiki.archlinux.org/title/Pacman) 是 Arch Linux 的包管理器，它是一个非常流行的 Linux 发行版。
我们希望能够像 Arch Linux 中的 pacman 一样，让第三方包在 Go 项目中使用起来更加简单。
2. Pacman 也是一个 1980 年代非常流行的视频游戏角色。
我们希望能够像 Pacman 吃掉迷宫中所有的点一样，让 Golang 应用的开发更加简单。

## 更新记录

`v0.0.5`

基本的文档编写，以及部分代码格式的更改

`v0.0.4(20220829)`

- 入口文件处理,将每个模块各自进行初始化后交由 pacman.App 统一引用

`v0.0.3(20220823)`

- 配置文件获取使用 `viper` 包 ,方便后期扩展
- 更新脚手架初始化依赖使用配置文件调用方法
- 移除不必要的开发文件

## 设计原则

Golang 在后端开发已经相对的成熟，但是在开发过程中往往会困扰我们的：

1. 我们需要选择合适的类库或者解决方案，或者我们想同时尝试多种的方案，但不想过于依赖；
2. 第三方依赖的升级和替换使用会非常的麻烦，而且有可能会失败（例如第三方类库不支持大版本更新等）；
3. 不同的微服务其实从复用角度上说有很多通用的场景，这些不必须重复去开发。

种种上述的问题，这需要我们去考虑如何引用、组织以及管理第三方的引用的依赖。其实 Java 已经提供了非常好的思路，例如 Spring Framework 中的 Bean 等等，这些都是 IoC 思想的一种实现。

我们的想法是可能不需要实现那么「笨重」或者「臃肿」，但是能够从 Golang 的角度和思维上来说解决这些问题。因此，就有了这个框架以及对应的代码组织形式。

我们的代码组织形式首先是使用 interface 约定，然后使用 contrib 去实现，从依赖的角度上来说，调用者只只需要关心 interface 的接口以及返回即可，不需要去关心如何实现。因此，可以完整的解决上面提出的两点问题。

## 如何使用

根据设计原则，我们可以调用者角度上来使用这个框架类库，例如我们使用 Cache 组件使用的方式：

首先，我们引入 cache 的 interface，它定义了我们 Cache 接口的形式

```go
import "github.com/segmentfault/pacman/cache/v2"
```

然后，我们引入 contrib 包下面的具体实现（如果有必要，您也可以使用自己的方案实现，参见下章节）

```go
import "github.com/segmentfault/pacman/contrib/cache/v2"
```

然后就是具体的使用，例如我们定义了个非常简单的字符串形式的 Cache

```go
cache := v2.NewCache[int, string]()
cache.Set(1, "hello")

got, ok := cache.Get(1)
assert.Equal(t, ok, true)
assert.Equal(t, got, "hello")
```

对，目前使用方面就是那么的简单！如果想查看详细的细节，可以参考各个模块的具体 README 说明。

## 参考资源

- <https://golangexample.com/simple-and-yet-powerful-dependency-injection-for-golang/>
- <https://zhuanlan.zhihu.com/p/141204279>
