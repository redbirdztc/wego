package db

import "gorm.io/gorm"

var connectionKeeper ConnectionKeeper

type ConnectionKeeper interface {
	GetConnection() *gorm.DB
}

func SetConnectionKeeper(keeper ConnectionKeeper) {
	connectionKeeper = keeper
}

func GetConnection() *gorm.DB {
	if connectionKeeper == nil {
		panic("connection keeper is not set")
	}
	return connectionKeeper.GetConnection()
}
