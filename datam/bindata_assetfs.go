// Code generated by go-bindata.
// sources:
// index.html
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x3a\x6d\x6f\xdb\x36\xb7\xdf\xf3\x2b\x78\xd5\x8b\x6b\xbb\xb3\xe5\x74\xed\x80\xc1\x49\x3c\xa4\x49\xb6\x75\x48\x1b\xa0\xed\xf6\x25\x0d\x0a\x5a\xa4\x6d\xb6\xb2\xe4\x52\xb4\x93\xdc\x22\xff\xfd\x39\x87\x6f\x22\x65\x39\x71\x9a\x3d\x29\xb6\x48\xe4\x79\x3f\x87\xe7\x85\xca\xe1\x5c\x2d\xf2\xf1\xde\xde\xe1\x9c\x53\x36\xde\x23\xe4\x50\x09\x95\xf3\xf1\x29\x55\x94\xbc\xa5\x05\x9d\x71\x79\x38\x34\x6b\xb8\x5b\x65\x52\x2c\x15\x3e\x12\x22\xa6\xa4\xab\x6e\x97\xbc\x9c\x92\x45\xc9\x56\x39\x27\x47\x47\x47\xa4\x53\x4e\xbe\xf0\x4c\x75\x7a\xe4\xbb\x86\x22\xe4\x5a\x14\xac\xbc\x4e\x1d\x8c\x05\x3e\xb0\xbb\x7e\x79\x55\x30\x3e\x15\x05\x67\x66\xe7\x0e\xd9\x0d\x6b\x7e\x96\x35\x41\x8e\x47\x89\xe2\x37\x6a\xf8\x85\xae\xa9\x59\x4d\x48\x25\xb3\xa3\x64\x98\x8b\xc9\xf0\xcb\xb7\x15\x97\xb7\xe9\x42\x14\xe9\x97\x2a\x19\x3f\x81\xc6\x60\x25\xfe\x0d\x32\xa9\xa2\x93\x9c\xb3\x82\x35\x08\x3d\x86\xd2\xa4\x2c\x55\xa5\x24\x5d\x3e\x4d\xa0\x59\x99\x33\x5e\xe4\xf4\xb6\x5c\xa9\x16\x4a\x8f\x21\xc5\x5e\xa6\xeb\x57\x4f\x93\x86\xbd\x1c\x80\x81\x56\x3c\x5d\xff\xfc\x64\x42\x92\x03\xa9\x4a\xa5\xeb\x17\x4f\x23\x25\x18\x2f\x67\x92\x2e\x9e\x20\x0c\x55\xe9\xec\x81\xc8\x31\xd0\xcb\x7c\x35\x13\x45\x35\xe4\x4c\xe8\x10\x99\x49\xc1\xa2\x97\xfb\x68\xd4\x67\x30\x3a\x5f\xbd\xfa\x40\x45\xeb\x07\xd1\x71\xc2\x97\x5c\x14\x5f\x89\xe4\xf9\x51\x52\xa9\xdb\x9c\x57\x73\xce\x41\x8f\xb9\xe4\xd3\x9d\x44\xcb\xaa\x2a\x09\x8d\xa1\xdf\x17\x00\x42\x81\x60\x26\x39\x2f\x92\xf1\x83\x6c\x00\xa9\x11\xda\x48\x66\x37\xbc\x30\x96\x07\x13\x5a\xf1\x1f\xc5\xcd\xc5\x6c\xae\x06\x6a\xce\x17\x8f\x21\x11\xa7\x88\xdd\xf1\x7c\x80\xed\x8e\x02\x11\x35\x80\x88\x6a\x17\x54\xe3\x98\x60\x48\xd9\x8c\xa4\x19\x11\xc5\x72\xa5\x2e\xb5\x6b\xd0\x33\x57\x90\x88\x6d\xae\x05\x4e\x7c\x30\xe7\x48\x66\x44\x8a\x52\x2e\x68\xee\x92\x6d\x48\x80\x89\xb5\x4f\xde\x93\xf2\x66\x50\x89\xff\x17\xc5\x6c\x44\xb2\xb2\x50\xbc\x00\x63\x97\x37\x11\xda\xfc\x85\x07\x9f\x02\x08\xc2\xf3\x11\x79\xf1\xcb\x32\x06\xa3\x1e\x2a\x2b\xf3\x52\x8e\xc8\xb3\x5f\x5e\xfd\xca\xb2\x57\xae\x12\xa0\xb0\x03\xc6\xb3\x52\x52\x25\xca\x02\x25\x2c\x78\x4c\x61\x34\x2f\xd7\x5c\x7a\x3a\x1b\x18\x58\x43\x24\x6a\x19\xa1\xe9\xa0\x4d\xc1\x8a\x14\x23\x37\x50\x4d\x02\xf4\x00\x84\xc9\xe9\xb2\xe2\xa8\x9f\x79\x3a\x88\x00\x40\x93\xe5\x0d\xa9\xca\x1c\x50\x9f\x9d\x9c\xbc\x3e\xf0\x65\x8d\xa9\xf9\x88\xfc\xba\xbf\xdf\xd0\xb3\xc1\x4e\xb1\x7e\xeb\xf2\xdc\x0b\xb2\xa4\x8c\x69\x03\x7b\x8b\xb5\xf2\x3e\xdb\xc7\x7f\xf7\xb2\xaa\x69\x4e\x68\xf6\x75\x26\x4b\x30\x08\xd8\xf9\xec\x17\xfc\x17\xd9\x99\x42\x30\x81\xc1\x72\x3e\x55\x11\x45\x1d\x3c\xa9\x28\xd6\x34\x0f\x4d\x15\x50\x93\xae\x42\xd7\x7e\xfc\xfd\x14\xff\x45\x75\xdb\x45\xe5\xfd\xd1\xad\x97\xf6\x87\x55\x01\x79\xca\x44\xf4\xe1\xd0\x74\x22\x7b\x87\x93\x92\xdd\x6a\x0a\xff\x33\x18\x90\xb7\x25\xa3\x39\x19\x0c\xf4\x02\xc6\x67\x96\xd3\xaa\x3a\x4a\x16\x7a\x7d\x4a\x19\x4f\x88\x60\xf0\x7e\xab\x21\x13\x22\xcb\x1c\x32\x13\x64\xa3\xbc\x9c\x25\xe6\x78\x6c\xe0\x0d\xfc\xb6\xd5\x27\x60\xe5\x62\x7d\x30\x76\x7b\x1b\xd8\x16\x24\x71\x10\x6d\x30\xa8\x0c\x97\x01\x08\x00\x4d\x56\x4a\x95\x85\xcd\x9d\xe6\x25\x71\x58\x59\x5e\x56\xa0\x0a\x7a\x14\xa4\xab\x16\xc2\x93\x4a\xc6\xff\xa7\xc4\x82\x57\x07\x87\x43\x83\x13\xd1\x9c\xbf\x8a\xf9\xea\x8e\x2d\x19\x1b\x5d\xfe\xd4\x42\x80\x65\x5f\x05\xa2\x0e\x41\xd6\xfb\x24\x47\xf3\xc7\x72\xeb\x60\x73\x40\xfa\x25\xda\x47\x08\x19\x2f\xe0\x12\x1b\x9f\xbe\x86\x0e\x92\xb5\xed\x1c\xea\x68\x0b\xaa\x88\x75\x22\x0a\x70\x3a\xc1\xb2\xa7\x01\x36\x50\x37\xc8\xc1\x8a\xdc\x49\x96\x73\x3a\xe1\xf9\x8f\x88\xa3\x11\xff\x1b\x12\xfd\xfd\xfe\xfc\x47\xe4\xf9\x5b\x8a\x27\x48\x03\xef\xe8\x3f\x1f\xf8\x3b\xc4\xc3\x14\xca\xf4\x63\x22\x79\xa2\x0a\x02\xff\x41\x7a\x9e\xd2\x55\xae\xda\x63\xba\xd6\xe7\x98\xb1\x64\x0c\xff\x6b\x8d\xee\x27\xf3\x19\x9f\xe0\xb9\xda\xa4\x1d\x69\x6d\x5f\xf6\xa2\x8d\xfa\x01\x73\x03\x2f\x18\x31\xe3\x8e\xcd\x46\x2e\x1d\xa1\x1e\xd7\xd0\xc3\x2c\xb5\x89\xea\x74\xa3\xf5\xe3\xc5\xea\x04\x72\x05\x85\xba\x24\xa1\x53\xc4\x9c\x77\x94\xb8\xa4\xff\x72\x79\xe3\x8d\xea\x14\x45\xac\x19\x57\x7f\x55\xa0\xe7\xf8\xb4\xbc\xce\x4b\xca\x08\xad\x08\x2e\xc4\x4a\x04\x0a\x78\x7e\xa6\xad\xd9\xe4\x88\xa5\x73\x9a\x97\xd7\x83\xdb\x11\x34\x68\x50\xeb\x0e\x5a\x65\xa8\xc5\xae\x2a\x18\xff\x92\x86\x79\x1a\x50\x68\xed\xb7\x74\xd9\x06\x15\x3e\x52\x03\x5c\x5e\x17\xa8\xcb\x71\x91\xcd\x4b\x79\x96\xf3\x85\x17\x0e\x1c\xb6\x04\xc1\x47\x58\xf3\x31\xb0\x69\x60\x7c\x50\x59\xd7\x82\xa0\xf3\xb5\xbd\x6d\xa9\xc9\x41\xaf\x3b\x5d\x15\x19\xb6\x00\x5d\x37\x76\xae\xa9\x24\xa2\x3a\xae\x94\xc8\x61\x1a\x95\x60\x54\x00\xa2\x79\xc5\xf7\xfc\x36\xb2\xa4\x2a\x9b\xc3\x0e\x0c\x33\xee\xad\x9b\x54\xe0\xe5\xa4\x4f\x92\x72\xc9\x8b\x29\xa0\xe3\x33\x83\xc3\x0f\xbf\x16\x60\xc3\xa4\xa7\x29\x58\x09\x42\x22\xf6\x51\x6f\xb3\x32\x5b\x81\xe3\x55\x0a\x26\x3e\x5b\xc3\xc3\xb9\xa8\xa0\x5a\x70\xd9\xed\xd0\x5a\x28\x18\x5b\x28\xbb\xed\xf4\x37\xe5\x27\x64\x38\x24\x1f\xe7\xa2\x02\x46\x79\x0e\x9d\x1b\x62\x13\x55\x12\xeb\x96\x8a\x80\x98\x8a\x4c\x6e\xc9\x1f\x17\x16\xa1\xa9\xae\x92\x2b\xdf\xca\x04\x3c\x53\x43\xab\xeb\x59\x5a\x8a\x35\xe7\x40\x7a\x88\x42\x74\x13\x3c\xbe\xbe\x7d\xc3\xba\x1d\x0b\xdb\xe9\x41\x8b\x00\xca\xfc\xf9\xf1\xed\x39\x4e\xf4\x66\xd5\x65\x93\xbb\x9e\x63\xeb\x4c\x02\x8e\xf2\x76\xf5\x7c\x59\xc8\x31\x14\x10\x01\x61\x77\x83\x9a\x71\x1a\x96\x6c\xe0\x89\x03\x9f\x90\xbc\x9b\x70\x8b\x96\xf4\x52\xc9\x17\x90\xa4\x52\x03\x73\xe0\xc4\x19\x3e\x27\xce\x95\xa4\x4f\x40\x23\x68\xb8\xa0\x53\x9a\xca\x72\xe1\xc8\xf5\x09\x85\x93\x9d\xd1\x3c\xc7\x4e\x87\x3c\x1f\xb6\xc9\x1f\xc4\x83\xd7\xe1\x73\x64\x35\x4d\x2c\xad\xe6\xe5\xf5\x05\xc0\x9e\xea\xd7\xee\xf7\x20\x93\x2d\x25\x10\x91\x4a\xf0\x6a\x44\x2e\x3b\x48\xf0\x77\x20\xd8\xb9\xea\x07\x30\xc0\x02\x12\x2d\x02\x7c\x8f\x72\x38\xe4\x7f\x5e\x54\xc0\x54\xe3\xaa\x1b\xd5\xb9\x0a\xf6\xef\xea\x97\xbb\x6d\x36\xc6\x86\xad\x80\x66\x92\xa7\x28\x97\x37\x30\xfe\x7c\x0e\x5f\xef\x6a\xd3\x1b\x1b\xda\x85\xc8\x1a\xe6\x44\xb4\x31\x0a\x99\x58\x30\x4f\x1c\x5d\x58\xd0\x05\x8e\xa7\x2c\x65\x13\x7c\x04\xaf\xc1\xc9\xcf\xc0\x93\x9f\xb1\x89\xfd\x8c\xe7\x2c\xe9\xb5\xd0\x42\xe0\x5e\xd3\x31\xe8\x33\x1f\x5a\x0a\xce\x4b\x9f\xfc\xf5\xe1\xe2\x5d\x0a\x13\x25\x64\x36\x31\xbd\x0d\xec\x9f\x64\x25\xf4\x89\x23\x7b\x9a\xeb\x65\x23\x07\x6c\xe0\xaf\x3e\x1c\xbc\x8f\x17\xa7\x17\xe0\x2b\x3e\x15\x37\xbf\xd5\x60\x82\x01\x08\x4b\x05\xf4\xf3\xde\x3e\xbd\xad\xf6\xd1\xa9\xa2\xdd\x40\x30\x3c\x83\x20\x60\x02\x03\xd4\x66\xb6\xb7\x17\xff\x9c\x85\x76\x7b\x84\xbe\xac\x16\xca\xa7\xba\x05\x57\x14\x33\x35\xf0\xbc\xbc\x32\xc7\xc9\x2d\xa5\xcb\x55\x35\xf7\x46\x42\x0b\x80\x81\x20\x84\xa5\x32\x2d\x8f\xb3\x53\x8e\x6f\xb0\x75\x7e\xfc\xfa\xec\xdc\xaf\x22\x05\x2c\xcb\x88\xa3\x25\xf0\x3b\xee\xb2\x60\xa4\xf3\x90\x95\x68\x07\xd6\x2b\xb9\xc9\x13\x7a\xa4\x47\x71\x34\x89\x3e\xf0\xcc\xbd\x1c\xa9\x76\xd0\x06\xd3\xe3\x93\x8f\x6f\x2e\xde\xb5\xf1\xc5\x3b\xd3\x87\xb8\xd6\xb6\x2f\xf8\xf5\x1f\x38\x9c\x05\x45\x0a\xe3\xfc\x9d\x8e\x35\x8d\x8a\x8f\xcd\xeb\xd2\x4b\x07\x73\x05\x88\x40\x82\x9c\x59\x3e\x48\xab\xeb\xd1\xfa\xc1\xf1\xd6\x8b\xe7\x50\x0f\x39\x4c\x69\x2d\x05\x65\x23\xc8\x02\x84\x24\x38\xfd\xfd\x98\xe2\x7b\x8e\x43\xf5\xa3\x68\x3a\x94\xc4\xa7\x6e\x4d\x0d\xa2\x15\x72\xbb\x3a\xe1\x79\x6e\x21\x64\xd7\x5b\x5f\x2b\x19\x6d\xc5\xe9\x4f\xea\xe5\x40\x88\x0c\x60\xfb\x60\xe1\x7c\xc5\x63\x79\xf0\xc7\x9e\xe1\xf7\x1c\x4f\x18\x39\x87\x3e\x88\xfc\x8e\xf9\x3e\xbc\xd7\x4e\x1b\x38\x48\x30\x2a\x6b\x09\xb4\x2d\x65\x91\xe5\x22\xfb\x7a\xf4\x29\xc1\xab\x36\x50\x75\x2a\xe4\xa2\xdb\x39\x96\x9c\x40\x97\x45\xaa\x95\x7d\xb8\xa6\x50\x8d\xa1\x3c\x43\x6a\xe1\x8a\x6b\x5d\xa1\x5b\x54\xf2\xf6\x37\xd2\xe9\x81\x7c\xf1\x01\xee\x00\x58\xa7\xaf\x8f\xef\x77\x93\x7d\x46\x9d\x84\xfc\x44\x9c\xd3\xe1\x31\xe9\xf4\x05\x33\xab\xcd\x88\xc0\xaa\xfc\x0f\xea\x7d\xac\xb4\x15\x52\x59\x5e\xbf\x01\xeb\xdc\xf4\xc9\x7e\x4f\xa3\xc2\x49\x73\x58\x48\xa1\xa1\x28\x89\xf9\x5c\xe9\xba\xb9\x86\x04\x0c\xaf\x11\x3d\xdc\xee\xdd\x1d\x7c\x72\x5d\xda\xa7\x24\x5b\xc9\x0a\x46\xfe\x65\x29\x60\xf6\x95\x9f\x92\x31\xe2\x24\x87\xa0\x5b\x01\xe3\x8a\xed\xc9\x3b\xb3\xfc\x76\x39\x17\x60\x2c\xe2\x9f\x06\x86\x47\x07\x67\x16\xbc\xce\xd4\x08\xd0\xe5\x25\x07\x91\x70\x77\x61\x49\xeb\x05\xe1\x13\x44\x25\x74\xdf\x3c\x3f\x99\xd3\x62\x16\x05\x25\x0a\xcd\xc0\x04\x59\x99\xeb\xdf\x65\xce\xb4\x91\x74\x64\xd9\x27\x80\x89\x63\xe5\xf3\xff\x76\x7d\xa3\xdb\x6b\x38\x7f\xe9\xd4\x29\xbf\x76\xc6\xef\x20\x3a\x75\xac\x41\x9b\x45\xb4\x57\x1c\x55\x6d\xc3\xc3\xe1\x32\x56\xa5\x91\xb0\x6d\x35\x30\x09\x3b\x0e\xd6\xba\xfc\x38\xaf\xb4\x57\xc4\x7e\x8c\x85\xed\x0b\xe0\x78\x2d\xe3\x5d\x55\x62\x39\xb3\x12\x46\x36\x6d\xb3\xe8\x12\xd4\xff\x60\x2e\xed\xf6\xf7\x83\xb6\xcb\x24\xb3\xbd\x20\x99\xf9\x22\xe2\x57\x32\xa9\x30\x25\x1d\xf8\x85\x4a\x2c\x96\xfa\xfe\xb9\x25\x5d\x68\x84\x39\x95\x2a\xdc\xae\x74\x07\x07\x4f\xa1\x6b\xfc\x62\xca\x29\x34\xe5\xdb\x1a\x1a\x24\x68\xee\x58\x4c\x17\x6f\xd0\xba\x68\xe7\x5e\x8a\xe3\x18\x74\x92\x30\x5b\xac\xc1\xb9\xda\x99\x1c\x5e\xed\x9d\x4c\x5f\x17\xa6\xb0\x09\x32\x1b\x1e\x6d\xfe\x32\xc4\x32\x37\x2a\x16\x29\xc5\x69\xbc\xcb\xd2\xb0\x25\xf1\x26\x82\xc9\xe8\x21\x59\xa8\x52\x90\xfa\x84\x1e\x33\x3e\xc3\xd2\x67\x0c\xa8\xad\xe4\x6c\xc3\xd4\x84\x6c\x00\xfa\xb6\xaa\xb1\x67\xeb\x8f\x0f\x26\x8f\x8f\xbc\x75\xba\xbe\x87\xfb\xcc\x54\x2e\x9b\x4b\x36\x48\x5c\x35\xc0\x17\xa6\xaf\x83\x20\x69\x6a\x11\xc5\x4d\xbd\x2e\x60\x71\x3f\x6c\x7e\x4b\x49\xba\xb8\xf1\x95\x88\x82\x2c\x9a\x89\x7d\xb3\x86\xbb\x1f\xc8\x95\x40\xec\x27\xf2\xa2\xdf\xd8\xd0\x87\x16\x5a\xe6\x26\x02\x84\x98\x6f\x72\x46\xe4\x6b\x13\x8d\x10\x68\x44\x46\x64\x71\xf9\xf5\xaa\xb1\x73\x17\x27\xac\xd0\x62\xe1\x96\xfe\x48\x82\xd3\x69\x43\xd4\xc4\xf5\x22\x70\x40\xdd\x63\xe3\xf0\xda\xdd\x86\x19\xef\xa2\x62\xaa\xc9\x9b\x9a\xa8\xbd\x6b\xc3\xa4\x6f\xb0\x71\x37\x2e\xbe\x68\x53\x73\x85\xb3\xfb\x41\xb1\x57\x3e\x9b\x07\x05\x89\xc1\x5c\xfb\x5a\xe9\xa9\x5a\x43\x79\x1a\xf6\x7a\xc6\x1e\x91\xe4\x98\x31\x82\x33\x4e\x12\xe2\x0f\x9f\x47\x0a\x3f\x74\xbf\x63\xef\x74\x54\x39\x9b\xe1\x55\x81\xbd\x3a\x32\x6b\x54\x42\x31\x3c\x4a\x9e\xb9\x8b\xdf\x31\x4e\x5e\xe6\x0e\xb7\xed\x2a\xc9\x4f\x76\x5a\x8c\x21\xf9\x00\xc9\x88\x33\x72\x01\xfd\x81\xc6\x69\x9a\x0c\xb4\xfc\xfb\xfd\x9b\x7b\x15\x8d\x54\xb1\x27\x1b\x35\xc1\xb3\x7d\x1f\x4c\xa0\x93\xb9\x5a\x40\xf1\xef\x81\xd4\x9a\x22\xa4\xd7\xb5\x01\x5c\x1b\x1c\x44\x6e\x6e\xe2\x48\xa2\x3b\x99\xa4\xf5\xb2\xc1\xfd\x40\x39\x7c\xe6\xae\x5f\x7b\x29\x9c\x9d\x96\x24\x87\x3f\x36\xeb\x83\x55\x66\x3e\xf9\xbb\x9f\xe8\x48\x0c\x87\x1c\x6f\x3f\x08\x93\x62\x6d\x2e\x30\x70\x00\x37\x17\x68\xa2\x98\xe2\x37\x28\x9d\xe5\xf7\x02\x1c\x13\x5a\xbb\x89\xdc\xa8\xb3\xc1\x8c\x6e\x6a\xed\xb6\xc2\x61\x64\x8b\x5a\x57\x37\xfe\x27\x7d\xd6\x54\x57\xe7\x57\x3c\x38\x97\xfb\x57\x69\xb5\xcc\x05\x18\x7a\x98\x34\xa1\xf4\x99\xb4\xe1\xa1\x9f\x75\xba\xca\x79\x31\x53\xf3\x66\xd5\xc7\x9f\x4e\x9d\x81\x3a\x50\xad\x2f\x0b\x0b\x4b\x06\xe4\xc5\xd5\x66\x42\xea\x40\x46\x02\x38\x94\xa1\x99\x92\xec\x11\x3d\x68\xac\xff\xd0\x94\xec\x7e\xfc\xb4\x0c\xee\x48\x36\xa5\xd9\x71\x6a\xf6\xe0\x7a\x7a\x7e\x50\x47\x18\xff\x44\x62\x74\x6c\x6e\xba\x61\xbb\x7e\xdf\xd2\xd7\xf8\xdb\x0b\xfb\x5b\x72\xb5\x92\x85\xe9\x3b\xea\x6f\x54\x41\xac\xe3\xdd\x73\xef\xa1\x70\xd3\x55\x6c\x02\x31\xb0\x79\x44\xc2\xcb\x0d\x7d\x26\x3c\x88\x19\xa1\x37\xa1\x40\xcb\x10\x0c\x2f\xf3\x63\xa0\xa7\xdd\x6f\xc4\x1e\xab\x3d\xc5\x26\x0f\xdc\x6e\x08\x16\xa0\x19\x57\xc0\xff\xfb\x7b\x4d\x0f\xd8\xe3\xef\x82\xdd\xbd\x6e\x89\xf7\x38\xce\x43\x1e\x36\xa4\x81\x87\x63\x11\x45\xb2\x75\xe5\x70\x88\x36\x33\x12\x57\x75\xa3\x99\xb9\x1b\xee\xa8\xa2\x25\xcf\xdc\x95\x74\xaf\xee\x5a\xb1\x2f\x31\x7d\x69\xd7\xb5\xb5\xb8\xae\x6f\x5f\xc3\x56\x94\x4b\x59\x4a\x18\x13\x78\xb5\xca\x55\xd5\x7e\xa5\xe5\x36\xed\x16\x8e\x85\x1a\xad\x07\xee\x81\xf9\x82\xe8\x97\x83\x66\xd4\xb8\x6b\x17\xe2\x88\xa7\x90\xfb\xce\xe2\xd6\xb6\x4f\x44\x74\x9d\x38\xd9\x6c\x75\x12\xef\x49\x7c\xa8\x2e\x45\x74\x4c\x7c\xe7\xd0\x72\x18\xc2\xf0\x83\xae\x0f\x44\xf2\xf6\xf3\x25\x6d\x95\x87\xe5\xbf\xa0\x6b\xdf\xef\x86\x8b\x50\x8a\x26\x55\xa3\x2b\xd0\xa9\xd4\x52\xd3\x21\x0b\x30\x0d\x8e\x15\x5e\x49\xd3\x89\xf5\xd3\x31\x86\x75\x2e\x82\x1c\xaa\xa3\xa7\xcb\x26\xc1\x0a\xc7\x21\xb3\x1b\x2c\x38\x41\x63\xc4\x36\x89\x5d\x3d\x41\xc6\x1e\x8d\x86\x58\xa6\xb6\xe2\x87\xe8\xad\x97\xd2\x3e\x77\x24\xcf\x5a\xda\xee\xbb\x0d\x62\x8d\x92\x0e\xbc\x43\x86\xba\x3e\x3f\xc0\x68\x1b\x0b\xad\x86\xb9\x19\xde\x1e\x30\x9e\x8a\xd0\x7f\x89\xb7\xdf\x2a\xa9\x37\x16\x5e\xbc\xac\x79\xd3\x5e\xee\x60\x61\xcd\x6e\x0b\x91\x46\x8b\x08\x72\xf9\xef\xe0\x71\x44\xd8\xd9\xa5\xb2\x54\xf0\x63\x42\xe0\x79\x4d\xe6\x87\x5c\xdf\xc0\x8c\x24\x59\xd2\x82\x6f\xb4\xab\x41\x3f\x8b\x7f\x26\xb0\xb1\x1d\xcc\x62\x3b\x39\xe7\xa0\xdd\xac\x18\xf4\x6c\x51\xe7\xef\xf5\xbf\xed\x2f\x51\xdc\xa3\x5a\xc3\x99\x9a\x48\x9d\xfe\x4c\xb2\x30\x82\x43\xa6\xfc\x52\x61\xad\x1b\xe2\x77\xa0\x1d\xae\xf0\xa3\xbb\xfb\x6f\x26\xd7\xea\x3f\xdb\xeb\xfe\xec\xb5\x35\x0c\x70\x2f\x5e\x69\x49\x72\xa1\x01\xbe\xa5\x8c\x4f\xc1\x42\x56\x26\x38\x32\x43\x7d\xce\xf0\x4a\x65\x98\x57\xc9\x46\xf2\xfa\x96\xd2\x6b\x2a\x74\x08\xe9\xdc\x1d\xdf\x74\x0f\x9f\x3f\x27\xa7\xe5\x75\x71\x1e\x7c\x2e\x25\xfa\xb3\x8f\x6d\xfb\xf5\x44\x6b\xbe\xab\x6e\xbf\x9f\x28\x57\x18\xf9\xdf\xef\xc2\x03\x31\xd9\xa5\x48\x10\x44\xbd\x34\x39\xf8\x4a\x7f\xa1\xd2\x00\x7e\x37\x34\x2a\x40\xd6\x6e\xac\xe8\x9a\x7f\x2c\x71\x4a\x32\x29\x44\x1b\x23\xd9\x28\xf4\x88\xd3\xec\x6d\x02\x6f\x56\x5b\x03\x58\x8b\x85\xb3\xa6\x16\xab\x2e\x0e\xdb\x1d\xba\xf5\x4b\x11\x7b\xc0\xa3\xdb\x7d\x6a\xbf\xf6\x34\x5d\x1b\x05\x7b\xe0\xde\x6c\xd2\x70\x7e\x1d\xce\xb5\xb9\x42\xaf\x4c\x81\x55\xd6\x68\xd7\xc0\x96\x1f\x14\xb6\x07\xda\xac\x23\xf3\x47\x9b\x20\xd4\x01\xf6\x82\x15\xcc\x8d\x2b\x35\x1d\xfc\xda\x47\xf9\x78\x81\xcd\x13\x0c\x4e\x27\xe5\x62\x59\x16\x90\x7e\xba\x59\xfc\xe1\x31\xaf\xbf\x5a\xa3\xcd\xb6\x7d\x21\xdd\xfc\xc6\xdd\xa9\x3f\x87\x06\x34\xf0\x26\xfc\x18\x52\x8f\x80\x19\x91\xfb\x3a\x64\x45\xde\x05\xc3\x31\x42\xb7\x17\xed\x08\xba\x9d\xed\xf6\xc2\xa6\x37\x6c\x94\xdc\x9f\x18\xe8\xde\xb7\xa3\x81\xdb\x3f\x45\x5b\xc0\x6e\xf0\xb5\xeb\x6e\x2f\xf8\x6b\xd2\xc3\xa1\xf9\x33\xf2\xff\x04\x00\x00\xff\xff\xa7\x93\x3b\xc6\x4e\x2e\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 11854, mode: os.FileMode(436), modTime: time.Unix(1502337389, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.html": indexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}