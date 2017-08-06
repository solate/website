package redis

import (
	"encoding/json"
	"strings"
	"github.com/garyburd/redigo/redis"
)

type Hash struct {
	Name string
}

func NewHash(name string) *Hash {
	Do("PING")
	return &Hash{name}
}

func (this *Hash) SetExpire(second int) error {
	return Send("EXPIRE", this.Name, second)
}

func (this *Hash) PutObject(k string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return Send("HSET", this.Name, k, b)
}

func (this *Hash) Put(k, v string) error {
	return Send("HSET", this.Name, k, []byte(v))
}

func (this *Hash) GetObject(k string, clazz interface{}) error {
	b, err := redis.Bytes(Do("HGET", this.Name, k))
	if err != nil {
		return err
	}
	return json.Unmarshal(b, clazz)
}

func (this *Hash) Get(k string) (string, error) {
	b, err := redis.String(Do("HGET", this.Name, k))
	if err != nil {
		return "", err
	}
	return b, nil
}

//func (orm *Hash) ScanPK(output interface{}) *Model {
//  if reflect.TypeOf(reflect.Indirect(reflect.ValueOf(output)).Interface()).Kind() == reflect.Slice {
//      sliceValue := reflect.Indirect(reflect.ValueOf(output))
//      sliceElementType := sliceValue.Type().Elem()
//      for i := 0; i < sliceElementType.NumField(); i++ {
//          bb := reflect.ValueOf(sliceElementType.Field(i).Tag)
//          if bb.String() == "PK" {
//              orm.PrimaryKey = sliceElementType.Field(i).Name
//          }
//      }
//  } else {
//      tt := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(output)).Interface())
//      for i := 0; i < tt.NumField(); i++ {
//          bb := reflect.ValueOf(tt.Field(i).Tag)
//          if bb.String() == "PK" {
//              orm.PrimaryKey = tt.Field(i).Name
//          }
//      }
//  }
//  return orm

//}

//func (this *Hash) GetObjectList(k []string, objs []interface{}) error {
//  args := []interface{}{}
//  args = append(args, this.Name)
//  for _, v := range k {
//      args = append(args, v)
//  }
//  b, err := redis.MultiBulk(Do("HMGET", args...))
//  if err != nil {
//      return err
//  }
//  for i, v := range b {
//      bb, err := redis.Bytes(v, nil)
//      if err != nil {
//          break
//      }
//      err = json.Unmarshal(bb, objs[i])
//      if err != nil {
//          break
//      }
//  }
//  return err
//}

func (this *Hash) PutString(k string, v string) error {
	return Send("HSET", this.Name, k, v)
}

func (this *Hash) GetString(k string) (string, error) {
	str, err := redis.String(Do("HGET", this.Name, k))
	if err == nil {
		str = strings.Trim(str, "\"")
	}
	return str, err
}

func (this *Hash) GetStringList(k []string) ([]string, error) {
	args := []interface{}{}
	args = append(args, this.Name)
	for _, v := range k {
		args = append(args, v)
	}
	reply, err := redis.MultiBulk(Do("HMGET", args...))
	if err != nil {
		return nil, err
	}
	var list = make([]string, 0)
	for _, v := range reply {
		s, err := redis.String(v, nil)
		if err != nil {
			break
		}
		s = strings.Trim(s, "\"")
		list = append(list, s)
	}
	return list, err
}

func (this *Hash) MultiGet(k []string) ([]string, error) {
	args := []interface{}{}
	args = append(args, this.Name)
	for _, v := range k {
		args = append(args, v)
	}
	reply, err := redis.MultiBulk(Do("HMGET", args...))
	if err != nil {
		return nil, err
	}
	var list = make([]string, 0)
	for _, v := range reply {
		b, err := redis.Bytes(v, nil)
		if err != nil {
			break
		}
		list = append(list, string(b))
	}
	return list, err
}

func (this *Hash) Size() (int, error) {
	return redis.Int(Do("HLEN", this.Name))
}

func (this *Hash) Remove(k string) error {
	return Send("HDEL", this.Name, k)
}

func (this *Hash) Exists(k string) bool {
	v, err := redis.Bool(Do("HEXISTS", this.Name, k))
	if err != nil {
		return false
	}
	return v
}

func (this *Hash) Clear() error {
	return Send("DEL", this.Name)
}

//根据key获取所有
func (this *Hash) HgetAll() ([]interface{}, error)  {
	return redis.Values(Do("hgetall", this.Name))
}