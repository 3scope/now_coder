package designpatterns

type File struct {
	Context string
}

func (t *File) Clone() *File {
	clone := *t
	return &clone
}

func (t *File) CloneNew() *File {
	return &File{Context: t.Context}
}
