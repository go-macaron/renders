package renders

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	cache               []*namedTemplate
	regularTemplateDefs []string
	basePath            string
	exts                []string
	lock                sync.Mutex
	re_defineTag        *regexp.Regexp = regexp.MustCompile("{{ ?define \"([^\"]*)\" ?\"?([a-zA-Z0-9]*)?\"? ?}}")
	re_templateTag      *regexp.Regexp = regexp.MustCompile("{{ ?template \"([^\"]*)\" ?([^ ]*)? ?}}")
)

type namedTemplate struct {
	Name string
	Src  string
}

// Load prepares and parses all templates from the passed basePath
func Load(opt Options) (map[string]*template.Template, error) {
	basePath = opt.Directory
	exts = opt.Extensions
	return loadTemplates(nil)
}

// LoadWithFuncMap prepares and parses all templates from the passed basePath and injects
// a custom template.FuncMap into each template
func LoadWithFuncMap(opt Options) (map[string]*template.Template, error) {
	basePath = opt.Directory
	exts = opt.Extensions
	return loadTemplates(opt.Funcs)
}

func loadTemplates(funcMap template.FuncMap) (map[string]*template.Template, error) {
	lock.Lock()
	defer lock.Unlock()

	templates := make(map[string]*template.Template)

	err := filepath.Walk(basePath, func(path string, fi os.FileInfo, err error) error {
		r, err := filepath.Rel(basePath, path)
		if err != nil {
			return err
		}

		ext := getExt(r)

		for _, extension := range exts {
			if ext == extension {
				if err := add(path); err != nil {
					panic(err)
				}

				// Now we find all regular template definitions and check for the most recent definiton
				for _, t := range regularTemplateDefs {
					found := false
					defineIdx := 0
					// From the beginning (which should) most specifc we look for definitions
					for _, nt := range cache {
						nt.Src = re_defineTag.ReplaceAllStringFunc(nt.Src, func(raw string) string {
							parsed := re_defineTag.FindStringSubmatch(raw)
							name := parsed[1]
							if name != t {
								return raw
							}
							// Don't touch the first definition
							if !found {
								found = true
								return raw
							}

							defineIdx += 1

							return fmt.Sprintf("{{ define \"%s_invalidated_#%d\" }}", name, defineIdx)
						})
					}
				}

				var (
					baseTmpl *template.Template
					i        int
				)

				for _, nt := range cache {
					var currentTmpl *template.Template
					if i == 0 {
						baseTmpl = template.New(nt.Name)
						currentTmpl = baseTmpl
					} else {
						currentTmpl = baseTmpl.New(nt.Name)
					}

					template.Must(currentTmpl.Funcs(funcMap).Parse(nt.Src))
					i++
				}
				tname := generateTemplateName(basePath, path)
				templates[tname] = baseTmpl

				// Make sure we empty the cache between runs
				cache = cache[0:0]

				break
				//return nil
			}
		}
		return nil

	})

	return templates, err
}

func add(path string) error {
	// Get file content
	tplSrc, err := file_content(path)
	if err != nil {
		return err
	}

	tplName := generateTemplateName(basePath, path)

	// Make sure template is not already included
	alreadyIncluded := false
	for _, nt := range cache {
		if nt.Name == tplName {
			alreadyIncluded = true
			break
		}
	}
	if alreadyIncluded {
		return nil
	}

	// Add to the cache
	nt := &namedTemplate{
		Name: tplName,
		Src:  tplSrc,
	}
	cache = append(cache, nt)

	// Check for any template block
	for _, raw := range re_templateTag.FindAllString(nt.Src, -1) {
		parsed := re_templateTag.FindStringSubmatch(raw)
		templatePath := parsed[1]
		ext := getExt(templatePath)
		if !strings.Contains(templatePath, ext) {
			regularTemplateDefs = append(regularTemplateDefs, templatePath)
			continue
		}

		// Add this template and continue looking for more template blocks
		add(filepath.Join(basePath, templatePath))
	}

	return nil
}
