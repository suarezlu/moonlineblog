<div class="layui-container">  
  <div class="layui-row">
    <div class="layui-col-md8">
		<!-- 图片轮播 -->
		<div class="blog-img-carousel">
			<div class="layui-carousel" id="img-carousel">
				<div carousel-item>
					<div><img src="https://www.sunmale.cn/static/blog/images/1.jpg"></img></div>
					<div><img src="https://www.sunmale.cn/static/blog/images/3.jpg"></img></div>
					<div><img src="https://www.sunmale.cn/static/blog/images/1.jpg"></img></div>
					<div><img src="https://www.sunmale.cn/static/blog/images/3.jpg"></img></div>
				</div>
			</div>
		</div>
		<!-- 内容 -->
		<div class="blog-list">
			{{{range $i,$item := .List}}}
				<div class="blog-item">
					<h2 class="blog-title">
						<span class="layui-badge">{{{$item.Category.Name}}}</span>
						<a href="#">{{{$item.Title}}}</a>
					</h2>
					<div class="blog-info">{{{$item.Info}}}</div>
					<div class="blog-foot">
						<i class="layui-icon" style="font-size: 14px;">&#xe60e;</i> {{{date $item.ReleaseTime "Y-m-d H:i"}}}
						<i class="layui-icon" style="font-size: 14px;">&#xe612;</i> by {{{$item.User.Username}}} 
					</div>
				</div>
			{{{end}}}
		</div>
    	<!-- 分页 -->
		<div id="page"></div>
	</div>
    <div class="layui-col-md4">
		<div class="blog-item-right">
			<fieldset class="layui-elem-field layui-field-title">
			  <legend>字段集区块 - 横线风格</legend>
			  <div class="layui-field-box">
			    内容区域
			  </div>
			</fieldset>
		</div>
		<div class="blog-item-right">
			<ul class="layui-timeline">
			  <li class="layui-timeline-item">
			    <i class="layui-icon layui-timeline-axis">&#xe63f;</i>
			    <div class="layui-timeline-content layui-text">
			      <h3 class="layui-timeline-title">8月18日</h3>
			      <p>
			        layui 2.0 的一切准备工作似乎都已到位。发布之弦，一触即发。
			        <br>不枉近百个日日夜夜与之为伴。因小而大，因弱而强。
			        <br>无论它能走多远，抑或如何支撑？至少我曾倾注全心，无怨无悔 <i class="layui-icon"></i>
			      </p>
			    </div>
			  </li>
			  <li class="layui-timeline-item">
			    <i class="layui-icon layui-timeline-axis">&#xe63f;</i>
			    <div class="layui-timeline-content layui-text">
			      <div class="layui-timeline-title">过去</div>
			    </div>
			  </li>
			</ul>
		</div>
    </div>
  </div>
</div>
<script>
//一般直接写在一个js文件中
layui.use(['layer', 'form', 'carousel', 'laypage'], function(){
	
  	var layer = layui.layer, form = layui.form, laypage = layui.laypage;
	layui.carousel.render({
		elem: '#img-carousel',
	    width: '100%',
	    arrow: 'hover',
		anim: 'default',
		indicator: 'none'
	});
	
	laypage.render({
		elem: 'page',
		curr: {{{.PageInfo.Page}}},
		limit: {{{.PageInfo.Limit}}},
		count: {{{.PageInfo.Count}}},
		jump: function(obj, first){
			if (!first){
				var categoryId = {{{.PageInfo.CategoryId}}};
				var url = "";
				if(categoryId>0){
					url = "/?cat="+categoryId+"&page="+obj.curr;
				}else{
					url = "/?page=" + obj.curr;
				}
				location.href = url;
			}
		}
	});
});

</script> 