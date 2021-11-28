package urlshort

import (
	"net/http"
  "gopkg.in/yaml.v2"
)

type UrlMap struct {
  Path string
  Url string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this... 
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
      redirect_url := pathsToUrls[r.URL.Path]
      if redirect_url != "" {
        http.Redirect(w, r, redirect_url, http.StatusSeeOther)
      } else {
        fallback.ServeHTTP(w, r)
      }
    })
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
  parsedYaml, err := parseYAML(yml)
  if err != nil {
    return nil, err
  }
  pathMap := buildMap(parsedYaml)
  return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]UrlMap, error) {
  var maps []UrlMap
  err := yaml.Unmarshal(yml, &maps)
  if err != nil {
    return maps, err
  }
  return maps, nil
}

func buildMap(data []UrlMap) map[string] string {
  path_to_url := make(map[string] string)
  for _, v := range data {
    path_to_url[v.Path] = v.Url
  }
  return path_to_url
}