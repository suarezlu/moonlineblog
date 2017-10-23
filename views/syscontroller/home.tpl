<!DOCTYPE html>
<html>
<head>
	 <meta charset="utf-8">
	 <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
	 <title>BLOG管理后台</title>
	 <link rel="stylesheet" href="/static/layui/css/layui.css">
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
  <div class="layui-header">
    <div class="layui-logo">管理后台</div>
    <!-- 头部区域（可配合layui已有的水平导航） -->
    <ul class="layui-nav layui-layout-right">
      <li class="layui-nav-item">
        <a href="javascript:;">
          <img src="/static/img/4743876.jpg" class="layui-nav-img">
          {{{.Auth.User.Username}}}
        </a>
        <dl class="layui-nav-child">
          	<dd><a href="/sys/logout">退出</a></dd>
        </dl>
      </li>
    </ul>
  </div>
  
  <div class="layui-side layui-bg-black">
    <div class="layui-side-scroll">
      <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
      <ul class="layui-nav layui-nav-tree"  lay-filter="side-nav">
			<li class="layui-nav-item layui-nav-itemed">
				<a href="javascript:;">
					<i class="layui-icon">&#xe60a;</i>内容管理
				</a>
				<dl class="layui-nav-child">
					<dd><a href="javascript:;" data-src="/sys/category">&nbsp;&nbsp;&nbsp;&nbsp;分类管理</a></dd>
					<dd><a href="javascript:;" data-src="/sys/articles">&nbsp;&nbsp;&nbsp;&nbsp;文章管理</a></dd>
					<dd><a href="javascript:;" data-src="/sys/article/0">&nbsp;&nbsp;&nbsp;&nbsp;添加文章</a></dd>
				</dl>
			</li>
			<li class="layui-nav-item layui-nav-itemed">
				<a href="javascript:;">
					<i class="layui-icon">&#xe614;</i>设置
				</a>
				<dl class="layui-nav-child">
					<dd><a href="javascript:;" data-src="/sys/pwd">&nbsp;&nbsp;&nbsp;&nbsp;修改密码</a></dd>
					<dd><a href="javascript:;" data-src="/sys/config">&nbsp;&nbsp;&nbsp;&nbsp;站点配置</a></dd>
				</dl>
			</li>
      </ul>
    </div>
  </div>
  
  <div class="layui-body">
    <!-- 内容主体区域 -->
<!--    <div style="padding: 15px;">
	</div>-->
<!--	<iframe src="/sys/category" style="width:100%;height:99%;border:0px none;">
	</iframe>-->
  </div>
  
  <div class="layui-footer">
    <!-- 底部固定区域 -->
    © layui.com - 底部固定区域
  </div>
</div>
<script src="/static/layui/layui.js"></script>
<script>
layui.use(['element','jquery'], function(){
	var element = layui.element;
	var $ = layui.jquery;
	element.on('nav(side-nav)', function(elem){
		var src=$(elem).find('a').attr("data-src");
		var html ='	<iframe src="' + src + '" style="width:100%;height:99%;border:0px none;"></iframe>';
		$(".layui-body").html(html);
	});
});
</script>
</body>
</html>