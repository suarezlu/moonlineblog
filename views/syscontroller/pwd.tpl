<fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
	<legend>修改密码</legend>
</fieldset>

<form class="layui-form">
	<div class="layui-form-item">
		<label class="layui-form-label">原密码</label>
		<div class="layui-input-inline">
			<input name="pwd" lay-verify="required" placeholder="请输入密码" class="layui-input" type="password">
		</div>
	</div>
	<div class="layui-form-item">
		<label class="layui-form-label">新密码</label>
		<div class="layui-input-inline">
			<input name="newpwd" lay-verify="required" placeholder="请输入新密码" class="layui-input" type="password">
		</div>
	</div>
	<div class="layui-form-item">
		<label class="layui-form-label">重复新密码</label>
		<div class="layui-input-inline">
			<input name="repeatpwd" lay-verify="required" autocomplete="off" placeholder="请输入新密码" class="layui-input" type="password">
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
	//var form = layui.form,layer = layui.layer,laydate = layui.laydate;
	layui.form.on('submit(*)', function(data){
		console.log(data.field);
		layui.jquery.ajax({
			url:"/sys/pwd",
			type:'post',
			data:data.field,
			dataType:'json',
			success:function(resp){
				if(resp.code==0){
					layui.layer.msg("修改成功！");
				}else{
					layui.layer.msg(resp.msg);
				}
				console.log(resp);
			}
		});
		return false;
	});
});
</script>