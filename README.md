# Distributed uuid generation
Inspired by twitter snowflake

```
+--------------------------------------------------------------------------+
| 1 Bit Unused | 41 Bit Timestamp |  10 Bit MachineIDID  |   12 Bit Sequence ID |
+--------------------------------------------------------------------------+
```

## Install
```shell
$ go get github.com/willxm/snowflake
```

## Usage

```golang
uuid, err := snowflake.NewSnowflack().NextID()
```