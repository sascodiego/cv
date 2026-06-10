package main

import (
	"testing"
	"time"
)

func parseDate(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}

func TestCalculateMonths(t *testing.T) {
	tests := []struct {
		start string
		end   string
		want  int
	}{
		{"2019-01-01", "2019-01-31", 1},
		{"2019-01-01", "2019-12-31", 12},
		{"2019-01-01", "2020-01-01", 13}, // Ene 2019 a Ene 2020 inclusive es 13 meses
		{"2020-01-15", "2020-01-16", 1},  // Mismo mes
		{"2020-05-01", "2020-04-01", 0},  // Fecha de inicio después de fin
	}

	for _, tt := range tests {
		start := parseDate(tt.start)
		end := parseDate(tt.end)
		got := CalculateMonths(start, end)
		if got != tt.want {
			t.Errorf("CalculateMonths(%s, %s) = %d; want %d", tt.start, tt.end, got, tt.want)
		}
	}
}

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		want      []Interval
	}{
		{
			name: "Sin solapamiento",
			intervals: []Interval{
				{Start: parseDate("2019-01-01"), End: parseDate("2019-06-30")},
				{Start: parseDate("2020-01-01"), End: parseDate("2020-06-30")},
			},
			want: []Interval{
				{Start: parseDate("2019-01-01"), End: parseDate("2019-06-30")},
				{Start: parseDate("2020-01-01"), End: parseDate("2020-06-30")},
			},
		},
		{
			name: "Solapamiento parcial",
			intervals: []Interval{
				{Start: parseDate("2019-01-01"), End: parseDate("2019-12-31")},
				{Start: parseDate("2019-06-01"), End: parseDate("2020-06-30")},
			},
			want: []Interval{
				{Start: parseDate("2019-01-01"), End: parseDate("2020-06-30")},
			},
		},
		{
			name: "Totalmente anidado",
			intervals: []Interval{
				{Start: parseDate("2019-01-01"), End: parseDate("2019-12-31")},
				{Start: parseDate("2019-03-01"), End: parseDate("2019-09-30")},
			},
			want: []Interval{
				{Start: parseDate("2019-01-01"), End: parseDate("2019-12-31")},
			},
		},
		{
			name: "Continuos en el mismo mes",
			intervals: []Interval{
				{Start: parseDate("2019-01-01"), End: parseDate("2019-06-30")},
				{Start: parseDate("2019-06-15"), End: parseDate("2019-12-31")},
			},
			want: []Interval{
				{Start: parseDate("2019-01-01"), End: parseDate("2019-12-31")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeIntervals(tt.intervals)
			if len(got) != len(tt.want) {
				t.Fatalf("Longitud del resultado incorrecta. Got %d; want %d", len(got), len(tt.want))
			}
			for i := range got {
				if !got[i].Start.Equal(tt.want[i].Start) || !got[i].End.Equal(tt.want[i].End) {
					t.Errorf("Intervalo [%d] incorrecto. Got [%s, %s]; want [%s, %s]",
						i, got[i].Start.Format("2006-01-02"), got[i].End.Format("2006-01-02"),
						tt.want[i].Start.Format("2006-01-02"), tt.want[i].End.Format("2006-01-02"))
				}
			}
		})
	}
}

func TestAboutMePassThrough(t *testing.T) {
	raw := &CVRaw{
		PersonalInfo: PersonalInfo{
			Name:    "Test",
			Title:   "Dev",
			AboutMe: "This is my professional summary.",
		},
	}
	processed := ProcessCV(raw)
	if processed.PersonalInfo.AboutMe != "This is my professional summary." {
		t.Errorf("AboutMe passthrough failed: got %q", processed.PersonalInfo.AboutMe)
	}
}

func TestAboutMeEmptyString(t *testing.T) {
	raw := &CVRaw{
		PersonalInfo: PersonalInfo{
			Name:  "Test",
			Title: "Dev",
		},
	}
	processed := ProcessCV(raw)
	if processed.PersonalInfo.AboutMe != "" {
		t.Errorf("Empty AboutMe should produce empty string, got %q", processed.PersonalInfo.AboutMe)
	}
}

func TestHiddenProjectFiltered(t *testing.T) {
	skillGo := SkillRaw{ID: "golang", Name: "Go", Category: "Languages"}
	raw := &CVRaw{
		PersonalInfo: PersonalInfo{Name: "Test"},
		Skills:       []SkillRaw{skillGo},
		WorkExperience: []WorkExperience{
			{
				Company:      "Corp",
				Role:         "Dev",
				StartDate:    CVDate{Time: parseDate("2020-01-01")},
				EndDate:      CVDate{Time: parseDate("2020-12-31")},
				Technologies: []string{"golang"},
				Projects:     []string{"vis-proj", "hid-proj"},
			},
		},
		Projects: []Project{
			{
				ID:           "vis-proj",
				Name:         "Visible Project",
				StartDate:    CVDate{Time: parseDate("2020-01-01")},
				EndDate:      CVDate{Time: parseDate("2020-06-30")},
				Technologies: []string{"golang"},
				VisibleWeb:   true,
			},
			{
				ID:           "hid-proj",
				Name:         "Hidden Project",
				StartDate:    CVDate{Time: parseDate("2020-07-01")},
				EndDate:      CVDate{Time: parseDate("2020-12-31")},
				Technologies: []string{"golang"},
			},
		},
	}
	processed := ProcessCV(raw)

	// Only 1 project should render (the visible one)
	if len(processed.WorkExperience) != 1 {
		t.Fatalf("Expected 1 work entry, got %d", len(processed.WorkExperience))
	}
	if len(processed.WorkExperience[0].Projects) != 1 {
		t.Errorf("Expected 1 rendered project, got %d", len(processed.WorkExperience[0].Projects))
	}
	if processed.WorkExperience[0].Projects[0].Name != "Visible Project" {
		t.Errorf("Expected 'Visible Project', got %q", processed.WorkExperience[0].Projects[0].Name)
	}

	// Skill duration should include both visible + hidden intervals (12 months total)
	if len(processed.Skills) != 1 {
		t.Fatalf("Expected 1 skill, got %d", len(processed.Skills))
	}
	if processed.Skills[0].MonthsExperience != 12 {
		t.Errorf("Expected 12 months for golang (both projects), got %d", processed.Skills[0].MonthsExperience)
	}
}

func TestHiddenIndependentProject(t *testing.T) {
	skillGo := SkillRaw{ID: "golang", Name: "Go", Category: "Languages"}
	raw := &CVRaw{
		PersonalInfo: PersonalInfo{Name: "Test"},
		Skills:       []SkillRaw{skillGo},
		WorkExperience: []WorkExperience{
			{
				Company:      "Corp",
				Role:         "Dev",
				StartDate:    CVDate{Time: parseDate("2020-01-01")},
				EndDate:      CVDate{Time: parseDate("2020-06-30")},
				Technologies: []string{"golang"},
			},
		},
		Projects: []Project{
			{
				ID:           "hidden-indie",
				Name:         "Hidden Indie",
				StartDate:    CVDate{Time: parseDate("2021-01-01")},
				EndDate:      CVDate{Time: parseDate("2021-06-30")},
				Technologies: []string{"golang"},
			},
		},
	}
	processed := ProcessCV(raw)

	// No independent projects should render
	if len(processed.Projects) != 0 {
		t.Errorf("Expected 0 rendered independent projects, got %d", len(processed.Projects))
	}

	// Skill duration should include work (6 months) + hidden indie (6 months) = 12 months
	if len(processed.Skills) != 1 {
		t.Fatalf("Expected 1 skill, got %d", len(processed.Skills))
	}
}

func TestProcessCV(t *testing.T) {
	// Crear un caso ficticio donde la tecnología "csharp" se repite:
	// Trabajo 1: 2019-01-01 a 2019-12-31 (12 meses) -> usa csharp
	// Trabajo 2: 2020-06-01 a 2021-06-01 (13 meses) -> usa csharp en un proyecto de 6 meses (2020-06-01 a 2020-11-30)
	// Total meses esperados para csharp: 12 + 6 = 18 meses (1 año y 6 meses)

	skillCsharp := SkillRaw{ID: "csharp", Name: "C#", Category: "Backend"}

	raw := &CVRaw{
		PersonalInfo: PersonalInfo{Name: "Test Developer"},
		Skills:       []SkillRaw{skillCsharp},
		WorkExperience: []WorkExperience{
			{
				Company:      "Company 1",
				Role:         "Dev",
				StartDate:    CVDate{Time: parseDate("2019-01-01")},
				EndDate:      CVDate{Time: parseDate("2019-12-31")},
				Technologies: []string{"csharp"},
			},
			{
				Company:      "Company 2",
				Role:         "Dev",
				StartDate:    CVDate{Time: parseDate("2020-06-01")},
				EndDate:      CVDate{Time: parseDate("2021-06-01")},
				Technologies: []string{"golang"}, // csharp no está global
				Projects:     []string{"proj-a"},
			},
		},
		Projects: []Project{
			{
				ID:           "proj-a",
				Name:         "Project A",
				StartDate:    CVDate{Time: parseDate("2020-06-01")},
				EndDate:      CVDate{Time: parseDate("2020-11-30")},
				Technologies: []string{"csharp"},
			},
		},
	}

	processed := ProcessCV(raw)

	if len(processed.Skills) != 1 {
		t.Fatalf("Se esperaba 1 skill procesado, se obtuvieron %d", len(processed.Skills))
	}

	csharpSkill := processed.Skills[0]
	if csharpSkill.ID != "csharp" {
		t.Fatalf("ID de skill inesperado: %s", csharpSkill.ID)
	}

	expectedMonths := 18 // 12 meses del primer trabajo + 6 meses del proyecto en el segundo
	if csharpSkill.MonthsExperience != expectedMonths {
		t.Errorf("Se esperaban %d meses de experiencia en csharp, se obtuvieron %d", expectedMonths, csharpSkill.MonthsExperience)
	}

	expectedText := "1 año y 6 meses"
	if csharpSkill.ExperienceText != expectedText {
		t.Errorf("Texto de experiencia inesperado. Got %q, want %q", csharpSkill.ExperienceText, expectedText)
	}
}
