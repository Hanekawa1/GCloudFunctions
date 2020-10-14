package automovelgooglecloud

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"net/http"
	"os"
)

type Automovel struct {
	Placa  string `json:"placa"`
	Cor    string `json:"cor"`
	Preco  string `json:"preco"`
	Modelo string `json:"modelo"`
	Marca  string `json:"marca"`
}

var ctx context.Context

func InsereDadosAutomovel(w http.ResponseWriter, r *http.Request) {
	ctx = r.Context()
	var automovel Automovel

	automovel.Placa = r.URL.Query().Get("placa")
	if len(automovel.Placa) == 0 {
		w.WriteHeader(400)
		return
	}
	automovel.Cor = r.URL.Query().Get("cor")
	if len(automovel.Cor) == 0 {
		w.WriteHeader(400)
		return
	}
	automovel.Preco = r.URL.Query().Get("preco")
	if len(automovel.Preco) == 0 {
		w.WriteHeader(400)
		return
	}
	automovel.Modelo = r.URL.Query().Get("modelo")
	if len(automovel.Modelo) == 0 {
		w.WriteHeader(400)
		return
	}
	automovel.Marca = r.URL.Query().Get("marca")
	if len(automovel.Marca) == 0 {
		w.WriteHeader(400)
		return
	}
	log.Printf("Automovel: %v", automovel)

	err := insereAutomovelNoFireStore(automovel, ctx)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	log.Println("Inserido com sucesso")
	w.WriteHeader(200)
}

func insereAutomovelNoFireStore(automovel Automovel, contexto context.Context) error {
	client, err := firestore.NewClient(contexto, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Println("Erro ao criar o cliente do firestore: %v", err)
	}
	_, _, err = client.Collection("automoveis").Add(contexto, map[string]interface{}{
		"Placa":  automovel.Placa,
		"Cor":    automovel.Cor,
		"Preco":  automovel.Preco,
		"Modelo": automovel.Modelo,
		"Marca":  automovel.Marca,
	})
	if err != nil {
		log.Fatalf("Falha ao inserir dados: %v", err)
	}

	if err != nil {
		log.Println("Erro ao inserir no firestore: %v", err)
	}
	return nil
}
