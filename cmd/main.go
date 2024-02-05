package main

import (
	"POST_SERVICE/config"
	pb "POST_SERVICE/genproto/post-service"
	"POST_SERVICE/pkg/db"
	"POST_SERVICE/pkg/logger"
	"POST_SERVICE/service"
	grpcclient "POST_SERVICE/service/grpc_client"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "POST_SERVICE")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err, _ := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	grpcClient, err := grpcclient.New(cfg)

	if err != nil {
		log.Fatal("grpc client dail error", logger.Error(err))
	}

	postService := service.NewPostService(connDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
