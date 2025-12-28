# PassGen 🔐  
A secure, Windows-friendly CLI password generator written in Go.

PassGen is a lightweight command-line tool that generates cryptographically secure passwords with customizable rules. It supports short and long flags, combined flags, clipboard copy, and Go-native installation using `go install`.

---

## ✨ Features

- 🔒 Cryptographically secure password generation (`crypto/rand`)
- 📏 Custom password length
- 🔠 Select character types:
  - Uppercase letters
  - Lowercase letters
  - Digits
  - Symbols
- 🚫 Exclude ambiguous characters (`0 O 1 l I`)
- 📋 Copy password directly to clipboard (Windows supported)
- 🔢 Generate multiple passwords at once
- ⚡ Short flags, long flags, and combined flags (`-ulsA`)
- 🌍 Installable globally using Go’s native tooling

---

## 📦 Installation (Go Native – Recommended)

### Prerequisites
- Go **1.20+**
- `$GOPATH/bin` (or `$HOME/go/bin`) added to your PATH

### Install using `go install`

```bash
go install github.com/laatu08/passgen/cmd/passgen@latest
```

Verify installation:

```bash
passgen --version
```

---

## 🚀 Usage

### Generate a password

```bash
passgen -uld -L 12
```

### Strong password + clipboard

```bash
passgen -ulsA -L 16 -C
```

### Multiple passwords

```bash
passgen -ud -L 10 -c 5
```

---

## 🧾 Flags

| Short | Long | Description |
|------|------|-------------|
| `-L` | `--length` | Password length |
| `-u` | `--upper` | Uppercase letters |
| `-l` | `--lower` | Lowercase letters |
| `-d` | `--digits` | Digits |
| `-s` | `--symbols` | Symbols |
| `-A` | `--no-ambiguous` | Exclude ambiguous characters |
| `-c` | `--count` | Number of passwords |
| `-C` | `--clipboard` | Copy to clipboard |
| `-v` | `--version` | Show version |
|  | `--help` | Help menu |

---

## 🔐 Security

- Uses `crypto/rand`
- Enforces selected rules
- Shuffles output
- No password storage

---

## 🛠 Project Structure

```text
passgen/
├── cmd/
│   └── passgen/
│       └── main.go
├── internal/
│   └── generator/
│       └── password.go
├── go.mod
└── README.md
```

---

## 📄 License

MIT License

---

## 🙌 Author

Partha Borah  
https://github.com/laatu08
