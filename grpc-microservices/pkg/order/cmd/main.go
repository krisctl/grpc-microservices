package main

import (
	"log"

	"github.com/krisctl/grpc-microservices/pkg/order/config"
	"github.com/krisctl/grpc-microservices/pkg/order/internal/adapters/db"
	"github.com/krisctl/grpc-microservices/pkg/order/internal/adapters/grpc"
	"github.com/krisctl/grpc-microservices/pkg/order/internal/application/core/api"
)

// The DB adapter needs a data source URL to connect and return an instance for a DB
// reference. The core application needs this DB adaptor to modify order objects in the
// database. Finally, the gRPC adapter needs a core application and a specific port to get
// the gRPC up and running via the Run method:

func main() {
	dbAdapter, err := db.NewDbAdapter(config.GetDataSourceUrl())
	if err != nil {
		log.Fatalf("Failed to connect to database, error %v", err)
	}

	app := api.NewApplication(dbAdapter)
	grpc := grpc.NewGrpcAdapter(app, config.GetApplicationPort())
	grpc.Run()
}
