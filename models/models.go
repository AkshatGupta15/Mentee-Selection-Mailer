package models

type Config struct {
	Mail    MailConfig
	Club    ClubConfig
	Project ProjectConfig
}

type MailConfig struct {
	Host string
	Port string
	User string
	Pass string
}

type ClubConfig struct {
	Name             string
	Site             string
	Discord          string
	Instagram        string
	Email            string
	Cordinator1      string
	CordinatorEmail1 string
	Cordinator2      string
	CordinatorEmail2 string
	Cordinator3      string
	CordinatorEmail3 string
	Cordinator4      string
	CordinatorEmail4 string
	TemplateText     string

	Coordinators []Coordinator
}

type Coordinator struct {
	Name  string
	Email string
}

type ProjectConfig struct {
	Name         string
	WhatsappLink string
	Mentors      []string
}

type Recipient struct {
	Name  string
	Email string
}
