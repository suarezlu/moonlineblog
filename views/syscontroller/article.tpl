<script type="text/javascript" src="/static/ueditor/ueditor.config.js"></script>
<script type="text/javascript" src="/static/ueditor/ueditor.all.min.js"></script>

<fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
	<legend>添加文章</legend>
</fieldset>

<form class="layui-form" action="" style="margin-right:50px;">
	<input id="articleId" type="text" name="id" value="{{{.Article.Id}}}" style="display:none;">
	<div class="layui-form-item">
	  	<label class="layui-form-label">标题</label>
	  	<div class="layui-input-block">
	    	<input type="text" name="title" value="{{{.Article.Title}}}" placeholder="请输入" autocomplete="off" class="layui-input" lay-verify="required">
	  	</div>
	</div>
    <div class="layui-form-item layui-form-text">
        <label class="layui-form-label">简介</label>
        <div class="layui-input-block">
            <textarea name="info" class="layui-textarea" lay-verify="required">{{{.Article.Info}}}</textarea>
        </div>
    </div>
    <div class="layui-form-item layui-form-text">
        <label class="layui-form-label">内容</label>
        <div class="layui-input-block">
            <textarea id="content" name="content">{{{.Article.Content}}}</textarea>
        </div>
    </div>
  	<div class="layui-form-item">
    	<label class="layui-form-label">分类</label>
   		<div class="layui-input-inline">
    		<select name="category_id" lay-filter="aihao" lay-verify="required">
		        <option value="">请选择分类</option>
				{{{range $i,$item := .Categories}}}
				<option value="{{{$item.Id}}}"  {{{if eq $item.Id $.Article.CategoryId}}}selected{{{end}}}>{{{$item.Name}}}</option>
				{{{end}}}
      		</select>
    	</div>
  	</div>
  	<div class="layui-form-item">
    	<label class="layui-form-label">发布时间</label>
    	<div class="layui-input-inline">
      		<input id="form_rt" type="text" name="release_time" autocomplete="off" class="layui-input" lay-verify="required" readonly="readonly">
    	</div>
  	</div>
  	<div class="layui-form-item">
    	<div class="layui-input-block">
      		<button class="layui-btn" lay-submit lay-filter="*">立即提交</button>
    	</div>
  	</div>
</form>

<script type="text/javascript" charset="utf-8">
window.UEDITOR_HOME_URL = "/static/ueditor/";
var options = {
	"fileUrl":"/sys/upload",
	"filePath":"",
	"imageUrl":"/sys/upload",
	"imagePath":"",
	"initialFrameHeight":"400",
};
var ue = UE.getEditor("content",options);

layui.use(['form', 'layedit', 'laydate', 'jquery'], function(){
	var form = layui.form,layer = layui.layer,laydate = layui.laydate;
	laydate.render({ 
	  	elem: '#form_rt',
		type: 'datetime',
		value: {{{if eq 0 .Article.Id}}}new Date(){{{else}}}new Date({{{.Article.ReleaseTime}}}){{{end}}},
	});
	form.on('submit(*)', function(data){
		data.field.content = ue.getContent();
		layui.jquery.ajax({
			url:'/sys/articlesave',
			type:'post',
			data:data.field,
			dataType:'json',
			success:function(resp){
				if(resp.code==0){
					layui.jquery('#articleId').val(resp.data)
					layer.msg('保存成功！');
				}else{
					layer.msg(resp.msg);
				}
			}
		});
		return false;
	});
});
</script>