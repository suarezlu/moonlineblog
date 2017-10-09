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
          <img src="https://avatars3.githubusercontent.com/u/4743876?v=4&s=400&u=a2068c53e864a9cb023a34567386512312109f17" class="layui-nav-img">
          {{{.Auth.User.Username}}}
        </a>
        <dl class="layui-nav-child">
			<dd><a href="#">修改密码</a></dd>
          	<dd><a href="/sys/logout">退出</a></dd>
        </dl>
      </li>
    </ul>
  </div>
  
  <div class="layui-side layui-bg-black">
    <div class="layui-side-scroll">
      <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
      <ul class="layui-nav layui-nav-tree"  lay-filter="test">
			<li class="layui-nav-item layui-nav-itemed">
				<a href="javascript:;">
					<i class="layui-icon">&#xe60a;</i>内容管理
				</a>
				<dl class="layui-nav-child">
					<dd><a href="javascript:;" data-src="/sys/category">&nbsp;分类管理</a></dd>
				</dl>
			</li>
<!--        <li class="layui-nav-item layui-nav-itemed">
          <a class="" href="javascript:;">所有商品</a>
          <dl class="layui-nav-child">
            <dd><a href="javascript:;">列表一</a></dd>
            <dd><a href="javascript:;">列表二</a></dd>
            <dd><a href="javascript:;">列表三</a></dd>
            <dd><a href="">超链接</a></dd>
          </dl>
        </li>
        <li class="layui-nav-item">
          <a href="javascript:;">解决方案</a>
          <dl class="layui-nav-child">
            <dd><a href="javascript:;">列表一</a></dd>
            <dd><a href="javascript:;">列表二</a></dd>
            <dd><a href="">超链接</a></dd>
          </dl>
        </li>
        <li class="layui-nav-item"><a href="">云市场</a></li>
        <li class="layui-nav-item"><a href="">发布商品</a></li>-->
      </ul>
    </div>
  </div>
  
  <div class="layui-body">
    <!-- 内容主体区域 -->
<!--    <div style="padding: 15px;">
	</div>-->
	<iframe src="/sys/category" style="width:100%;height:99%;border:0px none;">
	</iframe>
  </div>
  
  <div class="layui-footer">
    <!-- 底部固定区域 -->
    © layui.com - 底部固定区域
  </div>
</div>
<script src="/static/layui/layui.js"></script>
<script>
//JavaScript代码区域
layui.use(['element','jquery'], function(){
  var element = layui.element;
  var $ = layui.jquery;
  
});
</script>
</body>
</html>