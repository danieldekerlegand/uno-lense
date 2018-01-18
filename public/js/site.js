var repo;

function pullRepoHandler(e) {
  $.ajax({
    type: "POST",
    url: '/repository/pull',
    data: {'image': e.currentTarget.dataset.name}
  })
}

$('#search-repo').submit(function(e) {
  e.preventDefault();
  e.stopPropagation();

  var options = {
    'repo': $('#search-repo input[name="repo"]').val(),
    'username': $('#search-repo input[name="username"]').val(),
    'password': $('#search-repo input[name="password"]').val()
  };

  $.ajax({
    type: "POST",
    url: '/repository/images',
    data: options
  })
  .done(function(res) {
    console.log(res)
    if (res && res.repositories) {
      res.repositories.forEach(function(repo) {
        $('#repo-search-results').append('<p>' + repo + ' - <button class="btn btn-success pull-repo" data-name="' + repo + '">Pull</button></p>');
      });

      $('.pull-repo').click(pullRepoHandler);
    }
  })
  .error(function(err) {
    console.log(err);
  });
});

$('#choose-repo').click(function(e) {
  console.log("#choose-repo")
  repo = $('#repository-name').val();
  $.ajax({
    type: "GET",
    url: '/local/images'
  })
  .done(function(res) {
    console.log(res)
    if (res) {
      //var repos = JSON.parse(res);
      console.log(res)
      res.forEach(function(repo) {
        $('#image').append('<option value="' + repo.Id + '">' + repo.RepoTags + '</option>');
      });
      $("#create-repo").removeClass("hidden");
    }
  })
  .error(function(err) {
    console.log(err);
  });
});

$('#connect-to-server').submit(function(e) {
  e.preventDefault();
  e.stopPropagation();
  var endpoint = 'http://' + $('#server-ip').val() + ':8080/connect';
  console.log(endpoint)
  $.ajax({
    type: "POST",
    url: endpoint,
    data: {
      name: $('#connect-to-server input[name="name"]').val(),
      s_id: $('#connect-to-server input[name="s_id"]').val()
    }
  })
  .done(function(res) {
    console.log(res)
  })
  .error(function(err) {
    console.log(err);
  });
});

var studentIPs = new Set();

function controlContainer(cmd) {
  studentIPs.forEach(function(ip) {
    var endpoint = 'http://' + ip + ":8080" + cmd;
    console.log('endpoint', endpoint)
    $.ajax({
      type: "POST",
      url: endpoint,
      data: {
        name: "lesson4-comp1"
      }
    })
    .done(function(res) {
      console.log(res)
    })
    .error(function(err) {
      console.log(err);
    });
  });
}

$('.student input[type="checkbox"]').change(function(e) {
  var uri = e.currentTarget.value.split(":")[0];
  if (studentIPs.has(uri)) {
    studentIPs.delete(uri);
  } else {
    studentIPs.add(uri);
  }
  console.log(studentIPs);
});

$('#stop-remote-container').click(function() {
  console.log("test")
  controlContainer("/container/stop")
});

$('#pause-remote-container').click(function() {
  controlContainer("/container/pause");
});

$('#start-remote-container').click(function() {
  controlContainer("/container/start");
});

$('#restart-remote-container').click(function() {
  controlContainer("/container/restart");
});

$('#download-remote-container').click(function() {
  controlContainer("/container/download");
});

var currentImages;

function getLocalImages() {
  $.ajax({
    type: "GET",
    url: "/local/images"
  })
  .done(function(images) {
    if (Array.isArray(images)) {
      currentImages = images;

      images.forEach(function(image) {
        var option = $('<option value="' + image.RepoTags[0] + '">' + image.RepoTags[0] + '</option>');
        $('#image-select').append(option);
      });

      $('#image-select').removeClass('hidden');

      $('#image-select').change(function(e) {
        $('#lesson-name').text(e.currentTarget.value);
      });
    }
  })
  .error(function(err) {
    console.log(err);
  });
}

$(document).ready(function() {
  getLocalImages();
});
