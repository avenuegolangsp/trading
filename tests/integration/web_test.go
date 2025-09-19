package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	restful "github.com/emicklei/go-restful/v3"

	"trading/internal/services/web/handlers"
)

// TestWebServiceEndpoints testa se todos os endpoints estão funcionando
func TestWebServiceEndpoints(t *testing.T) {
	// Setup
	container := handlers.NewInternalWebRestfulContainer()
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.Add(container.GetWS())

	// Testa health check
	t.Run("HealthCheck", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/health", nil)
		resp := httptest.NewRecorder()
		restful.DefaultContainer.ServeHTTP(resp, req)

		if resp.Code != 200 {
			t.Errorf("Esperado status 200, obtido %d", resp.Code)
		}

		body := resp.Body.String()
		if body != "OK - HealthCheck" {
			t.Errorf("Esperado 'OK - HealthCheck', obtido '%s'", body)
		}
	})

	// Testa order book
	t.Run("GetOrderBook", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/orderbook/AAPL", nil)
		resp := httptest.NewRecorder()
		restful.DefaultContainer.ServeHTTP(resp, req)

		if resp.Code != 200 {
			t.Errorf("Esperado status 200, obtido %d", resp.Code)
		}

		body := resp.Body.String()
		if body != "OK - GetOrderBook" {
			t.Errorf("Esperado 'OK - GetOrderBook', obtido '%s'", body)
		}
	})

	// Testa portfolio
	t.Run("GetPortfolio", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/portfolio/ana-silva", nil)
		resp := httptest.NewRecorder()
		restful.DefaultContainer.ServeHTTP(resp, req)

		if resp.Code != 200 {
			t.Errorf("Esperado status 200, obtido %d", resp.Code)
		}

		body := resp.Body.String()
		if body != "OK - GetPortfolio" {
			t.Errorf("Esperado 'OK - GetPortfolio', obtido '%s'", body)
		}
	})

	// Testa stocks
	t.Run("GetStocks", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/stocks", nil)
		resp := httptest.NewRecorder()
		restful.DefaultContainer.ServeHTTP(resp, req)

		if resp.Code != 200 {
			t.Errorf("Esperado status 200, obtido %d", resp.Code)
		}

		body := resp.Body.String()
		if body != "OK - GetStocks" {
			t.Errorf("Esperado 'OK - GetStocks', obtido '%s'", body)
		}
	})

	t.Logf("✅ Todos os endpoints estão respondendo corretamente!")
}
