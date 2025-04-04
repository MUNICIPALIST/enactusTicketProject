package main

import (
	"bytes"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/fogleman/gg"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	http.HandleFunc("/ticket", TicketHandler)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TicketHandler генерирует изображение с билетом и возвращает его как PNG
func TicketHandler(w http.ResponseWriter, r *http.Request) {
	// Предположим, что имя и уникальная информация для QR приходят в query-параметрах
	guestName := r.URL.Query().Get("name")
	if guestName == "" {
		guestName = "Айдана МАЛАДЕС"
	}
	qrData := r.URL.Query().Get("qr")
	if qrData == "" {
		qrData = "vincera.tech" // fallback
	}

	// Генерация итогового изображения
	ticketBytes, err := GenerateTicket(guestName, qrData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем как PNG
	w.Header().Set("Content-Type", "image/png")
	w.Write(ticketBytes)
}

// GenerateTicket создает изображение с QR-кодом и именем гостя
func GenerateTicket(guestName, qrData string) ([]byte, error) {
	// 1. Загружаем шаблон
	bgFile, err := os.Open("assets/template.png")
	if err != nil {
		return nil, err
	}
	defer bgFile.Close()

	bgImg, err := png.Decode(bgFile)
	if err != nil {
		return nil, err
	}

	// 2. Генерируем QR-код (Размер, Уровень коррекции ошибок и т.д.)
	qr, err := qrcode.New(qrData, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	// Превращаем в image.Image
	qrImg := qr.Image(540)

	// 3. Создаем новое "полотно" с gg на базе нашего шаблона
	// Получаем размеры шаблона
	width := bgImg.Bounds().Dx()
	height := bgImg.Bounds().Dy()

	dc := gg.NewContext(width, height)

	// Сначала рисуем фон
	dc.DrawImage(bgImg, 0, 0)

	// 4. Накладываем QR-код (примерно где-то внизу/справа/слева)
	//    Допустим, хотим разместить QR-код в нижней части
	offsetX := width - 256 - 752
	offsetY := height - 256 - 485
	dc.DrawImage(qrImg, offsetX, offsetY)

	// 5. Пишем имя гостя
	// Загружаем шрифт (если нужен кастомный)
	if err := dc.LoadFontFace("assets/georgia.ttf", 100); err != nil {
		return nil, err
	}

	dc.SetRGB(1, 1, 1) // Белый цвет

	// Позиция текста – выберите подходящее место
	dc.DrawStringAnchored(guestName, float64(width/2), float64(height-1300), 0.5, 0.5)
	dc.DrawStringAnchored("ID 1234", float64(width/2), float64(height-1170), 0.5, 0.5)

	// 6. Генерируем итоговое изображение в память
	outBuf := new(bytes.Buffer)
	if err := png.Encode(outBuf, dc.Image()); err != nil {
		return nil, err
	}

	return outBuf.Bytes(), nil
}
