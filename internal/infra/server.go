package infra

import (
	"context"
	"fmt"
	"go-file-share/internal/controller"
	"go-file-share/internal/service"
	"go-file-share/internal/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	port      string
	dirPath   string
	recursive bool

	ginEngine      *gin.Engine
	pathController controller.PathController
	pathService    service.PathService
}

func (s *server) Run() {
	s.setup()
	ips := utils.GetLocalIPs()

	fmt.Println("--------------------------------------------")
	fmt.Printf("🚀 File Server Started!\n")
	fmt.Printf("📂 Sharing: %s\n", s.dirPath)
	fmt.Printf("📂 Recursive: %v\n", s.recursive)
	fmt.Println("--------------------------------------------")

	srv := &http.Server{
		Addr:    ":" + s.port,
		Handler: s.ginEngine.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println("Access the server at:")

	fmt.Printf("  🏠 Local:   http://localhost:%s\n", s.port)
	for _, ip := range ips {
		fmt.Printf("  🌐 Network: http://%s:%s\n", ip, s.port)
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("Press Ctrl+C to stop.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server shutdown:", err)
	}
	<-ctx.Done()

	log.Println("Server exiting")

}

func (s *server) setup() {
	mode := os.Getenv("GIN_MODE")
	if len(mode) == 0 {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	s.ginEngine = gin.Default()
	s.ginEngine.Use(gin.Recovery())

	s.ginEngine.LoadHTMLGlob("templates/*")
	s.ginEngine.Static("/static", "./static")
	s.ginEngine.StaticFile("/favicon.ico", "./static/favicon.ico")

	s.instances()
	s.routes()
}

func (s *server) instances() {
	s.pathService = service.NewPathService()
	s.pathController = controller.NewPathController(s.dirPath, s.recursive, s.pathService)
}

func (s *server) routes() {
	s.ginEngine.Handle("GET", "/files/*path", s.pathController.DownloadFile)
	s.ginEngine.Handle("POST", "/files/*path", s.pathController.UploadFile)

	s.ginEngine.Handle("GET", "/", s.pathController.ShowFolder)
	listPathInfo := s.pathService.ListPathInfo(s.dirPath, "/", s.recursive)
	for _, pathInfo := range listPathInfo.Paths {
		if pathInfo.IsDir {
			s.ginEngine.Handle("GET", "/"+pathInfo.Path+"/*path", s.pathController.ShowFolder)
		}
	}

}

func NewServer(port string, dirPath string, recursive bool) Server {
	return &server{
		port:      port,
		dirPath:   dirPath,
		recursive: recursive,
	}
}
