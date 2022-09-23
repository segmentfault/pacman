# Redis Cache
> cache data in redis implemented by [go-redis](https://github.com/go-redis/redis)

## Usage
```go
    redisCache := NewCache(addr, username, password)
    err := redisCache.SetString(context.TODO(), key, val, -1)
```
