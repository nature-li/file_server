// 错误提示框
$.showErr = function (str, func) {
    BootstrapDialog.show({
        type: BootstrapDialog.TYPE_DANGER,
        title: '错误 ',
        message: str,
        size: BootstrapDialog.SIZE_SMALL,
        draggable: true,
        buttons: [{
            label: '关闭',
            action: function (dialogItself) {
                dialogItself.close();
            }
        }],
        onhide: func
    });
};

// 确认对话框
$.showConfirm = function (str, func_ok, func_close) {
    BootstrapDialog.show({
        title: '确认',
        message: str,
        cssClass: 'bootstrap_center_dialog',
        type: BootstrapDialog.TYPE_WARNING,
        size: BootstrapDialog.SIZE_SMALL,
        draggable: true,
        closable: false,
        buttons: [{
            label: '取消',
            action: function (dialogItself) {
                dialogItself.close();
            }
        },{
            label: '确定',
            cssClass: 'btn-warning',
            action: function (dialogItself) {
                dialogItself.close();
                func_ok();
            }
        }],
        onhide: func_close,
    });
};

// 隐藏侧边栏
$(".side-bar-hidden").click(function (e) {
    $("#wrapper").removeClass("toggled");
    $("#expand-side-bar").removeClass("hidden-self");
    $.cookie("pin_nav", 0);
});

// 点击菜单条件隐藏侧边栏
$(".side-bar-condition-hidden").click(function (e) {
    if ($("#lock-side-bar").hasClass("glyphicon-pushpin")) {
        $("#wrapper").removeClass("toggled");
        $("#expand-side-bar").removeClass("hidden-self");
        $.cookie("pin_nav", 0);
    }
});

// 展示侧边栏
$(".side-bar-show").click(function (e) {
    $("#wrapper").addClass("toggled");
    $("#expand-side-bar").addClass("hidden-self");
    $.cookie("pin_nav", 1);
});

// 切换浮动锁
$("#lock-side-bar").click(function (e) {
    if ($("#lock-side-bar").hasClass("glyphicon-pushpin")) {
        $("#lock-side-bar").removeClass("glyphicon-pushpin");
        $("#lock-side-bar").addClass("glyphicon-lock");
        $.cookie("pin_lock", 1);
    } else {
        $("#lock-side-bar").removeClass("glyphicon-lock");
        $("#lock-side-bar").addClass("glyphicon-pushpin");
        $.cookie("pin_lock", 0);
    }
});

// 点击非菜单条件隐藏侧边栏
$(document).click(function (event) {
    if ($("#lock-side-bar").hasClass("glyphicon-pushpin")) {
        if (!$(event.target).closest("#sidebar-wrapper, #expand-side-bar").length) {
            if ($("#wrapper").hasClass("toggled")) {
                $("#wrapper").removeClass("toggled");
                $("#expand-side-bar").removeClass("hidden-self");
                $.cookie("pin_nav", 0);
            }
        }
    }
});

// 点击文件列表
$("#file_list_menu").click(function (e) {
    window.parent.location.replace("/list_file");
});

// 点击上传文件
$("#upload_file_menu").click(function (e) {
    window.parent.location.replace("/upload_file");
});

// 点击下载文件
$("#download_file_menu").click(function (e) {
    $(".side-bar-menu").removeClass("chosen-menu");
    $("#download_file_menu").addClass("chosen-menu");
});

// 点击其它功能
$("#other_function_menu").click(function (e) {
    $(".side-bar-menu").removeClass("chosen-menu");
    $("#other_function_menu").addClass("chosen-menu");
});

// html 编码
function htmlEncode(value) {
    return $("<div/>").text(value).html();
}

// 返回首页
$("#back_to_home").click(function (e) {
    window.location.replace("/");
});