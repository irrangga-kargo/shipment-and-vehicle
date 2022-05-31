package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irrangga-kargo/kargo-trucks/graph"
	"github.com/irrangga-kargo/kargo-trucks/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := &graph.Resolver{}
	resolver.Init()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}
	f, err := os.Create("users.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	// Sender data.
	from := "irranggaa@gmail.com"
	password := os.Getenv("EMAILPASSWORD")
	// Receiver email address.
	to := []string{
		"irranggaa@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	// Message.
	message := []byte("This is a test email message.")
	// Authentication.
	auth := smtp.PlainAuth("", from, password,
		smtpHost)
	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth,
		from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
