# PDF to JPG Converter

A simple command-line tool written in Go that converts PDF pages into high-quality JPG images. This tool extracts each page of a PDF file and saves it as a separate JPG image with high resolution.

## Installation

### Using Go Install

To install the tool using `go install`, run the following command:

```bash
go install github.com/stilyanpetrovv/pdf-to-jpg@latest
```
```bash
./pdf-to-jpg <pdf-file>
```

or clone the repo and use it as it 
```bash
git clone https://github.com/stilyanpetrovv/pdf-to-jpg.git
cd pdf-to-jpg
go run pdf-to-jpg.go <pdf-file>
```
replace pdf-file with the name of the pdf file you want to convert

## License

This project is licensed under the [MIT License](https://github.com/stilyanpetrovv/pdf-to-jpg/blob/main/LICENSE).
