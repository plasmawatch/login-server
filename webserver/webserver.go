package webserver

import (
	"bnet-mock/network/account_manager"
	"bnet-mock/network/client/rest/login"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var accMgr = &account_manager.AccountMgr{} // Global instance

func StartWebServer(cer tls.Certificate) {
	form := new(login.FormInputs)
	form.Type = (*login.FormType)(proto.Int32(int32(login.FormType_LOGIN_FORM)))

	accountName := new(login.FormInput)
	accountName.InputId = proto.String("account_name")
	accountName.Type = proto.String("text")
	accountName.Label = proto.String("Email or Phone")
	accountName.MaxLength = proto.Uint32(320)

	password := new(login.FormInput)
	password.InputId = proto.String("password")
	password.Type = proto.String("password")
	password.Label = proto.String("Password")
	password.MaxLength = proto.Uint32(128)

	logInSubmit := new(login.FormInput)
	logInSubmit.InputId = proto.String("log_in_submit")
	logInSubmit.Type = proto.String("submit")
	logInSubmit.Label = proto.String("Log In")

	form.Inputs = append(form.Inputs, accountName)
	form.Inputs = append(form.Inputs, password)
	form.Inputs = append(form.Inputs, logInSubmit)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/battlenet/login", func(c *fiber.Ctx) error {
		//c.Cookies("web.id", "US-b330f9b9-7896-44e7-aa80-84532857ab7b")
		//c.Cookies("JSESSIONID", "23007ab6-fc44-4b96-bd61-69e5eda7cdc6")
		SetCommonHeaders(c)
		// c.Set("Content-Length", fmt.Sprintf("%d", len(form.)))

		c.Cookie(&fiber.Cookie{
			Name:  "web.id",
			Value: "US-b330f9b9-7896-44e7-aa80-84532857ab7b",
		})

		c.Cookie(&fiber.Cookie{
			Name:  "JSESSIONID",
			Value: "23007ab6-fc44-4b96-bd61-69e5eda7cdc6",
		})

		f := protojson.MarshalOptions{UseProtoNames: true}
		bytes, err := f.Marshal(form)
		if err != nil {
			log.Print(err)
		}

		log.Printf("%s", string(bytes))

		c.Set("Content-Length", fmt.Sprintf("%d", len(bytes)))

		return c.Send(bytes)
	})

	app.Post("/battlenet/login", func(c *fiber.Ctx) error {
		var l login.LoginForm
		if err := c.BodyParser(&l); err != nil {
			log.Print(err)
			return SendErrorResponse(c, "There was an internal error, failed to parse login form.")
		}

		/*
			f := protojson.MarshalOptions{UseProtoNames: true}
			bytes, err := f.Marshal(successLogin)
			if err != nil {
				log.Print(err)
			}*/

		log.Print(l.String())

		if len(l.GetInputs()) < 2 {
			log.Print("Too few login form inputs!")
			return SendErrorResponse(c, "There was an internal error, too few login form inputs.")
		}

		// Get user credentials
		accountNameStr := l.GetInputs()[0].GetValue()
		passwordStr := l.GetInputs()[1].GetValue()

		err := accMgr.Authenticate(accountNameStr, passwordStr)
		if err != nil {
			return SendErrorResponse(c, err.Error())
		}

		return SendSuccessResponse(c)
	})

	log.Println("Starting webserver on :6969")
	log.Fatal(app.Listen(":6969"))
}

// Sets common HTTP headers for the response
func SetCommonHeaders(c *fiber.Ctx) {
	c.Set("Content-Type", "application/json;charset=utf-8")
	c.Set("X-Frame-Options", "DENY")
	c.Set("Server", "Apache")
}

// Returns a successful authentication response
func SendSuccessResponse(c *fiber.Ctx) error {
	successPayload := `{
	"authentication_state": "DONE", 
	"login_ticket": "US-6b566d8e88863148abe3f872f1c1e33b-867364322"
	}`
	SetCommonHeaders(c)
	c.Set("Content-Length", fmt.Sprintf("%d", len(successPayload)))
	return c.Status(fiber.StatusOK).SendString(successPayload)
}

// Returns an error response
func SendErrorResponse(c *fiber.Ctx, errorMessage string) error {
	errorPayload := fmt.Sprintf(`{
        "authentication_state": "LOGIN",
        "error_code": "INVALID_ACCOUNT_OR_CREDENTIALS",
        "error_message": "%s",
        "input_id": "password",
        "support_error_code": "BLZBNTTAS00000002",
        "error_status": "WARNING",
        "error_message_helper": "Please enter the correct email and password."
    }`, errorMessage)

	SetCommonHeaders(c)
	c.Set("Content-Length", fmt.Sprintf("%d", len(errorPayload)))
	return c.Status(fiber.StatusUnauthorized).SendString(errorPayload)
}

/*
		mux := http.NewServeMux()

		mux.HandleFunc("/battlenet/login", func(w http.ResponseWriter, r *http.Request) {
			log.Print("login")

			// w.Write([]byte("login"))

			// redirect to /battlenet/login/1

			// set web.id cookie to US-b330f9b9-7896-44e7-aa80-84532857ab7b
			cookie := http.Cookie{
				Name:  "web.id",
				Value: "US-b330f9b9-7896-44e7-aa80-84532857ab7b",
			}

			http.SetCookie(w, &cookie)

			cookie = http.Cookie{
				Name:  "JSESSIONID",
				Value: "23007ab6-fc44-4b96-bd61-69e5eda7cdc6",
			}

			http.SetCookie(w, &cookie)

			f := `{"type":"LOGIN_FORM","inputs":[{"input_id":"account_name","type":"text","label":"Email or Phone","max_length":320},{"input_id":"password","type":"password","label":"Password","max_length":128},{"input_id":"log_in_submit","type":"submit","label":"Log In"}],"srp_url":"https://us.battle.net/login/srp","srp_js":"https://us.battle.net/login/static/js/login/srp-client.min.js"}`

			bytes, err := json.Marshal(f)
			if err != nil {
				log.Print(err)
			}

			w.Write(bytes)
		})

		// start the server on 1120
		log.Fatalf("Error starting server: %s", http.ListenAndServe(":6969", mux))

		log.Printf("Starting server on port 1120")
		/*
			server := &http.Server{
				Addr:    ":6969",
				Handler: mux,
				TLSConfig: &tls.Config{
					Certificates:       []tls.Certificate{cer},
					MinVersion:         tls.VersionTLS12,
					InsecureSkipVerify: true,
				},
			}

			log.Fatal(server.ListenAndServeTLS("", ""))
}
*/
