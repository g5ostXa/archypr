# shellcheck disable=SC2148
#      __            __
#     / /  ___ ____ / /  ________
#    / _ \/ _ `(_-</ _ \/ __/ __/
# (_)_.__/\_,_/___/_//_/_/  \__/
#
# ~/.bashrc
#
# =============================================

# // ===== Colors =====
PS1='[\u@\h \W]\$'
MAGENTA='\033[1;35m'
RC='\033[0m'

# // ===== General aliases =====
alias ls="eza --icons=always --color=always"
alias ll="ls -a"
alias bwrap-btop="~/src/Scripts/sandboxes/bwrap-btop.sh"
alias Hy2clean="~/src/Scripts/cleanup.sh"
alias errcheck="~/src/Scripts/checkerrors.sh"
alias cw="cliphist wipe"
alias mirrors-update="~/src/Scripts/mirrors.sh"
alias reboot="~/src/Scripts/reboot.sh"
alias poweroff="~/src/Scripts/poweroff.sh"
alias Hy2in="~/src/Scripts/hypr/start-hypr.sh"
alias Hy2out="~/src/Scripts/hypr/killhypr.sh"
alias lumine="~/src/Scripts/lumineV3.sh"
alias r4in="unimatrix -n -s 96 -l o"

# // ===== General ======
export PATH="$HOME/src:$HOME/.local/bin:$HOME/go/bin:$PATH"
export ARCHYPR_VERSION_FILE="$HOME/.config/archypr/.version/latest"
if [[ -f "$ARCHYPR_VERSION_FILE" ]]; then
	# shellcheck disable=SC2155
	export ARCHYPR_VERSION=$(cat "$ARCHYPR_VERSION_FILE")
else
	export ARCHYPR_VERSION="unknown"
fi

eval "$(starship init bash)"
cat ~/.cache/wal/sequences

if [[ $(tty) == *"pts"* ]]; then
	if [[ -f "$HOME/go/bin/ghostshell" ]]; then
		"$HOME/go/bin/ghostshell" --title "archypr" --version-file "$HOME/.config/archypr/.version/latest"
	else
		echo -e "${MAGENTA}"
		cat <<"EOF"
               __
 ___ _________/ /  __ _____  ____
/ _ `/ __/ __/ _ \/ // / _ \/ __/
\_,_/_/  \__/_//_/\_, / .__/_/
                 /___/_/
EOF
		echo ""
		echo "$ARCHYPR_VERSION"
		echo -e "${RC}"
	fi
fi

# // ===== Set fish interactively =====
if [[ $(ps --no-header --pid=$PPID --format=comm) != "fish" && -z ${BASH_EXECUTION_STRING} ]]; then
	shopt -q login_shell && LOGIN_OPTION='--login' || LOGIN_OPTION=''
	# shellcheck disable=SC2086
	exec fish $LOGIN_OPTION
fi

# // ===== Set UWSM for hyprland management =====
if uwsm check may-start && uwsm select; then
	exec systemd-cat -t uwsm_start uwsm start default
fi
