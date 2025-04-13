package adapters

// import (
// 	"context"

// 	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/dtos"
// 	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/infra/adapters/mongo/mappers"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func (adapter *UserAdapter) CreateUser(ctx context.Context, user *dtos.UserDTO) (string, error) {
// 	collection := adapter.MongoDatabase.Collection(adapter.collection)

// 	res, err := collection.InsertOne(ctx, mappers.MapCreateUser(user))
// 	if err != nil {
// 		return "", err
// 	}

// 	return res.InsertedID.(primitive.ObjectID).Hex(), nil
// }
