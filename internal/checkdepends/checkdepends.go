package checkdepends

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/g5ostXa/archypr/internal/core"
	"github.com/g5ostXa/archypr/internal/installpackages"
	"github.com/g5ostXa/archypr/internal/styles"
)

var dependInstallPromptMsg = "→ Do you want to install all missing dependencies now? (Yy/Nn):"

var packages = []string{
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

type packageChecker func(string) bool

func Validate() {
	missingPackages := missingDependencies(packages, isPackageInstalled)
	if len(missingPackages) == 0 {
		core.Logger.Info("All dependencies are already installed.")
		return
	}

	core.Logger.Warn(fmt.Sprintf("Missing dependencies: %s", strings.Join(missingPackages, ", ")))

	if !confirmDependencyInstall(os.Stdin, os.Stdout) {
		fmt.Println()
		core.Logger.Info("Installation cancelled by user.")
		os.Exit(0)
	}

	installpackages.Needed(missingPackages)
}

func missingDependencies(packages []string, installed packageChecker) []string {
	missing := make([]string, 0)
	for _, pkg := range packages {
		if !installed(pkg) {
			missing = append(missing, pkg)
		}
	}
	return missing
}

func isPackageInstalled(pkg string) bool {
	cmd := exec.Command("paru", "-Qi", pkg)
	return cmd.Run() == nil
}

func fprint(output io.Writer, message string) {
	fmt.Fprint(output, styles.CommonPromptStyle.Render(message))
}

func confirmDependencyInstall(input io.Reader, output io.Writer) bool {
	reader := bufio.NewReader(input)

	for {
		fmt.Fprintln(output)
		fprint(output, dependInstallPromptMsg)

		answer, _ := reader.ReadString('\n')
		answer = strings.ToLower(strings.TrimSpace(answer))

		switch answer {
		case "y":
			return true
		case "n", "":
			return false
		default:
			core.Logger.Warn("Invalid input. Please type y or n ...")
		}
	}
}
