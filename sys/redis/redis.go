package redis
//
//import (
//	"time"
//	"github.com/garyburd/redigo/redis"
//	"git.oschina.net/YPTArch/processcache/sys/logs"
//	"fmt"
//)
//
//var (
//	pool *RedisPool
//)
//
//func SetPool(netType, server, password string, dbIndex, maxIdle, maxActive, idleTimeout int)  {
//	pool = &RedisPool{newPool(netType, server, password, dbIndex,  maxIdle, maxActive, idleTimeout )}
//}
//
//
//type RedisPool struct {
//	*redis.Pool
//}
//
////redigo连接池实例化
//func newPool(netType, server, password string, dbIndex, maxIdle, maxActive, idleTimeout int) *redis.Pool  {
//	return &redis.Pool{
//		MaxIdle: maxIdle,  //空闲连接的最大数目
//		MaxActive:maxActive, //给定时间内最大连接数，为0则连接数没有限制
//		IdleTimeout:time.Duration(idleTimeout) * time.Second, //过期时间，为0则不关闭，超时时间应小服务器的超时时间
//		Dial: func() (redis.Conn, error) {
//			//连接
//			c, err := redis.Dial(netType, server)
//			if err != nil {
//				logs.Logger.Error("redis connect err:" + err.Error())
//				return nil, err
//			}
//			//auth验证
//			if _, err := c.Do("AUTH", password); err != nil {
//				c.Close()
//				logs.Logger.Error(err.Error())
//				return nil, err
//			}
//			//选择数据库
//			if _, err := c.Do("SELECT", dbIndex); err != nil {
//				c.Close()
//				logs.Logger.Error(err.Error())
//				return nil, err
//			}
//			logs.Logger.Debug("new redis pool success!")
//			return c, err
//		},
//		//连接池检查功能
//		//连接被再次使用之前的空闲连接的健康情况
//		// t 返回连接的时间，大于1分钟就发送ping检查
//		TestOnBorrow: func(c redis.Conn, t time.Time) error {
//			if time.Since(t) < time.Minute {
//				return nil
//			}
//			_, err := c.Do("PING")
//			return err
//		},
//	}
//}
//
////获取redis连接
//func GetRedis() redis.Conn {
//	return pool.Get()
//}
//
//
//func Send(cmd string, args ...interface{}) error {
//	conn := GetRedis()
//	defer conn.Close()
//	err := conn.Send(cmd, args...)
//	logs.Logger.Debug(fmt.Sprintf("Send %v %v %v", cmd, args, err))
//	if err != nil {
//		return err
//	}
//	return conn.Flush()
//}
//
//func Do(cmd string, args ...interface{}) (interface{}, error) {
//	conn := GetRedis()
//	defer conn.Close()
//	logs.Logger.Debug(fmt.Sprintf("Do %v %v", cmd, args))
//	return conn.Do(cmd, args...)
//}
//
////删除指定key
//func Del(key string) (int, error)  {
//	return redis.Int(Do("DEL", key))
//}
//
////删除指定key
//func Exist(key string) (int, error)  {
//	return redis.Int(Do("EXISTS", key))
//}
//
////转换成struct
//func ScanStruct(src []interface{}, dest interface{}) error {
//	return redis.ScanStruct(src, dest)
//}