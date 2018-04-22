var submitting = 0;

// 页面加载时
$(document).ready(function () {
    $("#upload_file_btn").click(function (e) {
        submitting += 1;
        if (submitting > 1) {
            e.preventDefault();
            return;
        }

        var file_info = $("#choose_file_for_upload").val();
        if (file_info == "") {
            $("#upload_file_error_label").removeClass("hidden-self");
            $("#upload_file_error_label").find("span").html("请选择文件");
            e.preventDefault();
            return;
        }

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
                if (data.code == "200") {
                    submitting = 0;
                    $("#upload_file_form").addClass("no-display");
                    $("#upload_file_form_again").removeClass("no-display");
                } else {
                    submitting = 0;
                    $("#upload_file_progress").addClass("hidden-self");
                    $("#upload_file_error_label").removeClass("hidden-self");
                    $("#upload_file_error_label").find("span").html("上传失败");
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                submitting = 0;
                $("#upload_file_progress").addClass("hidden-self");
                $("#upload_file_error_label").removeClass("hidden-self");
                $("#upload_file_error_label").find("span").html("上传失败");
            }
        });
    });

    // 再传一个
    $("#upload_again_btn").click(function (e) {
       window.location.reload();
    });

    // 检测文件大小
    $("#choose_file_for_upload").change(function () {
        if (this.files == null) {
            return;
        }

        if (this.files.length < 1) {
            return;
        }

        $("#upload_file_error_label").addClass("hidden-self");
        $("#upload_file_error_label").find("span").html("");

        var file = this.files[0];
        var length = file.size / 1024.0 / 1024.0;
        length = length.toFixed(2);
        var text = "该文件大小为: " + length + "M";
        $("#check_file_size_label").html(text);
    });
});