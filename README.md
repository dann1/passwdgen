# passwdgen

[![Go](https://github.com/dann1/passwdgen/actions/workflows/go.yml/badge.svg)](https://github.com/dann1/passwdgen/actions/workflows/go.yml)

CLI password generator

## Usage

```text
passwdgen <length> [--uppercase] [--numbers] [--specials]
```

Examples

```text
 ~/P/g/passwdgen    passwdgen 8
lcjwaiad
 ~/P/g/passwdgen    passwdgen 15 --uppercase
LNbchmDxYvQrDfs
 ~/P/g/passwdgen    passwdgen 15 --uppercase --specials
VyYHgw!b-]wxUGs
 ~/P/g/passwdgen    passwdgen 15 --uppercase --specials --numbers
.=.oIzACeV9J.q?
```

## Build

```text
make
```
