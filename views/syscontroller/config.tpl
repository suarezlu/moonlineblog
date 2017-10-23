<fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
	<legend>站点配置</legend>
</fieldset>

<form class="layui-form">
	<div class="layui-form-item">
		<label class="layui-form-label">标题</label>
		<div class="layui-input-inline">
			<input name="title" lay-verify="required" class="layui-input" value="{{{.Configs.title}}}" type="text">
		</div>
	</div>
	<div class="layui-form-item">
		<label class="layui-form-label">链接</label>
		<div class="layui-input-inline">
			<input name="url" class="layui-input" type="text" value="{{{.Configs.url}}}">
		</div>
	</div>
  	<div class="layui-form-item">
    	<div class="layui-input-block">
      		<button class="layui-btn" lay-submit lay-filter="*">提交</button>
    	</div>
  	</div>
</form>
<script>
layui.use(['form', 'jquery'], function(){
	layui.form.on('submit(*)', function(data){
		layui.jquery.ajax({
			url:"/sys/config",
			type:'post',
			data:data.field,
			dataType:'json',
			success:function(resp){
				if(resp.code==0){
					layui.layer.msg("修改成功！");
				}else{
					layui.layer.msg(resp.msg);
				}
			}
		});
		return false;
	});
});
</script>