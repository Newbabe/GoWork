package util

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 声明一些全局变量
var (
	RedisPool     *redis.Pool
	redisServer   = "r-2zehbb71li2d1wvunx.redis.rds.aliyuncs.com:6379"
	redisPassword = "aYI!Ec_dNH@Frh%1N*4uivHc$jBex"
)

func init() {
	RedisPool = newPool(redisServer, redisPassword)
}

// 初始化一个pool
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     20,
		MaxActive:   150,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func AddTest(id int, name string, source string) {

	conn := RedisPool.Get()

	defer conn.Close()
	saveRedisReqLog("AddTest-SET-" + source)

	//redis操作
	v, err := conn.Do("SET", id, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(conn.Do("GET", id))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

}

func GetRedisConn() redis.Conn {
	conn := RedisPool.Get()
	return conn
}

func Zadd(key string, expriedTime int, score int, value interface{}, source string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	conn := RedisPool.Get()
	defer conn.Close()
	saveRedisReqLog("Zadd-ZADD-" + source)
	conn.Do("ZADD", key, score, value)
	if expriedTime > 0 {
		saveRedisReqLog("Zadd-EXPIRE-" + source)
		conn.Do("EXPIRE", key, expriedTime)
	}

}

type Pipeline struct {
	conn redis.Conn
}

func GetPipeline() Pipeline {
	tStart := getTimeMS()

	p := Pipeline{}

	t2 := getTimeMS()
	p.conn = RedisPool.Get()
	t3 := getTimeMS()

	t3_2 := t3 - t2
	t3_2Str := strconv.FormatInt(t3_2, 10)

	p.conn.Send("MULTI")
	tEnd := getTimeMS()

	tEnd_t3 := tEnd - t3
	tEnd_t3Str := strconv.FormatInt(tEnd_t3, 10)

	tAll := tEnd - tStart
	tAllStr := strconv.FormatInt(tAll, 10)

	saveRedisReqLog("GetPipeline-MULTI(conn:" + t3_2Str + ",Do:" + tEnd_t3Str + ",allTime:" + tAllStr + ")")
	return p
}

func (p *Pipeline) Send(commend string, args ...interface{}) {
	//defer func() { // 必须要先声明defer，否则不能捕获到panic异常
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//		debug.PrintStack()
	//	}
	//}()
	//defer p.conn.Close()
	tStart := getTimeMS()
	p.conn.Send(commend, args...)
	tEnd := getTimeMS()

	tAll := tEnd - tStart
	tAllStr := strconv.FormatInt(tAll, 10)

	saveRedisReqLog("PipelineSend-Send(conn:0,Do:" + tAllStr + ",allTime:" + tAllStr + ")")

}

func (p *Pipeline) Execute() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	defer p.conn.Close()
	tStart := getTimeMS()
	p.conn.Do("EXEC")
	tEnd := getTimeMS()

	tAll := tEnd - tStart
	tAllStr := strconv.FormatInt(tAll, 10)
	saveRedisReqLog("Pipeline-EXEC(conn:0,Do:" + tAllStr + ",allTime:" + tAllStr + ")")
	//fmt.Println("err:",err,"reply：",reply)
}

func ZrevrangeByScore(key string, lastId string, pageSize int) []interface{} {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	conn := RedisPool.Get()
	defer conn.Close()
	values, err := redis.Values(conn.Do("ZREVRANGEBYSCORE", key, lastId, "-inf", "limit", 1, pageSize))
	if err != nil {
		fmt.Println(" err", err.Error())
	}
	return values
}

func Zrevrange(key string, start int, end int, source string) []interface{} {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	tStart := getTimeMS()

	conn := RedisPool.Get()
	defer conn.Close()

	t2 := getTimeMS()

	t2_tStart := t2 - tStart
	t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	values, err := redis.Values(conn.Do("ZREVRANGE", key, start, end))
	t3 := getTimeMS()
	t3_t2 := t3 - t2
	t3_t2Str := strconv.FormatInt(t3_t2, 10)

	tAll := t3 - tStart
	tAllStr := strconv.FormatInt(tAll, 10)

	saveRedisReqLog("Zrevrange-ZREVRANGE-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr + ")")

	if err != nil {
		fmt.Println(" err", err.Error())
	}
	return values
}

func Zscore(key string, value interface{}, source string) int {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	tStart := getTimeMS()

	conn := RedisPool.Get()
	defer conn.Close()

	t2 := getTimeMS()
	t2_tStart := t2 - tStart
	t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	data, _ := redis.Int(conn.Do("ZSCORE", key, value))

	t3 := getTimeMS()
	t3_t2 := t3 - t2
	t3_t2Str := strconv.FormatInt(t3_t2, 10)

	tAll := t3 - tStart
	tAllStr := strconv.FormatInt(tAll, 10)

	saveRedisReqLog("Zscore-ZSCORE-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr + ")")

	return data
}

func ExistRedisKey(key string, source string) bool {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	tStart := getTimeMS()

	conn := RedisPool.Get()
	defer conn.Close()

	t2 := getTimeMS()
	t2_tStart := t2 - tStart
	t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	existKey, _ := redis.Bool(conn.Do("EXISTS", key))
	t3 := getTimeMS()
	t3_t2 := t3 - t2
	t3_t2Str := strconv.FormatInt(t3_t2, 10)

	tAll := t3 - tStart
	tAllStr := strconv.FormatInt(tAll, 10)

	saveRedisReqLog("ExistRedisKey-EXISTS-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr + ")")

	return existKey
}

func Zcard(key string, source string) int {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	tStart := getTimeMS()

	conn := RedisPool.Get()
	defer conn.Close()

	t2 := getTimeMS()
	t2_tStart := t2 - tStart
	t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	data, _ := redis.Int(conn.Do("ZCARD", key))

	t3 := getTimeMS()
	t3_t2 := t3 - t2
	t3_t2Str := strconv.FormatInt(t3_t2, 10)

	tAll := t3 - tStart
	tAllStr := strconv.FormatInt(tAll, 10)

	saveRedisReqLog("Zcard-ZCARD-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr + ")")

	return data
}

func DelRedisKey(key string, source string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	tStart := getTimeMS()

	conn := RedisPool.Get()
	defer conn.Close()

	t2 := getTimeMS()
	t2_tStart := t2 - tStart
	t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	redis.Bytes(conn.Do("DEL", key))

	t3 := getTimeMS()
	t3_t2 := t3 - t2
	t3_t2Str := strconv.FormatInt(t3_t2, 10)

	tAll := t3 - tStart
	tAllStr := strconv.FormatInt(tAll, 10)

	saveRedisReqLog("DelRedisKey-DEL-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr + ")")
}

func SetStringToRedis(key string, expriedTime int, args string, source string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//tStart := getTimeMS()

	conn := RedisPool.Get()
	defer conn.Close()

	//t2 := getTimeMS()
	//t2_tStart := t2 - tStart
	//t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	conn.Do("set", key, args)
	//t3 := getTimeMS()
	//t3_t2 := t3 - t2
	//t3_t2Str := strconv.FormatInt(t3_t2, 10)

	//tAll1 := t3 - tStart
	//tAllStr1 := strconv.FormatInt(tAll1, 10)

	//saveRedisReqLog("SetStringToRedis-set-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr1 + ")")
	if expriedTime > 0 {

		//t22 := getTimeMS()
		conn.Do("EXPIRE", key, expriedTime)
		//t33 := getTimeMS()
		//t33_t22 := t33 - t22
		//t33_t22Str := strconv.FormatInt(t33_t22, 10)

		//tAll2 := t33_t22 + t2_tStart
		//tAllStr2 := strconv.FormatInt(tAll2, 10)

		//saveRedisReqLog("SetStringToRedis-EXPIRE-" + source + "(conn:" + t2_tStartStr + ",Do:" + t33_t22Str + ",allTime:" + tAllStr2 + ")")
	}
}

func SetStringToRedisFiled(key string, field string, expriedTime int, args string, source string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	tStart := getTimeMS()

	conn := RedisPool.Get()
	defer conn.Close()

	t2 := getTimeMS()
	t2_tStart := t2 - tStart
	t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	conn.Do("hset", key, field, args)

	t3 := getTimeMS()
	t3_t2 := t3 - t2
	t3_t2Str := strconv.FormatInt(t3_t2, 10)

	tAll1 := t3 - tStart
	tAllStr1 := strconv.FormatInt(tAll1, 10)

	saveRedisReqLog("SetStringToRedisFiled-hset-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr1 + ")")

	if expriedTime > 0 {

		t22 := getTimeMS()
		conn.Do("EXPIRE", key, field, expriedTime)
		t33 := getTimeMS()
		t33_t22 := t33 - t22
		t33_t22Str := strconv.FormatInt(t33_t22, 10)

		tAll2 := t33_t22 + t2_tStart
		tAllStr2 := strconv.FormatInt(tAll2, 10)

		saveRedisReqLog("SetStringToRedisFiled-EXPIRE-" + source + "(conn:" + t2_tStartStr + ",Do:" + t33_t22Str + ",allTime:" + tAllStr2 + ")")

	}
}

func GetStringFromRedis(key string, source string) string {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	tStart := getTimeMS()

	conn := RedisPool.Get()
	defer conn.Close()

	t2 := getTimeMS()
	t2_tStart := t2 - tStart
	t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	//value, err := redis.String(conn.Do("GET", key))
	//content,err := conn.Do("get", key)

	res, err := conn.Do("GET", key)
	t3 := getTimeMS()
	t3_t2 := t3 - t2
	t3_t2Str := strconv.FormatInt(t3_t2, 10)

	tAll1 := t3 - tStart
	tAllStr1 := strconv.FormatInt(tAll1, 10)

	saveRedisReqLog("GetStringFromRedis-GET-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr1 + ")")

	if res != nil {
		value, err := redis.String(res, err)

		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
		return value
	}
	return ""

}

func GetStringFromRedisFiled(key string, filed string, source string) string {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	tStart := getTimeMS()
	conn := RedisPool.Get()
	defer conn.Close()

	t2 := getTimeMS()
	t2_tStart := t2 - tStart
	t2_tStartStr := strconv.FormatInt(t2_tStart, 10)

	//value, err := redis.String(conn.Do("GET", key))
	//content,err := conn.Do("get", key)
	res, err := conn.Do("HGET", key, filed)

	t3 := getTimeMS()
	t3_t2 := t3 - t2
	t3_t2Str := strconv.FormatInt(t3_t2, 10)

	tAll1 := t3 - tStart
	tAllStr1 := strconv.FormatInt(tAll1, 10)

	saveRedisReqLog("GetStringFromRedisFiled-HGET-" + source + "(conn:" + t2_tStartStr + ",Do:" + t3_t2Str + ",allTime:" + tAllStr1 + ")")

	if res != nil {
		value, err := redis.String(res, err)

		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
		return value
	}
	return ""

}

func Lpush(key, value string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	conn := RedisPool.Get()
	defer conn.Close()

	conn.Do("lpush", key, value)

}

func getTimeMS() int64 {
	ms := time.Now().UnixNano() / 1e6
	return ms
}
func saveRedisReqLog(source string) {
	//调用这个函数需要在redis关闭之后
	//logPerday("redisReqLog2", source)
}
func logPerday(fileName string, log string) {
	dir := "/usr/local/golang_app/HeroOkWeb/static/log/" + fileName
	if !Exists(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}

	fileName = dir + "/" + fileName + "_" + time.Now().Format("20060102")
	//"/"+fileName+"/"+
	var file *os.File
	defer file.Close()
	var err error

	result := Exists(fileName)
	if result { //文件存在
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, os.ModePerm) //打开文件
	} else { //
		file, err = os.Create(fileName) //创建文件
	}
	check(err)
	//buf :=[]byte(log)
	//_, err1 :=file.Write(buf)
	log += "---" + time.Now().Format("2006-01-02 15:04:05")

	_, err1 := io.WriteString(file, log+"\n") //写入文件(字符串)
	check(err1)

}

// io操作容易出错的检查
func check(e error) {

}

// 判断文件是否存在
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true

}
func GetStringFromRedisByKey(key string) string {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	conn := RedisPool.Get()
	defer conn.Close()

	//value, err := redis.String(conn.Do("GET", key))
	//content,err := conn.Do("get", key)
	res, err := conn.Do("GET", key)

	if res != nil {
		value, err := redis.String(res, err)

		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
		return value
	}
	return ""

}
func GetIntsFromRedisByKey(key string) []string {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	conn := RedisPool.Get()
	defer conn.Close()
	//获取key集合中总元素个数

	// 使用 SRANDMEMBER 命令获取 set 中的随机元素
	value, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		fmt.Println(err)
	}
	return value
}

type userData struct {
	StageBgAnimId int    `json:"stageBgAnimId"`
	UserLevel     int    `json:"userLevel"`
	UserLiveLv    int    `json:"userLiveLv"`
	HeadDressUrl  string `json:"headDressUrl"`
	IconUrl       string `json:"iconUrl"`
}
