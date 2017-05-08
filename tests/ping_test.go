package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	_ "github.com/scmo/foodchain-backend/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
)


// Test ping endpoint
func TestPing(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/ping", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPing", "Code:", w.Code, "Response:", w.Body.String())


	Convey("Subject: Test Ping Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result be 'pong'", func() {
			pong, _ := json.Marshal("pong")
			So(w.Body.String(), ShouldEqual, string(pong))
		})
	})
}