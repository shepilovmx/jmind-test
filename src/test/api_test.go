package test

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"jmind-test/src/config"
	"jmind-test/src/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var router *mux.Router

type TestStruct struct {
	name         string
	req          []byte
	expectedErr  bool
	expectedCode int
	options      map[string]interface{}
}

func TestMain(m *testing.M) {
	os.Setenv("DB_ADDR", "localhost:5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "jmind-test")
	os.Setenv("ETHERSCAN_API_URL", "http://api.etherscan.io/api")
	os.Setenv("ETHERSCAN_API_KEY", "YourApiKeyToken")

	router = routes.Router()
	config.ServerCtx = config.InitServerContext()
	code := m.Run()
	os.Exit(code)
}

func TestBlockTotal(t *testing.T) {
	testCases := []TestStruct{
		{
			name:        "block/{block_number}/totals VALID",
			expectedErr: false,
			options: map[string]interface{}{
				"block_number": "11509797",
			},
		},
		{
			name: "block/{block_number}/totals INVALID",
			options: map[string]interface{}{
				"block_number": "qwe",
			},
			expectedErr:  true,
			expectedCode: 400,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("/api/block/%s/total", tc.options["block_number"]), nil)

			if err != nil {
				t.Errorf("Req error. Got %s", err)
				t.FailNow()
			}
			res := executeRequest(req)
			log.Println(res)

			if tc.expectedErr {
				checkResponseCode(t, tc.expectedCode, res.Code)
			} else {
				checkResponseCode(t, http.StatusOK, res.Code)
				body, _ := jsonDecode(res.Body)
				assertEqual(t, body["transactions"], json.Number("155"), "")
				assertEqual(t, body["amount"], json.Number("2.285404805647828"), "")
			}
		})
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
		t.FailNow()
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func jsonDecode(data io.Reader) (map[string]interface{}, error) {
	var responseData map[string]interface{}

	dec := json.NewDecoder(data)
	dec.UseNumber()
	err := dec.Decode(&responseData)

	return responseData, err
}
