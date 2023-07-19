package auth_test

import (
	"net/http"
	"net/http/httptest"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/security/auth"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func TestCriarTokenJWT(t *testing.T) {
	tokenString, err := auth.CriarTokenJWT(uint64(1))
	if err != nil {
		t.Errorf("Erro ao criar token JWT: %v", err)
	}

	if len(tokenString) == 0 {
		t.Error("Token JWT vazio")
	}
}

func TestValidarToken(t *testing.T) {
	// Cria um contexto com o token no cabeçalho de autorização
	token := jwt.New(jwt.SigningMethodHS256)

	permissoes := token.Claims.(jwt.MapClaims)
	permissoes["teste"] = "token_de_teste"
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()

	tokenString, err := token.SignedString([]byte(configs.SecretKeyJWT))
	if err != nil {
		t.Errorf("Erro ao criar o token: %v", err)
	}

	// Cria um contexto com o token no cabeçalho de autorização
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	c := echo.New().NewContext(req, nil)

	// Chama a função ValidarToken
	err = auth.ValidarToken(c)
	if err != nil {
		t.Errorf("Erro ao validar token: %v", err)
	}
}

func TestExtrairUsuarioID(t *testing.T) {
	usuarioID := uint64(1)

	// Cria um token válido com o ID do usuário
	tokenString, err := auth.CriarTokenJWT(usuarioID)
	if err != nil {
		t.Errorf("Erro ao extrair ID do usuário: %v", err)

	}
	// Cria um contexto com o token no cabeçalho de autorização
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	c := echo.New().NewContext(req, nil)

	// Chama a função ExtrairUsuarioID
	id, err := auth.ExtrairUsuarioID(c)

	if err != nil {
		t.Errorf("Erro ao extrair ID do usuário: %v", err)
	}

	if id != usuarioID {
		t.Errorf("ID do usuário incorreto. Esperado: %d, Obtido: %d", usuarioID, id)
	}
}

/*
// CriarToken retorna um token assinado com as permissões do usuário
func TestCriarTokenJWT(t *testing.T) {

}

// ValidarToken verifica se o token passado na requisição é válido
func TestValidarToken(t *testing.T) {

}

// ExtrairUsuarioID retorna o usuarioId que está salvo no token
func TestExtrairUsuarioID(t *testing.T) {

}

func TestextrairToken(c echo.Context) {

}

func TestRetornarChaveDeVerificacao(t *testing.T) {

}
*/
