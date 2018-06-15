$(document)
  .ready(function () {
    $('.ui.dropdown')
      .dropdown();

    $('.menu .item')
      .tab();

    $('#login-from')
      .form({
        fields: {
          username: 'empty',
          password: ['minLength[3]', 'empty'],
        }
      });

    $('#login-from').on('keyup keypress', function (e) {
      var keyCode = e.keyCode || e.which;
      if (keyCode === 13) {
        showLoading();
        return false;
      }
    });

    $('#add-pkg-from')
      .form({
        fields: {
          package_name: 'empty',
          root_repo_url: 'empty',
        }
      });

    $('#modify-pkg-from')
      .form({
        fields: {
          root_repo_url: 'empty',
        }
      });

    $('#setting-from')
      .form({
        fields: {
          old_password: 'empty',
          new_password: 'empty',
        }
      });

    $('#add-user-from')
      .form({
        fields: {
          username: 'empty',
          password: 'empty',
        }
      });

    $('#delete-modal-cancel')
      .click(function () {
        $('#delete-modal').modal('hide')
      });
  });

function showDeleteModal(id) {
  $('#delete-modal').find('input[name="id"]').val(id);
  $('#delete-modal').modal('show');
}

function showLoading() {
  $('#loading').dimmer('show')
}

function login() {
  showLoading();
  $('#login-from').form().submit()
}