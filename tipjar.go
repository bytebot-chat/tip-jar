package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

const ERR_NOT_ENOUGH_ARGS = "who needs their nose swatted?"
const ERR_COULD_NOT_TIP = "sorry, can't do that right now"

// Pattern borrowed from party pack
func suserSaidSomethingProblematic(ctx context.Context, m Message, rdb *redis.Client) (string, bool) {
	if !(strings.HasPrefix(m.Content, "!suser")) {
		return "", false
	}

	res, err := rdb.Incr(ctx, "problematic/suser").Result()

	if err != nil {
		return ERR_COULD_NOT_TIP, true
	}

	return fmt.Sprintf("suser has had his nose swatted %d times lately", res), true
}
