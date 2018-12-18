// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package swaggerjson

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Swagger statically implements the virtual filesystem provided to vfsgen.
var Swagger = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 12, 23, 58, 5, 76182838, time.UTC),
		},
		"/api.swagger.json": &vfsgen۰CompressedFileInfo{
			name:             "api.swagger.json",
			modTime:          time.Date(2018, 12, 18, 23, 7, 52, 33765133, time.UTC),
			uncompressedSize: 15179,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x73\xe3\xb6\x11\x7f\xf7\x5f\xb1\xc3\x76\xa6\x2f\x77\xd6\x25\x7d\xe9\xb8\x0f\xad\xc7\x76\x2f\x9a\x9c\x3f\xc6\xf2\x25\x0f\xbd\x1b\xcd\x8a\x5c\x91\x88\x49\x80\xc1\x82\x52\x35\x1d\xff\xef\x1d\x80\xa0\x45\x52\xa4\x44\x4b\xe7\x9c\x2e\x69\x5e\x72\x22\x80\xc5\x7e\xfc\xb0\x5f\x80\xff\x7b\x02\x10\xf0\x12\xe3\x98\x74\x70\x06\xc1\xf7\xa7\xef\x82\x37\xf6\x9b\x90\x73\x15\x9c\x81\x1d\x07\x08\x8c\x30\x29\xd9\xf1\x8b\xb4\x60\x43\x1a\xae\x51\x62\x4c\x1a\xce\xef\xc6\x30\x99\xfc\x00\xb9\x56\x0b\x11\x91\x76\x8b\x01\x82\x05\x69\x16\x4a\xda\x25\x8b\x77\xa7\xdf\x79\xaa\x00\x41\xa8\xa4\xc1\xd0\x3c\x93\x06\x08\x24\x66\x8e\xf6\x35\x6a\x34\xf0\x1e\x35\xce\x0b\x13\x09\xa9\x16\x7e\x11\x40\x50\xe8\xd4\x4e\x49\x8c\xc9\xf9\x6c\x34\x8a\x85\x49\x8a\xd9\x69\xa8\xb2\x11\x63\xc6\x85\x8c\xdf\x86\x32\x34\xa3\x30\xc3\xb7\xcc\xc9\x7a\x1d\x65\x28\xdc\xca\xcc\x12\x57\x22\xfa\x67\x6c\xbf\xd8\x95\x81\x9b\xf3\x74\x02\xf0\xe4\x24\xe6\x30\xa1\x8c\x38\x38\x83\x7f\x97\xac\xba\xcd\x2a\xbe\xed\x0f\xbb\xe2\xb3\x9b\x1b\x2a\xc9\x45\x63\x32\xe6\x79\x2a\x42\x34\x42\xc9\xd1\x2f\xac\xe4\x7a\x6e\xae\x55\x54\x84\x03\xe7\xa2\x49\x78\xad\xf6\x11\xe6\x62\xb4\xf8\x6e\x14\x96\x5a\xaf\x2b\x2d\xa6\xba\x0e\x2d\xfb\x45\x96\xa1\x5e\x59\x59\x7f\x16\x69\x0a\x9a\x8c\x16\xb4\x20\x30\x09\x01\x1b\x34\x05\x83\x9a\x03\x82\x27\x06\x28\x23\x10\x86\xe1\xb1\x98\x51\xa8\xe4\x5c\xc4\x30\x57\x1a\x42\x25\x25\x85\x46\x2c\x84\x59\x3d\xeb\x11\x20\x50\x39\x69\xc7\xf2\x38\xb2\x7b\xbc\x27\xe3\xb1\x50\x9f\xa4\x89\x73\x25\x99\xb8\xc1\x1b\x40\xf0\xfd\xbb\x77\xad\x4f\x00\x41\x44\x1c\x6a\x91\x1b\x0f\x94\x1a\xa1\x52\x22\x6b\x10\xdc\x58\x06\x10\xfc\x59\xd3\xdc\xae\xf8\xd3\x28\xa2\xb9\x90\xc2\x52\x60\x6b\x7c\xe6\x64\xcd\xd8\x3d\xe5\xe9\x2a\x68\xac\x7d\x3a\xe9\xfa\xf7\x53\x4d\x82\x1c\x35\x66\x64\x48\xaf\xed\x55\xfe\xd7\xe2\xbd\x02\xad\xfb\xff\x9b\xad\x72\xdd\x60\x46\x56\xf5\xd6\x10\x95\xf2\x8d\x82\x19\x41\xaa\xd4\x23\x45\x50\xe4\xa7\x6d\x12\xc2\xad\xfc\xb5\x20\xbd\x6a\x0f\x69\xfa\xb5\x10\x9a\xac\x15\xe6\x98\x32\xb5\x86\xcd\x2a\x77\x8c\xb1\xd1\x42\xc6\x41\xa7\xc0\x9f\x6b\x02\x1b\x8c\xdb\xa2\x56\xa7\x7c\xbd\xf8\xf3\x49\x4b\x53\x41\x44\x29\x19\xda\x0e\xc1\x72\xce\x1a\x72\x5b\xe0\x74\xe9\xa6\x1e\x27\xa2\x1a\xbc\x1d\x0b\xa8\x7e\x4e\xd0\x80\xe0\x3a\xa8\xfe\xc2\x60\x17\x5a\x6c\x45\xc4\x46\xab\xd5\xb7\x07\xab\x5c\xf1\x0e\xbf\xe6\x22\x8d\x8d\x2d\x83\x70\x75\xa1\x09\x8f\x15\x57\x0d\xde\x7e\x13\x5c\xcd\x54\xb4\x61\xf7\x12\x12\x5d\x23\x35\x44\x18\x5d\xb4\x01\xf1\x05\x64\xbe\xe6\x78\x88\xc4\x5f\x02\x56\xc5\x0e\x54\x61\xf4\x4b\xc1\x06\xf0\x85\xf0\x3a\x77\xcb\x3c\x03\x37\x2a\x22\x3e\x32\x8c\x35\x18\xfc\x83\x60\xac\x21\xf3\xab\x63\xec\xa4\xa6\xb0\x76\xc2\x36\x4a\x45\xc3\x9d\xbd\x20\x6b\x43\xb0\x6b\x6d\xda\xe0\x69\xf1\xa0\x64\xec\x83\xdd\xf0\xb8\x20\xd8\x64\x6e\x2f\x0c\x7e\x41\x8b\x14\x79\xac\x31\xa2\x97\x1a\xa5\xd0\x12\xfc\x52\x50\x4e\x41\xec\xb2\x65\x84\x58\x2c\x48\x0e\x70\x15\xef\xc9\x7c\x2c\x09\x78\xce\xc7\x72\xae\x74\xe6\x66\x1c\x9f\xc5\x7a\x59\x3d\xe2\xfc\x07\x8c\xfd\xb6\x24\x40\x4d\xb6\x8c\x61\x5b\x90\x0a\x59\x96\x35\xde\x78\xdf\x60\x46\xb4\x33\x74\x19\x43\x59\x6e\x6c\xda\x57\x21\x74\x48\xe8\x6a\x5a\xf8\xc8\x10\xd8\x64\xee\x0f\x12\xb7\x9a\x42\x7f\x9d\xc0\xb5\xee\xdb\xbc\xd8\x3d\xfa\xa5\x20\xd6\xce\x02\x70\xa6\x0a\x03\x98\x0b\x60\xd2\x8b\x5d\xfe\xf1\xa7\x92\xc2\x51\x3b\x46\xcf\xe3\x6f\x16\xc6\x9e\xdb\x53\x35\x6e\xd6\x0d\xa2\x76\xaa\x33\xe1\xe4\x9e\x32\xb5\xa0\x6b\x0c\x13\x21\x69\x92\x53\x58\xb7\x65\xe5\xb3\xd4\xec\x17\x0a\xd7\x89\x42\x90\x6b\x6b\x0d\x23\x5a\xca\x0d\x92\x76\x39\xb6\xe9\xf7\xde\x34\xc6\xaa\x3e\xe1\x43\x42\x60\x17\x3b\xdf\x3b\x99\xfc\x00\x18\x86\xc4\xbc\x16\xf3\xa9\x13\x86\x2d\x05\x77\x00\xe2\x00\x61\x62\x61\xa6\x9b\xf8\x7e\x99\x4c\x06\x63\x50\xd2\x45\x9d\x58\x18\xd0\x94\x2b\x16\x46\xe9\x1a\x16\xea\x16\xb7\x5b\x86\x2a\xcb\xc4\x01\x5a\x44\x4e\xaa\xee\x91\xdd\xd2\x93\xeb\xdd\xce\x68\xa2\x29\x1b\x6c\x35\x67\x86\x6e\xf9\x73\x42\x26\x21\x0d\x4a\x83\x54\xc6\xed\x6a\x29\xc2\x12\x19\xc2\x94\x50\xc2\x32\x21\x09\xb3\x42\xa4\x3d\x4c\xd8\xa1\x68\x1a\xed\xcb\xc0\x25\x1a\xd7\x2d\x73\x64\x7a\xc4\x54\x07\xd9\xd1\xa3\xca\x6e\x12\x2b\x28\x98\x22\x1b\x3b\x43\x95\xe5\x22\xa5\xee\x1d\xfd\xa0\xde\x6b\xbf\x0b\xbf\xd8\x6d\xd5\x4d\x3f\x4f\xd1\x58\x8c\xef\x45\xff\xce\x2f\x06\x61\x4a\x33\x95\xfb\x45\xee\xec\x8d\x40\x17\x52\xda\x2c\xa8\xec\x33\xfb\xbd\x3b\x4f\x5f\x4f\xf9\x74\xc0\x91\xf3\xc1\x75\x1f\x18\xee\x95\xde\x75\x6b\x17\xa3\x68\x2a\x5d\x71\xde\xc3\x0a\x6a\x8d\xcd\x10\x1f\x08\x43\x59\x7b\xfe\x8e\xe8\x30\xe1\xa4\xee\x76\xeb\xe1\xa0\x5b\x48\x3f\x9b\x61\x99\x88\x30\xb1\xa2\x2d\x51\xba\x44\x0e\x23\x87\xc9\x9a\xf4\xdd\x92\x69\xe7\xec\x5f\x57\xb8\x41\x31\xe6\x10\x61\x4b\x21\x60\xae\x55\xd6\x23\xf1\x50\xb4\x96\x51\xf9\x00\xbc\xaa\xc7\x3e\x2d\xce\x94\xb2\xde\xaf\xa9\xc7\x32\x2c\xf5\x0e\xaf\xd1\x8c\x16\xc9\x82\x01\x81\x0b\x17\x08\xe7\x85\xcd\x9b\x7e\x2d\x88\xcd\x30\x39\xbd\x84\x97\x64\x50\xa4\x63\x43\xd9\x21\x62\x8a\x68\xaf\x43\x39\xbe\x6c\x5d\x63\x74\x83\x72\xef\x43\xdf\x71\x51\xd2\xbd\x43\x79\x9d\x35\xcd\x88\x19\xe3\xfd\xf6\x3a\x8f\x22\x87\x6f\x4c\x3b\x12\xd6\xe6\x95\xd9\x4e\x76\xd6\x37\x68\x07\xfb\xba\xda\x65\x9c\x8b\x49\xee\x2e\x6e\x90\x2b\x28\xf9\x6d\x33\xb0\xb5\x17\x5b\x92\x9b\x94\x0b\xfb\x33\x90\x5d\x9a\xd8\x8d\xda\xff\xe3\xf5\xa8\xf0\x7a\xcc\x50\x99\xb4\x79\xeb\x53\x4f\x40\xb2\xc8\x1a\xd5\x54\x30\x79\x38\x7f\xf8\x38\x99\x7e\xbc\x99\xdc\x5d\x5d\x8c\xff\x35\xbe\xba\xac\xd7\x8f\x77\xf7\xb7\x3f\x8d\x27\xe3\xdb\x9b\xf1\xcd\xfb\xfa\xf7\xfb\x8f\x37\x1b\x9f\xae\x2e\x6e\x6f\x2e\xc6\x1f\x5a\x9f\x27\x0f\xb7\x77\x77\xad\x6f\x57\xf7\xf7\xb7\xf7\xf5\x0f\x97\x57\xef\xef\xcf\x2f\xaf\x2e\x2b\xa9\x3f\xd7\xae\x4c\xe7\x58\xa4\x2e\x5c\x6c\xe1\x74\xad\xd0\xb7\xb0\x39\xed\x0c\x6e\x94\x01\x26\xf3\x49\xc2\x5b\xa8\x8b\x74\x06\xd6\x00\xf5\x2f\xce\x1a\x04\x42\x46\x22\x44\x43\xcd\x64\x4a\x30\xcc\xc8\x66\x4f\xa1\xbb\x91\x89\x4e\x1d\x41\xaf\x8b\x92\x96\xff\xb1\x95\x4c\x82\x96\x0e\xc9\x8a\x4c\xf9\xaa\x80\x61\x5e\xa4\xe9\x0a\x0a\xc6\x59\x4a\x9e\xf4\x5a\xa7\x9e\xfc\xfa\x43\xc7\x16\x68\x80\x55\x46\xb0\x54\xfa\xd1\x12\xc4\xd0\x88\x05\xa5\x2b\xcf\x75\xa4\x24\x55\x85\x98\xe7\xe5\x8d\x8d\xad\x09\x20\xfb\x64\xd0\x4e\xb3\xc3\x19\x3a\x4e\x5d\x41\x13\x11\xb0\x9a\x9b\x25\x6a\xcf\x55\x65\xd2\x92\xa5\xea\xd7\x30\xcd\x95\xd7\xdb\x91\xa3\xe3\x60\x50\x12\x71\xff\xdc\x4a\x21\x43\x2b\x06\x14\xb2\xd4\x8f\x23\x50\xc1\xa6\xa4\x51\xfd\xda\x4a\xc6\xf7\xa0\xd8\xd6\x15\xda\x29\xc8\x2a\xc4\xe6\x53\x6c\x94\x26\x67\x03\x98\x17\x32\x2c\xfd\x86\x30\xbe\x48\x6d\x9d\xbb\xf6\x8d\xdc\xd7\x48\xf7\xfb\x9f\x48\x3c\xdf\xc7\xf5\x95\x4e\x8f\x7f\xe3\x7d\xaa\xc1\x56\x7f\xc8\x2a\x7d\xb1\xae\x0b\x7f\x2c\x66\xa4\x25\x59\x8d\xdb\xea\xc2\x82\x90\x4a\xfc\xf0\x29\x5c\x28\x69\xb4\x4a\x21\x4f\x51\x3e\xaf\x62\x57\x96\x44\x64\x48\x67\x42\x52\x04\xb3\x95\x93\xa6\x96\x24\x9f\x76\x0b\x90\x88\x38\x99\xe2\x02\x45\x8a\x33\xe1\xac\xf4\x2a\x29\xe8\x66\x5d\x5f\x03\xf4\x0f\xe7\x3d\xa1\x91\x8c\x95\x7d\x3a\xc7\x99\x16\xe1\xde\x3d\x8c\x72\xb9\xb7\x68\x7f\x15\x1c\x96\x8a\x9d\x3a\xc5\x7e\x0b\x15\x9b\xad\xb4\xb5\x60\xda\x1d\x72\x4b\x04\xfd\xbe\x64\xc2\x5c\x4c\x49\x46\xb9\x12\x72\xdf\xf6\x96\x60\xe0\x44\x15\x69\x64\x81\x81\xb0\xc0\xb4\x20\x48\xc5\x23\x81\xc8\xcf\x72\xa5\x8d\x2f\xfa\x45\x9a\xfa\x19\x42\x9b\x02\x53\x18\xdf\x8d\xec\xf0\x27\x79\x87\xcc\xf6\xb4\x61\xf8\x68\xf1\x45\xff\x31\xa4\x6d\x92\x14\x16\x6c\x54\x46\x9a\x3d\xea\xac\xa3\xf5\x6d\x9e\xac\x90\xce\x9b\x0e\x4a\xab\x73\x2d\x16\x68\x68\xfa\x48\xbd\xe7\x72\x7b\x73\xa6\x5c\x0f\x8f\xb4\x7a\x6e\x35\x31\x27\x20\xa4\x51\x90\x79\xd5\x0f\xcc\x90\x36\xdf\x8b\x1c\x61\xa9\xdb\xef\x67\x96\xc8\x75\x77\x5e\x7a\x48\xc1\x1d\x75\x70\xdd\x27\x6c\x3c\x7d\xdc\x79\x02\xea\x45\x47\xbf\x4f\x8a\x5c\x25\xdd\xce\x55\x2b\x66\xa0\xba\x6d\x18\x66\x9a\x8e\x27\x62\xc7\x67\x9a\x0b\x77\xd0\xea\xc2\xce\xa8\x7a\x29\xd6\xe7\x92\xbb\x4b\x85\x41\xc8\x9f\x34\x6a\x81\x17\x35\x3b\xba\x9e\x0c\x1c\x9f\x3e\xc7\xad\xbc\xb0\xcc\x46\x79\xc5\x16\x78\xdb\xc0\xfc\xca\x11\xa0\x8e\xff\xdd\xee\xff\x43\xfb\x99\xc9\x0b\x2d\xf4\xbb\xb3\xce\x1e\xae\xa6\xd6\x95\x1b\xaa\xbd\x1d\x0f\x2c\x8e\x4f\x9f\x17\x28\xdb\xbe\xc3\x37\xdd\x7b\x5c\x47\x95\x18\x1f\x0e\xf6\x2d\x2f\x2d\xb6\x74\xd4\x1a\x89\x79\xae\x98\x85\x4d\x01\xb4\x88\x13\x03\x52\x2d\x07\x5b\xaa\x71\xe3\x7b\x7c\x76\x19\xcf\xe1\xf9\x26\xdd\x85\xd8\xdb\x1f\xb7\xda\x63\x2a\x3a\x6f\x53\xa1\x1f\xea\xbb\x2f\x65\xb7\xdf\xb3\xd5\x67\x6e\x2a\x7d\xb3\xef\xe1\x76\x71\x75\x57\x55\x91\x35\x82\x47\xd3\x40\xeb\x4a\xed\x03\xce\x28\xfd\x2a\xd5\xab\x4d\x27\xa4\xaf\x60\x11\x52\xc7\x47\xb7\x09\x6c\x7e\xbb\xf7\x16\x65\x76\xdc\xb9\xc7\x16\x04\xb7\xea\x81\x03\xf4\x53\xb0\x4d\xad\x0f\xd0\x51\x45\xa0\xff\x51\x40\xbd\x2c\x7e\xbd\xa7\x07\xf5\xf4\x5e\xe9\xde\x5d\x84\x34\x14\x37\xde\xac\x34\x8e\xa5\x90\xe6\xaf\xdf\x6f\xe1\xc1\x95\x2f\x83\x78\x40\xe6\xa5\xd2\x2f\xea\x71\x77\xb4\x2f\x2a\x32\xad\x3d\x4f\xc1\x95\x59\x82\x5d\x3e\x2e\xb2\x3c\xa5\x8c\xa4\xa1\x08\x96\xc2\x24\xa2\xe1\xd4\x31\x17\x9f\xe4\x8c\x42\x2c\x98\xdc\xb0\x2a\x0c\x3c\x4a\xb5\x94\x53\xa7\x51\x2e\x72\x27\x15\xc2\xf5\xf8\xe1\x1a\x42\x94\x10\x53\x99\xe4\x57\xbb\x9f\xc2\x79\x39\x28\xf8\x93\x64\xe3\x1e\xa8\xd9\xac\x7f\x96\x52\xe6\x38\xb3\x85\xd0\x0c\x6d\x25\x84\x85\x49\x48\x1a\xff\x97\x50\x7f\x07\x5a\x90\x04\xe1\xd2\xc5\x15\x44\xca\xf1\xeb\xc9\x7f\x92\x76\x99\x1b\xb0\x9b\x96\x74\x45\x96\x93\x66\x25\x5d\x29\xe7\x1a\x7d\x0e\xe6\xa7\xf0\x70\x7b\x79\x7b\xb6\x16\xb5\x26\x01\xf7\xf4\x60\xdc\x79\x7a\xe5\xb4\xac\xed\xa9\x76\x87\x31\x6b\xd5\x92\x33\xa7\xb9\x9a\x8c\x7d\x48\x2a\x66\xa9\x08\xf7\x3e\x39\xe5\xc3\x56\x4c\xa1\xa4\x53\x9e\x22\x91\x0f\x72\xda\xae\xfb\x9f\x53\x28\xe6\xde\x9e\xfe\x79\x6c\xf5\xcd\x35\xcf\xba\x1c\xf8\xe6\xab\xb7\x6f\xfe\xbd\xc1\x21\x8f\x53\xea\xa9\x8b\xdb\x79\x49\xf5\x7e\xb6\xfa\xc7\x30\xaf\xdf\xf5\x80\xf2\xf8\x72\x97\xfd\xee\xc5\x9f\x9f\xc4\x55\xcd\x9e\x4b\x15\xd6\xde\xc4\xb5\x5c\xe3\xb5\xd2\xe4\xaf\xc9\x86\xfe\xf1\xea\x0b\xff\xd8\xd4\xf2\x73\xf2\x74\xf2\xbf\x00\x00\x00\xff\xff\x30\xa0\xd7\x07\x4b\x3b\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/api.swagger.json"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr: gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
