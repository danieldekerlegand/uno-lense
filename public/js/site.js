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

$('#choose-repo').click(function() {
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

$('#connect-to-server').click(function() {
  var serverIP = $('#server-ip').val();
  $.ajax({
    type: "POST",
    url: serverIP + '/connect'
  })
  .done(function(res) {
    console.log(res)
  })
  .error(function(err) {
    console.log(err);
  });
});

$('#start-remote-repo').click(function() {

});

$('#stop-remote-repo').click(function() {

});
