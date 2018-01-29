# EasyDoc

### Documentation

- [中文文档](https://wuyumin.github.io/easydoc/dist/zh-CN.html)
- [English Document](https://wuyumin.github.io/easydoc/dist/en.html)

### What is EasyDoc?

EasyDoc, `Easy` to generate `Documents`.

EasyDoc pronunciation [ˈiziˈdɑk] [sound file](https://wuyumin.github.io/easydoc/dist/static/EasyDoc.mp3)

### Binary File for Download

[link to download binary file](https://github.com/wuyumin/easydoc/releases) (Compressed package need to extract the binary file.)

Only one binary file, do not install, not to rely on others, support to Microsoft system computer, Apple system computer, Linux system computer.

EasyDoc is developed by Go language, open source software, you can use the source code to compile. In fact, you do not have to do that, we have compiled and optimized binary file for download.

### Source Directory Basic Structure

Use `-init` command automatically generate

```html
├─dist  //release directory
├─src   //store .md source files and template files(required, support to store in this directory and its subdirectories)
│ ├─static  //Static file directory, this directory will be completely copied to the release directory
│ ├─theme
│ │ ├─template
│ │ │ ├─doc.tpl //document template file (not required, if not, use the software default template)
│ ├─index.md    //Home(not required)
│ ├─SUMMARY.md  //use to generate menu(required)
├─easydoc       //binary file. Such as easydoc.exe on Windows system
```
It is recommended that `.md source file only in src directory` to avoid forgetting to modify the template, web link is error.

### Command Line to Use

> Ensure that easydoc binary file can executable!

Binary file in the current directory:  
Windows system `$ easydoc -version`  
Unix-like system (such as Mac, Linux. Note that in front of ./ ) `$ ./easydoc -version`  
You can put easydoc binary file in the global environment directory, directly using the `$ easydoc -version` for more convenient.

##### EasyDoc Currently Supported Command:

`-init` Init the document structure  
`-build` Build the document  
`-help` Help about EasyDoc  
`-version` Print the version number of EasyDoc  

Static documents which Generated by EasyDoc are placed in the `dist` directory,  **copy this directory** to someplace as a web directory to use.  
Note: `dist` directory will be emptied when it is generated again, please do not put important files in it.

### Contribution

Help us to improve: [submit an issue to us](https://github.com/wuyumin/easydoc/issues) or [pull request to us](https://github.com/wuyumin/easydoc/pulls).

### Who Use EasyDoc

Welcome to provide "Who Use EasyDoc".

- [EasyDoc Document Website](https://wuyumin.github.io/easydoc/dist)
