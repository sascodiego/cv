package main

import (
	"fmt"
	"sort"
	"time"
)

// Interval representa un rango de fechas.
type Interval struct {
	Start time.Time
	End   time.Time
}

// CalculateMonths calcula la cantidad inclusiva de meses entre dos fechas.
// Por ejemplo, de 2019-01-01 a 2019-01-31 es 1 mes.
func CalculateMonths(start, end time.Time) int {
	if start.After(end) {
		return 0
	}
	months := (end.Year()-start.Year())*12 + int(end.Month()-start.Month()) + 1
	if months < 0 {
		return 0
	}
	return months
}

// FormatDuration convierte un número de meses en una cadena legible (ej: "2 años y 3 meses" o "1 año").
func FormatDuration(months int) string {
	if months <= 0 {
		return "Sin experiencia"
	}
	years := months / 12
	remMonths := months % 12

	if years > 0 {
		var yearText string
		if years == 1 {
			yearText = "1 año"
		} else {
			yearText = fmt.Sprintf("%d años", years)
		}

		if remMonths > 0 {
			var monthText string
			if remMonths == 1 {
				monthText = "1 mes"
			} else {
				monthText = fmt.Sprintf("%d meses", remMonths)
			}
			return fmt.Sprintf("%s y %s", yearText, monthText)
		}
		return yearText
	}

	if remMonths == 1 {
		return "1 mes"
	}
	return fmt.Sprintf("%d meses", remMonths)
}

// MergeIntervals fusiona intervalos de fechas solapados.
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// Ordenar por fecha de inicio
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start.Before(intervals[j].Start)
	})

	merged := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := &merged[len(merged)-1]

		// Si el inicio del actual está antes o coincide con el fin del último
		// o cae dentro del mismo mes/año (para evitar huecos menores a un mes)
		if current.Start.Before(last.End) || current.Start.Equal(last.End) ||
			(current.Start.Year() == last.End.Year() && current.Start.Month() == last.End.Month()) {
			if current.End.After(last.End) {
				last.End = current.End
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

// ProcessCV ejecuta los cálculos del pipeline sobre el CV original
func ProcessCV(raw *CVRaw) *CVProcessed {
	processed := &CVProcessed{
		PersonalInfo: PersonalInfo{
			Name:    raw.PersonalInfo.Name,
			Title:   raw.PersonalInfo.Title,
			AboutMe: raw.PersonalInfo.AboutMe,
			// Contact fields (Email, Website, GitHub, LinkedIn) are deliberately
			// excluded from the public output. They are injected at PDF generation
			// time by Puppeteer from .env — never stored in the repo.
		},
		Skills:          []SkillProcessed{},
		Education:       []EducationProcessed{},
		Projects:        []ProjectProcessed{},
		Recommendations: raw.Recommendations,
	}
	if processed.Recommendations == nil {
		processed.Recommendations = []Recommendation{}
	}

	// Mapear proyectos globales por su ID para resolución relacional
	projectsMap := make(map[string]Project)
	for _, proj := range raw.Projects {
		projectsMap[proj.ID] = proj
	}

	// Registrar qué proyectos fueron asociados a alguna experiencia laboral
	associatedProjects := make(map[string]bool)

	// Mapear tecnologías a sus intervalos para calcular duraciones acumuladas
	techIntervals := make(map[string][]Interval)

	// Procesar Experiencia Laboral
	for _, work := range raw.WorkExperience {
		workStart := work.StartDate.Time
		workEnd := work.EndDate.Time

		// Período general de este trabajo
		periodText := fmt.Sprintf("%s - %s", work.StartDate.FormatMonthYear(), work.EndDate.FormatMonthYear())
		if work.StartDate.FormatYear() == work.EndDate.FormatYear() {
			if work.EndDate.IsPresent {
				periodText = fmt.Sprintf("%s - Presente", work.StartDate.FormatMonthYear())
			}
		}

		months := CalculateMonths(workStart, workEnd)
		durationText := FormatDuration(months)

		// Mapear y procesar proyectos internos asociados a esta experiencia
		var processedProjects []ProjectProcessed
		for _, projID := range work.Projects {
			proj, found := projectsMap[projID]
			if !found {
				fmt.Printf("⚠️ Advertencia: Proyecto '%s' referenciado en '%s' no fue encontrado en la lista global de proyectos.\n", projID, work.Company)
				continue
			}
			associatedProjects[projID] = true
			pStart := proj.StartDate.Time
			pEnd := proj.EndDate.Time
			pMonths := CalculateMonths(pStart, pEnd)

			pPeriodText := fmt.Sprintf("%s - %s", proj.StartDate.FormatMonthYear(), proj.EndDate.FormatMonthYear())
			if proj.EndDate.IsPresent {
				pPeriodText = fmt.Sprintf("%s - Presente", proj.StartDate.FormatMonthYear())
			}

			// Registrar las tecnologías en base a los intervalos de este proyecto
			for _, tech := range proj.Technologies {
				techIntervals[tech] = append(techIntervals[tech], Interval{Start: pStart, End: pEnd})
			}

			// Only include visible projects in rendered output
			if proj.VisibleWeb {
				processedProjects = append(processedProjects, ProjectProcessed{
					Name:         proj.Name,
					Description:  proj.Description,
					StartDate:    pStart.Format("2006-01-02"),
					EndDate:      pEnd.Format("2006-01-02"),
					PeriodText:   pPeriodText,
					DurationText: FormatDuration(pMonths),
					Technologies: proj.Technologies,
					URL:          proj.URL,
				})
			}
		}

		// Si el trabajo no tiene proyectos específicos, asignamos las tecnologías del trabajo a todo el período del trabajo
		if len(work.Projects) == 0 {
			for _, tech := range work.Technologies {
				techIntervals[tech] = append(techIntervals[tech], Interval{Start: workStart, End: workEnd})
			}
		} else {
			// También, si hay tecnologías declaradas globalmente en el trabajo pero no en los proyectos asociados,
			// las asociamos a todo el período del trabajo
			projTechs := make(map[string]bool)
			for _, projID := range work.Projects {
				if proj, found := projectsMap[projID]; found {
					for _, tech := range proj.Technologies {
						projTechs[tech] = true
					}
				}
			}
			for _, tech := range work.Technologies {
				if !projTechs[tech] {
					techIntervals[tech] = append(techIntervals[tech], Interval{Start: workStart, End: workEnd})
				}
			}
		}

		processed.WorkExperience = append(processed.WorkExperience, WorkExperienceProcessed{
			Company:      work.Company,
			Role:         work.Role,
			StartDate:    work.StartDate.Time.Format("2006-01-02"),
			EndDate:      work.EndDate.Time.Format("2006-01-02"),
			DurationText: durationText,
			PeriodText:   periodText,
			Description:  work.Description,
			Technologies: work.Technologies,
			Projects:     processedProjects,
		})
	}

	// Register tech intervals for non-visible independent projects (they still contribute to skill durations)
	for _, proj := range raw.Projects {
		if associatedProjects[proj.ID] || proj.VisibleWeb {
			continue
		}
		pStart := proj.StartDate.Time
		pEnd := proj.EndDate.Time
		for _, tech := range proj.Technologies {
			techIntervals[tech] = append(techIntervals[tech], Interval{Start: pStart, End: pEnd})
		}
	}

	// Procesar Proyectos Independientes (solo los que NO fueron asociados a experiencias de trabajo)
	for _, proj := range raw.Projects {
		if associatedProjects[proj.ID] {
			continue // Ya fue renderizado dentro de una compañía
		}
		if !proj.VisibleWeb {
			continue // Non-visible independent projects are not rendered
		}

		pStart := proj.StartDate.Time
		pEnd := proj.EndDate.Time
		pMonths := CalculateMonths(pStart, pEnd)

		// Registrar tecnologías del proyecto independiente
		for _, tech := range proj.Technologies {
			techIntervals[tech] = append(techIntervals[tech], Interval{Start: pStart, End: pEnd})
		}

		periodText := fmt.Sprintf("%s - %s", proj.StartDate.FormatMonthYear(), proj.EndDate.FormatMonthYear())
		if proj.EndDate.IsPresent {
			periodText = fmt.Sprintf("%s - Presente", proj.StartDate.FormatMonthYear())
		}

		processed.Projects = append(processed.Projects, ProjectProcessed{
			Name:         proj.Name,
			Description:  proj.Description,
			StartDate:    pStart.Format("2006-01-02"),
			EndDate:      pEnd.Format("2006-01-02"),
			PeriodText:   periodText,
			DurationText: FormatDuration(pMonths),
			Technologies: proj.Technologies,
			URL:          proj.URL,
		})
	}

	// Procesar Educación
	for _, edu := range raw.Education {
		periodText := fmt.Sprintf("%s - %s", edu.StartDate.FormatMonthYear(), edu.EndDate.FormatMonthYear())
		processed.Education = append(processed.Education, EducationProcessed{
			Institution: edu.Institution,
			Degree:      edu.Degree,
			StartDate:   edu.StartDate.Time.Format("2006-01-02"),
			EndDate:     edu.EndDate.Time.Format("2006-01-02"),
			PeriodText:  periodText,
			Status:      edu.Status,
		})
	}

	// Calcular tiempo real acumulado para cada Skill
	for _, skill := range raw.Skills {
		intervals, exists := techIntervals[skill.ID]
		var totalMonths int
		if exists {
			merged := MergeIntervals(intervals)
			for _, iv := range merged {
				totalMonths += CalculateMonths(iv.Start, iv.End)
			}
		}

		years := float64(totalMonths) / 12.0
		experienceText := FormatDuration(totalMonths)

		processed.Skills = append(processed.Skills, SkillProcessed{
			ID:               skill.ID,
			Name:             skill.Name,
			Category:         skill.Category,
			MonthsExperience: totalMonths,
			YearsExperience:  years,
			ExperienceText:   experienceText,
		})
	}

	return processed
}
