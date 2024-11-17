module user-service

go 1.22.3

require (
	github.com/golang-migrate/migrate/v4 v4.18.1
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.28.0
)

require (
	github.com/gabriel-vasile/mimetype v1.4.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240513163218-0867130af1f8 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
)

require (
	github.com/go-playground/validator/v10 v10.22.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/golang/protobuf v1.5.4
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.64.1
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	user-service/shared v0.0.0-00010101000000-000000000000
)

replace user-service/shared => ../shared
