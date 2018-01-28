# EasyDoc 中文文档

EasyDoc，简单、快速生成文档的工具。

### 可执行文件下载

[从这下载可执行文件](https://github.com/wuyumin/easydoc/releases)

仅一个可执行软件搞定，不用安装，更不用其它依赖，兼容微软系统电脑，苹果系统电脑，Linux系统电脑。

### 模板目录基本结构

```html
├─dist    //发布目录
├─src     //存放 .md 源文件及模板文件(必须，支持存放在此目录及其子目录)
│  ├─theme
│  │  ├─template
│  │  │  ├─doc.tpl  //文档模板文件(非必须，没有则使用系统默认模板)
│  ├─index.md       //首页(非必须)
│  ├─README.md      //生成菜单使用(必须)
├─easydoc  //可执行文件。如 Windows 系统下的 easydoc.exe
```

### 命令行使用

> 确保 easydoc 可执行文件有可执行权限！

可执行文件在当前目录下时：  
Windows系统如 `$ easydoc -version`  
类Unix系统(如Mac，Linux系统。注意前面有 ./ )如 `$ ./easydoc -version`  
你可以将 easydoc 可执行文件放在全局环境目录下，直接使用`$ easydoc -version`进行使用，更加方便。  

##### EasyDoc 目前支持的命令：  
`-init` 初始文档结构  
`-build` 生成文档  
`-help` 帮助文档  
`-version` 查看 EasyDoc 版本  

生成的静态文件都放在`dist`目录，直接**复制该目录**到适当地方当作网站目录来使用。  
注意：再次生成时`dist`目录都会被清空，请不要将重要文件放在里面。

_ _ _

![EasyDoc](https://sfault-image.b0.upaiyun.com/234/289/2342891614-5a6ce27d98153_articlex)
