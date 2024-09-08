package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: pdf-to-jpg <pdf-file>")
        os.Exit(1)
    }

    pdfPath := os.Args[1]
    outDir := "output"

    // Check if the input PDF file exists
    if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
        fmt.Printf("Error: The file %s does not exist.\n", pdfPath)
        os.Exit(1)
    }

    if err := os.MkdirAll(outDir, 0755); err != nil {
        fmt.Println("Error creating output directory:", err)
        os.Exit(1)
    }

    // Convert PDF to JPEG using ImageMagick
    err := convertPDFToJPEG(pdfPath, outDir)
    if err != nil {
        fmt.Println("Error converting PDF to JPEG:", err)
        os.Exit(1)
    }

    fmt.Println("Conversion completed. Check the output directory.")
}

func convertPDFToJPEG(pdfPath, outDir string) error {
      
    // Extract base name from the PDF file
    baseName := strings.TrimSuffix(filepath.Base(pdfPath), filepath.Ext(pdfPath))
    outputPath := filepath.Join(outDir, fmt.Sprintf("%s-page-%%d.jpg", baseName))

    cmd := exec.Command("convert", "-density", "200", pdfPath, outputPath)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to execute command: %s, %s", err, output)
    }

    return nil
}
