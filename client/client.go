package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"log"

	"github.com/ATM/ATMpd"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("eorrr 2222")
	}
	client := ATMpd.NewATMServiceClient(conn)
	value, err2 := client.SetMoney(context.Background(), &ATMpd.MoneyResquest{
		Account: &ATMpd.Account{
			ID:     strconv.FormatInt(rand.Int63(), 16),
			Number: 100,
		},
	})
	if err2 != nil {
		log.Fatalf("error : %v\n", err2)
	}
	fmt.Println(value)
}
