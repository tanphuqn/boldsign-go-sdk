package model

type DocumentFile struct {
	ContentType string `field:"content_type"`
	FilePath    string `field:"file_path"`
	Filesize    int64  `field:"file_size"`
}
