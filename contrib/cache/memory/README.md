# Memory Cache
> cache data in memory implemented by [go-cache](https://github.com/patrickmn/go-cache)

## Usage
```go
    memCache := NewCache()
    err := memCache.SetString(context.TODO(), key, val, -1)
```

## Notice
`Save` and `Load` function just for test because `go-cache` `SaveFile` is deprecated.