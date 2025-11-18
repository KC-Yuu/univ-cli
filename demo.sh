#!/bin/bash

# Script de démonstration pour la soutenance
# Appuyez sur ENTRÉE pour passer à la commande suivante

# Couleurs
CYAN='\033[0;36m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Fonction pour exécuter une commande avec pause
run_command() {
    local description="$1"
    local command="$2"

    echo ""
    echo "=========================================="
    echo "$description"
    echo "=========================================="
    echo ""
    echo -e "💻 Commande : ${CYAN}$command${NC}"
    echo ""
    read -p "Appuyez sur ENTRÉE pour exécuter..."
    echo ""

    # Exécuter la commande
    eval "$command"

    echo ""
    read -p "Appuyez sur ENTRÉE pour continuer..."
}

# Début de la démonstration
clear
echo -e "${GREEN}=========================================="
echo -e "   🚀 DÉMONSTRATION univ-cli"
echo -e "==========================================${NC}"
echo ""
echo "Ce script démontre toutes les fonctionnalités"
echo "de l'application univ-cli."
echo ""
read -p "Appuyez sur ENTRÉE pour commencer..."

# 1. Aide générale
run_command "1. Afficher l'aide générale" "./univ-cli --help"

# 2. Aide de fs
run_command "2. Afficher l'aide de la commande fs" "./univ-cli fs --help"

# === COMMANDE LS ===

# 3. Liste répertoire courant
run_command "3. [ls] Lister le répertoire courant" "./univ-cli fs ls"

# 4. Liste avec chemin relatif
run_command "4. [ls] Lister avec un chemin relatif (internal)" "./univ-cli fs ls internal"

# 5. Erreur ls - dossier inexistant
run_command "5. [ls] Gestion d'erreur (dossier inexistant)" "./univ-cli fs ls dossier_inexistant"

# === COMMANDE CAT ===

# 6. Préparer un fichier de test
clear
echo -e "${YELLOW}Préparation : Création d'un fichier de test${NC}"
echo "Contenu de test pour démonstration" > demo_test.txt
echo "Ligne 2 du fichier" >> demo_test.txt
echo "Ligne 3 du fichier" >> demo_test.txt
echo -e "${GREEN}✓ Fichier demo_test.txt créé dans le répertoire courant${NC}"
read -p "Appuyez sur ENTRÉE pour continuer..."

run_command "6. [cat] Afficher le contenu d'un fichier" "./univ-cli fs cat demo_test.txt"

run_command "7. [cat] Afficher le README.md" "./univ-cli fs cat README.md | head -20"

run_command "8. [cat] Gestion d'erreur (fichier inexistant)" "./univ-cli fs cat fichier_inexistant.txt"

# === COMMANDE CP ===

run_command "9. [cp] Copier un fichier" "./univ-cli fs cp demo_test.txt demo_copy.txt"

run_command "10. [cp] Vérifier que la copie a fonctionné" "./univ-cli fs cat demo_copy.txt"

run_command "11. [cp] Voir les fichiers créés dans le répertoire" "./univ-cli fs ls | grep demo"

run_command "12. [cp] Gestion d'erreur (fichier source inexistant)" "./univ-cli fs cp fichier_inexistant.txt dest.txt"

# === COMMANDE MKDIR ===

run_command "13. [mkdir] Créer un répertoire" "./univ-cli fs mkdir demo_nouveau_dossier"

run_command "14. [mkdir] Vérifier que le dossier a été créé" "./univ-cli fs ls | grep demo"

run_command "15. [mkdir] Créer des répertoires imbriqués" "./univ-cli fs mkdir demo_parent/demo_enfant/demo_petit_enfant"

run_command "16. [mkdir] Vérifier la structure créée" "./univ-cli fs ls demo_parent/demo_enfant"

run_command "17. [mkdir] Lister pour voir tous les fichiers de démo" "./univ-cli fs ls"

run_command "18. [mkdir] Gestion d'erreur (dossier existe déjà)" "./univ-cli fs mkdir demo_nouveau_dossier"

# === COMMANDE CUSTOM ===

clear
echo ""
echo "=========================================="
echo "COMMANDES CUSTOM"
echo "=========================================="
echo ""
echo -e "${YELLOW}Note importante:${NC}"
echo "Les commandes custom sont disponibles en CLI et en TUI."
echo ""
echo "• ${GREEN}Mode CLI${NC}: Idéal pour les scripts, automation, et utilisation rapide"
echo "  Exemple: ./univ-cli custom sysinfo"
echo ""
echo "• ${GREEN}Mode TUI${NC}: Interface interactive pour exploration et usage manuel"
echo "  Accessible via le menu interactif"
echo ""
read -p "Appuyez sur ENTRÉE pour continuer..."

run_command "19. [custom] Afficher l'aide de custom" "./univ-cli custom --help"

run_command "20. [custom quote] Citation aléatoire (usage CLI)" "./univ-cli custom quote"

run_command "21. [custom calc] Calculer 42 + 8 (usage CLI)" "./univ-cli custom calc \"42 + 8\""

run_command "22. [custom calc] Calculer 100 / 4" "./univ-cli custom calc \"100 / 4\""

run_command "23. [custom calc] Calculer 15 * 3" "./univ-cli custom calc \"15 * 3\""

run_command "24. [custom calc] Gestion d'erreur (division par zéro)" "./univ-cli custom calc \"10 / 0\""

clear
echo ""
echo "=========================================="
echo "Démonstration: Custom CLI vs TUI"
echo "=========================================="
echo ""
echo -e "${CYAN}Cas d'usage CLI:${NC}"
echo "Les commandes custom en CLI sont parfaites pour:"
echo "  • Scripts d'automation"
echo "  • Intégration CI/CD"
echo "  • Monitoring système"
echo "  • Utilisation dans des pipelines"
echo ""
echo -e "${CYAN}Exemple pratique:${NC}"
echo "Vous pouvez scripter les informations système:"
echo ""
read -p "Appuyez sur ENTRÉE pour exécuter: ./univ-cli custom sysinfo"
echo ""

./univ-cli custom sysinfo

echo ""
echo -e "${GREEN}✓${NC} Les informations sont affichées directement, parfait pour l'automation!"
echo ""
read -p "Appuyez sur ENTRÉE pour continuer..."

# === COMMANDE TUI ===

clear
echo ""
echo "=========================================="
echo "25. [tui] Démonstration TUI interactive"
echo "=========================================="
echo ""
echo -e "💻 Commande : ${CYAN}./univ-cli tui${NC}"
echo ""
echo -e "${YELLOW}Note:${NC} Le TUI va se lancer. Vous pourrez:"
echo "  - Naviguer avec ↑/↓ ou j/k"
echo "  - Accéder aux 5 options du menu principal"
echo "  - Explorer le sous-menu Custom avec 5 fonctionnalités"
echo "  - Tester les 6 thèmes de couleurs disponibles"
echo "  - Appuyer sur 'q' pour quitter et revenir à la démo"
echo ""
echo -e "${CYAN}Cas d'usage TUI:${NC}"
echo "Le mode TUI est idéal pour:"
echo "  • Exploration interactive"
echo "  • Découverte des fonctionnalités"
echo "  • Interface utilisateur conviviale"
echo "  • Pas besoin de mémoriser les commandes"
echo ""
read -p "Appuyez sur ENTRÉE pour lancer le TUI..."
echo ""

./univ-cli tui

echo ""
echo -e "${GREEN}✓ Retour de l'interface TUI${NC}"
echo ""
echo -e "${YELLOW}Résumé:${NC}"
echo "CLI et TUI offrent deux approches complémentaires:"
echo "  • ${GREEN}CLI${NC}: Rapide, scriptable, pour experts"
echo "  • ${GREEN}TUI${NC}: Intuitif, guidé, pour tous les utilisateurs"
echo ""
read -p "Appuyez sur ENTRÉE pour continuer..."

# === NETTOYAGE ===
clear
echo -e "${YELLOW}Nettoyage des fichiers de test...${NC}"
rm -f demo_test.txt demo_copy.txt
rm -rf demo_nouveau_dossier demo_parent
echo -e "${GREEN}✓ Nettoyage terminé${NC}"
echo ""
echo -e "${YELLOW}Vérification : lister le répertoire après nettoyage${NC}"
./univ-cli fs ls | grep demo || echo -e "${GREEN}✓ Aucun fichier de démo restant${NC}"
echo ""
read -p "Appuyez sur ENTRÉE pour voir le résumé..."

# Fin
clear
echo ""
echo -e "${GREEN}=========================================="
echo -e "   ✅ DÉMONSTRATION TERMINÉE"
echo -e "==========================================${NC}"
echo ""
echo -e "${YELLOW}Résumé des fonctionnalités démontrées :${NC}"
echo ""
echo -e "${GREEN}Commande fs (Système de fichiers) :${NC}"
echo -e "  ${GREEN}✓${NC} ls  : Listage de répertoires"
echo -e "  ${GREEN}✓${NC} cat : Affichage de fichiers texte"
echo -e "  ${GREEN}✓${NC} cp  : Copie de fichiers avec préservation des permissions"
echo -e "  ${GREEN}✓${NC} mkdir : Création de répertoires imbriqués"
echo -e "  ${GREEN}✓${NC} Gestion complète d'erreurs"
echo ""
echo -e "${GREEN}Commande custom (Fonctionnalités personnalisées) :${NC}"
echo -e "  ${GREEN}✓${NC} quote   : Citations aléatoires inspirantes"
echo -e "  ${GREEN}✓${NC} calc    : Calculatrice simple (+, -, *, /)"
echo -e "  ${GREEN}✓${NC} sysinfo : Informations système et Go runtime"
echo -e "  ${GREEN}✓${NC} Gestion d'erreurs (division par zéro, etc.)"
echo ""
echo -e "${GREEN}Commande tui (Interface utilisateur textuelle) :${NC}"
echo -e "  ${GREEN}✓${NC} Menu interactif avec navigation (↑/↓, j/k)"
echo -e "  ${GREEN}✓${NC} Affichage date et heure avec timestamp"
echo -e "  ${GREEN}✓${NC} Message de bienvenue personnalisé"
echo -e "  ${GREEN}✓${NC} Sous-menu Custom avec 5 fonctionnalités"
echo -e "  ${GREEN}✓${NC} Citations aléatoires, calculatrice, sysinfo, mini-jeu"
echo -e "  ${GREEN}✓${NC} 6 thèmes de couleurs avec dégradés"
echo -e "  ${GREEN}✓${NC} Interface Bubble Tea avec Lipgloss styling"
echo ""
echo -e "${GREEN}Double accès CLI et TUI :${NC}"
echo -e "  ${GREEN}✓${NC} Mode CLI : Automation et scripts (./univ-cli custom sysinfo)"
echo -e "  ${GREEN}✓${NC} Mode TUI : Interface interactive et exploration"
echo -e "  ${GREEN}✓${NC} Même fonctionnalités, deux approches complémentaires"
echo ""
echo -e "${GREEN}Technologies utilisées :${NC}"
echo -e "  ${GREEN}✓${NC} Go 1.25.1"
echo -e "  ${GREEN}✓${NC} Cobra (framework CLI)"
echo -e "  ${GREEN}✓${NC} Bubble Tea (framework TUI)"
echo -e "  ${GREEN}✓${NC} Lipgloss (styling terminal)"
echo -e "  ${GREEN}✓${NC} Bubbles (composants TUI)"
echo ""
echo -e "${GREEN}Bonnes pratiques :${NC}"
echo -e "  ${GREEN}✓${NC} Architecture modulaire (cmd/, internal/)"
echo -e "  ${GREEN}✓${NC} Gestion d'erreurs avec wrapping (%w)"
echo -e "  ${GREEN}✓${NC} Messages clairs en français"
echo -e "  ${GREEN}✓${NC} Interface colorée avec emojis"
echo -e "  ${GREEN}✓${NC} Documentation complète (README, NOTES, CAHIER)"
echo ""
echo -e "${GREEN}==========================================${NC}"
echo ""
