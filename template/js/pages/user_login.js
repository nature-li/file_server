$(document).ready(function () {
    $("#btn-guest").click(function () {
        $.cookie("pin_nav", 0);
        $.cookie("pin_lock", 0);
        window.location.replace("/");
    });
});