package jira

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Estructura para extraer el estado del ticket
type IssueResponse struct {
	Fields struct {
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
	} `json:"fields"`
}

func CheckState() {
	// 🔧 Configura tus datos
	jiraDomain := ""
	issueKey := "AE-125"
	email := ""
	apiToken := ""

	// 📦 Armar la URL
	url := fmt.Sprintf("%s/rest/api/3/issue/%s", jiraDomain, issueKey)

	// 🔐 Crear encabezado Authorization
	authString := fmt.Sprintf("%s:%s", email, apiToken)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))

	// 📨 Crear solicitud HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creando solicitud:", err)
		os.Exit(1)
	}

	req.Header.Add("Authorization", "Basic "+encodedAuth)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Error: código %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		os.Exit(1)
	}

	// 📥 Leer y parsear la respuesta
	body, _ := io.ReadAll(resp.Body)
	var issue IssueResponse
	json.Unmarshal(body, &issue)

	fmt.Printf("El estado del ticket %s es: %s\n", issueKey, issue.Fields.Status.Name)
}
