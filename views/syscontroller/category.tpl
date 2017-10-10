<fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
	<legend>分类管理</legend>
</fieldset>

<button id="categoryAddBtn" class="layui-btn layui-btn-big">添加分类</button>

<table class="layui-hide" id="table_category" lay-filter="categoryEvent"></table> 

<script type="text/html" id="createdTpl">
	{{d.Created.slice(0,19).replace(/T/," ")}}
</script>

<script type="text/html" id="updatedTpl">
	{{d.Updated.slice(0,19).replace(/T/," ")}}
</script>

<script type="text/html" id="barCategory">
	<a class="layui-btn layui-btn-mini" lay-event="setName">编辑</a>
	<a class="layui-btn layui-btn-danger layui-btn-mini" lay-event="delCategory">删除</a>
</script>

<script>
layui.use('table', function(){
	var table = layui.table;
	var $ = layui.jquery;
	
	$('#categoryAddBtn').on('click',function(){
		layer.prompt({
				formType:3,
				title:'添加分类',
				value:''
		},function(value,index){
			$.ajax({
				url:'/sys/categoryAdd',
				data:{Name:value},
				type:'post',
				dataType:'json',
				success:function(resp){
					layer.msg('添加成功！');
				}
			});
		});
	});
  
	// 分类列表渲染
	table.render({
	    elem: '#table_category',
		url: '/sys/categorylist',
		page: true,
		limit:20,
		cols: [[
	    	{field:'Id', title: 'ID', sort: true, width:100},
	    	{field:'Name', title: '分类', width:200},
			{field:'Created', title:'创建时间', width:300, templet:'#createdTpl'},
			{field:'Updated', title:'最后更新时间', width:300, templet:'#updatedTpl'},
			{fixed: 'right', width:150, align:'center', toolbar: '#barCategory'}
		]]
	});

	// 分类名列表事件绑定
	table.on('tool(categoryEvent)',function(obj){
		var data = obj.data;
		// 修改
		if(obj.event==='setName'){
			layer.prompt({
				formType:3,
				title:'修改分类名',
				value:data.Name
			},function(value,index){
				$.ajax({
					url:'/sys/categoryupdate',
					data:{Id:data.Id,Name:data.Name},
					type:'post',
					dataType:'json',
					success:function(resp){
						obj.update({
							Name:value,
							Updated:resp.data.Updated
						});
						layer.close(index);
						layer.msg('修改成功！');
					}
				});
			});
		}
		// 删除
		if(obj.event==='delCategory'){
			layer.confirm('是否删除分类？', {icon: 3, title:'提示'}, function(index){
				$.ajax({
					url:'/sys/categorydel',
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