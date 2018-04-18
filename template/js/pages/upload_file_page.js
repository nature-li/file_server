// 页面加载时
$(document).ready(function () {
    $("#upload_file_btn").click(function (e) {
        var bar = $('#upload_file_progress_bar');
        var percent = $('#upload_file_process_percent');

        $('form').ajaxForm({
            dataType: 'json',
            beforeSend: function() {
                $("#upload_file_progress").removeClass("hidden-self");
                var percentVal = '0%';
                bar.width(percentVal);
                percent.html(percentVal);
            },
            uploadProgress: function(event, position, total, percentComplete) {
                var percentVal = percentComplete + '%';
                bar.width(percentVal);
                percent.html(percentVal);
            },
            success: function(data) {
                if (data.code == "0") {
                    $("#upload_file_form").addClass("no-display");
                    $("#upload_file_form_again").removeClass("no-display");
                } else {
                    $("#upload_file_error_label").removeClass("hidden-self");
                    $("#upload_file_error_label").find("span").html("上传失败")
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                $("#upload_file_error_label").removeClass("hidden-self");
                $("#upload_file_error_label").find("span").html("上传失败")
            }
        });
    });

    $("#upload_again_btn").click(function (e) {
       window.location.reload();
    });

    $("#back_to_home").click(function (e) {
        window.location.replace("/");
    })
});