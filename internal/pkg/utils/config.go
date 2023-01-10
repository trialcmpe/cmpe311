package utils

import "fmt"

func GetConnectionString(user, pw string) string {
	return fmt.Sprintf("mongodb+srv://%s:%s@cluster0.bupce6d.mongodb.net/?retryWrites=true&w=majority", user, pw)
}

func GetDBName() string {
	return "Cluster0"
}
