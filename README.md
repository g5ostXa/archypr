<div align="center">

<img src=".github/logos/archypr.png" height="350" width="400">
</div>

Archypr is an installer for a pre-configured [hyprland](https://hypr.land) setup on top of [Archlinux](https://archlinux.org).

We welcome contributions, ideas, advice and even opinions. \
Communicate with us: g5ostX2@proton.me

> [!NOTE]
> - archypr is NOT production ready. 
> - This is a slow project that might get discontinued really quick, or not. 
> - My real hyprland dotfiles repo lives [here](https://github.com/g5ostXa/hyprarch2)
> - I'm 100% self-taught and definitly have skill issues, so this is also meant to improve my skills.
<br>

### Installation
First you'll need to make sure you're on a brand new [Arch](https://archlinux.org) base system. 

Then, install the pre-install dependencies:
```bash
sudo pacman -S --needed --noconfirm go curl git reflector xdg-utils xdg-user-dirs vim networkmanager network-manager-applet wireless_tools wpa_supplicant dialog os-prober mtools dosfstools base-devel linux-headers
```
<br>

Also, make sure you have user directories and `~/.config` in your `HOME` directory:
```bash
$ xdg-user-dirs-update && mkdir -p "$HOME/.config"
```
<br>

Now, to download the latest git for [archypr](https://github.com/g5ostXa/archypr):
```bash
git clone --depth=1 https://github.com/g5ostXa/archypr.git
```
<br>

To get the latest release, get the binrary [here](https://github.com/g5ostXa/archypr/releases/download/0.0-1/archypr-bin) and it's `sha256` [here](https://github.com/g5ostXa/archypr/releases/download/0.0-1/archypr-bin.sha256)
  or, simply use `curl`:
```bash
curl -L -o archypr-bin https://github.com/g5ostXa/archypr/releases/download/0.0-1/archypr-bin && curl -L -o archypr-bin.sha256 https://github.com/g5ostXa/archypr/releases/download/0.0-1/archypr-bin.sha256
```
<br>

### Verify the binary
If you downloaded the project via `curl` or the release page via your browser, \
you can verify the binrary with the `archypr-bin.sha256` file:
```bash
sha256sum -c ./archypr-bin.sha256
```
Expected output:
```bash
archypr-bin: OK
```
<br>

### Usage
Make the binary executable and run it:
```bash
chmod +x ./archypr-bin && ./archypr-bin 
```
