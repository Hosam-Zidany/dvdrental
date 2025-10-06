package main

import (
	"fmt"
	"log"

	"github.com/Hosam-Zidany/dvdrental/internal/config"
	"github.com/Hosam-Zidany/dvdrental/internal/database"
	"github.com/Hosam-Zidany/dvdrental/internal/routers"
)

func main() {
	cfg := config.LoadConfig()
	database.InintDB(cfg)
	r := routers.SetupRouter()

	addr := fmt.Sprintf(":%s", cfg.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
