package imageUtils_test

import (
	"testing"

	imageUtils "github.com/anushka/consumer/pkg/imageUtils"
)

const (
	imgURL = "https://raw.githubusercontent.com/anushka2911/images/main/uploads/headphone.jpg"
)

func BenchmarkDownloadImage(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := imageUtils.DownloadImage(imgURL)
		if err != nil {
			b.Fatalf("Error downloading image: %v", err)
		}
	}
}

func BenchmarkResizeAndCompressImage(b *testing.B) {
	img, err := imageUtils.DownloadImage(imgURL)
	if err != nil {
		b.Fatalf("Error downloading image: %v", err)
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := imageUtils.ResizeAndCompressImage(img, 80)
		if err != nil {
			b.Fatalf("Error resizing and compressing image: %v", err)
		}
	}
}

func BenchmarkSaveImageToLocal(b *testing.B) {
	img, err := imageUtils.DownloadImage(imgURL)
	if err != nil {
		b.Fatalf("Error downloading image: %v", err)
	}
	compressedImage, err := imageUtils.ResizeAndCompressImage(img, 80)
	if err != nil {
		b.Fatalf("Error resizing and compressing image: %v", err)
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := imageUtils.SaveImageToLocal("test_image.jpg", compressedImage, "test_dir")
		if err != nil {
			b.Fatalf("Error saving image to local: %v", err)
		}
	}
}
