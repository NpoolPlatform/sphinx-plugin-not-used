package db

import (
	"context"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
	"github.com/NpoolPlatform/go-service-framework/pkg/app"
)

var myClient *ent.Client

func Init() error {
	myClient = ent.NewClient(ent.Driver(app.Mysql().Driver))
	return myClient.Schema.Create(context.Background())
}

func Client() *ent.Client {
	return myClient
}

