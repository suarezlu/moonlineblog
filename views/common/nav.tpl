<div class="layui-header header header-doc">
  <div class="layui-main">
    <a class="logo" href="/">
      	<img src="/static/img/logo.png" alt="layui">
    </a>
	<ul class="layui-nav" lay-filter="">
	  	<li class="layui-nav-item"><a href="/"><i class="layui-icon">&#xe68e;</i> 网站首页</a></li>
		{{{range $i,$item := .Categories}}}
		<li class="layui-nav-item"><a href="/">{{{$item.Name}}}</a></li>
		{{{end}}}
	</ul>
  </div>
</div>