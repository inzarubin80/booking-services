package config
import (
    "log"
	"github.com/joho/godotenv"
)

func init() {
    Load()
}

func Load(){
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}