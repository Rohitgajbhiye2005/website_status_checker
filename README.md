# 🚀 Website Status Checker in Go

This is a simple **Website Status Checker** built with **Go**! It uses **goroutines** and **channels** to check if websites are **up or down** concurrently.

## 🟠 Features
- **Concurrent status checks** for multiple websites
- **Goroutines** for fast, non-blocking execution
- **Channels** for communication between routines

## 🟢 How It Works
- You provide a **list of website URLs**.
- Each website is checked in a **separate Goroutine**.
- The **HTTP response** is analyzed to determine if the site is up or down.

## 🟠 Installation
1. **Clone the repo:**
```bash
git clone https://github.com/Rohitgajbhiye2005/website_status_checker.git
cd website_status_checker
```

2. **Run the code:**
```bash
go run main.go
```

## 🟢 Example Output
```
http://google.com is up!
http://facebook.com is up!
http://stackoverflow.com is up!
http://golang.org is up!
http://amazon.com is up!
```

Or if a site is down:
```
http://example.com might be down
```

## 🟠 How to Contribute
1. **Fork the repository**.
2. **Create a branch** for your feature:
```bash
git checkout -b feature-name
```
3. **Commit your changes:**
```bash
git commit -m "Add new feature"
```
4. **Push and create a pull request:**
```bash
git push origin feature-name
```

## 🟢 Ideas for Improvement
- Add **retry logic** for failed requests.
- Implement a **timeout** for slow websites.
- Create a **CLI tool** with flags for customization.

---

Enjoy building with Go! 🚀 If you like this project, **give it a star on GitHub**! 🌟

Would you like me to help you implement any of the suggested improvements? Let me know! 🛠️

