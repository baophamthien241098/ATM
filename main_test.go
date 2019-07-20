package ATM

import (
	"context"


	"math/rand"
	"strconv"
	"testing"

	"github.com/ATM/ATMpd"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func BenchmarkATM(b *testing.B) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := ATMpd.NewATMServiceClient(conn)
	b.ResetTimer()

	b.Run("Set", func(p *testing.B) {
		rand.Seed(56789)
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_, err2 := client.SetMoney(context.Background(), &ATMpd.MoneyResquest{
					Account: &ATMpd.Account{
						ID:     strconv.FormatInt(rand.Int63(), 16),
						Number: 100,
					},
				})
				//expected err2 == nil
				assert.Nil(b, err2)
			}
		})
	})
	b.Run("Get", func(p *testing.B) {
		rand.Seed(56789)
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				value, err2 := client.GetMoney(context.Background(), &ATMpd.GetMoneyRequest{
					ID: strconv.FormatInt(rand.Int63(), 16),
				})
				//expected err2 ==nil
				assert.Nil(b, err2)
				//
				assert.Contains(b, []int64{0, 100}, value.GetNumber())
			}
		})
	})
	b.Run("Incre", func(p *testing.B) {
		rand.Seed(56789)
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_, err := client.IncreMoney(context.Background(), &ATMpd.MoneyResquest{
					Account: &ATMpd.Account{
						ID:     strconv.FormatInt(rand.Int63(), 16),
						Number: 100,
					},
				})
				assert.Nil(b, err)
			}
		})
	})
	b.Run("Descre", func(p *testing.B) {
		rand.Seed(56789)
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_, err := client.DecreMoney(context.Background(), &ATMpd.MoneyResquest{
					Account: &ATMpd.Account{
						ID:     strconv.FormatInt(rand.Int63(), 16),
						Number: 100,
					},
				})
				assert.Nil(b, err)
			}
		})
	})

}
