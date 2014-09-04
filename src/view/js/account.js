require(["jquery", "common"], function($, commonFunction) {
	//时间选择插件
	// $('#datetimepicker').datetimepicker({
	//     format: 'yyyy-mm-dd hh:ii'
	// });
	
	//重复的代码（与index.js的代码一样，到时候得统一引用）
	var COOKIE_NAME = "testMyCookie";
	//判断用户是否已登录
    var cookieValue=commonFunction.getCookie(COOKIE_NAME);
    if(cookieValue.length>0){
        $("#login-logout a").text("退出"); 
    }else{
        $("#login-logout a").text("登录"); 
    }

    //用户退出删除cookie
    $("#login-logout a").click(function(){
        if($(this).text()==="退出"){
            commonFunction.setCookie(COOKIE_NAME, "fail", -1);
        }
    });

});