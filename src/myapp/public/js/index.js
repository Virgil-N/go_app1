require(["jquery", "common"], function($, commonFunction) {
	var COOKIE_NAME = "testMyCookie";

    //判断是否有cookie，动态显示状态，跳转到登录页面
    var cookieValue=commonFunction.getCookie(COOKIE_NAME);
    if(cookieValue.length>0){
        $("#nickname").html("欢迎您，<a href='/accountAction'>"+cookieValue+"</a>");
        $("#login-logout a").text("退出"); 
    }else{
        $("#nickname").html("<a href='/accountAction'></a>");
        $("#login-logout a").text("登录"); 
    }

    //用户退出删除cookie
    $("#login-logout a").click(function(){
        if($(this).text()==="退出"){
            commonFunction.setCookie(COOKIE_NAME, "fail", -1);
        }
    });
    //跳转页面
    $("#go").click(function(event){
        var pageValue=~~($("#pageValue").val());
        if(pageValue.length===0||pageValue<=0){
            alert("请输入正整数！");
            return false;
        }
        var dataObj={
            "pageValue": pageValue
        };
        $.ajax({
           type: "POST",
           url: "goToPage",
           data: dataObj,
           success: function(msg){
                alert( "Data Send: " + msg );
            }
        });
        event.preventDefault();
    });

});
