package cloudinary

import (
	"io"

	gocloud "github.com/WitzHsiao/go-cloudinary"
	"golang.org/x/net/context"
)

func UploadStaticImage(ctx context.Context, fileName string, data io.Reader) (string, error) {
	c, _ := FromContext(ctx)
	return c.UploadStaticImage(fileName, data, "")
}

func Resources(ctx context.Context) ([]*gocloud.Resource, error) {
	c, _ := FromContext(ctx)
	return c.Resources(gocloud.ImageType)
}

func ResourceURL(ctx context.Context, fileName string) string {
	c, _ := FromContext(ctx)
	return c.Url(fileName, gocloud.ImageType)
}

func DeleteStaticImage(ctx context.Context, fileName string) error {
	c, _ := FromContext(ctx)
	return c.Delete(fileName, "", gocloud.ImageType)
}
