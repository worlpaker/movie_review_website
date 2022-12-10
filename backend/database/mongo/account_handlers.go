package MongoDB

import (
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"mime/multipart"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (s *service) Login(user *Models.Account) (*Models.Token, error) {
	var dbUser *Models.Account
	collection := client.Database("GODB").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser); ErrHandler.Log(err) {
		return nil, err
	}
	if err := VerifyPassword(dbUser.Password, user.Password); ErrHandler.Log(err) {
		return nil, err
	}
	data_token := &Models.Token{Save_Token: true,
		Email: dbUser.Email, First_name: dbUser.First_name, Last_name: dbUser.Last_name,
		Birth_date: dbUser.Birth_date, Gender: dbUser.Gender, Pp_id: fmt.Sprint(dbUser.Profile_picture)}
	return data_token, nil
}

func (s *service) CreateUser(user *Models.Account) (*Models.Account, error) {
	var dbUser *Models.Account
	user.Password = HashPassword([]byte(user.Password))
	collection := client.Database("GODB").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ifexist := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser); ifexist == nil {
		ErrHandler.Log(fmt.Errorf("user already exist"))
		return nil, fmt.Errorf("user already exist")
	}
	user, err := InitAccountBSON(user)
	if ErrHandler.Log(err) {
		return nil, err
	}
	if _, err := collection.InsertOne(ctx, user); ErrHandler.Log(err) {
		return nil, err
	}
	return user, nil
}

func (s *service) ChangePassword(user *Models.ChangePassword) error {
	var dbUser *Models.Account
	collection := client.Database("GODB").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser); ErrHandler.Log(err) {
		return err
	}
	if err := VerifyPassword(dbUser.Password, user.Old_Password); ErrHandler.Log(err) {
		return err
	}
	New_Pw := HashPassword([]byte(user.New_Password))
	filter := bson.M{"email": user.Email}
	update := bson.M{"$set": bson.M{"Password": New_Pw}}
	if _, err := collection.UpdateOne(ctx, filter, update); ErrHandler.Log(err) {
		return err
	}
	return nil
}

func (s *service) UpdateProfile(user *Models.UpdateAccount) (*Models.Token, error) {
	var dbUser *Models.Account
	collection := client.Database("GODB").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser); ErrHandler.Log(err) {
		return nil, err
	}
	filter := bson.M{"email": user.Email}
	update := bson.M{"$set": bson.M{"First_name": user.First_name, "Last_name": user.Last_name, "Birth_date": user.Birth_date, "Gender": user.Gender}}
	if _, err := collection.UpdateOne(ctx, filter, update); ErrHandler.Log(err) {
		return nil, err
	}
	data_token := &Models.Token{Save_Token: true,
		Email: user.Email, First_name: user.First_name, Last_name: user.Last_name,
		Birth_date: user.Birth_date, Gender: user.Gender, Pp_id: fmt.Sprint(dbUser.Profile_picture)}
	return data_token, nil
}

func (s *service) UploadPhoto(file multipart.File, filename string) error {
	data, err := jpeg.Decode(file)
	if ErrHandler.Log(err) {
		return err
	}
	f, err := os.Create(filename)
	if ErrHandler.Log(err) {
		return err
	}
	defer f.Close()
	err = jpeg.Encode(f, data, nil)
	if ErrHandler.Log(err) {
		return err
	}
	bucket, err := gridfs.NewBucket(client.Database("GODB"))
	if ErrHandler.Log(err) {
		return err
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	if ErrHandler.Log(err) {
		return err
	}
	defer uploadStream.Close()
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, data, nil)
	if ErrHandler.Log(err) {
		return err
	}
	send_s3 := buf.Bytes()
	if _, err := uploadStream.Write(send_s3); ErrHandler.Log(err) {
		return err
	}
	return nil
}

// DeleteData removes data
func (s *service) DeleteData(d *Models.DeleteDataModel) error {
	collection := client.Database("GODB").Collection(d.CollName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := collection.DeleteOne(ctx, bson.M{d.Filter: d.Data}); ErrHandler.Log(err) {
		return err
	}
	return nil
}
