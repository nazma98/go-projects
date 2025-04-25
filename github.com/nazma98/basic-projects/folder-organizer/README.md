# 📁 Folder Organizer

A simple, efficient CLI tool written in Go that automatically organizes your messy folders by moving files into subfolders based on their file types!
Perfect for cleaning up your Downloads, Desktop, or Projects directory in one go. 🧹

🚀 Features
- ✅ Detects file types by extension (e.g., .pdf, .jpg, .mp3)
- ✅ Automatically creates folders if they don't exist
- ✅ Moves files into the right places
- ✅ Super lightweight and fast
- ✅ Written in beginner-friendly Go code!

## 📸 Before & After

```
mathematica
Copy
Edit
Before:
📂 Downloads/
├── report.pdf
├── music.mp3
├── picture.jpg
├── notes.txt

After running the tool:
📂 Downloads/
├── Documents/
│   └── report.pdf
├── Music/
│   └── music.mp3
├── Images/
│   └── picture.jpg
├── Text/
│   └── notes.txt

```

## 🛠️ Requirements
Go installed (v1.18 or newer)

## 🧑‍💻 How to Run
- Clone the repo:
```
git clone https://github.com/your-username/folder-organizer.git

cd folder-organizer
```
- Edit the target path:

In main.go, change this line to your target folder:

```
basePath := "/home/yourname/Downloads"
```
- Run it:
```
go run main.go
```
✨ Boom! Your files are sorted like magic.

## 🧠 How It Works
Uses `os.Stat` to check if folders exist.

Uses `filepath.Ext()` to detect file extensions.

If the target folder doesn't exist, it creates one with `os.Mkdir`.

Then moves the file using `os.Rename`.

Code is kept simple and readable so you can learn from it!

## ✨ Future Ideas (if you're feeling spicy)

- 🔄 Watch the folder in real-time (auto-organize as files arrive)
- 📅 Organize by date (e.g., 2025/April)
- 💻 GUI version using web or desktop framework
- 🚫 Ignore certain file types or hidden files

