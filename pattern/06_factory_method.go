package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Фабричный метод - порождающий шаблон, который предоставляет интерфейс
		 для создания объектов в суперклассе, но позволяет подклассам изменять тип создаваемых объектов.

Плюсы - 1. Гибкость подклассов. Паттерн позволяет подклассам выбирать
			тип создаваемого продукта, что делает систему более гибкой и расширяемой
		2. Разделение обязанностей: Создание объекта выносится в отдельный метод,
			что способствует соблюдению принципа единственной ответственности.
Минусы - 1. Усложнение структуры кода: Паттерн может создать большое количество классов,
			что усложняет структуру кода, особенно если существует множество семейств продуктов.
		 2. Неудобство в случае изменения иерархии продуктов: Если иерархия продуктов изменится,
		 	то придется изменять иерархию создателей, что может быть неудобным.
Реализация данного патерна представлена в виде приложения, которое обрабатывает изображения, и мы хотим
			добавить сохранение изображений в различных форматах.
*/

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// ImageSaver interface
type ImageSaver interface {
	Save(image image.Image, filename string) error
}

// JPEGImageSaver is a concrete product for saving images in JPEG format
type JPEGImageSaver struct{}

func (saver *JPEGImageSaver) Save(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, nil)
}

// PNGImageSaver is a concrete product for saving images in PNG format
type PNGImageSaver struct{}

func (saver *PNGImageSaver) Save(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// ImageSaverFactory is the creator interface
type ImageSaverFactory interface {
	CreateSaver() ImageSaver
}

// JPEGImageSaverFactory is a concrete creator for JPEG images
type JPEGImageSaverFactory struct{}

func (factory *JPEGImageSaverFactory) CreateSaver() ImageSaver {
	return &JPEGImageSaver{}
}

// PNGImageSaverFactory is a concrete creator for PNG images
type PNGImageSaverFactory struct{}

func (factory *PNGImageSaverFactory) CreateSaver() ImageSaver {
	return &PNGImageSaver{}
}

// func main() {
// 	jpegFactory := &JPEGImageSaverFactory{}
// 	jpegSaver := jpegFactory.CreateSaver()

// 	// simulate an image
// 	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

// 	err := jpegSaver.Save(img, "output.jpg")
// 	if err != nil {
// 		fmt.Println("Error saving JPEG image:", err)
// 	}

// 	// client code using PNG saver
// 	pngFactory := &PNGImageSaverFactory{}
// 	pngSaver := pngFactory.CreateSaver()

// 	err = pngSaver.Save(img, "output.png")
// 	if err != nil {
// 		fmt.Println("Error saving PNG image:", err)
// 	}
// }
