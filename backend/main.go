package main

import (
	"web/internal/handler"
	"web/internal/repository"
	"web/internal/server"
	"web/internal/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const configPath = "config/config.json"

func main() {
	// config, err := config.ParseConfig(configPath)
	// if err != nil {
	// 	logger.Error("Config error")
	// 	return
	// }

	// connection := database.NewDBConnection()
	// err = connection.Connect(config)
	// defer connection.Close()

	// fmt.Println(connection.Connection.Ping())

	// // MARK: USER SERVICE
	// userRepo := repository.NewUserRepo(connection)
	// userService := services.NewUserService(userRepo)

	// playerRepo := repository.NewPlayerRepo(connection)
	// playerService := services.NewPlyerService(playerRepo)

	// teamRepo := repository.NewTeamRepo(connection)
	// teamService := services.NewTeamService(teamRepo)

	// go server.SetupServer(config, *userService, *playerService, *teamService).Run(":5007")

	// listenErr := make(chan error, 1)
	// select {
	// case err = <-listenErr:
	// 	if err != nil {
	// 		logger.Error(fmt.Sprintln("%v", err))
	// 		os.Exit(1)
	// 	}

	// }
	// playerRepo := repository.NewPlayerRepo(connection)
	// var player model.Player = model.Player{
	// 	ID:      1,
	// 	FName:   "Leo",
	// 	LNname:  "Messi",
	// 	Country: "Port1",
	// 	Dob:     common.GetCustomTime(1980, 1, 25),
	// }
	// fmt.Println(player)
	// playerRepo.CreatePlayer(player)
	// players, _ := playerRepo.GetPlayers()
	// fmt.Println(len(players))

	// dateString := "2100-13-02"
	// date, error := time.Parse("2006-11-02", dateString)

	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }

	// fmt.Printf("Type of dateString: %T\n", dateString)
	// fmt.Printf("Type of date: %T\n", date)
	// fmt.Println()
	// fmt.Printf("Value of dateString: %v\n", dateString)
	// fmt.Printf("Value of date: %v", date)

	// users, _ := service.GetUsers()

	// for _, user := range users {
	// 	fmt.Println(user.Login)
	// }

	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalln("error init configs:", err.Error())
	}

	db, err := repository.NewPostrgersDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalln("failed to init db:", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalln("error occured while running http server:", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
