<p align="center">
  <img src="images/dancing-gopher.gif" width="200" alt="Dancing Gopher" />
</p>

# goHTTP 🍩

A tiny but mighty HTTP file server written in Go — serves files from any directory with colorful request logging, minimal setup, and zero dependencies.

This is a passion project of mine because honestly?  
Typing `python3 -m http.server` every time gets old — and this is faster.  
Yes, I'm aware `socat` exists — but let's not talk about *that* syntax.

---

## 🚀 Features

- 🔥 Blazing fast file server written in Go
- 🌈 Color-coded logs for request methods (GET, POST, etc.)
- ⏱️ Logs client IP, request time, method, and path
- 📂 Easily serve any directory you want
- 🔒 Port availability check before binding
- ⚙️ Simple CLI with optional directory override
- 🧠 Clean and hackable — perfect for red teaming, demos, or local dev

---

## 📸 Demo

<p align="center">
  <img src="images/Demo.png" alt="goHTTP Demo Output" />
</p>

---

## 🧠 Usage

### 1. Build it

```bash
go build -o gohttp
```

### 2. (Optional) Move it

```bash
sudo mv gohttp /usr/bin/
```

### 3. Run it

```bash
# Serve current directory on default port 8000
./gohttp

# Serve current directory on custom port
./gohttp 9000

# Serve specific directory on custom port
./gohttp 9000 /path/to/dir
```

> If u moved the binary to `/usr/bin` to act a system binary, or followed the 2nd step then follow these 

```bash
# Serve current directory on default port 8000
gohttp

# Serve current directory on custom port
gohttp 9000

# Serve specific directory on custom port
gohttp 9000 /path/to/dir
```
