package constants

var Branches []string // will load this from env . Migrate is removed

var ENV = []string{
	"REDIS_URI",
	"POSTGRES_USER",
	"POSTGRES_PASSWORD",
	"POSTGRES_DB",
	"SALTVALUE",
	// for course registeration and adding new branches quickly
	"BRANCHES",
}
