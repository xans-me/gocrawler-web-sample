package configuration

// AppConfig struct model, the base configuration struct for this web service
type AppConfig struct {
	App     App
	Mongodb Mongodb
}

// App struct model, the configuration file related to basic webservice info such as name environtment
type App struct {
	Name        string
	Environment string
	Debug       bool
	Host        string
	Port        string
	Protocol    string
}

// Mongodb struct model, contains configuration for mongodb. MongoDB is a cross-platform document-oriented database program (noSQL)
type Mongodb struct {
	DbName     string
	Connection string
}
