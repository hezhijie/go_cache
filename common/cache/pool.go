package cache

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
)

var (
	// 定义常量
	RedisClient     *redis.Pool
	REDIS_HOST string
	REDIS_DB int
	MaxIdle int  	//最大空闲连接数,等候的连接
	MaxActive int	//最大激活连接,被占用的连接
	IdleTimeout time.Duration  //连接超时时间
)

func init(){

	REDIS_HOST = "localhost:6379"
	REDIS_DB = 0
	MaxIdle = 1
	MaxActive = 10
	IdleTimeout = 180 * time.Second

	RedisClient = &redis.Pool{
		MaxIdle:MaxIdle,
		MaxActive:MaxActive,
		IdleTimeout:IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

func ReadStr(cmd string, key string)  (string) {

	rc:=RedisClient.Get()

	v,err:=redis.String(rc.Do(cmd,key))
	if(err != nil){
		fmt.Println("read str error",err)
		return ""
	}

	return v
}

func Write(cmd string){

}



