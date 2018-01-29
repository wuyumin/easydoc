package easydoc

import (
	"github.com/wuyumin/easydoc/utils"
	"github.com/russross/blackfriday"
	"github.com/mostafah/fsync"
	"io/ioutil"
	"text/template"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"path"
	"net/http"
	"path/filepath"
)

var (
	srcStr, distStr, themeStr, staticStr string
	templateDefaultDoc                   = `<!doctype html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<title>{{.dataTitle}}</title>
<meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no">
<link rel="stylesheet" href="https://cdn.bootcss.com/semantic-ui/2.2.13/semantic.min.css">
<!-- <link rel="stylesheet" href="static/css/style.css"> -->
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
            <a href="index.html" class="item">Home</a>
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
<!-- <script src="static/js/app.js"></script> -->
<script src="https://wuyumin.github.io/easydoc/dist/static/js/app.js"></script>
</body>
</html>
`
	markdownSummary = `- [Home](index.md)
`
	markdownIndex = `# Home

This is content.  
You can use markdown to write, EasyDoc will be converted to html content.
`

	cssDefault = `@charset "utf-8";
a{color: #009c95;}
.menu{}
.menu ul{list-style: none;padding-left: 0.5em;}
.menu a{display: block;padding: 0.2em;border-radius:0.5em;font-size: 1.2em;color: rgba(0,0,0,.87);word-break : break-all;}
.menu a:hover{background-color: rgba(0,0,0,.05);color: rgba(0,0,0,.95);}
.menu .made-by{padding: 0.5em;}
.menu .made-by a{text-align: center;}
.content{}
.ui .new-grid{margin-left: 0;margin-right: 0;}

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
	jsDefault = `$(function(){
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
)

func init() {
	// Current directory
	pwd, err := os.Getwd()
	utils.CheckErr(err)

	// Various paths
	srcStr = fmt.Sprint(pwd, "/src/")
	distStr = fmt.Sprint(pwd, "/dist/")
	themeStr = fmt.Sprint(srcStr, "theme/")
	staticStr = fmt.Sprint(srcStr, "static/")
}

func GenerateInit() error {
	// "src" directory
	err := os.MkdirAll(srcStr, 0777)
	if err != nil {
		return err
	}
	// menu file
	err = ioutil.WriteFile(fmt.Sprint(srcStr, "SUMMARY.md"), []byte(markdownSummary), 0777)
	if err != nil {
		return err
	}
	// index file
	err = ioutil.WriteFile(fmt.Sprint(srcStr, "index.md"), []byte(markdownIndex), 0777)
	if err != nil {
		return err
	}
	// css directory
	err = os.MkdirAll(fmt.Sprint(staticStr, "css/"), 0777)
	if err != nil {
		return err
	}
	// css file
	err = ioutil.WriteFile(fmt.Sprint(staticStr, "css/style.css"), []byte(cssDefault), 0777)
	if err != nil {
		return err
	}
	// js directory
	err = os.MkdirAll(fmt.Sprint(staticStr, "js/"), 0777)
	if err != nil {
		return err
	}
	// js file
	err = ioutil.WriteFile(fmt.Sprint(staticStr, "js/app.js"), []byte(jsDefault), 0777)
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
	// static directory
	err = os.MkdirAll(fmt.Sprint(srcStr, "static/"), 0777)
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

func GenerateDoc(isEmptydist bool) error {
	// "src" directory is exist?
	if _, err := os.Stat(srcStr); err != nil {
		return err
	}

	// Empty dist directory
	if isEmptydist {
		err := EmptyDist()
		utils.CheckErr(err)
	}
	// Make dist directory
	err := os.MkdirAll(distStr, 0777)
	utils.CheckErr(err)

	// copy static directory
	err = os.MkdirAll(fmt.Sprint(distStr, "static/"), 0777)
	utils.CheckErr(err)
	err = fsync.Sync(fmt.Sprint(distStr, "static/"), fmt.Sprint(srcStr, "static/"))
	utils.CheckErr(err)

	// Menu content
	markdownMenu, err := ioutil.ReadFile(fmt.Sprint(srcStr, "SUMMARY.md"))
	if err != nil {
		return err
	}
	markdownMenuHtml := strings.Replace(string(blackfriday.MarkdownCommon(markdownMenu)), ".md", ".html", -1) // Menu html content

	// Template content
	var templateFile, templateContent string
	templateFile = fmt.Sprint(themeStr, "template/doc.tpl")
	if _, err := os.Stat(templateFile); err != nil {
		templateContent = templateDefaultDoc //Default template content
	} else {
		templateNewDoc, err := ioutil.ReadFile(templateFile)
		if err != nil {
			return err
		}
		templateContent = string(templateNewDoc) //File template content
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
		if strings.HasPrefix(v[2], "https:") || strings.HasPrefix(v[2], "http:") {
			continue
		}
		markdownDoc, err := ioutil.ReadFile(fmt.Sprint(srcStr, v[2]))
		if err != nil {
			return err
		}
		markdownDocHtml := string(blackfriday.MarkdownCommon(markdownDoc)) // Document html content
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

func StartServer(port string, path string) error {
	output := ""

	if port == "" {
		port = "80"
	}
	output = fmt.Sprint(output, "Your port on: ", port, "\n")
	port = fmt.Sprint(":", port)

	if path == "" {
		path = distStr
	}
	absPath, _ := filepath.Abs(path)
	output = fmt.Sprint(output, "Your path on: ", absPath, "\n")

	output = fmt.Sprint(output, "URL is: ", "*", port, "  For example  localhost", port, "\n")
	output = fmt.Sprint(output, "\nPress Ctrl + C to exit.", "\n")

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
	if _, err := os.Stat(distStr); err != nil {
		return err
	}
	err := os.RemoveAll(distStr)
	return err
}
