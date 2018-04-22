$(document).ready(function () {
    $("#btn-guest").click(function () {
        window.location.replace("/");
    });

    $("#btn-login").click(function () {
        $("#login-alert").addClass("no-display");
        var user_email = $("#login-user-email").val();
        var user_password = $("#login-password").val();

        if (!user_email || !user_email.trim()) {
            show_alert_msg("用户名不能为空");
            return;
        }

        if (!user_password || !user_password.trim()) {
            show_alert_msg("密码不能为空");
            return;
        }

        $.ajax({
                url: '/user_login_api',
                type: "post",
                data: {
                    'user_email': user_email,
                    'user_password': user_password
                },
                dataType: 'json',
                success: function (response) {
                    login_callback(response);
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    if (jqXHR.status == 302) {
                        window.parent.location.replace("/");
                    } else {
                        show_alert_msg("登录失败")
                    }
                }
            }
        );
    });

    $("#btn-oa-login").click(function () {
        window.parent.location.replace("/user_login_auth");
    });
});

function login_callback(data) {
    if (data.success != true) {
        show_alert_msg(data.message);
        return;
    }

    window.parent.location.replace("/");
}

function show_alert_msg(data) {
    $("#login-alert").html(data);
    $("#login-alert").removeClass("no-display");
}

