// TODO - Predict multiple files by file-group
$(document).ready(function () {
    $(".file-group" ).each(function(index) {
        var exist = $(this).children('#file-exist');
        var self = $(this);
        if(exist.length) {
            $(this).children('#file-input').hide();
        }

        $(this).find('#file-remove').click(function () {
            self.find('input[name=file-remove]').val("true");
            self.children('#file-input').show();
            self.children('#file-exist').hide();
        });
    });
});