package utils

import (
	"Personalised_Mailer/models"
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"text/template"

	"github.com/spf13/viper"
)

func LoadConfig() models.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	var cfg models.Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	var coordinators []models.Coordinator
	if cfg.Club.Cordinator1 != "" {
		coordinators = append(coordinators, models.Coordinator{cfg.Club.Cordinator1, cfg.Club.CordinatorEmail1})
	}
	if cfg.Club.Cordinator2 != "" {
		coordinators = append(coordinators, models.Coordinator{cfg.Club.Cordinator2, cfg.Club.CordinatorEmail2})
	}
	if cfg.Club.Cordinator3 != "" {
		coordinators = append(coordinators, models.Coordinator{cfg.Club.Cordinator3, cfg.Club.CordinatorEmail3})
	}
	if cfg.Club.Cordinator4 != "" {
		coordinators = append(coordinators, models.Coordinator{cfg.Club.Cordinator4, cfg.Club.CordinatorEmail4})
	}
	cfg.Club.Coordinators = coordinators

	return cfg
}

func ReadEmailsFromCSV(filename string) ([]models.Recipient, error) {
	var recipients []models.Recipient

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for i, record := range records {
		if i == 0 {
			continue // Skip the header row
		}
		if len(record) > 1 {
			recipients = append(recipients, models.Recipient{
				Name:  record[0],
				Email: record[1],
			})
		}
	}

	return recipients, nil
}

func BuildEmailHTML(cfg models.Config, recipientName string) string {
	// BuildWithLove := "Akshat23"
	// By := BuildWithLove + " : SPO WEB HEAD"
	var body bytes.Buffer
	tmpl := template.Must(template.ParseFiles("template/template.html"))
	data := struct {
		ClubName      string
		Name          string
		ProjectName   string
		WhatsappLink  string
		Mentors       []string
		ClubSite      string
		ClubDiscord   string
		ClubEmail     string
		ClubInstagram string
		TemplateText  string
		Coordinators  []models.Coordinator
	}{
		ClubName:      cfg.Club.Name,
		Name:          recipientName,
		ProjectName:   cfg.Project.Name,
		WhatsappLink:  cfg.Project.WhatsappLink,
		Mentors:       cfg.Project.Mentors,
		ClubSite:      cfg.Club.Site,
		ClubDiscord:   cfg.Club.Discord,
		ClubEmail:     cfg.Club.Email,
		TemplateText:  cfg.Club.TemplateText,
		ClubInstagram: cfg.Club.Instagram,
		Coordinators:  cfg.Club.Coordinators,
	}

	err := tmpl.Execute(&body, data)
	if err != nil {
		log.Fatalf("Template exec error: %v", err)
	}
	return body.String()
}
