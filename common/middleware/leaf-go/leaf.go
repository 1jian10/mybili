package leaf_go

import (
	"bilibili/common/middleware/leaf-go/segment"
	"bilibili/common/middleware/leaf-go/snowflake"
	"context"
	"errors"
	"time"
)

type Config struct {
	Model           int
	SegmentConfig   *SegmentConfig
	SnowflakeConfig *SnowflakeConfig
}

var (
	Segment   = 1
	Snowflake = 2
)

type Core interface {
	GetId() (int64, bool)
}

func Init(c *Config) (Core, error) {
	if c.Model == Segment {
		return segment.NewCreator(c.SegmentConfig)
	} else if c.Model == Snowflake {
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		return snowflake.NewCreator(timeout, c.SnowflakeConfig)
	}

	return nil, errors.New("please select id model")
}

type SnowflakeConfig struct {
	CreatorName string
	Addr        string
	EtcdAddr    []string
}

type SegmentConfig struct {
	Name string

	UserName string
	Password string
	Address  string
}
