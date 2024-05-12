package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Estructura para el cuerpo de la solicitud POST
type PublishRequest struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	// Configurar opciones de CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Permitir todas las solicitudes CORS desde cualquier origen
	r.Use(cors.New(config))

	// Ruta para la página HTML
	r.GET("/", serveStaticHTML)

	// Ruta para la API POST de publicación (requiere autenticación)
	r.POST("/publish", publishMessage)

	// Inicia el servidor en el puerto 8080
	log.Fatal(r.Run(":8080"))
}

// Función para servir una página HTML estática con un formulario
func serveStaticHTML(c *gin.Context) {
	html := `
	<!DOCTYPE html>
<html>
<head>
    <title>Enviar Mensaje a Pub/Sub</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            padding: 20px;
            background-color: #f8f8f8;
        }
        h1 {
            color: #333;
            text-align: center;
        }
        #publishForm {
            background-color: #fff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
            width: 80%;
            margin: 0 auto;
        }
        label {
            font-weight: bold;
        }
        textarea {
            width: 100%;
            padding: 10px;
            box-sizing: border-box;
            border: 1px solid #ccc;
            border-radius: 4px;
            resize: vertical;
        }
        input[type="submit"] {
            background-color: #4CAF50;
            color: white;
            padding: 10px 20px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            float: right;
        }
        input[type="submit"]:hover {
            background-color: #45a049;
        }
    </style>
</head>
<body>
    <h1>Enviar Mensaje a Pub/Sub</h1>
    <form id="publishForm">
        <label for="message">Mensaje:</label><br>
        <textarea id="message" name="message" rows="4" cols="50"></textarea><br><br>
        <input type="submit" value="Enviar">
    </form>

    <script>
        document.getElementById("publishForm").addEventListener("submit", function(event) {
            event.preventDefault();
            const formData = new FormData(event.target);
            const message = formData.get("message");

            fetch("/publish", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ message })
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Error al enviar la solicitud");
                }
                return response.json();
            })
            .then(data => {
                console.log(data);
                // TODO: manejar la respuesta del servidor aquí si es necesario.
            })
            .catch(error => {
                console.error("Error:", error);
            });
        });
    </script>
</body>
</html>

	`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

// Función para publicar un mensaje en Google Cloud Pub/Sub
func publishMessage(c *gin.Context) {
	var req PublishRequest

	// Decodifica el cuerpo JSON de la solicitud
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error al decodificar la solicitud JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud JSON"})
		return
	}

	// Validación de mensaje
	if req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere un mensaje en el cuerpo de la solicitud"})
		return
	}

	// Llama a la función publish con los datos de la solicitud
	if err := publish(c.Request.Context(), req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al publicar mensaje: %v", err)})
		return
	}

	// Responde con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Mensaje publicado correctamente"})
}

// Función para publicar un mensaje en Google Cloud Pub/Sub
func publish(ctx context.Context, msg string) error {
	// Reemplaza con tu lógica de publicación en Pub/Sub
	projectID := "tss-1s2024"
	topicID := "my-topic"

	// Imprime información sobre la publicación
	fmt.Printf("Publicando mensaje: %s\n", msg)

	// Crea un cliente de Google Cloud Pub/Sub
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %w", err)
	}
	defer client.Close()

	// Obtiene el tema
	t := client.Topic(topicID)

	// Publica un mensaje en el tema
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	// Espera el resultado de la publicación y obtiene el ID del mensaje
	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %w", err)
	}

	return nil
}
