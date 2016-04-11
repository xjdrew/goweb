(function(){
    $("input[type=\"checkbox\"], input[type=\"radio\"]").not("[data-switch-no-init]").bootstrapSwitch();

    // set status
    var val = $('[name="Status"]').val();
    var boxes = $("[data-status-value]");
    for (var i = 0; i < boxes.length; i++) {
        var box = $(boxes[i]);
        box.bootstrapSwitch("state", val&box.data("status-value"));
    }

    $("[data-status-value]").on("switchChange.bootstrapSwitch",function(){
        var val = 0;
        var boxes = $("[data-status-value]");
        for (var i = 0; i < boxes.length; i++) {
            var box = $(boxes[i]);
            if(box.bootstrapSwitch("state")) {
                val = val | box.data("status-value");
            }
        }
        $('[name="Status"]').val(val);
    });
})(this);
