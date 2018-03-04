package main

import "fmt"
import "github.com/go-redis/redis"

func main(){
  var client *redis.Client
  client = ExampleNewClient()
  defer client.Close()

  err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

  val, err := client.Get("key").Result()
  if err != nil {
    panic(err)
  }
  fmt.Println("key", val)

  dsize, err2 := client.DBSize().Result()
  if err2 != nil {
    panic(err2)
  }
  fmt.Println(dsize)
  //psub := client.PSubscribe("hello", "crazy")
  psub := client.PSubscribe("batshit", "crazy")

  go PushMsg(client)



  ReadMsg(psub)
}

func ExampleNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.168:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
  //fmt.Printf( "%T", client)

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
  return client
}

func PushMsg(client *redis.Client) {
  var j int
  for i:= 1; i<20; i++ {
    j=i+100
    client.Publish("batshit", j)
    client.Publish("crazy", i)

  }
  client.Publish("batshit" , "done")
}
func ReadMsg(ps *redis.PubSub) {

  for true  {
    msg, err := ps.ReceiveMessage()
    if err != nil  {
      panic(err)
    }
    fmt.Println("msg = ", msg)

    fmt.Println("channel =", msg.Channel)
    fmt.Println("payload =", msg.Payload)

    if (msg.Payload == "done") {
      return
    }
  }
}
