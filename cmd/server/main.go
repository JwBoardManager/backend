package main

import (
	out "backend/db/sqlc/out"
	"backend/pkg/assignment"
	"backend/routes"
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// 🔹 Carregar variáveis de ambiente do .env (se existir)
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Arquivo .env não encontrado, usando variáveis do sistema")
	}

	// 🔹 Definir conexão correta com o banco de dados
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL não foi configurado")
	}

	// 🔹 Testar conexão com o banco
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Não foi possível pingar o banco de dados: %v", err)
	}

	log.Println("✅ Conexão com o banco de dados estabelecida!")

	// 🔹 Criar instância do SQLC Queries
	queries := out.New(db)

	// 🔹 Criar serviços
	assignmentService := assignment.NewAssignmentService(queries)

	r := gin.Default()

	// 🔹 Configurar rotas
	routes.SetupRouter(r, assignmentService)

	// 🔹 Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}

	log.Printf("🚀 Servidor rodando na porta %s", port)
	r.Run(":" + port)
}
