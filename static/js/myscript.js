$(document).ready(function(){
        $('#uploaddataform').bind('submit', function(){
            $(this).ajaxSubmit({ 
                success: function(msg) { 
                        if (msg.result) {
                            alert("登录成功-"+msg.msg);
                            Store.save("username",user_name);
                            window.location = msg.refer;
                        } else{
                            alert("登录失败-"+msg.msg);
                        };
                     }
                //$(this).resetForm(); // 提交后重置表单
            });
            return false
        });
    });