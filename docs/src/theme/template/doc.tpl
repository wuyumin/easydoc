<!doctype html>
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
        <div class="ui inverted menu">
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

<script src="https://cdn.bootcss.com/jquery/2.2.3/jquery.min.js"></script>
<script src="https://cdn.bootcss.com/semantic-ui/2.2.13/semantic.min.js"></script>
<!-- <script src="static/js/app.js"></script> -->
<script src="https://wuyumin.github.io/easydoc/dist/static/js/app.js"></script>
</body>
</html>
