require(["jquery"], function($){
	//点击发起ajax请求
	$("#login-button").click(function(){
		var username=$("#username").val();
		var password=$("#password").val();
		var formData={
			username: username,
			password: password
		};
		$.ajax({
			type: "POST",
			url: "/login",
			data: formData,
			success: function(responseText){
				if(responseText==="success"){
					window.location.href="/";
				}else{
					alert("用户名或密码错误！");
				}
			}
		});
	});
	$("#register-button").click(function(){
		var username=$("#username").val();
		var password=$("#password").val();
		var formData={
			username: username,
			password: password
		};
		$.ajax({
			type: "POST",
			url: "/register",
			data: formData,
			success: function(responseText){
				if(responseText==="success"){
					alert("注册成功！");
				}else if(responseText==="fail"){
					alert("注册失败，名字已存在！");
				}else if(responseText==="noName"){
					alert("请输入名称！");
				}else{
					alert("error: "+responseText);
				}
			}
		});
	});
});