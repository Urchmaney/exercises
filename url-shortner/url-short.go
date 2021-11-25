package urls-hort

import (
	"net/http"
  "gopkg.in/yaml.v2"
)

type UrlMap struct {
  path string `yaml:path`
  url string `yaml:url`
}

type UrlMaps struct {
  maps []UrlMap 
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return htt.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    redirect_url := pathToUrls[r.URL.Path]
    if redirect_url != nil {
      http.Redirect(w, r, redirect_url, http.StatusSeeOther)
    } else {
      fallback
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
  parsedYaml, err := parseYAML(yaml)
  if err != nil {
    return nil, err
  }
  pathMap := buildMap(parsedYaml)
  return MapHandler(pathMap, fallback), nil
}

func parsedYaml(yml []byte) (UrlMaps, err) {
  var maps UrlMaps
  err = yaml.Unmarshal(yamlFile, &map)
  if err != nil {
    return nil, err
  }
  return maps, nil
}

func buildMap(data UrlMaps) map[string] string {
  if data == nil {
    return nil
  }
  path_to_url = make(map[string] string)
  for _, v := range data.maps {
    path_to_url[v.path] = v.url
  }
  return path_to_url
}