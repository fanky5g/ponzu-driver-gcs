package storage

import (
	"errors"
	"net/url"
	"strings"
)

var ErrInvalidPath = errors.New("path must be in the format bucket/path/to/file.ext")

func parsePath(p string) (string, string, error) {
	u, err := url.Parse(p)
	if err != nil {
		return "", "", err
	}

	if strings.HasPrefix(p, "gs://") {
		return u.Host, trimPath(u.Path), nil
	}

	cfg, err := getConfig()
	if err != nil {
		return "", "", err
	}

	if cfg.Bucket != "" {
		return trimPath(cfg.Bucket), trimPath(u.Path), nil
	}

	paths := strings.SplitN(u.Path, "/", 2)
	if len(paths) == 2 {
		bucket, key := paths[0], paths[1]
		return trimPath(bucket), trimPath(key), nil
	}

	return "", "", ErrInvalidPath
}

func trimPath(p string) string {
	return strings.TrimPrefix(strings.TrimSuffix(p, "/"), "/")
}
