package easydoc

import (
	"github.com/wuyumin/easydoc/utils"
	"github.com/mostafah/fsync"
	"github.com/BurntSushi/toml"
	"github.com/russross/blackfriday"
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"net/http"
	"strings"
	"text/template"
	"bytes"
	"errors"
	"path"
)

type Config struct {
	FixLink      string
	LanguageCode string
	Theme        string
	SuffixTitle  string
	HomeTitle    string
	ScanFile     [][]string
}

var (
	err error
	// Current directory
	// Windows E:\GoPath\github.com\wuyumin\easydoc
	// Unix-like /GoPath/github.com/wuyumin/easydoc
	pwd                      string
	configFileDefaultContent string
	conf                     Config
	curConfigDir             string
	configFile               string
	curDistDir               string
	curSrcDir                string
	curStaticDir             string
	curThemeDir              string
	cssDefaultContent        string
	jsDefaultContent         string
	docDefaultContent        string
	menuDefaultContent       string
)

func init() {
	pwd, err = os.Getwd()
	utils.CheckErr(err)

	// Current directory config
	curConfigDir = fmt.Sprint(pwd, "/config")
	configFile = fmt.Sprint(curConfigDir, "/config.toml")

	// Current directory dist
	curDistDir = fmt.Sprint(pwd, "/dist")

	// Current directory src
	curSrcDir = fmt.Sprint(pwd, "/src")

	// Current directory static
	curStaticDir = fmt.Sprint(pwd, "/static")

	// Current directory theme
	curThemeDir = fmt.Sprint(pwd, "/theme")

	configFileDefaultContent = `fixLink = ""
languageCode = "zh-CN"
theme = "default"
suffixTitle = ""
homeTitle = "Home Title"
# SCANFILE: The local path begins with src(base on src directory). support external link.
scanFile = [
    ["Home", "src/index.md"],
]
`

	cssDefaultContent = `@charset "utf-8";
a{color: #009c95;}
.menu{}
.menu ul{list-style: none;padding-left: 0.5em;}
.menu a{display: block;padding: 0.2em;border-radius:0.5em;font-size: 1.2em;color: rgba(0,0,0,.87);word-break : break-all;}
.menu a:hover{background-color: rgba(0,0,0,.05);color: rgba(0,0,0,.95);}
.menu .made-by{padding: 0.5em;}
.menu .made-by a{text-align: center;}
.content{margin-top: 10px;}
.ui .new-grid{margin-left: 0;margin-right: 0;}
.ui.menu{border-radius: 0 !important;}

/* markdown2html style */
.content img{max-width: 100%;height: auto;}
code{padding: 2px 4px;font-size: 90%;color: #c7254e;background-color: #f9f2f4;border-radius: 4px;}
pre{padding: 2px 5px;background-color: #f2f2f2;border-radius: 3px;max-width: 100%;overflow-x: scroll;}
pre code{padding: 0;font-size: 90%;color: #333;background-color: #f2f2f2;border-radius: 0;}
blockquote{margin: 5px 0;padding: 5px 10px;border-left: 2px solid #00b5ad;background-color: #f6f6f6;color: #555;font-size: 1em;}

/* back2top */
#back2top{position: fixed;bottom: 5px;right: 5px;display: none;width: 30px;height: 30px;border-radius: 30px;line-height: 30px;text-align: center;background: #222;color: #fff;font-weight: bold;cursor: pointer;-webkit-transition: 1s;-moz-transition: 1s;-ms-transition: 1s;-o-transition: 1s;transition: 1s;}
#back2top:hover{background: #555;}
`
	jsDefaultContent = `$(function(){
    /* btn-sidebar */
    $('#btn-sidebar').on('click', function(){
        $('.ui.sidebar').sidebar('toggle');
    });

    /* back2top */
    $(window).on('scroll', $.throttle(250, function(){
        if($(this).scrollTop() >= 100){
            $('#back2top').fadeIn();
        } else {
            $('#back2top').fadeOut();
        }
    }));
    $('#back2top').on('click', $.throttle(250, true, function(){
        $('body,html').animate({
            scrollTop: 0
        }, 800);
    }));
});
`

	docDefaultContent = `<!doctype html>
<html lang="{{.languageCode}}">
<head>
<meta charset="utf-8">
<title>{{if (eq .newUrlPath "index.html")}}{{.homeTitle}}{{else}}{{.dataTitle}}{{.suffixTitle}}{{end}}</title>
<meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no">
<meta name="format-detection" content="telephone=no,email=no,adress=no">
<link rel="stylesheet" href="https://cdn.bootcss.com/semantic-ui/2.2.13/semantic.min.css">
<!-- <link rel="stylesheet" href="{{.fixLink}}asset/css/style.css"> -->
<link rel="stylesheet" href="https://wuyumin.github.io/easydoc/dist/static/css/style.css">
</head>
<body>

<div class="ui left vertical menu sidebar">
    <div class="menu">
        {{.dataMenu}}

        <div></div>
        <ul class="made-by">
            <li><a href="https://github.com/wuyumin/easydoc" target="_blank" title="EasyDoc">EasyDoc</a></li>
        </ul>
    </div>
</div>

<div class="pusher">
    <div class="ui vertical">
        <div class="ui inverted borderless menu">
            <a href="javascript:;" class="item" id="btn-sidebar"><i class="sidebar icon"></i></a>
            <a href="{{.fixLink}}index.html" class="item">Home</a>
            <div class="right menu">
                <a  href="https://github.com/wuyumin/easydoc" class="item" target="_blank" title="EasyDoc">EasyDoc</a>
            </div>
        </div>

        <div class="ui grid new-grid">
            <div class="sixteen wide column">
                <div class="ui raised segment">
                    <strong class="ui teal ribbon label">{{.dataTitle}}</strong>
                    <div class="content">
                        {{.dataDoc}}
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<p id="back2top">&and;</p>

<script src="https://cdn.bootcss.com/jquery/2.2.3/jquery.min.js"></script>
<script src="https://cdn.bootcss.com/semantic-ui/2.2.13/semantic.min.js"></script>
<script src="https://cdn.bootcss.com/jquery-throttle-debounce/1.1/jquery.ba-throttle-debounce.min.js"></script>
<!-- <script src="{{.fixLink}}asset/js/app.js"></script> -->
<script src="https://wuyumin.github.io/easydoc/dist/static/js/app.js"></script>
</body>
</html>
`

	menuDefaultContent = ``

	var configData string
	configFileData, err := ioutil.ReadFile(configFile)
	if err != nil {
		configData = configFileDefaultContent // Default
	} else {
		configData = string(configFileData)
	}
	if _, err := toml.Decode(configData, &conf); err != nil {
		utils.CheckErr(err)
	}
}

func GenerateInit() error {
	indexFileDefaultContent := `# Home

This is content.  
You can use markdown to write, EasyDoc will be converted to html content.
`
	// Current directory config
	err = os.MkdirAll(curConfigDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(configFile, []byte(configFileDefaultContent), os.ModePerm)
	if err != nil {
		return err
	}

	// Current directory dist
	err = os.MkdirAll(curDistDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Current directory src
	err = os.MkdirAll(curSrcDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprint(curSrcDir, "/index.md"), []byte(indexFileDefaultContent), os.ModePerm)
	if err != nil {
		return err
	}
	// Tips
	if err = generateTips(); err != nil {
		return err
	}

	// Current directory static
	err = os.MkdirAll(curStaticDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Current directory theme
	err = os.MkdirAll(curThemeDir, os.ModePerm)
	if err != nil {
		return err
	}
	// theme css
	themeCssFileDir := fmt.Sprint(curThemeDir, "/default/css")
	if err = utils.ExistsOrMkdir(themeCssFileDir); err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprint(themeCssFileDir, "/style.css"), []byte(cssDefaultContent), os.ModePerm)
	if err != nil {
		return err
	}

	// theme js
	themeJsFileDir := fmt.Sprint(curThemeDir, "/default/js")
	if err = utils.ExistsOrMkdir(themeJsFileDir); err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprint(themeJsFileDir, "/app.js"), []byte(jsDefaultContent), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprint(curThemeDir, "/default/doc.tpl"), []byte(docDefaultContent), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprint(curThemeDir, "/default/menu.tpl"), []byte(menuDefaultContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func GenerateDoc(isEmptydist bool) error {
	// src exist?
	if _, err := os.Stat(curSrcDir); err != nil {
		return err
	}
	// dist exist?
	if err = utils.ExistsOrMkdir(curDistDir); err != nil {
		return err
	}

	// Empty dist directory
	if isEmptydist {
		err := EmptyDist()
		return err
	}

	////////////Source processing////////////
	absCurSrcPath, _ := filepath.Abs(curSrcDir)
	// store post source
	postSourceById = make(map[int]*PostSource)
	if _, err := os.Stat(configFile); err == nil && len(conf.ScanFile) > 0 {
		// source from config file
		for k, v := range conf.ScanFile {
			// Is External Link
			if utils.IsExternalLink(v[1]) {
				postSourceById[k+1] = &PostSource{Id: k + 1, Title: v[0], AbsPath: "", UrlPath: v[1]}
				continue
			}
			if _, err := os.Stat(v[1]); err != nil {
				return err
			}
			filePathAbs, _ := filepath.Abs(v[1])
			postSourceById[k+1] = &PostSource{Id: k + 1, Title: v[0], AbsPath: filePathAbs, UrlPath: strings.TrimLeft(strings.Replace(strings.Replace(filePathAbs, absCurSrcPath, "", 1), "\\", "/", -1), "/")}
		}
	} else {
		// source from automatic scanning
		sourceKey := 0
		filepath.Walk(absCurSrcPath, func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() || (!strings.HasSuffix(f.Name(), ".md")) {
				return nil
			}
			postSourceById[sourceKey+1] = &PostSource{Id: sourceKey + 1, Title: strings.TrimRight(f.Name(), ".md"), AbsPath: path, UrlPath: strings.TrimLeft(strings.Replace(strings.Replace(path, absCurSrcPath, "", 1), "\\", "/", -1), "/")}
			sourceKey++
			return nil
		})
	}

	if len(postSourceById) < 1 {
		return errors.New("no markdown files to EasyDoc")
	}

	// Menu template content
	var menuTemplate, menuTemplateContent string
	menuTemplate = fmt.Sprint(curThemeDir, "/", conf.Theme, "/menu.tpl")
	if _, err := os.Stat(menuTemplate); err == nil {
		menuTemplateRead, err := ioutil.ReadFile(menuTemplate)
		if err != nil {
			return err
		}
		menuTemplateContent = string(menuTemplateRead)
		if strings.TrimSpace(menuTemplateContent) == "" {
			menuTemplateContent = generateMenuByMap(postSourceById)
		}
	} else {
		menuTemplateContent = generateMenuByMap(postSourceById)
	}

	// Doc template content
	var docTemplate, docTemplateContent string
	docTemplate = fmt.Sprint(curThemeDir, "/", conf.Theme, "/doc.tpl")
	if _, err := os.Stat(docTemplate); err != nil {
		docTemplateContent = docDefaultContent //Default template content
	} else {
		docTemplateRead, err := ioutil.ReadFile(docTemplate)
		if err != nil {
			return err
		}
		docTemplateContent = string(docTemplateRead) //File template content
	}
	docTemplateName := template.New("doc")
	docTemplateName, err = docTemplateName.Parse(docTemplateContent)
	if err != nil {
		return err
	}

	// Each document content
	for _, v := range postSourceById {
		markdownDoc, err := ioutil.ReadFile(v.AbsPath)
		if err != nil {
			return err
		}
		if utils.IsExternalLink(v.UrlPath) {
			continue
		}
		var bufDoc bytes.Buffer
		markdown2Html := string(blackfriday.MarkdownCommon(markdownDoc)) // Document html content
		newUrlPath := fmt.Sprint(strings.TrimRight(v.UrlPath, ".md"), ".html")
		newAbsPath := fmt.Sprint(curDistDir, "/", newUrlPath) // The target path to be generated
		docTemplateName.Execute(&bufDoc, map[string]interface{}{"dataTitle": v.Title, "dataMenu": menuTemplateContent, "dataDoc": markdown2Html, "fixLink": conf.FixLink, "languageCode": conf.LanguageCode, "suffixTitle": conf.SuffixTitle, "homeTitle": conf.HomeTitle, "newUrlPath": newUrlPath})
		err = utils.ExistsOrMkdir(path.Dir(newAbsPath))
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(newAbsPath, bufDoc.Bytes(), os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println(v.AbsPath, "--->", "OK")
	}
	////////////Source processing////////////

	// dist css
	themeCssFile := fmt.Sprint(curThemeDir, "/", conf.Theme, "/css/style.css")
	distCssFileDir := fmt.Sprint(curDistDir, "/asset/css")
	var distCssContent []byte
	if _, err := os.Stat(themeCssFile); err == nil {
		distCssContent, err = ioutil.ReadFile(themeCssFile)
		if err != nil {
			return err
		}
	} else {
		distCssContent = []byte(cssDefaultContent)
	}
	if err = utils.ExistsOrMkdir(distCssFileDir); err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprint(distCssFileDir, "/style.css"), distCssContent, os.ModePerm)
	if err != nil {
		return err
	}

	// dist js
	themeJsFile := fmt.Sprint(curThemeDir, "/", conf.Theme, "/js/app.js")
	distJsFileDir := fmt.Sprint(curDistDir, "/asset/js")
	var distJsContent []byte
	if _, err := os.Stat(themeJsFile); err == nil {
		distJsContent, err = ioutil.ReadFile(themeJsFile)
		if err != nil {
			return err
		}
	} else {
		distJsContent = []byte(jsDefaultContent)
	}
	if err = utils.ExistsOrMkdir(distJsFileDir); err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprint(distJsFileDir, "/app.js"), distJsContent, os.ModePerm)
	if err != nil {
		return err
	}

	// static
	if err = utils.ExistsOrMkdir(curStaticDir); err != nil { // Current directory static
		return err
	}
	distStaticDir := fmt.Sprint(curDistDir, "/static") // dist directory static
	if err = utils.ExistsOrMkdir(distStaticDir); err != nil {
		return err
	}
	if err = fsync.Sync(distStaticDir, curStaticDir); err != nil {
		return err
	}

	// Tips
	if err = generateTips(); err != nil {
		return err
	}

	return nil
}

func generateMenuByMap(myMap map[int]*PostSource) string {
	var buf bytes.Buffer
	buf.WriteString("<ul>")
	for _, v := range myMap {
		buf.WriteString("<li><a href=\"")
		if !utils.IsExternalLink(v.UrlPath) {
			buf.WriteString(conf.FixLink)
			buf.WriteString(fmt.Sprint(strings.TrimRight(v.UrlPath, ".md"), ".html"))
		} else {
			buf.WriteString(v.UrlPath)
		}
		buf.WriteString("\">")
		buf.WriteString(v.Title)
		buf.WriteString("</a></li>")
	}
	buf.WriteString("</ul>")
	return buf.String()
}

func generateTips() error {
	err = ioutil.WriteFile(fmt.Sprint(curSrcDir, "/NO-asset-folder.txt"), []byte(""), os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprint(curSrcDir, "/NO-static-folder.txt"), []byte(""), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func StartServer(port string, path string) error {
	output := ""

	if port == "" {
		port = "80"
	}
	output = fmt.Sprint(output, "Your port on: ", port, "\n")
	port = fmt.Sprint(":", port)

	if path == "" {
		path = curDistDir
	}
	absPath, _ := filepath.Abs(path)
	output = fmt.Sprint(output, "Your path on: ", absPath, "\n")

	output = fmt.Sprint(output, "URL is: ", "*", port, "  For example  localhost", port, "\n")
	output = fmt.Sprint(output, "\nPress Ctrl+C to quit.", "\n")

	srv := &http.Server{
		Addr:    port,
		Handler: http.FileServer(http.Dir(path)),
	}
	fmt.Println(output)
	// fmt.Println("Start server is OK.")
	// Print info before ListenAndServe()
	status := srv.ListenAndServe()
	return status
}

func EmptyDist() error {
	if _, err := os.Stat(curDistDir); err != nil {
		return err
	}

	var submit string
	fmt.Print("Empty dist directory? (y or n. Press Ctrl+C to quit):")
	fmt.Scan(&submit)
	if submit == "Y" || submit == "y" {
		err := os.RemoveAll(curDistDir)
		utils.CheckErr(err)
		fmt.Println("Empty dist is OK.")
	} else {
		fmt.Println("No empty dist.")
	}

	return nil
}
