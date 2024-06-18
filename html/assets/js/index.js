
var entityMap = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#39;',
    '/': '&#x2F;',
    '`': '&#x60;',
    '=': '&#x3D;'
};

function escapeHtml(string) {
    return String(string).replace(/[&<>"'`=\/]/g, function (s) {
      return entityMap[s];
    });
}

function isRootPage(pathname) {
    return pathname == "/";
}

function isGamesPage(pathname) {
    return pathname == "/games" || pathname == "/games/";
}

function isServicesPage(pathname) {
    return pathname == "/services" || pathname == "/services/";
}

function isTeamPage(pathname) {
    return pathname == "/teams" || pathname == "/teams/";
}

function getHumanTimeHasPassed(dt_end) {
    var dt = dt_end.getTime();
    var dt_now = new Date().getTime();
    var diff = dt_now - dt;
    diff = Math.floor(diff / 1000); // sec
    if (diff < 60) {
        return diff + " seconds ago";
    }
    diff = Math.floor(diff / 60);
    if (diff < 60) {
        return diff + " minutes ago";
    }
    diff = Math.floor(diff / 60);
    if (diff < 24) {
        return diff + " hours ago";
    }
    diff = Math.floor(diff / 24);
    return diff + " days ago";
}

function showSuccessNotification(text) {
    $.notify({
        icon: "add_alert",
        message: text
    }, {
        type: 'success',
        timer: 4000,
        placement: {
            from: 'top',
            align: 'right'
        }
    });
  }

function updateGameTeams(el_id, game_id) {
    window.ctf01d_tp_api.game_info(game_id).fail(function(res) {
        $('#games_page_error').css({
            "display": "block"
        });
        $('#games_page_error').html("Error load info about game #" + game_id);
        console.error(res);
    }).done(function(res) {
        console.log(res)
        var gameTeamsInfoHtml = '<ul class="list-group">';


        for (var i in res['team_details']) {
            var team_info = res['team_details'][i];
            console.log("team_info", team_info);
            gameTeamsInfoHtml += '<li class="list-group-item"><strong>' + escapeHtml(team_info.name) + '</strong> - ' + escapeHtml(team_info.description) + '</li>';
            console.log(team_info)
        }
        gameTeamsInfoHtml += '</ul>';
        $('#' + el_id).html(gameTeamsInfoHtml);
    })
}

function gameCreate() {
    var startTimeInput = $('#game_create_start_time').val();
    var endTimeInput = $('#game_create_end_time').val();
    // convert to ISO 8601
    var startTime = new Date(startTimeInput).toISOString();
    var endTime = new Date(endTimeInput).toISOString();
    var description = $('#game_create_description').val()
    window.ctf01d_tp_api.game_create({
        start_time: startTime,
        end_time: endTime,
        description: description,
    }).fail(function(res) {
        $('#games_page_error').css({
            "display": "block"
        });
        $('#games_page_error').html("Error creating game");
        console.error(res);
    }).done(function(res) {
        console.log(res)
        $('#modal_create_game').modal('hide');
        showSuccessNotification('Game created successfully!')
        renderGamePage();
    })
}

function showLoginForm() {
    $('#sign_error_info').html('')
    $('#sign_error_info').css({"display": "none"});
    $('#signin_username').focus();
    $('#signin_username').unbind();
    $('#signin_username').keypress(function (e) {
        if (e.which == 13) {
            $('#signin_password').focus();
            return false; // <---- Add this line
        }
    });
    $('#signin_password').unbind();
    $('#signin_password').keypress(function (e) {
        if (e.which == 13) {
            doSignin();
            return false; // <---- Add this line
        }
    });
    $('#modal_signin').modal('show');
}

function doSignin() {
    var username = $('#signin_username').val();
    var password = $('#signin_password').val();
    $('#sign_error_info').html('')
    $('#sign_error_info').css({"display": "none"});
    window.ctf01d_tp_api.auth_signin({
        user_name: username,
        password: password,
    }).fail(function(res) {
        $('#sign_error_info').css({"display": ""});
        $('#sign_error_info').html('Signin failed. Check username and password.');
        console.error(res);
    }).done(function(res) {
        console.log(res);
        showSuccessNotification('Login successful!');
        $('#modal_signin').modal('hide');
    })
}

function renderGamePage() {
    $('.spa-web-page').css({"display": ""})
    $('#games_page').css({"display": "block"})
    if (window.location.pathname != "/games/") {
        window.location.href = "/games/";
    }
    window.ctf01d_tp_api.games_list().fail(function(res) {
        $('#games_page_error').css({
            "display": "block"
        });
        $('#games_page_error').html("Error loading games");
        console.error(res);
    }).done(function(res) {
        var gamesHtml = ""
        for (var i in res) {
            var game_info = res[i];
            // console.log("game_info", game_info);
            gamesHtml += '<div href="#" class="list-group-item list-group-item-action flex-column align-items-start">';
            gamesHtml += '  <div class="d-flex w-100 justify-content-between">';
            gamesHtml += '    <h5 class="mb-1">#' + game_info.id + '</h5>';
            gamesHtml += '    <small>' + getHumanTimeHasPassed(new Date(game_info.end_time)) + '</small>';
            gamesHtml += '  </div>';
            gamesHtml += '  <p class="mb-1">' + escapeHtml(game_info.description) + '</p>';
            gamesHtml += '  <small>Начало: ' + new Date(game_info.start_time) + '</small><br>';
            gamesHtml += '  <small>Конец: ' + new Date(game_info.end_time) + '</small><br><br>';
            gamesHtml += '  <div id="game_teams_' + game_info.id + '"> ' + new Date(game_info.end_time) + '</div>';
            gamesHtml += '</div>';
            updateGameTeams('game_teams_' + game_info.id, game_info.id)
        }
        $('#games_page_list').html(gamesHtml);
    })
}

function renderPage(pathname) {
    console.log("pathname", pathname)
    if (pathname == "/") {
        $('.spa-web-page').css({
            "display": ""
        })
        $('#root_page').css({"display": "block"})
    } else if (isGamesPage(pathname)) {
        renderGamePage();
    } else if (isServicesPage(pathname)) {
        $('.spa-web-page').css({
            "display": ""
        })
        $('#services_page').css({"display": "block"})
    } else if (isTeamPage(pathname)) {
        $('#teams_page').css({"display": "block"})
        $('.spa-web-page').css({
            "display": ""
        })
    } else {
        $('.spa-web-page').css({
            "display": ""
        })
        $('#unknown_page').css({"display": "block"})
    }
}


$(document).ready(function () {
    console.log("Ready")
    renderPage(window.location.pathname)

    window.ctf01d_tp_api.auth_session().fail(function(res) {
        console.error(res);
        $('#btn_signin').css({"display": "inline-block"});
        $('#btn_signout').css({"display": "none"});
        $('#btn_profile').css({"display": "none"});
    }).done(function(res) {
        $('#btn_signin').css({"display": "none"});
        $('#btn_signout').css({"display": "inline-block"});
        $('#btn_profile').css({"display": "inline-block"});
        $('#btn_profile').html(res.username + "(" + res.userrole + ")");
    })
})