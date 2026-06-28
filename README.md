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

### Installation
Download tbe latest git:
```bash
git clone --depth=1 https://github.com/g5ostXa/archypr.git /path/to/target/dir
```

To run, cd into [archypr](https://github.com/g5ostXa/archypr) and simply run:
```bash
go run ./cmd/archypr
```

Or, build the binary and move it to your `$GOBIN` directory:
```bash
go build -o "binrary-name" ./cmd/archypr
mv ./binary-name "$GOBIN" 
```
