package app

import (
	"database/sql"
	"fmt"

	"github.com/SamanShafigh/lambda-go-boilerplate/util"
)

// Config app configourations
type Config struct {
	DbUser     string `json:"dbUser"`
	DbPassword string `json:"dbPassword"`
	DbHost     string `json:"dbHost"`
	DbName     string `json:"dbName"`
	Pagination int    `json:"pagination"`
}

// App provides the app structure
type App struct {
	Config *Config
	Model  *Model
}

// Model provides access to models
type Model struct {
	db *sql.DB
}

// New initialise the app
func New(path string) (*App, error) {
	config := Config{}
	// Load App configs from config json
	err := util.LoadConfig(path, &config)
	if err != nil {
		return nil, err
	}

	// Load DB credentials from env variable
	dbCredentials, _ := util.KmsDecrypt(util.Getenv("dbCredentials"))
	err = util.JSONDecode(dbCredentials, &config)
	if err != nil {
		return nil, err
	}

	// Open DB connection
	db, err := util.OpenDB(
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbName)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: &config,
		Model: &Model{
			db: db,
		}}, nil
}

// Run execute the main app
func (app *App) Run() (string, error) {

	userQuery := UserQuery{
		Username: "admin",
	}

	userModel := app.Model.GetUserModel()
	user, err := userModel.GetUser(userQuery)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("ID: %s Username: %s", user.Id, user.Username), nil
}
