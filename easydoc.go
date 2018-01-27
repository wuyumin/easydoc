package easydoc

import (
	"github.com/wuyumin/easydoc/utils"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"text/template"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"path"
)

var (
	srcStr, distStr, themeStr string
	templateDefaultDoc        = `<!doctype html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<title>{{.dataTitle}}</title>
<meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no">
<style>
</style>
</head>
<body>
<div id="wrap">

    <div class="menu">
        {{.dataMenu}}
    </div>

    <div class="content">
        {{.dataDoc}}
    </div>
    
</div>
</body>
</html>
`
	markdownReadme = `- [Home](index.md)
`
	markdownIndex = `# Home

This is content.  
You can use markdown to write, EasyDoc will be converted to html content.
`
)

func init() {
	// Current directory
	pwd, err := os.Getwd()
	utils.CheckErr(err)

	// Various paths
	srcStr = fmt.Sprint(pwd, "./src/")
	distStr = fmt.Sprint(pwd, "./dist/")
	themeStr = fmt.Sprint(srcStr, "theme/")
}

func GenerateInit() error {
	// "src" directory
	err := os.MkdirAll(srcStr, 0777)
	if err != nil {
		return err
	}
	// menu file
	err = ioutil.WriteFile(fmt.Sprint(srcStr, "README.md"), []byte(markdownReadme), 0777)
	if err != nil {
		return err
	}
	// index file
	err = ioutil.WriteFile(fmt.Sprint(srcStr, "index.md"), []byte(markdownIndex), 0777)
	if err != nil {
		return err
	}
	// theme directory
	err = os.MkdirAll(themeStr, 0777)
	if err != nil {
		return err
	}
	// template directory
	err = os.MkdirAll(fmt.Sprint(themeStr, "template/"), 0777)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprint(themeStr, "template/doc.tpl"), []byte(templateDefaultDoc), 0777)
	if err != nil {
		return err
	}

	// "dist" directory
	err = os.MkdirAll(distStr, 0777)
	if err != nil {
		return err
	}

	return nil
}

func GenerateDoc() error {
	// "src" directory is exist?
	if _, err := os.Stat(srcStr); err != nil {
		return err
	}

	// Empty the generated file
	err := os.RemoveAll(distStr)
	utils.CheckErr(err)
	err = os.MkdirAll(distStr, 0777)
	utils.CheckErr(err)

	// Menu content
	markdownMenu, err := ioutil.ReadFile(fmt.Sprint(srcStr, "README.md"))
	if err != nil {
		return err
	}
	markdownMenuHtml := strings.Replace(string(blackfriday.MarkdownCommon(markdownMenu)), ".md", ".html", -1)  // Menu html content

	// Template content
	var templateFile, templateContent string
	templateFile = fmt.Sprint(themeStr, "template/doc.tpl")
	if _, err := os.Stat(templateFile); err != nil {
		templateContent = templateDefaultDoc  //Default template content
	} else {
		templateNewDoc, err := ioutil.ReadFile(templateFile)
		if err != nil {
			return err
		}
		templateContent = string(templateNewDoc)  //File template content
	}

	template_doc := template.New("Doc")
	template_doc, err = template_doc.Parse(templateContent)
	if err != nil {
		return err
	}

	//Each document content
	var slice [][]string
	slice = regexp.MustCompile(`\[(.*)\]\((.*)\)`).FindAllStringSubmatch(string(markdownMenu), -1)
	for _, v := range slice {
		markdownDoc, err := ioutil.ReadFile(fmt.Sprint(srcStr, v[2]))
		if err != nil {
			return err
		}
		markdownDocHtml := string(blackfriday.MarkdownCommon(markdownDoc))  // Document html content
		var buf bytes.Buffer
		template_doc.Execute(&buf, map[string]interface{}{"dataTitle": v[1], "dataMenu": markdownMenuHtml, "dataDoc": markdownDocHtml})
		if _, err := os.Stat(fmt.Sprint(distStr, path.Dir(v[2]))); err != nil {
			err = os.MkdirAll(fmt.Sprint(distStr, path.Dir(v[2])), 0777)
			utils.CheckErr(err)
		}
		err = ioutil.WriteFile(fmt.Sprint(distStr, strings.Replace(v[2], ".md", ".html", 1)), buf.Bytes(), 0777)
		if err != nil {
			return err
		}
		fmt.Println(v[2], "--->", "OK")
	}

	return nil
}
