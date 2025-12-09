#!/bin/bash

# Script de démonstration pour la soutenance
# Appuyez sur ENTRÉE pour passer à la commande suivante

CYAN='\033[0;36m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

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

    eval "$command"

    echo ""
    read -p "Appuyez sur ENTRÉE pour continuer..."
}

clear
echo -e "${GREEN}=========================================="
echo -e "   🚀 DÉMONSTRATION univ-cli"
echo -e "==========================================${NC}"
echo ""
echo "Ce script démontre toutes les fonctionnalités"
echo "de l'application univ-cli."
echo ""
read -p "Appuyez sur ENTRÉE pour commencer..."

run_command "1. Afficher l'aide générale" "./univ-cli --help"

run_command "2. Afficher l'aide de la commande fs" "./univ-cli fs --help"

# === COMMANDE LS ===

run_command "3. [ls] Lister le répertoire courant" "./univ-cli fs ls"

run_command "4. [ls] Lister avec un chemin relatif (internal)" "./univ-cli fs ls internal"

run_command "5. [ls] Gestion d'erreur (dossier inexistant)" "./univ-cli fs ls dossier_inexistant"

# === COMMANDE CAT ===

clear
echo -e "${YELLOW}Préparation : Création d'un fichier de test${NC}"
echo "Contenu de test pour démonstration" > demo_test.txt
echo "Ligne 2 du fichier" >> demo_test.txt
echo "Ligne 3 du fichier" >> demo_test.txt
echo -e "${GREEN}✓ Fichier demo_test.txt créé${NC}"
read -p "Appuyez sur ENTRÉE pour continuer..."

run_command "6. [cat] Afficher le contenu d'un fichier" "./univ-cli fs cat demo_test.txt"

run_command "7. [cat] Gestion d'erreur (fichier inexistant)" "./univ-cli fs cat fichier_inexistant.txt"

# === COMMANDE CP ===

run_command "8. [cp] Copier un fichier" "./univ-cli fs cp demo_test.txt demo_copy.txt"

run_command "9. [cp] Vérifier que la copie a fonctionné" "./univ-cli fs cat demo_copy.txt"

run_command "10. [cp] Gestion d'erreur (fichier source inexistant)" "./univ-cli fs cp fichier_inexistant.txt dest.txt"

# === COMMANDE MKDIR ===

run_command "11. [mkdir] Créer un répertoire" "./univ-cli fs mkdir demo_dossier"

run_command "12. [mkdir] Créer des répertoires imbriqués" "./univ-cli fs mkdir demo_parent/enfant/petit_enfant"

run_command "13. [mkdir] Vérifier la structure créée" "./univ-cli fs ls demo_parent/enfant"

run_command "14. [mkdir] Gestion d'erreur (dossier existe déjà)" "./univ-cli fs mkdir demo_dossier"

# === COMMANDE CUSTOM ===

run_command "15. [custom] Afficher l'aide de custom" "./univ-cli custom --help"

run_command "16. [custom sysinfo] Informations système" "./univ-cli custom sysinfo"

# === COMMANDE TUI ===

clear
echo ""
echo "=========================================="
echo "17. [tui] Interface interactive"
echo "=========================================="
echo ""
echo -e "💻 Commande : ${CYAN}./univ-cli tui${NC}"
echo ""
echo -e "${YELLOW}Navigation :${NC}"
echo "  ↑/↓ ou j/k : Naviguer dans les menus"
echo "  Enter      : Sélectionner"
echo "  q ou Esc   : Retour / Quitter"
echo ""
echo -e "${YELLOW}À démontrer :${NC}"
echo "  1. Date et Heure"
echo "  2. Message de bienvenue"
echo "  3. Custom > Mini-Jeu (deviner un nombre)"
echo "  4. Custom > Informations système"
echo "  5. Changer de thème (4 thèmes disponibles)"
echo ""
read -p "Appuyez sur ENTRÉE pour lancer le TUI..."
echo ""

./univ-cli tui

# === NETTOYAGE ===

clear
echo -e "${YELLOW}Nettoyage des fichiers de test...${NC}"
rm -f demo_test.txt demo_copy.txt
rm -rf demo_dossier demo_parent
echo -e "${GREEN}✓ Nettoyage terminé${NC}"
echo ""

# === RÉSUMÉ ===

echo -e "${GREEN}=========================================="
echo -e "   ✅ DÉMONSTRATION TERMINÉE"
echo -e "==========================================${NC}"
echo ""
echo -e "${YELLOW}Fonctionnalités démontrées :${NC}"
echo ""
echo -e "${GREEN}fs${NC} - Système de fichiers :"
echo "  ✓ ls    : Lister répertoires"
echo "  ✓ cat   : Afficher fichiers"
echo "  ✓ cp    : Copier fichiers"
echo "  ✓ mkdir : Créer répertoires"
echo ""
echo -e "${GREEN}custom${NC} - Fonctionnalité personnalisée :"
echo "  ✓ sysinfo : Informations système"
echo ""
echo -e "${GREEN}tui${NC} - Interface interactive :"
echo "  ✓ Date et Heure"
echo "  ✓ Message de bienvenue"
echo "  ✓ Mini-Jeu (deviner un nombre)"
echo "  ✓ Informations système"
echo "  ✓ 4 thèmes de couleurs"
echo ""
echo -e "${GREEN}Technologies :${NC}"
echo "  ✓ Cobra      - Framework CLI"
echo "  ✓ Bubble Tea - Framework TUI"
echo "  ✓ Lipgloss   - Styling terminal"
echo ""
