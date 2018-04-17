// 初始化全局变量
function reset_save_data() {
    window.save_data = {
        'item_list': [],
        'db_total_item_count': 0,
        'db_return_item_count': 0,
        'db_max_page_idx': 0,
        'view_max_page_count': 5,
        'view_item_count_per_page': 15,
        'view_start_page_idx': 0,
        'view_current_page_idx': 0,
        'view_current_page_count': 0
    };
}

// 查询并更新页面
function query_and_update_view() {
    var data = {};
    data.success = "true";
    data.item_count = 100;
    data.content = [];
    for (var i = 0; i < 10; i++) {
        var item = {};
        item['id'] = i;
        item['file_name'] = 'name';
        item['version'] = 'version';
        item['md5'] = '1232';
        item['user_name'] = 'lyg';
        item['create_time'] = '2018';
        item['desc'] = 'hello';
        data.content.push(item);
    }
    save_data_and_update_page_view(data);

    // // 获取 ad_network_id
    // var ad_network_id = $("#ad_network_id_selector option:selected").text();
    // if (ad_network_id == 'all_ad_network_id') {
    //     ad_network_id = "all";
    // }
    //
    // // 获取开始日期
    // var start_date = $("#start_date").datepicker('getDate');
    // var start_year = start_date.getFullYear();
    // var start_month = start_date.getMonth() + 1;
    // var start_day = start_date.getDate();
    // var start_dt = start_year + "-" + start_month + "-" + start_day;
    //
    // // 获取结束日期
    // var end_date = $("#end_date").datepicker('getDate');
    // var end_year = end_date.getFullYear();
    // var end_month = end_date.getMonth() + 1;
    // var end_day = end_date.getDate();
    // var end_dt = end_year + "-" + end_month + "-" + end_day;
    //
    // var off_set = window.save_data.view_current_page_idx * window.save_data.view_item_count_per_page;
    // var limit = window.save_data.view_item_count_per_page;
    //
    // // 加载数据
    // $.ajax({
    //         url: '/query_day_page',
    //         type: "post",
    //         data: {
    //             'ad_network_id': ad_network_id,
    //             'start_dt': start_dt,
    //             'end_dt': end_dt,
    //             'off_set': off_set,
    //             'limit': limit
    //         },
    //         dataType: 'json',
    //         success: function (data) {
    //             save_data_and_update_page_view(data);
    //         },
    //         error: function (jqXHR, textStatus, errorThrown) {
    //             if (jqXHR.status == 302) {
    //                 window.parent.location.replace("/");
    //             } else {
    //                 $.showErr("查询失败");
    //             }
    //         }
    //     }
    // );
}

// 更新表格和分页
function update_page_view(page_idx) {
    // 更新表格
    var html = "";
    for (var i = 0; i < window.save_data.item_list.length; i++) {
        var item = window.save_data.item_list[i];
        html += "<tr><td>" + item.id + "</td>" +
            "<td>" + item.file_name + "</td>" +
            "<td>" + item.version + "</td>" +
            "<td>" + item.md5 + "</td>" +
            "<td>" + item.user_name + "</td>" +
            "<td>" + item.create_time + "</td>" +
            "<td>" + item.desc + "</td></tr>";
    }
    $("#file_list_result").find("tr:gt(0)").remove();
    $("#file_list_result").append(html);

    // 更新分页标签
    update_page_partition(page_idx);
}

// 页面加载时
$(document).ready(function () {
    // 定义全局变量
    if (!window.save_data) {
        reset_save_data();
    }

    // 查询数据并更新页面
    query_and_update_view();
});