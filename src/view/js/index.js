require(["jquery", "common"], function($, commonFunction) {
	var COOKIE_NAME = "testMyCookie";

    //判断是否有cookie，动态显示状态，跳转到登录页面
    var cookieValue=commonFunction.getCookie(COOKIE_NAME);
    if(cookieValue.length>0){
        $("#nickname").html("欢迎您，<a href='/account'>"+cookieValue+"</a>");
        $("#login-logout a").text("退出"); 
    }else{
        $("#nickname").html("<a href='/account'></a>");
        $("#login-logout a").text("登录"); 
    }

    //用户退出删除cookie
    $("#login-logout a").click(function(){
        if($(this).text()==="退出"){
            commonFunction.setCookie(COOKIE_NAME, "fail", -1);
        }
    });

});
