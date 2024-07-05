package services

import (
	context "context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userServer struct {
	db *mongo.Collection
}

func NewUserServer(db *mongo.Collection) *userServer {
	return &userServer{db: db}
}

func (sv *userServer) mustEmbedUnimplementedUserServiceServer() {}

func (sv *userServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	var result struct {
		UserID string `bson:"user_id"`
	}

	filter := bson.M{"user_id": req.UserId}
	err := sv.db.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	res := &GetUserResponse{
		UserId: result.UserID,
	}
	return res, nil
}
