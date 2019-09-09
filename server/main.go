package main

import (
	"context"
	"log"
	"net"
	"sync"

	ATMpd "github.com/ATM/ATMpd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var arr = map[string]int64{}
var sw sync.Mutex

func (*server) SetMoney(ctx context.Context, req *ATMpd.MoneyResquest) (*ATMpd.MoneyResponse, error) {
	defer sw.Unlock()
	account := req.GetAccount()
	sw.Lock()
	arr[account.GetID()] = account.GetNumber()
	return &ATMpd.MoneyResponse{
		Number: account.GetNumber(),
	}, nil
}
func (*server) GetMoney(ctx context.Context, req *ATMpd.GetMoneyRequest) (*ATMpd.MoneyResponse, error) {
	ID := req.GetID()
	value, err := arr[ID]
	if err == false {
		return &ATMpd.MoneyResponse{
			Number: 0,
		}, nil
	}
	return &ATMpd.MoneyResponse{
		Number: value,
	}, nil
}
func (*server) IncreMoney(ctx context.Context, req *ATMpd.MoneyResquest) (*ATMpd.MoneyResponse, error) {
	defer sw.Unlock()
	sw.Lock()
	account := req.GetAccount()
	value, err := arr[account.GetID()]
	if err == false {
		arr[account.GetID()] = 0
		return &ATMpd.MoneyResponse{
			Number: 0,
		}, nil
	}
	arr[account.GetID()] = value + account.GetNumber()
	return &ATMpd.MoneyResponse{Number: arr[account.GetID()]}, nil
}
func (*server) DecreMoney(ctx context.Context, req *ATMpd.MoneyResquest) (*ATMpd.MoneyResponse, error) {
	defer sw.Unlock()
	sw.Lock()
	account := req.GetAccount()
	value, err := arr[account.GetID()]
	if err == false {
		arr[account.GetID()] = 0
		return &ATMpd.MoneyResponse{
			Number: 0,
		}, nil
	}
	arr[account.GetID()] = account.GetNumber() - value
	return &ATMpd.MoneyResponse{Number: arr[account.GetID()]}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}
	s := grpc.NewServer()
	ATMpd.RegisterATMServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error listening client %v\n", err)
	}

}
