package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Advanced
	VeryAdvanced
	Extreme
)

var difficultyCharsets = map[Difficulty]string{
	Easy:         "abcdefghijklmnopqrstuvwxyz0123456789",
	Advanced:     "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	VeryAdvanced: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()",
	Extreme:      "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+[]{}|;:,.<>?/~`",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Configuración para servir archivos estáticos
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handleRequest)

	log.Println("Servidor web iniciado en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()
		lengthStr := r.FormValue("length")
		difficultyStr := r.FormValue("difficulty")

		length, err := strconv.Atoi(lengthStr)
		if err != nil || length < 6 {
			length = 6
		}

		difficulty, err := parseDifficulty(difficultyStr)
		if err != nil {
			difficulty = Easy
		}

		password, err := generatePassword(length, difficulty)
		if err != nil {
			http.Error(w, "Error al generar la contraseña", http.StatusInternalServerError)
			return
		}

		data := struct {
			Password   string
			Length     int
			Difficulty string
		}{
			Password:   password,
			Length:     length,
			Difficulty: difficultyStr,
		}

		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func parseDifficulty(difficultyStr string) (Difficulty, error) {
	switch difficultyStr {
	case "easy":
		return Easy, nil
	case "advanced":
		return Advanced, nil
	case "very_advanced":
		return VeryAdvanced, nil
	case "extreme":
		return Extreme, nil
	default:
		return Easy, fmt.Errorf("dificultad desconocida")
	}
}

func generatePassword(length int, difficulty Difficulty) (string, error) {
	if length < 6 {
		return "", fmt.Errorf("la longitud mínima es 6")
	}

	charset, ok := difficultyCharsets[difficulty]
	if !ok {
		return "", fmt.Errorf("dificultad desconocida")
	}

	var password []byte
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		password = append(password, charset[randomIndex])
	}

	return string(password), nil
}
