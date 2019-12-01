package main

import infra "github.com/hukurou-s/user-auth-api-sample/infrastructure"

func main() {
	infra.Echo.Logger.Fatal(infra.Echo.Start(":1323"))
}
