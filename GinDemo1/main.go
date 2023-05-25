package main

import (
	"GinDemo1/conf"
	"GinDemo1/util"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"runtime/debug"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func mainbbb() {
	/*r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")*/

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello WORLD", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
func isInt(str string) bool {
	for _, s := range str {
		digit := unicode.IsDigit(s)
		if !digit { //有一个不是10进制就直接返回false
			return false
		}
	}
	return true
}

func main() {
	isInt("Kiki")
}
func mainccc() {
	dataList := GetHomePageRecommendSongList(1, 100)
	marshal, err := json.Marshal(dataList)
	if err != nil {
		fmt.Println(err)

	}
	//存值
	util.SetStringToRedis("HomePageRecommendSongData", 1800, string(marshal))
	//util.GetStringFromRedis()
}
func GetHomePageRecommendSongList(page, pageSize int) []Record {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	sql := "select  y.id, y.song_id, y.song_name, y.singer, y.album, y.user_id, y.nick_name, " +
		"  y.score, y.update_time, y.time_length, y.status, y.channelId, y.vote_count, " +
		"  y.vote_total, y.type, y.chorus_id, y.weight, y.chorus_weight, y.lrc_url, y.lrc_type, " +
		"  y.backup, y.first_total, y.org_chorus_id, y.del_date, y.is_mv, y.time_difference" +
		" from hero_ok_web.exam_record_top x right join hero_ok_web.record y on x.record_id=y.id " +
		"  where  x.level in(4,5) and x.status=1 and y.status=1  order by x.record_time desc limit ?,? "
	query, err := conf.GetDbHerookWebRead().Query(sql, (page-1)*pageSize, pageSize)

	if err != nil {
		return nil
	}
	defer query.Close()
	var recordList []Record
	for query.Next() {
		var recordNew Record
		query.Scan(&recordNew.Id, &recordNew.SongId, &recordNew.SongName, &recordNew.Singer, &recordNew.Album, &recordNew.UserId, &recordNew.NickName,
			&recordNew.Score, &recordNew.UpdateTime, &recordNew.TimeLength, &recordNew.Status, &recordNew.ChannelId, &recordNew.VoteCount,
			&recordNew.VoteTotal, &recordNew.TypeInt, &recordNew.ChorusId, &recordNew.Weight, &recordNew.ChorusWeight, &recordNew.LrcUrl, &recordNew.LrcType,
			&recordNew.Backup, &recordNew.FirstTotal, &recordNew.OrgChorusId, &recordNew.DelDate, &recordNew.IsMv, &recordNew.TimeDifference)
		recordList = append(recordList, recordNew)
	}
	return recordList
}

type Record struct {
	Id             int            `json:"id"`
	SongId         int            `json:"songId"`
	SongName       string         `json:"songName"`
	Singer         string         `json:"singer"`
	Album          string         `json:"album"`
	UserId         int            `json:"userId"`
	NickName       string         `json:"nickName"`
	Score          int            `json:"score"`
	UpdateTime     time.Time      `json:"updateTime"`
	TimeLength     int            `json:"timeLength"`
	Status         int8           `json:"status"`
	ChannelId      int            `json:"channelId"`
	VoteCount      int            `json:"voteCount"`
	VoteTotal      int            `json:"voteTotal"`
	TypeInt        int            `json:"TypeInt"`
	ChorusId       int            `json:"chorusId"`
	Weight         int8           `json:"weight"`
	ChorusWeight   int8           `json:"ChorusWeight"`
	LrcUrl         string         `json:"lrcUrl"`
	LrcType        int            `json:"lrcType"`
	Backup         int8           `json:"backup"`
	FirstTotal     int            `json:"firstTotal"`
	OrgChorusId    int            `json:"orgChorusId"`
	DelDate        sql.NullString `json:"delDate"`
	IsMv           int            `json:"isMv"`
	TimeDifference string         `json:"timeDifference"`
	Enable         int            `json:"enable"`
}
