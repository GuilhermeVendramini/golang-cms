// TODO - Predict  more than one file-group by content type
$(document).ready(function () {
    var fGroup = $('#file-exist');

    if(fGroup.length) {
        $('#file-input').hide();
    }

    $('#file-remove').click(function () {
        var pGroup = $(this).parents('.file-group');
        pGroup.find('input[name=file-remove]').val("true")
        pGroup.find('#file-input').show();
        pGroup.find('#file-exist').hide();
    });
});