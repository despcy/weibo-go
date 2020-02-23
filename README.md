# weibo-go:

This repo aims to provide apis for weibo that do not require login, IP may be blocked after several consecutive requests, but will back to normal after about 3 minutes.

## Installation&Setup:

For picture downloading and crawling:

```shell
go get -u github.com/yangchenxi/weibo-go/album
```

For user searching according to geo location:

```shell
go get -u github.com/yangchenxi/weibo-go/radar
```

## Documentation:

Album:

   https://godoc.org/github.com/yangchenxi/weibo-go/album
    
Radar:

   https://godoc.org/github.com/yangchenxi/weibo-go/radar

## Demo:

See test file on each package

## TODO:

Full Blog crawl package will be released soon
