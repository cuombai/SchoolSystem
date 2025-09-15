package utils

import (
    "bytes"

    "github.com/go-pdf/fpdf"
)

type PDFBuilder struct {
    doc *fpdf.Fpdf
    buf *bytes.Buffer
}

func NewPDF(title string) *PDFBuilder {
    pdf := fpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, title)
    pdf.Ln(12)
    pdf.SetFont("Arial", "", 12)

    return &PDFBuilder{
        doc: pdf,
        buf: new(bytes.Buffer),
    }
}

func (p *PDFBuilder) AddLine(text string) {
    p.doc.Cell(0, 10, text)
    p.doc.Ln(8)
}

func (p *PDFBuilder) Output() ([]byte, error) {
    err := p.doc.Output(p.buf)
    if err != nil {
        return nil, err
    }
    return p.buf.Bytes(), nil
}
