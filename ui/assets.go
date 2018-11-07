package ui

import (
	"path"
	"strconv"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Assets struct {
	rootPath   string
	fontStore  map[string]*ttf.Font
	imageStore map[string]*sdl.Surface
}

func NewAssetManager(filepath string) *Assets {
	return &Assets{filepath, make(map[string]*ttf.Font, 0), make(map[string]*sdl.Surface, 0)}
}

func (a *Assets) GetFont(name string, size int) (*ttf.Font, error) {
	filepath := path.Join(a.rootPath, name)
	hash := strconv.Itoa(size)
	if font, ok := a.fontStore[filepath+hash]; ok {
		return font, nil
	}

	font, err := ttf.OpenFont(filepath, size)
	if err != nil {
		return nil, err
	}

	a.fontStore[filepath+hash] = font
	return font, nil
}

func (a *Assets) GetImage(name string) (*sdl.Surface, error) {
	filepath := path.Join(a.rootPath, name)
	if image, ok := a.imageStore[filepath]; ok {
		return image, nil
	}
	tex, err := img.Load(filepath)
	if err != nil {
		return nil, err
	}
	a.imageStore[filepath] = tex
	return tex, nil
}
