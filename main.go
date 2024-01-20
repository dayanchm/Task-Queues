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
	m.SetHeader("From", "example@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Test E-posta")
	m.SetBody("text/html", "Merhaba, bu bir test e-postasıdır.")

	d := gomail.NewDialer("smtp.gmail.com", 587, "example@gmail.com", "Password")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error:", err)
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
			fmt.Println("Error:", err)
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
		email := "recipient.email@example.com"
		if err := client.RPush(ctx, taskQueueName, email).Err(); err != nil {
			fmt.Println("Görev kuyruğuna görev eklenirken hata:", err)
		}
	}

	time.Sleep(time.Second * 5)
}
