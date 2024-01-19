package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/gomail.v2"
)

const (
	redisURL      = "localhost:6379"
	taskQueueName = "emailQueue"
)

func sendEmail(email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "muhammetegeldiyevdayanc@gmail.com") // Gönderen e-posta adresi
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Test E-posta")
	m.SetBody("text/html", "Merhaba, bu bir test e-postasıdır.")

	d := gomail.NewDialer("smtp.gmail.com", 587, "muhammetegeldiyevdayanc@gmail.com", "Genetik19970825.*/")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("E-posta gönderiminde hata:", err)
		return err
	}

	fmt.Printf("E-posta gönderildi: %s\n", email)
	return nil
}

func processTask(ctx context.Context, client *redis.Client) {
	for {
		result, err := client.BLPop(ctx, 0, taskQueueName).Result()
		if err != nil {
			fmt.Println("Görev kuyruğundan görev alınamadı:", err)
			continue
		}

		email := result[1]
		if err := sendEmail(email); err != nil {
			fmt.Println("E-posta gönderiminde hata:", err)
		}
	}
}

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	client.Del(ctx, taskQueueName)

	go processTask(ctx, client)

	for i := 1; i <= 5; i++ {
		email := "recipient.email@example.com" // Gerçek alıcı e-posta adresi
		if err := client.RPush(ctx, taskQueueName, email).Err(); err != nil {
			fmt.Println("Görev kuyruğuna görev eklenirken hata:", err)
		}
	}

	time.Sleep(time.Second * 5)
}
