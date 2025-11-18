# univ-cli

## 📋 Description

`univ-cli` est une application en ligne de commande développée en Go qui offre trois modules principaux :
- **fs** : Manipulation du système de fichiers (ls, cat, cp, mkdir)
- **custom** : Fonctionnalités personnalisées (citations, calculatrice, infos système)
- **tui** : Interface utilisateur textuelle interactive avec en bonus 6 thèmes de couleurs

## 🚀 Installation

### Prérequis
- Go 1.21 ou supérieur

### Compilation
```bash
# Télécharger les dépendances
go mod download

# Compiler l'application
go build -o univ-cli

# Rendre exécutable (Linux/macOS)
chmod +x univ-cli
```

## 📖 Utilisation

### Aide générale
```bash
./univ-cli --help
```

### Module `fs` - Système de fichiers

```bash
# Lister les fichiers et dossiers
./univ-cli fs ls
./univ-cli fs ls internal

# Afficher le contenu d'un fichier
./univ-cli fs cat README.md

# Copier un fichier
./univ-cli fs cp source.txt dest.txt

# Créer un répertoire
./univ-cli fs mkdir nouveau_dossier
./univ-cli fs mkdir parent/enfant/petit-enfant
```

### Module `custom` - Fonctionnalités personnalisées

```bash
# Citation aléatoire inspirante
./univ-cli custom quote

# Calculatrice simple
./univ-cli custom calc "42 + 8"
./univ-cli custom calc "100 / 4"

# Informations système
./univ-cli custom sysinfo
```

### Module `tui` - Interface interactive

```bash
# Lancer l'interface textuelle
./univ-cli tui
```

**Fonctionnalités du TUI :**
- Menu principal avec 5 options
- Sous-menu Custom avec citations, calculatrice, sysinfo et mini-jeu
- 6 thèmes de couleurs personnalisables
- Logo ASCII avec dégradés
- Navigation : ↑/↓ ou j/k, Enter pour sélectionner, q pour quitter

## 🎨 Thèmes disponibles

- Bleu Océan (thème par défaut)
- Vert Forêt
- Violet Galaxie
- Orange Sunset
- Rouge Cardinal
- Cyan Arctique

## 🛠️ Technologies utilisées

| Bibliothèque | Usage |
|--------------|-------|
| [Cobra](https://github.com/spf13/cobra) | Framework CLI |
| [Bubble Tea](https://github.com/charmbracelet/bubbletea) | Framework TUI |
| [Lipgloss](https://github.com/charmbracelet/lipgloss) | Styling terminal |
| [Bubbles](https://github.com/charmbracelet/bubbles) | Composants TUI |

## 📁 Architecture

```
univ-cli/
├── main.go                 # Point d'entrée
├── cmd/                    # Définitions des commandes
│   ├── root.go
│   ├── fs.go
│   ├── custom.go
│   └── tui.go
├── internal/               # Logique métier
│   ├── filesystem/         # Fonctions fs
│   ├── custom/             # Fonctions custom
│   └── tui/                # Interface TUI
└── demo.sh                 # Script de démonstration
```

## 🧪 Démonstration

Un script de démonstration complet est disponible :

```bash
chmod +x demo.sh
./demo.sh
```

Le script démontre :
- Toutes les commandes `fs` avec gestion d'erreurs
- Toutes les commandes `custom` (via les commandes CLI et via l'interface TUI)
- L'interface TUI interactive

## 💡 Cas d'usage

### Mode CLI
- **Scripts d'automation** : `./univ-cli custom sysinfo > rapport.txt`
- **Intégration CI/CD** : Monitoring système dans des pipelines
- **Utilisation rapide** : Pas besoin de naviguer dans des menus

### Mode TUI
- **Exploration interactive** : Découvrir les fonctionnalités
- **Interface conviviale** : Pas besoin de mémoriser les commandes
- **Expérience utilisateur** : Navigation intuitive avec thèmes

## 🔍 Bonnes pratiques

- ✅ Architecture modulaire (cmd/, internal/)
- ✅ Gestion d'erreurs avec wrapping (`%w`)
- ✅ Messages clairs en français
- ✅ Code commenté et documenté
- ✅ Séparation des responsabilités
- ✅ Helper functions pour réutilisabilité

## 👤 Auteur

Maxime Caron - Projet CLI Go - ESGI

## 📄 Licence

Projet académique ESGI

---

**Note** : Ce projet démontre une application CLI complète avec double accès (CLI + TUI).
