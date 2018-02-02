# EasyDoc v1 旧版本中文文档

<iframe src="https://ghbtns.com/github-btn.html?user=wuyumin&repo=easydoc&type=star&count=true" frameborder="0" scrolling="0" width="170px" height="20px"></iframe>

EasyDoc，简单、快速生成文档的工具。

EasyDoc 读音 [ˈiziˈdɑk] [语音文件](https://wuyumin.github.io/easydoc/dist/static/EasyDoc.mp3)

### 互动·交流

- QQ交流群：群号码 80998448 [加入QQ群](https://shang.qq.com/wpa/qunwpa?idkey=e8c0258f779fa73a7d503871d2ff0f8da5698233b79f4e29836471a1d7491494)
- GitHub: <https://github.com/wuyumin/easydoc> 欢迎 star 它

### 软件下载

[从这下载软件](https://github.com/wuyumin/easydoc/releases) (压缩包需要解压出软件文件。)

仅一个软件文件搞定，不用安装，更不用其它依赖，支持微软系统电脑，苹果系统电脑，Linux系统电脑。

EasyDoc 使用 Go 语言开发，是开源软件，你可以自行使用源码进行编译。其实你可不必这么做，我们已经有编译并优化好的软件来下载。

### 命令行的使用

> 确保 easydoc 软件文件有可执行权限！

软件文件在当前目录下时：  
Windows系统 `$ easydoc -version`  
类Unix系统(如Mac，Linux系统。注意前面有 ./ ) `$ ./easydoc -version`  
你可以将 easydoc 软件文件放在全局环境目录下(推荐此做法)，直接使用`$ easydoc -version`进行使用。  

##### EasyDoc 目前支持的命令：  

`-init` 初始化文档结构  
`-build` 生成文档  
`-server` 启动 web 服务(可以配合[或不配合]端口`-port`和路径`-path`一起使用，默认端口是 80 ，默认路径是 dist 目录)  
`-emptydist` 清空 dist 目录  
`-help` 帮助文档  
`-version` 查看 EasyDoc 版本  

生成的静态文件都放在`dist`目录，直接使用或复制该目录当网站目录。

### 源目录基本结构

使用 `-init` 命令自动生成

```html
├─dist    //发布目录
├─src     //存放 .md 源文件及模板文件(必须，支持存放在此目录及其子目录)
│  ├─static  //静态文件目录，此目录会完整地复制到发布目录
│  ├─theme
│  │  ├─template
│  │  │  ├─doc.tpl  //文档模板文件(非必须，没有则使用软件默认模板)
│  ├─index.md       //首页(非必须，但推荐)
│  ├─SUMMARY.md     //生成菜单使用(必须)
├─easydoc.exe           //软件文件(推荐放在全局环境目录下)
```
- `建议将 .md 源文件仅放在 src 根目录下`，避免忘记模板修改而生成的网页链接路径不对，造成无法访问。
- 源文件使用 Markdown 语法编写。

### 贡献

GitHub: <https://github.com/wuyumin/easydoc> 欢迎star它  
建议或帮我们改进：[提交 issue 给我们](https://github.com/wuyumin/easydoc/issues) 或者 [提交 pull request 给我们](https://github.com/wuyumin/easydoc/pulls)。

### 命令操作示例动图

![EasyDoc](https://wuyumin.github.io/easydoc/dist/static/EasyDoc.gif)

### 谁在使用 EasyDoc

欢迎你提供使用 EasyDoc 的网站，方便我们收录。

- [EasyDoc 文档中心](https://wuyumin.github.io/easydoc)
