# Distributed uuid generation
Inspired by twitter snowflake

```
+-------------------------------------------------------------------------------+
| 1 Bit Unused | 39 Bit Timestamp |  16 Bit MachineIDID  |   8 Bit Sequence ID |
+-------------------------------------------------------------------------------+
```

## Install
```shell
$ go get github.com/willxm/snowflake
```

## Usage

```golang
uuid, err := snowflake.NewSnowflack().UUID()
```