package mongodbtest

import (
	"context"
	"fmt"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DockerClient creates a MongoDB server via TestContainers.
// A client connected to the server and a cleanup function are returned.
// The cleanup function will close the client and shut down the server.
func DockerClient() (client *mongo.Client, cleanup func() error, err error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo:4.0.8",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("waiting for connections"),
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": "admin",
			"MONGO_INITDB_ROOT_PASSWORD": "admin",
		},
	}
	mongoServer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, err
	}
	ip, err := mongoServer.Host(ctx)
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	port, err := mongoServer.MappedPort(ctx, "27017/tcp")
	if err != nil {
		return nil, nil, err
	}

	client, err = mongo.Connect(
		ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://admin:admin@%v:%d", ip, port.Int())),
	)
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, nil, err
	}
	return client, func() error {
		if err = client.Disconnect(context.Background()); err != nil {
			return err
		}
		mongoServer.Terminate(context.Background())
		return nil
	}, nil
}
