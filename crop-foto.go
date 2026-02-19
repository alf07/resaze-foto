package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

type arrImage struct {
	Image_path string
	Class_name string
}

func generate_json(images []arrImage) {
	result, err := json.MarshalIndent(images, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(images)

	os.WriteFile("annotations.json", result, 0644)
}

func main() {

	//inputDir := аргумент командной строки 1
	//outputDir := аргумент командной строки 2
	inputDir := os.Args[1]
	outputDir := os.Args[2]
	//folder := os.Args[3]
	//folder := "1"

	files, err := os.ReadDir(inputDir)
	if err != nil {
		panic(err)
	}

	var resultJson []arrImage

	for _, file := range files {

		inputPath := filepath.Join(inputDir, file.Name())
		outputPath := filepath.Join(outputDir, file.Name())

		img, err := imaging.Open(inputPath)
		if err != nil {
			log.Printf("Ошибка обработки %s: %v", file.Name(), err)
			continue
		}

		// Координаты области обрезки, тут условие какая камера
		// 1 - алюминий микс
		// 14 - медь микс
		//x0, y0 := 500, 560   // начальная точка 1
		//x1, y1 := 1035, 1057 // конечная точка 1
		x0, y0 := 876, 604   // начальная точка 14
		x1, y1 := 1404, 1080 // конечная точка 14

		// Создаем прямоугольник для обрезки
		cropRect := image.Rect(x0, y0, x1, y1)
		// обрезаем
		img = imaging.Crop(img, cropRect)

		if err := imaging.Save(img, outputPath); err != nil {
			log.Printf("Ошибка сохранения %s: %v", file.Name(), err)
		}

		res := arrImage{
			Image_path: outputPath,
			Class_name: file.Name(),
		}

		resultJson = append(resultJson, res)

	}

	generate_json(resultJson)

}
