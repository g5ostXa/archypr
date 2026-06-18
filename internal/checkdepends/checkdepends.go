package checkdepends

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/g5ostXa/archypr/internal/core"
	"github.com/g5ostXa/archypr/internal/styles"
)

var dependInstallPromptMsg = "→ Do you want to install all dependencies now? (Yy/Nn):"

func Validate() {
	Packages := []string{
		"hyprland",
		"hyprpolkitagent",
		"uwsm",
		"ghostty",
		"aquamarine",
		"waybar",
		//"rofi ",
		"dunst",
		"libnotify",
		"cliphist",
		//"wlogout-git",
		"xdg-desktop-portal-hyprland",
		"xdg-desktop-portal-gtk",
		"qt5-wayland",
		"waypaper-git",
		"hyprpicker",
		"hyprlock",
		"hyprcursor",
		"hypridle",
		"hyprgraphics",
		"hyprlang",
		"hyprls-git",
		"hyprwayland-scanner",
		"cantarell-fonts",
		"otf-font-awesome",
		"woff2-font-awesome",
		"ttf-fira-sans",
		"ttf-fira-code",
		"ttf-firacode-nerd",
		"ttf-nerd-fonts-symbols",
		"ttf-nerd-fonts-symbols-mono",
		"ttf-jetbrains-mono-nerd",
		"gnu-free-fonts",
		"brightnessctl",
		"neovim",
		"nautilus",
		"fastfetch",
		"pavucontrol",
		"pipewire",
		"pipewire-pulse",
		"pipewire-alsa",
		"pipewire-jack",
		"wireplumber",
		"bibata-cursor-theme",
		"dracula-icons-theme",
		"tokyonight-gtk-theme-git",
		"python-pywal16",
		"gtk3",
		"gtk4",
		"awww",
		"fish",
		"starship",
		"python-pip",
		"eza",
		"swappy",
		"firefox-nightly-bin",
		"ccache",
		"jq",
		"pacman-contrib",
		"fzf",
		"ttf-0xproto-nerd",
		"grim",
		"bubblewrap",
		"btop",
		"gum",
		"figlet",
		"lua",
	}

	for _, pkg := range Packages {
		cmd := exec.Command("paru", "-Qi", pkg)

		if err := cmd.Run(); err != nil {

			core.TimeLogger.Warn("Some dependencies are missing...")

			// Uncomment to test weather the cmd fails or if really missing a package:
			//core.TimeLogger.Warn(fmt.Sprintf("Missing dependency: %s", pkg))

			reader := bufio.NewReader(os.Stdin)
			fmt.Println()

			lipgloss.Print(styles.CommonPromptStyle.Render(dependInstallPromptMsg))

			input, _ := reader.ReadString('\n')
			input = strings.ToLower(strings.TrimSpace(input))

			if input == "n" || input == "" {
				fmt.Println()
				core.Logger.Info("Installation cancelled by user.")
				os.Exit(0)
			} else if input == "y" {
				break
			}

			core.Logger.Warn("Invalid input. Please type y or n ...")
		}
	}
}
