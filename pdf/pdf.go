package pdf

import (
	"log"

	"github.com/signintech/gopdf"
)

func TestPDF() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "times.ttf")
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
	}
	pdf.Cell(nil, "您好")
	pdf.WritePdf("hello.pdf")
}
