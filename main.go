package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/pflag"
	"github.com/valyala/fasthttp"
)

const DevMode = false

var port = pflag.Int("port", 9090, "Port to listen on")

func main() {
	pflag.Parse()

	engine := html.NewFileSystem(http.FS(index), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(requestid.New())
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: configureLogger(),
		Fields: []string{"ip", "latency", "status", "method", "url", "error", "request_id"},
	}))
	configureRouter(app)

	if DevMode {
		configureDevMode(app)
	} else {
		configureProductionMode(app)
	}

	err := app.Listen(fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
}

func configureDevMode(app *fiber.App) {
	proxy.WithClient(&fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		DisablePathNormalizing:   true,
		DialDualStack:            true,
	})
	app.Get("*", proxy.BalancerForward([]string{"http://localhost:5173"}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "Test",
			"DevMode": DevMode,
		})
	})
}

func configureProductionMode(app *fiber.App) {
	// Parse JSON Vite manifest
	manifest := map[string]Chunk{}
	data, err := content.ReadFile("frontend/dist/.vite/manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &manifest)
	if err != nil {
		log.Fatal(err)
	}

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(content),
		PathPrefix: "frontend/dist/assets",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "Test",
			"DevMode": DevMode,
			"MainCSS": manifest["src/main.tsx"].Css[0],
			"MainJS":  manifest["src/main.tsx"].File,
		})
	})
}
