package postgres

import (
	pb "POST_SERVICE/genproto/post-service"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(post *pb.Post) (*pb.Post, error) {

	if post.Id == "" {
		id, err := uuid.NewUUID()
		if err != nil {
			return nil, err
		}
		post.Id = id.String()
	}

	query := `INSERT INTO posts (
		id, 
		title, 
		image_url,
		owner_id
	) 
	VALUES ($1, $2, $3, $4) 
	RETURNING 
		id, 
		title, 
		image_url,
		owner_id
	`
	var respPost pb.Post
	err := r.db.QueryRow(
		query,
		post.Id,
		post.Title,
		post.ImageUrl,
		post.OwnerId,
	).Scan(
		&respPost.Id,
		&respPost.Title,
		&respPost.ImageUrl,
		&respPost.OwnerId,
	)
	if err != nil {
		return nil, err
	}
	return &respPost, nil
}

func (r *postRepo) GetPost(id string) (*pb.Post, error) {
	query := `SELECT id, title, image_url, owner_id FROM posts WHERE id = $1`
	var respPost pb.Post
	err := r.db.QueryRow(query, id).Scan(&respPost.Id, &respPost.Title, &respPost.ImageUrl, &respPost.OwnerId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &respPost, nil
}

// func (r *postRepo) Update(user *pb.UserRequest) (*pb.User, error) {
// 	query := `
// 	UPDATE
// 		users
// 	SET
// 		name = $1,
// 		last_name = $2
// 	WHERE
// 		id = $3
// 	RETURNING
// 		id,
// 		name,
// 		last_name
// 	`
// 	var respUser pb.User
// 	err := r.db.QueryRow(
// 		query,
// 		"New Name",
// 		"New Last Name",
// 		user.UserId,
// 	).Scan(
// 		&respUser.Id,
// 		&respUser.Name,
// 		&respUser.LastName,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &respUser, nil
// }

// func (r *postRepo) Delete(user *pb.UserRequest) (*pb.User, error) {
// 	query := `DELETE FROM users WHERE id = $1 RETURNING id, name, last_name`
// 	var respUser pb.User
// 	err := r.db.QueryRow(query, user.UserId).Scan(&respUser.Id, &respUser.Name, &respUser.LastName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &respUser, nil
// }

// func (r *postRepo) GetAll(req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
// 	offset := req.Limit * (req.Page - 1)
// 	query := `SELECT id, name, last_name FROM users LIMIT $1 OFFSET $2`
// 	rows, err := r.db.Query(query, req.Limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var allUsers pb.GetAllUsersResponse
// 	for rows.Next() {
// 		var user pb.User
// 		if err := rows.Scan(&user.Id, &user.Name, &user.LastName); err != nil {
// 			return nil, err
// 		}
// 		allUsers.AllUsers = append(allUsers.AllUsers, &user)
// 	}
// 	return &allUsers, nil
// }
