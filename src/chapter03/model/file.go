package model

import "fmt"

type File struct {
	//...
	FileName string
}

func (f *File) Read(buf []byte) (n int, err error)  {
	fmt.Println("Read")
	return 5, nil
}

func (f *File) Write(buf []byte) (n int, err error)  {
	fmt.Println("Write")
	return 5, nil
}

func (f *File) Seek(off int64, whence int) (pos int64, err error)  {
	fmt.Println("Seek")
	return 5, nil
}

func (f *File) Close() error {
	fmt.Println("Close")
	return nil
}


type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWriter interface {
	Write(buf []byte) (n int, err error)
}

type IClose interface {
	Close() error
}
