package main

import (
	"encoding/json"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// CVDate es un tipo personalizado para manejar fechas del CV que pueden ser YYYY-MM-DD o "present".
type CVDate struct {
	time.Time
	IsPresent bool
}

// UnmarshalYAML implementa la interfaz yaml.Unmarshaler
func (d *CVDate) UnmarshalYAML(value *yaml.Node) error {
	str := strings.TrimSpace(value.Value)
	if strings.ToLower(str) == "present" || str == "" {
		d.IsPresent = true
		d.Time = time.Now()
		return nil
	}

	// Soportamos YYYY-MM-DD y YYYY-MM
	layouts := []string{"2006-01-02", "2006-01"}
	var err error
	var t time.Time
	for _, layout := range layouts {
		t, err = time.Parse(layout, str)
		if err == nil {
			d.Time = t
			d.IsPresent = false
			return nil
		}
	}
	return err
}

// MarshalJSON implementa la interfaz json.Marshaler
func (d CVDate) MarshalJSON() ([]byte, error) {
	if d.IsPresent {
		return json.Marshal("Presente")
	}
	return json.Marshal(d.Format("2006-01-02"))
}

// FormatYear devuelve solo el año o "Presente"
func (d CVDate) FormatYear() string {
	if d.IsPresent {
		return "Presente"
	}
	return d.Format("2006")
}

// FormatMonthYear devuelve "Ene 2006" o "Presente"
func (d CVDate) FormatMonthYear() string {
	if d.IsPresent {
		return "Presente"
	}
	// Mapa en español rápido para nombres de meses
	meses := map[string]string{
		"January": "Ene", "February": "Feb", "March": "Mar", "April": "Abr",
		"May": "May", "June": "Jun", "July": "Jul", "August": "Ago",
		"September": "Sep", "October": "Oct", "November": "Nov", "December": "Dic",
	}
	engMonth := d.Format("January")
	esMonth := meses[engMonth]
	return esMonth + " " + d.Format("2006")
}

// CVRaw es la estructura original que leemos del YAML
type CVRaw struct {
	PersonalInfo    PersonalInfo     `yaml:"personal_info" json:"personal_info"`
	Skills          []SkillRaw       `yaml:"skills" json:"skills"`
	WorkExperience  []WorkExperience `yaml:"work_experience" json:"work_experience"`
	Education       []Education      `yaml:"education" json:"education"`
	Projects        []Project        `yaml:"projects" json:"projects"`
	Recommendations []Recommendation `yaml:"recommendations" json:"recommendations"`
}

type Recommendation struct {
	Author   string `yaml:"author" json:"author"`
	Role     string `yaml:"role" json:"role"`
	Relation string `yaml:"relation" json:"relation"`
	Text     string `yaml:"text" json:"text"`
}

type PersonalInfo struct {
	Name     string `yaml:"name" json:"name"`
	Title    string `yaml:"title" json:"title"`
	AboutMe  string `yaml:"about_me" json:"about_me,omitempty"`
	Email    string `yaml:"email" json:"email,omitempty"`
	Website  string `yaml:"website" json:"website,omitempty"`
	GitHub   string `yaml:"github" json:"github,omitempty"`
	LinkedIn string `yaml:"linkedin" json:"linkedin,omitempty"`
}

type SkillRaw struct {
	ID       string `yaml:"id" json:"id"`
	Name     string `yaml:"name" json:"name"`
	Category string `yaml:"category" json:"category"`
}

type WorkExperience struct {
	Company      string   `yaml:"company" json:"company"`
	Role         string   `yaml:"role" json:"role"`
	StartDate    CVDate   `yaml:"start_date" json:"start_date"`
	EndDate      CVDate   `yaml:"end_date" json:"end_date"`
	Description  string   `yaml:"description" json:"description"`
	Technologies []string `yaml:"technologies" json:"technologies"`
	Projects     []string `yaml:"projects" json:"projects"` // Lista de IDs de proyectos globales
}

type Education struct {
	Institution string `yaml:"institution" json:"institution"`
	Degree      string `yaml:"degree" json:"degree"`
	StartDate   CVDate `yaml:"start_date" json:"start_date"`
	EndDate     CVDate `yaml:"end_date" json:"end_date"`
	Status      string `yaml:"status" json:"status"`
}

type Project struct {
	ID           string   `yaml:"id" json:"id"`
	Name         string   `yaml:"name" json:"name"`
	Description  string   `yaml:"description" json:"description"`
	StartDate    CVDate   `yaml:"start_date" json:"start_date"`
	EndDate      CVDate   `yaml:"end_date" json:"end_date"`
	Technologies []string `yaml:"technologies" json:"technologies"`
	URL          string   `yaml:"url" json:"url"`
	VisibleWeb   bool     `yaml:"visible_web" json:"-"`
	VisiblePDF   bool     `yaml:"visible_pdf" json:"-"`
}

// CVProcessed es la estructura final calculada para Astro
type CVProcessed struct {
	PersonalInfo    PersonalInfo              `json:"personal_info"`
	Skills          []SkillProcessed          `json:"skills"`
	WorkExperience  []WorkExperienceProcessed `json:"work_experience"`
	Education       []EducationProcessed      `json:"education"`
	Projects        []ProjectProcessed        `json:"projects"`
	Recommendations []Recommendation          `json:"recommendations"`
}

type SkillProcessed struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Category         string  `json:"category"`
	MonthsExperience int     `json:"months_experience"`
	YearsExperience  float64 `json:"years_experience"`
	ExperienceText   string  `json:"experience_text"`
}

type WorkExperienceProcessed struct {
	Company      string             `json:"company"`
	Role         string             `json:"role"`
	StartDate    string             `json:"start_date"`
	EndDate      string             `json:"end_date"`
	DurationText string             `json:"duration_text"`
	PeriodText   string             `json:"period_text"` // Ej: "2019 - 2021" o "Ene 2019 - Presente"
	Description  string             `json:"description"`
	Technologies []string           `json:"technologies"`
	Projects     []ProjectProcessed `json:"projects"`
}

type EducationProcessed struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	PeriodText  string `json:"period_text"`
	Status      string `json:"status"`
}

type ProjectProcessed struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	PeriodText   string   `json:"period_text"`
	DurationText string   `json:"duration_text"`
	Technologies []string `json:"technologies"`
	URL          string   `json:"url"`
}
