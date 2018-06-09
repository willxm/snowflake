# Distributed uuid generation
Inspired by twitter snowflake

## Install
```shell
$ go get github.com/willxm/snowflake
```

## Usage

```golang
uuid, err := snowflake.NewSnowflack().NextID()
```