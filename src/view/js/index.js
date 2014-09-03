require(["jquery"], function($) {
	var COOKIE_NAME = "testMyCookie";
    var wrap = {

        //定义创建cookie函数
        setCookie: function(c_name, value, expiredays) {
            var exdate = new Date();
            exdate.setDate(exdate.getDate() + expiredays);
            document.cookie = c_name + "=" + escape(value) +
                ((expiredays == null) ? "" : ";expires=" + exdate.toGMTString());
        },

        //定义获取cookie函数
        getCookie: function(c_name) {
            if (document.cookie.length > 0) {
                c_start = document.cookie.indexOf(c_name + "=");
                if (c_start != -1) {
                    c_start = c_start + c_name.length + 1;
                    c_end = document.cookie.indexOf(";", c_start);
                    if (c_end == -1) c_end = document.cookie.length;
                    return unescape(document.cookie.substring(c_start, c_end));
                }
            }
            return "";
        },
        
    };
    //判断是否有cookie，动态显示状态，跳转到登录页面
    var cookieValue=wrap.getCookie("testMyCookie");
    if(cookieValue&&cookieValue==="success"){
        $("#login-link a").text("退出");
        //删除cookie
        wrap.setCookie("testMyCookie", "fail", -1);
    }

});
