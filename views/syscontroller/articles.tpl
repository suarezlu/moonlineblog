<fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
	<legend>文章管理</legend>
</fieldset>

<table class="layui-hide" id="table_articles" lay-filter="categoryEvent"></table> 

<script type="text/html" id="categoryName">
	<span class="layui-badge layui-bg-black">{{d.Category.Name}}</span>
</script>

<script type="text/html" id="createdTpl">
	{{d.Created.slice(0,19).replace(/T/," ")}}
</script>

<script type="text/html" id="updatedTpl">
	{{d.Updated.slice(0,19).replace(/T/," ")}}
</script>

<script type="text/html" id="barCategory">
	<a class="layui-btn layui-btn-mini" href="/sys/article/{{d.Id}}">编辑</a>
	<a class="layui-btn layui-btn-danger layui-btn-mini" lay-event="delArticle">删除</a>
</script>

<script>
layui.use('table', function(){
	var table = layui.table;
	var $ = layui.jquery;
  
	// 列表渲染
	table.render({
	    elem: '#table_articles',
		url: '/sys/articlelist',
		page: true,
		limit:20,
		cols: [[
	    	{field:'Id', title:'ID', width:80, sort: true},
	    	{field:'Title', title:'标题', width: 200},
			{field:'Category', title:'分类', width:150, templet:'#categoryName'},
			{field:'Created', title:'创建时间', width:164, templet:'#createdTpl'},
			{field:'Updated', title:'最后更新时间', width:164, templet:'#updatedTpl'},
			{fixed: 'right', width:150, align:'center', toolbar:'#barCategory'}
		]]
	});
	// 按钮事件绑定
	table.on('tool(categoryEvent)',function(obj){
		var data = obj.data;
		if(obj.event==='delArticle'){
			layer.confirm('是否删除文章？', {icon: 3, title:'提示'}, function(index){
				$.ajax({
					url:'/sys/articledel',
					data:{Id:data.Id},
					type:'post',
					dataType:'json',
					success:function(resp){
						obj.del();
						layer.close(index);
						layer.msg("删除成功！");
					}
				});
			}); 
		}

	});

});
</script>