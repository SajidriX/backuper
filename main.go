package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "time"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: backup <file_or_directory>")
        os.Exit(1)
    }

    source := os.Args[1]
    info, err := os.Stat(source)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }

    home, _ := os.UserHomeDir()
    backupDir := filepath.Join(home, "backups")
    os.MkdirAll(backupDir, 0755)

    baseName := filepath.Base(source)
    timestamp := time.Now().Format("2006-01-02_15-04-05")
    destName := baseName + "_" + timestamp
    destPath := filepath.Join(backupDir, destName)

    if info.IsDir() {
        err = copyDir(source, destPath)
    } else {
        err = copyFile(source, destPath)
    }

    if err != nil {
        fmt.Printf("Backup failed: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Backup created: %s\n", destPath)
}

func copyFile(src, dst string) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }

    info, err := os.Stat(src)
    if err != nil {
        return err
    }
    return out.Chmod(info.Mode())
}

func copyDir(src, dst string) error {
    return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        relPath, _ := filepath.Rel(src, path)
        destPath := filepath.Join(dst, relPath)

        if info.IsDir() {
            return os.MkdirAll(destPath, info.Mode())
        }
        return copyFile(path, destPath)
    })
}
