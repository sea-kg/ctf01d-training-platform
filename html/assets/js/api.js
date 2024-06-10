
if(!window.ctf01d_tp_api) window.ctf01d_tp_api = {};

window.ctf01d_tp_api.games_list = function() {
    return $.ajax({
        url: '/api/games',
        method: 'GET',
    });
}

window.ctf01d_tp_api.game_info = function(game_id) {
    return $.ajax({
        url: '/api/games/' + game_id,
        method: 'GET',
    });
}

window.ctf01d_tp_api.game_create = function(game_data) {
    return $.ajax({
        url: '/api/games',
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(game_data),
    });
}

window.ctf01d_tp_api.auth_signin = function(auth_data) {
    return $.ajax({
        url: '/api/login',
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(auth_data),
      });
}

window.ctf01d_tp_api.auth_session = function() {
    return $.ajax({
        url: '/api/auth/session',
        method: 'GET',
        contentType: 'application/json',
      });
}
