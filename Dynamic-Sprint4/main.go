package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	graph := NewGraph()

	// Simulando conexões entre armazéns e pontos
	graph.AddEdge("Armazém A", "Centro B", 15)
	graph.AddEdge("Armazém A", "Depósito C", 25)
	graph.AddEdge("Centro B", "Loja D", 12)
	graph.AddEdge("Depósito C", "Posto E", 40)
	graph.AddEdge("Loja D", "CD F", 30)
	graph.AddEdge("Posto E", "Estoque G", 20)
	graph.AddEdge("CD F", "Distribuidor H", 28)
	graph.AddEdge("Estoque G", "Retirada I", 33)
	graph.AddEdge("Distribuidor H", "Filial J", 18)
	graph.AddEdge("Retirada I", "Hub K", 10)
	graph.AddEdge("Filial J", "Depósito L", 35)
	graph.AddEdge("Hub K", "Centro M", 16)
	graph.AddEdge("Depósito L", "Galpão N", 26)
	graph.AddEdge("Centro M", "Unidade O", 22)
	graph.AddEdge("Galpão N", "Remessa P", 17)
	graph.AddEdge("Unidade O", "CD Q", 31)
	graph.AddEdge("Remessa P", "Armazém R", 11)
	graph.AddEdge("CD Q", "Base S", 19)
	graph.AddEdge("Armazém R", "Estoque T", 24)
	graph.AddEdge("Base S", "Logística U", 13)
	graph.AddEdge("Estoque T", "Loja V", 45)
	graph.AddEdge("Logística U", "Centro W", 12)
	graph.AddEdge("Loja V", "Filial X", 36)
	graph.AddEdge("Centro W", "Depósito Y", 20)
	graph.AddEdge("Filial X", "Galpão Z", 14)
	graph.AddEdge("Depósito Y", "Posto AA", 38)
	graph.AddEdge("Galpão Z", "Unidade AB", 21)
	graph.AddEdge("Posto AA", "CD AC", 27)
	graph.AddEdge("Unidade AB", "Ponto AD", 29)
	graph.AddEdge("CD AC", "Hub AE", 18)

	router.Static("/static", "./static")
	router.LoadHTMLFiles("templates/mapa.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mapa.html", nil)
	})

	router.GET("/rota", func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")

		if from == "" || to == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros 'from' e 'to' são obrigatórios"})
			return
		}

		path, distance := Dijkstra(graph, from, to)
		if path == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Rota não encontrada"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"rota":      path,
			"distância": distance,
		})
	})

	router.Run(":80")
}
