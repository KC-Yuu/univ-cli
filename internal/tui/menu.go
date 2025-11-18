package tui

import (
	"fmt"
	"math/rand"
	"runtime"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"univ-cli/internal/custom"
)

type state int

const (
	stateMenu state = iota
	stateDateTime
	stateWelcome
	stateCustom
	stateQuote
	stateCalc
	stateSysinfo
	stateGame
	stateTheme
	stateQuit
)

type Theme struct {
	name       string
	gradient   []string
	primary    string
	selectedBg string
	subtitle   string
}

var themes = []Theme{
	{
		name:       "Bleu Océan",
		gradient:   []string{"#66B3FF", "#3399FF", "#0080FF", "#0066CC", "#004C99", "#003366"},
		primary:    "#0066CC",
		selectedBg: "#0066CC",
		subtitle:   "#666666",
	},
	{
		name:       "Vert Forêt",
		gradient:   []string{"#90EE90", "#66CC66", "#4CAF50", "#2E7D32", "#1B5E20", "#0D3D0D"},
		primary:    "#4CAF50",
		selectedBg: "#2E7D32",
		subtitle:   "#666666",
	},
	{
		name:       "Violet Galaxie",
		gradient:   []string{"#E1BEE7", "#BA68C8", "#9C27B0", "#7B1FA2", "#6A1B9A", "#4A148C"},
		primary:    "#9C27B0",
		selectedBg: "#7B1FA2",
		subtitle:   "#666666",
	},
	{
		name:       "Orange Sunset",
		gradient:   []string{"#FFCC80", "#FFB74D", "#FF9800", "#F57C00", "#E65100", "#BF360C"},
		primary:    "#FF9800",
		selectedBg: "#F57C00",
		subtitle:   "#666666",
	},
	{
		name:       "Rouge Cardinal",
		gradient:   []string{"#EF9A9A", "#E57373", "#F44336", "#D32F2F", "#C62828", "#8E0000"},
		primary:    "#F44336",
		selectedBg: "#D32F2F",
		subtitle:   "#666666",
	},
	{
		name:       "Cyan Arctique",
		gradient:   []string{"#B2EBF2", "#4DD0E1", "#00BCD4", "#0097A7", "#00838F", "#006064"},
		primary:    "#00BCD4",
		selectedBg: "#0097A7",
		subtitle:   "#666666",
	},
}

type Model struct {
	state         state
	cursor        int
	choices       []string
	customChoices []string
	gameNumber    int
	gameAttempts  int
	textInput     textinput.Model
	message       string
	currentTheme  int
	calcResult    string
}

var (
	menuStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			Italic(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))
)

func getTitleStyle(theme Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color(theme.selectedBg)).
		Padding(0, 2)
}

func getHeaderStyle(theme Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.primary)).
		Bold(true)
}

func getSelectedStyle(theme Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color(theme.selectedBg)).
		Bold(true)
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Entrez un nombre..."
	ti.CharLimit = 3
	ti.Width = 20

	return Model{
		state:         stateMenu,
		cursor:        0,
		choices:       []string{"Date et Heure", "Message de bienvenue", "Custom", "Changer de thème", "Quitter"},
		customChoices: []string{"Citation aléatoire", "Calculatrice", "Informations système", "Mini-Jeu", "Retour au menu"},
		textInput:     ti,
		currentTheme:  0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case stateMenu:
			return m.updateMenu(msg)
		case stateDateTime, stateWelcome:
			if msg.String() == "q" || msg.String() == "esc" {
				m.state = stateMenu
				m.message = ""
			}
		case stateCustom:
			return m.updateCustomMenu(msg)
		case stateQuote, stateSysinfo:
			if msg.String() == "q" || msg.String() == "esc" {
				m.state = stateCustom
				m.message = ""
			}
		case stateCalc:
			return m.updateCalc(msg)
		case stateGame:
			return m.updateGame(msg)
		case stateTheme:
			return m.updateTheme(msg)
		case stateQuit:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) updateTheme(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.state = stateMenu
		return m, nil

	case "up", "k":
		if m.currentTheme > 0 {
			m.currentTheme--
		} else {
			m.currentTheme = len(themes) - 1
		}

	case "down", "j":
		if m.currentTheme < len(themes)-1 {
			m.currentTheme++
		} else {
			m.currentTheme = 0
		}

	case "enter":
		m.state = stateMenu
		return m, nil
	}
	return m, nil
}

func (m Model) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = stateQuit
		return m, tea.Quit

	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}

	case "enter":
		switch m.cursor {
		case 0: 
			m.state = stateDateTime
		case 1: 
			m.state = stateWelcome
		case 2: 
			m.state = stateCustom
			m.cursor = 0
		case 3: 
			m.state = stateTheme
		case 4: 
			m.state = stateQuit
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) updateCustomMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.state = stateMenu
		m.cursor = 0
		return m, nil

	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.customChoices)-1 {
			m.cursor++
		}

	case "enter":
		switch m.cursor {
		case 0: 
			m.state = stateQuote
		case 1: 
			m.state = stateCalc
			m.textInput.SetValue("")
			m.textInput.Focus()
			m.textInput.Placeholder = "Ex: 42 + 8"
			m.calcResult = ""
		case 2: 
			m.state = stateSysinfo
		case 3: 
			m.state = stateGame
			m.gameNumber = rand.Intn(100) + 1
			m.gameAttempts = 0
			m.textInput.SetValue("")
			m.textInput.Focus()
			m.textInput.Placeholder = "Entrez un nombre..."
			m.message = ""
		case 4: 
			m.state = stateMenu
			m.cursor = 0
		}
	}
	return m, nil
}

func (m Model) updateCalc(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc", "q":
		m.state = stateCustom
		m.calcResult = ""
		m.textInput.Blur()
		m.cursor = 0
		return m, nil

	case "enter":
		expression := m.textInput.Value()
		if expression == "" {
			return m, nil
		}

		result, err := custom.Calculate(expression)
		if err != nil {
			m.calcResult = fmt.Sprintf("ERREUR: %v", err)
		} else {
			m.calcResult = fmt.Sprintf("Résultat: %.2f", result)
		}

		m.textInput.SetValue("")
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) updateGame(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc", "q":
		m.state = stateCustom
		m.message = ""
		m.textInput.Blur()
		m.cursor = 0
		return m, nil

	case "enter":
		guess := m.textInput.Value()
		if guess == "" {
			return m, nil
		}

		var num int
		_, err := fmt.Sscanf(guess, "%d", &num)
		if err != nil {
			m.message = "ERREUR: Veuillez entrer un nombre valide."
			m.textInput.SetValue("")
			return m, nil
		}

		m.gameAttempts++

		if num == m.gameNumber {
			m.message = fmt.Sprintf("GAGNÉ ! Le nombre était %d. Trouvé en %d tentative(s).", m.gameNumber, m.gameAttempts)
		} else if num < m.gameNumber {
			m.message = fmt.Sprintf("Trop petit. Tentative #%d", m.gameAttempts)
		} else {
			m.message = fmt.Sprintf("Trop grand. Tentative #%d", m.gameAttempts)
		}

		m.textInput.SetValue("")
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	switch m.state {
	case stateMenu:
		return m.viewMenu()
	case stateDateTime:
		return m.viewDateTime()
	case stateWelcome:
		return m.viewWelcome()
	case stateCustom:
		return m.viewCustomMenu()
	case stateQuote:
		return m.viewQuote()
	case stateCalc:
		return m.viewCalc()
	case stateSysinfo:
		return m.viewSysinfo()
	case stateGame:
		return m.viewGame()
	case stateTheme:
		return m.viewTheme()
	case stateQuit:
		return "\nFermeture de l'application...\n\n"
	}
	return ""
}

func (m Model) viewMenu() string {
	theme := themes[m.currentTheme]

	logoLines := []string{
		"██╗   ██╗███╗   ██╗██╗██╗   ██╗      ██████╗██╗     ██╗",
		"██║   ██║████╗  ██║██║██║   ██║     ██╔════╝██║     ██║",
		"██║   ██║██╔██╗ ██║██║██║   ██║     ██║     ██║     ██║",
		"██║   ██║██║╚██╗██║██║╚██╗ ██╔╝     ██║     ██║     ██║",
		"╚██████╔╝██║ ╚████║██║ ╚████╔╝      ╚██████╗███████╗██║",
		" ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═══╝        ╚═════╝╚══════╝╚═╝",
	}

	s := "\n"
	for i, line := range logoLines {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(theme.gradient[i])).Bold(true)
		s += style.Render(line) + "\n"
	}

	s += "\n"

	s += "Menu principal:\n\n"

	selectedStyle := getSelectedStyle(theme)
	for i, choice := range m.choices {
		cursor := "   "
		if m.cursor == i {
			cursor = " > "
			s += selectedStyle.Render(fmt.Sprintf("%s %s ", cursor, choice)) + "\n"
		} else {
			s += menuStyle.Render(fmt.Sprintf("%s %s", cursor, choice)) + "\n"
		}
	}

	s += "\n" + helpStyle.Render("Navigation: ↑↓ ou j/k  |  Sélection: Enter  |  Quitter: q")
	return s
}

func (m Model) viewDateTime() string {
	now := time.Now()
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Date et Heure ") + "\n\n"
	s += getHeaderStyle(theme).Render("Informations temporelles") + "\n\n"
	s += fmt.Sprintf("  Date        : %s\n", formatDateFrench(now))
	s += fmt.Sprintf("  Heure       : %s\n", now.Format("15:04:05"))
	s += fmt.Sprintf("  Fuseau      : %s\n", now.Format("MST"))

	s += "\n" + helpStyle.Render("Appuyez sur 'q' ou 'Esc' pour revenir au menu")
	return s
}

func formatDateFrench(t time.Time) string {
	daysFr := map[time.Weekday]string{
		time.Monday:    "Lundi",
		time.Tuesday:   "Mardi",
		time.Wednesday: "Mercredi",
		time.Thursday:  "Jeudi",
		time.Friday:    "Vendredi",
		time.Saturday:  "Samedi",
		time.Sunday:    "Dimanche",
	}

	monthsFr := map[time.Month]string{
		time.January:   "janvier",
		time.February:  "février",
		time.March:     "mars",
		time.April:     "avril",
		time.May:       "mai",
		time.June:      "juin",
		time.July:      "juillet",
		time.August:    "août",
		time.September: "septembre",
		time.October:   "octobre",
		time.November:  "novembre",
		time.December:  "décembre",
	}

	return fmt.Sprintf("%s %02d %s %d", daysFr[t.Weekday()], t.Day(), monthsFr[t.Month()], t.Year())
}

func (m Model) viewWelcome() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Message de bienvenue ") + "\n\n"

	welcomeStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.primary)).
		Bold(true)
	s += welcomeStyle.Render("Bienvenue dans univ-cli !") + "\n\n"
	s += "Nous sommes ravis de vous accueillir dans notre application CLI.\n"
	s += "Explorez les différentes fonctionnalités et découvrez tout ce que vous pouvez faire.\n\n"

	s += getHeaderStyle(theme).Render("Technologies utilisées") + "\n\n"
	s += "  - Cobra        : Framework CLI\n"
	s += "  - Bubble Tea   : Interface TUI\n"
	s += "  - Lipgloss     : Styling terminal\n\n"
	s += "Fonctionnalités:\n\n"
	s += "  - fs           : Gestion de fichiers (ls, cat, cp, mkdir)\n"
	s += "  - tui          : Interface textuelle interactive\n"
	s += "  - custom       : Commandes personnalisées\n"

	s += "\n" + helpStyle.Render("Appuyez sur 'q' ou 'Esc' pour revenir au menu")
	return s
}

func (m Model) viewGame() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Mini-Jeu ") + "\n\n"
	s += getHeaderStyle(theme).Render("Devinez le nombre (1-100)") + "\n\n"
	s += "Un nombre aléatoire a été généré entre 1 et 100.\n"
	s += "Essayez de le deviner en un minimum de tentatives.\n\n"

	s += "Votre proposition:\n"
	s += m.textInput.View() + "\n\n"

	if m.message != "" {
		s += infoStyle.Render(m.message) + "\n\n"
	}

	s += helpStyle.Render("Enter pour valider  |  q ou Esc pour revenir au menu")
	return s
}

func (m Model) viewTheme() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Sélection du thème ") + "\n\n"
	s += getHeaderStyle(theme).Render("Choisissez votre palette de couleurs") + "\n\n"

	for i, t := range themes {
		cursor := "   "
		if i == m.currentTheme {
			cursor = " > "
			s += getSelectedStyle(t).Render(fmt.Sprintf("%s %s ", cursor, t.name)) + "\n"

			s += "     "
			for _, color := range t.gradient {
				colorBlock := lipgloss.NewStyle().
					Foreground(lipgloss.Color(color)).
					Bold(true)
				s += colorBlock.Render("███ ")
			}
			s += "\n\n"
		} else {
			s += menuStyle.Render(fmt.Sprintf("%s %s", cursor, t.name)) + "\n\n"
		}
	}

	s += helpStyle.Render("Navigation: ↑↓ ou j/k  |  Appliquer: Enter  |  Annuler: q/Esc")
	return s
}

func (m Model) viewCustomMenu() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Menu Custom ") + "\n\n"
	s += getHeaderStyle(theme).Render("Fonctionnalités personnalisées") + "\n\n"

	selectedStyle := getSelectedStyle(theme)
	for i, choice := range m.customChoices {
		cursor := "   "
		if m.cursor == i {
			cursor = " > "
			s += selectedStyle.Render(fmt.Sprintf("%s %s ", cursor, choice)) + "\n"
		} else {
			s += menuStyle.Render(fmt.Sprintf("%s %s", cursor, choice)) + "\n"
		}
	}

	s += "\n" + helpStyle.Render("Navigation: ↑↓ ou j/k  |  Sélection: Enter  |  Retour: q/Esc")
	return s
}

func (m Model) viewQuote() string {
	theme := themes[m.currentTheme]
	quote := custom.GetRandomQuote()

	s := "\n"
	s += getTitleStyle(theme).Render(" Citation aléatoire ") + "\n\n"
	s += getHeaderStyle(theme).Render("Inspiration du jour") + "\n\n"

	quoteStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.primary)).
		Italic(true).
		Width(80)

	s += quoteStyle.Render(quote) + "\n"

	s += "\n" + helpStyle.Render("Appuyez sur 'q' ou 'Esc' pour revenir au menu Custom")
	return s
}

func (m Model) viewCalc() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Calculatrice ") + "\n\n"
	s += getHeaderStyle(theme).Render("Calculatrice simple") + "\n\n"
	s += "Opérateurs supportés: +, -, *, /\n"
	s += "Exemple: 42 + 8\n\n"

	s += "Votre calcul:\n"
	s += m.textInput.View() + "\n\n"

	if m.calcResult != "" {
		s += infoStyle.Render(m.calcResult) + "\n\n"
	}

	s += helpStyle.Render("Enter pour calculer  |  q ou Esc pour revenir au menu Custom")
	return s
}

func (m Model) viewSysinfo() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Informations Système ") + "\n\n"
	s += getHeaderStyle(theme).Render("Configuration système") + "\n\n"

	s += fmt.Sprintf("  OS              : %s\n", runtime.GOOS)
	s += fmt.Sprintf("  Architecture    : %s\n", runtime.GOARCH)
	s += fmt.Sprintf("  Nombre de CPUs  : %d\n", runtime.NumCPU())
	s += fmt.Sprintf("  Version de Go   : %s\n", runtime.Version())

	cwd, err := os.Getwd()
	if err == nil {
		s += fmt.Sprintf("  Répertoire      : %s\n", cwd)
	}

	s += "\n"
	s += getHeaderStyle(theme).Render("Variables d'environnement") + "\n\n"
	s += fmt.Sprintf("  HOME            : %s\n", os.Getenv("HOME"))
	s += fmt.Sprintf("  USER            : %s\n", os.Getenv("USER"))
	s += fmt.Sprintf("  SHELL           : %s\n", os.Getenv("SHELL"))

	s += "\n"
	s += getHeaderStyle(theme).Render("Statistiques mémoire Go") + "\n\n"
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	s += fmt.Sprintf("  Alloc           : %d KB\n", mem.Alloc/1024)
	s += fmt.Sprintf("  TotalAlloc      : %d KB\n", mem.TotalAlloc/1024)
	s += fmt.Sprintf("  Sys             : %d KB\n", mem.Sys/1024)
	s += fmt.Sprintf("  NumGC           : %d\n", mem.NumGC)

	s += "\n" + helpStyle.Render("Appuyez sur 'q' ou 'Esc' pour revenir au menu Custom")
	return s
}
