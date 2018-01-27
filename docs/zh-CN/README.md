# EasyDoc 中文文档

EasyDoc，简单、快速生成文档的工具。

### 模板目录基本结构

```html
├─dist    //发布目录
├─src     //存放 .md 源文件及模板文件(必须，支持此目录存放及其子目录存放)
│  ├─theme
│  │  ├─template
│  │  │  ├─doc.tpl  //文档模板文件(非必须，没有则使用系统默认模板)
│  ├─index.md       //首页(非必须)
│  ├─README.md      //生成菜单使用(必须)
├─easydoc  //二进制文件。如 Windows 系统下的 easydoc.exe
```

### 命令行使用

> 确保 easydoc 二进制文件有可执行权限！

二进制文件在当前目录下时：  
Windows系统 `$ easydoc -version`  
类Unix系统(如Mac，Linux系统。注意前面有 ./ ) `$ ./easydoc -version`  
你可以将 easydoc 二进制文件放在全局环境目录下，直接使用`$ easydoc -version`进行使用，更加方便。  

##### EasyDoc 目前支持的命令：  
`-init` 初始文档结构  
`-build` 生成文档  
`-help` 帮助文档  
`-version` 查看 EasyDoc 版本  

生成的静态文件都放在 `dist` 目录，这个目录可以直接拿来当作网站使用。
