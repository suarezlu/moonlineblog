<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>{{{.Title}}}</title>
  <link rel="stylesheet" href="/static/layui/css/layui.css">
<link rel="stylesheet" href="/static/css/site.css">
<!-- Global site tag (gtag.js) - Google Analytics -->
<!--<script async src="https://www.googletagmanager.com/gtag/js?id=UA-108629114-1"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-108629114-1');
</script>-->
</head>
<body> 
<script src="/static/layui/layui.js"></script>
	{{{.Nav}}}
	<!--内容-->
	{{{.LayoutContent}}}
	<div class="foot">
		<div class="layui-container">
		<p class="foot-info">版权所有 © 2017 Suarez 授权许可遵循</p>
		</div>
	</div>
</body>
</html>