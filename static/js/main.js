$(document).ready(function() {
  var submitButton = $("#submit");

  $("#loading").hide();

  submitButton.click(function(event)
  {
    event.preventDefault();
    console.log("Form Submitted.");

    // Parse URL
    pathArray = location.href.split('/');
    protocol = pathArray[0];
    host = pathArray[2];
    baseURL = protocol + '//' + host;
    postURL = baseURL + "/register_user";

    // Parse Form Data
    var formData = {
      "name"    : $('input[name=name]').val(),
      "email"   : $('input[name=email]').val(),
      "uType"   : $('select[name=uType]').val()
    }

    // Process form submit
    $.ajax({
      url : postURL,
      type : "POST",
      data : formData,
    }).done(function(data, textStatus, jqXHR) {
      var status = data.Status
      if(status == 200)
      {
        $("#form").replaceWith("<p class='center success'>Pledge Signed! You a boss.</p>");
        console.log("Added");
      } else if(status == 302) {
        $("#status-message").replaceWith("<p class='center error'>Ahhh! Not added. Already signed or invalid email!</p>");
        console.log("Not Added")
      } else {
        $("#status-message").replaceWith("<p class='center error'>Ahhh! One or more fields missing!</p>");
        console.log("Not Added")
      }

      if(status == 302)
        console.log("hsavda")
    });
  });
});

$(document).on({
    ajaxStart: function() {
       $("#loading").show();
       $("#form").hide();
    },
    ajaxStop: function() {
      $("#form").show();
      $("#loading").hide();
    }
});
