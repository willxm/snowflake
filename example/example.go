package main

import (
	"fmt"

	"github.com/willxm/snowflake"
)

func main() {
	uuid, err := snowflake.NewSnowflack().UUID()
	fmt.Println(uuid, err)
}
