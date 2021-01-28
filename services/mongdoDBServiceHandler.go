package services

import (
	"context"
	"fmt"
	"github.com/dchest/uniuri"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBServiceHandler struct {
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Hosts    []string `yaml:"hosts"`
	Database string   `yaml:"database"`
}

var ctx = context.TODO()

func (m *MongoDBServiceHandler) login() (*mongo.Client, error) {
	mongoConfig := options.Client()
	mongoConfig.SetAuth(options.Credential{
		Username: m.Username,
		Password: m.Password})
	mongoConfig.SetHosts(m.Hosts)
	client, err := mongo.Connect(ctx, mongoConfig)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *MongoDBServiceHandler) CreateCredentials() (map[string]interface{}, error) {
	client, err := m.login()
	if err != nil {
		return nil, err
	}
	bindingUser := uniuri.New()
	bindingPassword := uniuri.New()

	role := bson.D{
		{"role", "readWrite"}, {"db", m.Database},
	}

	command := bson.D{
		{"createUser", bindingUser},
		{"pwd", bindingPassword},
		{"roles", []bson.D{role}},
	}
	db := client.Database(m.Database)

	var result bson.M
	err = db.RunCommand(ctx, command).Decode(&result)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"uri":      m.createDbUri(bindingUser, bindingPassword),
		"user":     bindingUser,
		"password": bindingPassword,
		"database": m.Database}, nil
}

func (m *MongoDBServiceHandler) DeleteCredentials(json map[string]interface{}) error {
	client, err := m.login()
	if err != nil {
		return err
	}
	user := json["user"].(string)

	command := bson.D{{"dropUser", user}}

	var result bson.M
	err = client.Database(m.Database).RunCommand(context.TODO(), command).Decode(&result)
	return err
}

func (m *MongoDBServiceHandler) createDbUri(user, pwd string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s/%s", user, pwd, m.Hosts[0], m.Database)
}
