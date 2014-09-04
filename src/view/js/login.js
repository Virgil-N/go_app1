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
});