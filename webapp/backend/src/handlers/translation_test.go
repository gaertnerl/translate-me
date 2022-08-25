package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gaertnerl/translate-me.git/webserver/services"
	"github.com/gaertnerl/translate-me.git/webserver/testingutils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Post_NextTranslation(t *testing.T) {

	testingutils.Setup(services.DB)

	testProtocol := `"protocol"`
	testHost := `"host"`
	testKey := `"` + testingutils.GetDummyEnvVars().REGISTER_SIM_ENDPOINT_KEY + `"`

	testData := `{
		"TranslationId": ` + testKey + `,
		"TranslationType": ` + testProtocol + `,
		"Sentence1": ` + testHost + `
		"Sentence1": ` + testHost + `
	}`

	testDataBytes := []byte(testData)

	g := gin.Default()
	g.POST("/", Post_RegisterSimilarityEndpoint)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(testDataBytes))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
