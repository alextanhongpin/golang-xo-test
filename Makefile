include .env
export

CONNSTR=pgsql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

start:
	@go run main.go

install:
	@go get -u github.com/xo/xo

gen:
	xo ${CONNSTR} -o models

# We can pass envvars through custom script.
gen-custom:
	./custom.sh
