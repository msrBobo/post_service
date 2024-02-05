package repo

import (
	pb "POST_SERVICE/genproto/post-service"
)

// PostStorageI ...
type PostStorageI interface {
	Create(user *pb.Post) (*pb.Post, error)
	GetPost(id string) (*pb.Post, error)
}
