package tui

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// === ÉTATS DE L'APPLICATION ===
type state int

const (
	stateMenu     state = iota // Menu principal
	stateDateTime              // Affichage date/heure
	stateWelcome               // Message de bienvenue
	stateCustom                // Sous-menu Custom
	stateGame                  // Mini-jeu
	stateSysinfo               // Informations système
	stateTheme                 // Sélection du thème
	stateQuit                  // Quitter
)

// === THÈMES DE COULEURS ===
type Theme struct {
	name       string   
	gradient   []string
	primary    string
	selectedBg string
}

var themes = []Theme{
	{
		name:       "Bleu Océan",
		gradient:   []string{"#66B3FF", "#3399FF", "#0080FF", "#0066CC", "#004C99", "#003366"},
		primary:    "#0066CC",
		selectedBg: "#0066CC",
	},
	{
		name:       "Vert Forêt",
		gradient:   []string{"#90EE90", "#66CC66", "#4CAF50", "#2E7D32", "#1B5E20", "#0D3D0D"},
		primary:    "#4CAF50",
		selectedBg: "#2E7D32",
	},
	{
		name:       "Violet Galaxie",
		gradient:   []string{"#E1BEE7", "#BA68C8", "#9C27B0", "#7B1FA2", "#6A1B9A", "#4A148C"},
		primary:    "#9C27B0",
		selectedBg: "#7B1FA2",
	},
	{
		name:       "Orange Sunset",
		gradient:   []string{"#FFCC80", "#FFB74D", "#FF9800", "#F57C00", "#E65100", "#BF360C"},
		primary:    "#FF9800",
		selectedBg: "#F57C00",
	},
}

// === MODÈLE PRINCIPAL ===
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
}

// === STYLES ===
var (
	menuStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			Italic(true)
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

// === INITIALISATION ===
func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Entrez un nombre..."
	ti.CharLimit = 3
	ti.Width = 20

	return Model{
		state:         stateMenu,
		cursor:        0,
		choices:       []string{"Date et Heure", "Message de bienvenue", "Custom", "Changer de thème", "Quitter"},
		customChoices: []string{"Mini-Jeu", "Informations système", "Retour"},
		textInput:     ti,
		currentTheme:  0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// === GESTION DES ÉVÉNEMENTS ===
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case stateMenu:
			return m.updateMenu(msg)
		case stateDateTime, stateWelcome, stateSysinfo:
			if msg.String() == "q" || msg.String() == "esc" {
				if m.state == stateSysinfo {
					m.state = stateCustom
				} else {
					m.state = stateMenu
				}
			}
		case stateCustom:
			return m.updateCustomMenu(msg)
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

// Gestion du menu principal
func (m Model) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
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
			return m, tea.Quit
		}
	}
	return m, nil
}

// Gestion du sous-menu Custom
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
		case 0: // Mini-Jeu
			m.state = stateGame
			m.gameNumber = rand.Intn(100) + 1
			m.gameAttempts = 0
			m.textInput.SetValue("")
			m.textInput.Focus()
			m.message = ""
		case 1: // Sysinfo
			m.state = stateSysinfo
		case 2:
			m.state = stateMenu
			m.cursor = 0
		}
	}
	return m, nil
}

// Gestion du mini-jeu
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
			m.message = "Veuillez entrer un nombre valide."
			m.textInput.SetValue("")
			return m, nil
		}

		m.gameAttempts++

		if num == m.gameNumber {
			m.message = fmt.Sprintf("🎉 GAGNÉ ! C'était %d. Trouvé en %d tentative(s) !", m.gameNumber, m.gameAttempts)
		} else if num < m.gameNumber {
			m.message = fmt.Sprintf("⬆️  Trop petit ! (Tentative #%d)", m.gameAttempts)
		} else {
			m.message = fmt.Sprintf("⬇️  Trop grand ! (Tentative #%d)", m.gameAttempts)
		}

		m.textInput.SetValue("")
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// Gestion de la sélection du thème
func (m Model) updateTheme(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.state = stateMenu
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
	}
	return m, nil
}

// === AFFICHAGE ===
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
	case stateGame:
		return m.viewGame()
	case stateSysinfo:
		return m.viewSysinfo()
	case stateTheme:
		return m.viewTheme()
	case stateQuit:
		return "\nÀ bientôt !\n\n"
	}
	return ""
}

// Affichage du menu principal avec le logo ASCII
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
	s += "\nMenu principal:\n\n"

	// Affichage des options du menu
	selectedStyle := getSelectedStyle(theme)
	for i, choice := range m.choices {
		if m.cursor == i {
			s += selectedStyle.Render(fmt.Sprintf(" > %s ", choice)) + "\n"
		} else {
			s += menuStyle.Render(fmt.Sprintf("   %s", choice)) + "\n"
		}
	}

	s += "\n" + helpStyle.Render("↑↓ naviguer • Enter sélectionner • q quitter")
	return s
}

// Affichage du sous-menu Custom
func (m Model) viewCustomMenu() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Custom ") + "\n\n"
	s += getHeaderStyle(theme).Render("Fonctionnalités personnalisées") + "\n\n"

	selectedStyle := getSelectedStyle(theme)
	for i, choice := range m.customChoices {
		if m.cursor == i {
			s += selectedStyle.Render(fmt.Sprintf(" > %s ", choice)) + "\n"
		} else {
			s += menuStyle.Render(fmt.Sprintf("   %s", choice)) + "\n"
		}
	}

	s += "\n" + helpStyle.Render("↑↓ naviguer • Enter sélectionner • q retour")
	return s
}

// Affichage de la date et l'heure
func (m Model) viewDateTime() string {
	now := time.Now()
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Date et Heure ") + "\n\n"

	jours := []string{"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"}
	mois := []string{"", "janvier", "février", "mars", "avril", "mai", "juin",
		"juillet", "août", "septembre", "octobre", "novembre", "décembre"}

	s += fmt.Sprintf("  📅 Date : %s %d %s %d\n", jours[now.Weekday()], now.Day(), mois[now.Month()], now.Year())
	s += fmt.Sprintf("  🕐 Heure : %s\n", now.Format("15:04:05"))
	s += fmt.Sprintf("  🌍 Fuseau : %s\n", now.Location().String())

	s += "\n" + helpStyle.Render("q ou Esc pour revenir")
	return s
}

// Affichage du message de bienvenue
func (m Model) viewWelcome() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Bienvenue ") + "\n\n"

	welcomeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(theme.primary)).Bold(true)
	s += welcomeStyle.Render("Bienvenue dans univ-cli !") + "\n\n"

	s += "Cette application CLI a été développée en Go.\n\n"

	s += getHeaderStyle(theme).Render("Fonctionnalités :") + "\n"
	s += "  • fs     - Gestion de fichiers (ls, cat, cp, mkdir)\n"
	s += "  • tui    - Cette interface interactive\n"
	s += "  • custom - Commandes personnalisées\n\n"

	s += getHeaderStyle(theme).Render("Technologies :") + "\n"
	s += "  • Cobra      - Framework CLI\n"
	s += "  • Bubble Tea - Interface TUI\n"
	s += "  • Lipgloss   - Styling terminal\n"

	s += "\n" + helpStyle.Render("q ou Esc pour revenir")
	return s
}

// Affichage du mini-jeu
func (m Model) viewGame() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Mini-Jeu : Devinez le nombre ") + "\n\n"

	s += "Un nombre entre 1 et 100 a été choisi.\n"
	s += "Essayez de le deviner !\n\n"

	s += "Votre proposition : "
	s += m.textInput.View() + "\n\n"

	if m.message != "" {
		s += m.message + "\n\n"
	}

	s += helpStyle.Render("Enter pour valider • q ou Esc pour revenir")
	return s
}

// Affichage des informations système
func (m Model) viewSysinfo() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Informations Système ") + "\n\n"

	s += getHeaderStyle(theme).Render("Système") + "\n"
	s += fmt.Sprintf("  🖥️  OS           : %s\n", runtime.GOOS)
	s += fmt.Sprintf("  ⚙️  Architecture : %s\n", runtime.GOARCH)
	s += fmt.Sprintf("  🔢 CPUs         : %d\n", runtime.NumCPU())
	s += fmt.Sprintf("  🐹 Go version   : %s\n", runtime.Version())

	if cwd, err := os.Getwd(); err == nil {
		s += fmt.Sprintf("  📂 Répertoire   : %s\n", cwd)
	}

	s += "\n" + getHeaderStyle(theme).Render("Environnement") + "\n"
	s += fmt.Sprintf("  🏠 HOME  : %s\n", os.Getenv("HOME"))
	s += fmt.Sprintf("  👤 USER  : %s\n", os.Getenv("USER"))
	s += fmt.Sprintf("  🐚 SHELL : %s\n", os.Getenv("SHELL"))

	s += "\n" + helpStyle.Render("q ou Esc pour revenir")
	return s
}

// Affichage de la sélection du thème
func (m Model) viewTheme() string {
	theme := themes[m.currentTheme]

	s := "\n"
	s += getTitleStyle(theme).Render(" Thèmes ") + "\n\n"

	for i, t := range themes {
		if i == m.currentTheme {
			s += getSelectedStyle(t).Render(fmt.Sprintf(" > %s ", t.name)) + "\n"
			s += "     "
			for _, color := range t.gradient {
				colorBlock := lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Bold(true)
				s += colorBlock.Render("██ ")
			}
			s += "\n\n"
		} else {
			s += menuStyle.Render(fmt.Sprintf("   %s", t.name)) + "\n\n"
		}
	}

	s += helpStyle.Render("↑↓ naviguer • Enter appliquer • q annuler")
	return s
}
