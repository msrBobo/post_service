package service

import (
	pb "POST_SERVICE/genproto/post-service"
	pbu "POST_SERVICE/genproto/user-service"
	l "POST_SERVICE/pkg/logger"
	grpcClient "POST_SERVICE/service/grpc_client"
	"POST_SERVICE/storage"

	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Postervice ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	req.Id = id.String()

	post, err := s.storage.Post().Create(req)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) Get(ctx context.Context, req *pb.GetRequest) (*pb.PostResponse, error) {
	post, err := s.storage.Post().GetPost(req.PostId)
	if err != nil {
		return nil, err
	}
	user, err := s.client.UserService().Get(ctx, &pbu.UserRequest{
		UserId: post.OwnerId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{
		Id:       post.Id,
		Title:    post.Title,
		ImageUrl: post.ImageUrl,
		Owner: &pb.Owner{
			Id:       user.Id,
			Name:     user.Name,
			LastName: user.LastName,
		},
	}, nil
}

// func (s *UserService) Update(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
// 	user, err := s.storage.User().Update(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (s *UserService) Delete(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
// 	user, err := s.storage.User().Delete(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (s *UserService) GetAll(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
// 	users, err := s.storage.User().GetAll(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
