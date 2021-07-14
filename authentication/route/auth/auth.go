package auth

import (
	"fmt"
	"net/http"

	"github.com/0xedb/everything_go/authentication/route"
)

const (
	root string = "/auth"

	ext     string = "here.apx"
	name    string = "/name/"
	options string = "/:opt/abt/*more"
)

type auth struct{}

// Instance is an auth endpoint instance
var Instance route.Path

func init() {
	Instance = &auth{}
}

func handleOption(c route.Ctx) {
	fmt.Println("------", c.Param("opt"), c.Param("more"))
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, fmt.Sprintf(`<h1>
	Main: %s <br />
	Others: %s
	</>`, c.Param("opt"), c.Param("more")))
}
