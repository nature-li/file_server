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

// 解锁侧边栏
function unlock_pin() {
    $("#lock-side-bar").removeClass("glyphicon-triangle-bottom");
    $("#lock-side-bar").addClass("glyphicon-triangle-left");
}

// 锁住侧边栏
function lock_pin() {
    $("#lock-side-bar").removeClass("glyphicon-triangle-left");
    $("#lock-side-bar").addClass("glyphicon-triangle-bottom");
}

// 隐藏侧边栏
$(".side-bar-hidden").click(function (e) {
    e.preventDefault();
    $("#wrapper").removeClass("toggled");
    $("#expand-side-bar").removeClass("hidden-self");
    unlock_pin();
});

// 隐藏侧边栏
$(".side-bar-condition-hidden").click(function (e) {
    if ($("#lock-side-bar").hasClass("glyphicon-triangle-left")) {
        e.preventDefault();
        $("#wrapper").removeClass("toggled");
        $("#expand-side-bar").removeClass("hidden-self");
        unlock_pin();
    }
});

// 展示侧边栏
$(".side-bar-show").click(function (e) {
    e.preventDefault();
    $("#wrapper").addClass("toggled");
    $("#expand-side-bar").addClass("hidden-self");
});

// 加/解锁侧边栏
$("#lock-side-bar").click(function (e) {
    e.preventDefault();
    if ($("#lock-side-bar").hasClass("glyphicon-triangle-left")) {
        lock_pin();
    } else {
        unlock_pin();
    }
});

// 加/解锁侧边栏
$(document).click(function (event) {
    event.preventDefault();
    if ($("#lock-side-bar").hasClass("glyphicon-triangle-left")) {
        if (!$(event.target).closest("#sidebar-wrapper, #expand-side-bar").length) {
            if ($("#wrapper").hasClass("toggled")) {
                $("#wrapper").removeClass("toggled");
                $("#expand-side-bar").removeClass("hidden-self");
            }
        }
    }
});

// 点击文件列表
$("#file_list_menu").click(function (e) {
    e.preventDefault();
    $(".side-bar-menu").removeClass("chosen-menu");
    $("#file_list_menu").addClass("chosen-menu");
});

// 点击上传文件
$("#upload_file_menu").click(function (e) {
    e.preventDefault();
    $(".side-bar-menu").removeClass("chosen-menu");
    $("#upload_file_menu").addClass("chosen-menu");
});

// 点击下载文件
$("#download_file_menu").click(function (e) {
    e.preventDefault();
    $(".side-bar-menu").removeClass("chosen-menu");
    $("#download_file_menu").addClass("chosen-menu");
});

// 点击其它功能
$("#other_function_menu").click(function (e) {
    e.preventDefault();
    $(".side-bar-menu").removeClass("chosen-menu");
    $("#other_function_menu").addClass("chosen-menu");
});