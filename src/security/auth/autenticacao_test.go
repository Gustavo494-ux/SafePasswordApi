package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/security/auth"
	"testing"
	"time"

	"safePasswordApi/src/modules/logger"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func TestCriarTokenJWT(t *testing.T) {
	tokenString, err := auth.CriarTokenJWT(uint64(1))
	if err != nil {
		logger.Logger().Error("Erro ao criar token JWT:", err)
		t.Errorf("Erro ao criar token JWT: %v", err)
	}

	if len(tokenString) == 0 {
		logger.Logger().Error("Token JWT vazio", err)
		t.Error("Token JWT vazio")
	}

	logger.Logger().Info("Teste de CriarTokenJWT executado com sucesso!")
}

func TestValidarToken(t *testing.T) {
	// Cria um contexto com o token no cabeçalho de autorização
	token := jwt.New(jwt.SigningMethodHS256)

	permissoes := token.Claims.(jwt.MapClaims)
	permissoes["teste"] = "token_de_teste"
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()

	tokenString, err := token.SignedString([]byte(configs.SecretKeyJWT))
	if err != nil {
		logger.Logger().Error("Erro ao criar o token:", err)
		t.Errorf("Erro ao criar o token: %v", err)
	}

	// Cria um contexto com o token no cabeçalho de autorização
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	c := echo.New().NewContext(req, nil)

	// Chama a função ValidarToken
	err = auth.ValidarToken(c)
	if err != nil {
		logger.Logger().Error("Erro ao validar token:", err)
		t.Errorf("Erro ao validar token: %v", err)
	}

	logger.Logger().Info("Teste de ValidarToken executado com sucesso!")
}

func TestExtrairUsuarioID(t *testing.T) {
	usuarioID := uint64(1)
	// Cria um token válido com o ID do usuário
	tokenString, err := auth.CriarTokenJWT(usuarioID)
	if err != nil {
		logger.Logger().Error("Erro ao extrair ID do usuário:", err)
		t.Errorf("Erro ao extrair ID do usuário: %v", err)
	}

	// Cria um contexto com o token no cabeçalho de autorização
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	c := echo.New().NewContext(req, nil)

	// Chama a função ExtrairUsuarioID
	id, err := auth.ExtrairUsuarioID(c)

	if err != nil {
		logger.Logger().Error("Erro ao extrair ID do usuário:", err)
		t.Errorf("Erro ao extrair ID do usuário: %v", err)
	}

	if id != usuarioID {
		logger.Logger().Error(fmt.Sprintf("ID do usuário incorreto. Esperado: %d		btido: %d: ", usuarioID, id), err)
		t.Errorf("ID do usuário incorreto. Esperado: %d, Obtido: %d", usuarioID, id)
	}

	logger.Logger().Info("Teste de ExtrairUsuarioID executado com sucesso!")
}
