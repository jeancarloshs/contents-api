package model

import "errors"

type UploadToMinIO struct {
	BucketName  *string `json:"bucket_name"`
	ObjectName  *string `json:"object_name"`
	FilePath    *string `json:"file_path"`
	ContentType *string `json:"content_type"`
}

// Função para validar os campos obrigatórios da estrutura UploadImage
func (ui *UploadToMinIO) Validate() error {
	if ui.BucketName == nil || *ui.BucketName == "" {
		return errors.New("bucket_name é obrigatório")
	}
	if ui.ObjectName == nil || *ui.ObjectName == "" {
		return errors.New("object_name é obrigatório")
	}
	if ui.FilePath == nil || *ui.FilePath == "" {
		return errors.New("file_path é obrigatório")
	}
	// Se o ContentType for nil, o valor padrão pode ser aplicado (por exemplo, "image/png")
	if ui.ContentType == nil {
		defaultContentType := "image/png"
		ui.ContentType = &defaultContentType
	}
	return nil
}
