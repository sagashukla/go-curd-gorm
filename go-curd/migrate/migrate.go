package main 
 
import (
	"github.com/sagarshukla785/go-curd/initializers"
	"github.com/sagarshukla785/go-curd/models"
)
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToMyDB();
}

func main() {
	initializers.DBM.AutoMigrate(&models.Post{})
}