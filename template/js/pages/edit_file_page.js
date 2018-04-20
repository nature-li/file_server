var edit_file_page_interval_func = null;

// init_edit_page
function handle_delete_file_response(data) {
    if (data.success != true) {
        $.showErr("删除失败");
        return
    }

    $("#file_desc_info_rows").addClass("no-display");
    $("#back_to_home_tips_row").removeClass("no-display");
    $("#back_to_home_btn_row").removeClass("no-display");
    edit_file_page_interval_func = setInterval(changeReloadTips, 1000);
}

// 定时任务
function changeReloadTips() {
    var seconds = $("#jump_after_seconds").text();
    console.log(seconds);
    if (seconds == "1") {
        if (edit_file_page_interval_func != null) {
            clearInterval(edit_file_page_interval_func);
            edit_file_page_interval_func = null;

            window.location.replace("/");
            return
        }
    }

    seconds = seconds - 1;
    $("#jump_after_seconds").html(seconds);
}

// 页面加载时
$(document).ready(function () {
    $("#delete_file_btn").click(function (e) {
        var file_id = $("#edit_file_id").text();

        // 加载数据
        $.ajax({
                url: '/delete_file_api',
                type: "post",
                data: {
                    'file_id': file_id
                },
                dataType: 'json',
                success: function (data) {
                    handle_delete_file_response(data);
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    if (jqXHR.status == 302) {
                        window.parent.location.replace("/");
                    } else {
                        $.showErr("删除失败");
                    }
                }
            }
        );
    });
});