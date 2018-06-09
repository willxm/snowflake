package main

import (
	"fmt"

	"github.com/willxm/snowflake"
)

func main() {
	uuid, err := snowflake.NewSnowflack().NextID()
	fmt.Println(uuid, err)
}
