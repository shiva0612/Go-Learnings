package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	sm "github.com/shiva0612/grpc-flat/api/models"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ApierServer struct {
	sm.UnimplementedUserAPIServer
}

func newServer() *ApierServer {
	return &ApierServer{}
}

var (
	res *flatbuffers.Builder
)

func (s *ApierServer) Session_charging(ctx context.Context, req *sm.UserRequest) (*flatbuffers.Builder, error) {
	return res, nil
}

func getresponse() {
	var err error
	b, err := ioutil.ReadFile("op.json")
	UserResponseT := &sm.UserResponseT{}
	json.Unmarshal(b, UserResponseT)
	if err != nil {
		panic(err.Error())
	}
	res = flatbuffers.NewBuilder(0)
	res.Finish(UserResponseT.Pack(res))
}
