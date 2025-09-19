package main

import (
	"log"
	"net/http"
	"os"

	restful "github.com/emicklei/go-restful/v3"

	"trading/internal/services/web/handlers"
)

func main() {
	log.Println("🚀 Iniciando Sistema de Trading Web Service...")

	// Cria container RESTful
	ws := handlers.NewInternalWebRestfulContainer()

	// Configura router
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.Add(ws.GetWS())

	// Configurações globais
	restful.DefaultContainer.EnableContentEncoding(true)

	// Porta do servidor
	port := getEnv("PORT", "8080")
	log.Printf("📊 Web Service rodando na porta %s", port)
	log.Printf("🌐 Health Check: http://localhost:%s/api/health", port)
	log.Printf("📈 API: http://localhost:%s/api/orders", port)
	log.Printf("📊 Order Book: http://localhost:%s/api/orderbook/AAPL", port)
	log.Printf("👤 Portfolio: http://localhost:%s/api/portfolio/ana-silva", port)

	// Inicia servidor
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("❌ Erro ao iniciar web service:", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
