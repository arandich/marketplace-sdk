package image_cdn

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

type IClient interface {
	UploadImage(ctx context.Context, image []byte, name string) (string, error)
}

type Client struct {
	httpClient *http.Client
	cfg        Config
}

func New(httpClient *http.Client, cfg Config) IClient {
	return &Client{
		httpClient: httpClient,
		cfg:        cfg,
	}
}

func (c Client) UploadImage(ctx context.Context, imageData []byte, name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("empty name")
	}
	if len(imageData) == 0 {
		return "", errors.New("empty image")
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	defer writer.Close()

	part, err := writer.CreateFormFile("media", name)
	if err != nil {
		return "", err
	}

	_, err = part.Write(imageData)
	if err != nil {
		return "", err
	}
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("http://%s:%s/image", c.cfg.Host, c.cfg.Port), body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("bad status code on image upload: %s", resp.Status)
	}
	location := resp.Header.Get("Location")
	if location == "" {
		return "", errors.New("missing Location header in the response")
	}

	return location, nil
}
