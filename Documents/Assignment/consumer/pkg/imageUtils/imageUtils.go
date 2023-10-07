package imageUtils

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/anushka/producer/pkg/models"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
)

//todo code refactoring

func DownloadImage(productID int) (image.Image, error) {
	imgURL, err := models.GetProductImageByProductID(productID)
	if err != nil {
		logrus.Error("Error in getting image from db")
		return nil, err
	}

	if imgURL == "" {
		logrus.Info("No image found for this product")
		return nil, nil
	}

	resp, err := http.Get(imgURL)
	if err != nil {
		logrus.Error("Error in getting image from URL")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Failed to download image. Status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to download image, status code: %d", resp.StatusCode)
	}

	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Error reading image data")
		return nil, err
	}

	img, _, err := image.Decode(strings.NewReader(string(imgData)))
	if err != nil {
		logrus.Error("Error decoding image")
		return nil, err
	}

	logrus.Info("Image downloaded successfully from URL")
	return img, nil
}

func ResizeImage(img image.Image) (image.Image, error) {
	logrus.Info("Resizing image")
	imgResized := resize.Resize(1024, 0, img, resize.Lanczos3)
	op, err := os.Create("compressed_images.jpeg")
	if err != nil {
		logrus.Error("Error in creating resized image")
		return nil, err
	}
	defer op.Close()
	err = jpeg.Encode(op, imgResized, nil)
	if err != nil {
		logrus.Error("Error in encoding resized image")
		return nil, err
	}
	return imgResized, nil
}

func CompressImage(img image.Image, quality int) ([]byte, error) {
	buffer := new(strings.Builder)
	err := jpeg.Encode(buffer, img, &jpeg.Options{Quality: quality})
	if err != nil {
		logrus.Error("Error in compressing image")
		return nil, err
	}

	return []byte(buffer.String()), nil
}

func SaveImageToLocal(filename string, data []byte, dir string) (string, error) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(dir, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return "", err
	}

	logrus.Infof("Image saved to file %s", filePath)

	return filePath, nil
}

func DownloadAndCompressProductImages(productID int) (resp string, err error) {
	img, err := DownloadImage(productID)
	if err != nil {
		fmt.Println("Error in downloading image")
		return
	}
	if img == nil {
		fmt.Println("No image found for this product")
		return
	}
	fmt.Printf("Image downloaded successfully.")

	resizedImage, err := ResizeImage(img)
	if err != nil {
		fmt.Println("Error in resizing image")
		return
	}
	compressedImage, err := CompressImage(resizedImage, 80)
	if err != nil {
		fmt.Println("Error in compressing image")
		return
	}
	filename := fmt.Sprintf("compressed_image_%d.jpg", productID)
	filepath, err := SaveImageToLocal(filename, compressedImage, "compressedImages")
	if err != nil {
		logrus.Error("Error saving compressed image to file:", err)
		return "", err
	}
	logrus.Infof("Image saved to file %s", filepath)

	return filepath, nil

}
